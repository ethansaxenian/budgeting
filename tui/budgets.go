package tui

import (
	"context"
	"database/sql"
	"slices"
	"sort"
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
)

type budgetsState struct {
	expenseTable table.Model
	incomeTable  table.Model
	focusedTable database.TransactionType
}

func (m model) budgetsView() string {
	focusedTable := m.state.budgets.focusedTable
	expenseTable := m.state.budgets.expenseTable
	incomeTable := m.state.budgets.incomeTable

	expenseTableView := lipgloss.JoinVertical(lipgloss.Center, tableHeaderStyle.Render("Expenses"), expenseTable.View())
	incomeTableView := lipgloss.JoinVertical(lipgloss.Center, tableHeaderStyle.Render("Income"), incomeTable.View())

	if focusedTable == database.TransactionTypeExpense {
		return lipgloss.JoinHorizontal(lipgloss.Top, focusedModelStyle.Render(expenseTableView), modelStyle.Render(incomeTableView))
	} else {
		return lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(expenseTableView), focusedModelStyle.Render(incomeTableView))
	}
}

func (m model) budgetsUpdate(msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case onSwitchPageMsg:
		m = m.budgetsRefresh()

	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			m.state.budgets.expenseTable.Focus()
			m.state.budgets.incomeTable.Blur()
			m.state.budgets.focusedTable = database.TransactionTypeExpense
		case "l":
			m.state.budgets.incomeTable.Focus()
			m.state.budgets.expenseTable.Blur()
			m.state.budgets.focusedTable = database.TransactionTypeIncome
		case "r":
			m = m.budgetsRefresh()
		case "t":
			m, cmd = m.switchPage(transactionsPage, nil)
		case "n":
			m, cmd = m.switchPage(editorPage, nil)
		}
	}

	cmds := []tea.Cmd{cmd}
	m.state.budgets.expenseTable, cmd = m.state.budgets.expenseTable.Update(msg)
	cmds = append(cmds, cmd)
	m.state.budgets.incomeTable, cmd = m.state.budgets.incomeTable.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func budgetsInit(db *sql.DB, month database.Month) budgetsState {
	columns := []table.Column{
		{Title: "ID", Width: 0},
		{Title: "Category", Width: 15},
		{Title: "Budget", Width: 15},
		{Title: "Actual", Width: 15},
		{Title: "Difference", Width: 30},
	}

	expenses, income := getBudgets(db, month)

	expenseRows := formatBudgetsToRows(expenses, database.TransactionTypeExpense)
	incomeRows := formatBudgetsToRows(income, database.TransactionTypeIncome)

	keyMap := table.KeyMap{
		LineUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		LineDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		HalfPageUp: key.NewBinding(
			key.WithKeys("u", "ctrl+u"),
			key.WithHelp("u", "½ page up"),
		),
		HalfPageDown: key.NewBinding(
			key.WithKeys("d", "ctrl+d"),
			key.WithHelp("d", "½ page down"),
		),
		GotoTop: key.NewBinding(
			key.WithKeys("home", "g"),
			key.WithHelp("g/home", "go to start"),
		),
		GotoBottom: key.NewBinding(
			key.WithKeys("end", "G"),
			key.WithHelp("G/end", "go to end"),
		),
	}
	expenseTable := table.New(
		table.WithColumns(columns),
		table.WithRows(expenseRows),
		table.WithFocused(true),
		table.WithHeight(len(util.EXPENSE_CATEGORIES)+2),
		table.WithKeyMap(keyMap),
	)

	incomeTable := table.New(
		table.WithColumns(columns),
		table.WithRows(incomeRows),
		table.WithFocused(false),
		table.WithHeight(len(util.INCOME_CATEGORIES)+2),
		table.WithKeyMap(keyMap),
	)

	return budgetsState{expenseTable: expenseTable, incomeTable: incomeTable, focusedTable: database.TransactionTypeExpense}
}

func formatBudgetsToRows(budgets []database.BudgetItem, transactionType database.TransactionType) []table.Row {
	rows := []table.Row{}

	for _, b := range budgets {
		row := []string{
			strconv.Itoa(b.BudgetID),
			string(b.Category),
			util.FormatAmountWithDollarSign(b.Planned),
			util.FormatAmountWithDollarSign(b.Actual),
		}

		switch transactionType {
		case database.TransactionTypeExpense:
			differenceStyle := colorAmount(b.Planned - b.Actual)
			row = append(row, differenceStyle.Render(util.FormatAmountWithDollarSign(b.Planned-b.Actual)))
		case database.TransactionTypeIncome:
			differenceStyle := colorAmount(b.Actual - b.Planned)
			row = append(row, differenceStyle.Render(util.FormatAmountWithDollarSign(b.Actual-b.Planned)))
		}

		rows = append(rows, row)
	}

	sumPlanned, sumActual := 0.0, 0.0

	for _, b := range budgets {
		sumPlanned += b.Planned
		sumActual += b.Actual
	}

	totalRow := []string{
		"",
		bold.Render("Total"),
		bold.Render(util.FormatAmountWithDollarSign(sumPlanned)),
		bold.Render(util.FormatAmountWithDollarSign(sumActual)),
	}
	switch transactionType {
	case database.TransactionTypeExpense:
		differenceStyle := colorAmount(sumPlanned - sumActual)
		totalRow = append(totalRow, differenceStyle.Render(util.FormatAmountWithDollarSign(sumPlanned-sumActual)))
	case database.TransactionTypeIncome:
		differenceStyle := colorAmount(sumActual - sumPlanned)
		totalRow = append(totalRow, differenceStyle.Render(util.FormatAmountWithDollarSign(sumActual-sumPlanned)))
	}
	rows = append(rows, totalRow)

	return rows
}

func (m model) budgetsRefresh() model {
	expenses, income := getBudgets(m.db, m.month)

	expenseRows := formatBudgetsToRows(expenses, database.TransactionTypeExpense)
	incomeRows := formatBudgetsToRows(income, database.TransactionTypeIncome)

	m.state.budgets.expenseTable.SetRows(expenseRows)
	m.state.budgets.incomeTable.SetRows(incomeRows)

	return m
}

func getBudgets(db *sql.DB, month database.Month) ([]database.BudgetItem, []database.BudgetItem) {
	ctx := context.Background()

	q := database.New(db)

	finalExpenses := []database.BudgetItem{}

	expenses, err := q.GetBudgetItemsByMonthIDAndTransactionType(
		ctx,
		database.GetBudgetItemsByMonthIDAndTransactionTypeParams{MonthID: month.ID,
			TransactionType: database.TransactionTypeExpense,
		},
	)
	if err != nil {
		expenses = []database.BudgetItem{}
	}

	for _, e := range expenses {
		if slices.Contains(util.EXPENSE_CATEGORIES, e.Category) {
			finalExpenses = append(finalExpenses, e)
		}
	}

	finalIncome := []database.BudgetItem{}

	income, err := q.GetBudgetItemsByMonthIDAndTransactionType(
		ctx,
		database.GetBudgetItemsByMonthIDAndTransactionTypeParams{MonthID: month.ID,
			TransactionType: database.TransactionTypeIncome,
		},
	)
	if err != nil {
		income = []database.BudgetItem{}
	}

	for _, i := range income {
		if slices.Contains(util.INCOME_CATEGORIES, i.Category) {
			finalIncome = append(finalIncome, i)
		}
	}

	sort.Slice(expenses, func(i, j int) bool {
		return string(expenses[i].Category) < string(expenses[j].Category)
	})

	sort.Slice(income, func(i, j int) bool {
		return string(income[i].Category) < string(income[j].Category)
	})

	return finalExpenses, finalIncome
}
