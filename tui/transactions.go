package tui

import (
	"context"
	"database/sql"
	"sort"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
)

var (
	modelStyle        = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#555555"))
	tableHeaderStyle  = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Bold(true).PaddingBottom(1)
)

type transactionsState struct {
	expenseTable table.Model
	incomeTable  table.Model
	focusedTable database.TransactionType
}

func (m model) transactionsRefresh() model {
	expenseRows := []table.Row{}
	incomeRows := []table.Row{}

	expenses, income := getTransactions(m.db, m.month)

	for _, t := range expenses {
		expenseRows = append(expenseRows, []string{util.FormatDate(t.Date), util.FormatAmountWithDollarSign(t.Amount), t.Description, string(t.Category)})
	}

	for _, t := range income {
		incomeRows = append(incomeRows, []string{util.FormatDate(t.Date), util.FormatAmountWithDollarSign(t.Amount), t.Description, string(t.Category)})
	}

	m.state.transactions.expenseTable.SetRows(expenseRows)
	m.state.transactions.incomeTable.SetRows(incomeRows)

	return m
}

func (m model) transactionsView() string {
	focusedTable := m.state.transactions.focusedTable
	expenseTable := m.state.transactions.expenseTable
	incomeTable := m.state.transactions.incomeTable

	expenseTableView := lipgloss.JoinVertical(lipgloss.Center, tableHeaderStyle.Render("Expenses"), expenseTable.View())
	incomeTableView := lipgloss.JoinVertical(lipgloss.Center, tableHeaderStyle.Render("Income"), incomeTable.View())

	if focusedTable == database.TransactionTypeExpense {
		return lipgloss.JoinHorizontal(lipgloss.Top, focusedModelStyle.Render(expenseTableView), modelStyle.Render(incomeTableView))
	} else {
		return lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(expenseTableView), focusedModelStyle.Render(incomeTableView))
	}
}

func (m model) transactionsUpdate(msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case switchPageMsg:
		m = m.transactionsRefresh()

	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			m.state.transactions.expenseTable.Focus()
			m.state.transactions.incomeTable.Blur()
			m.state.transactions.focusedTable = database.TransactionTypeExpense
		case "l":
			m.state.transactions.incomeTable.Focus()
			m.state.transactions.expenseTable.Blur()
			m.state.transactions.focusedTable = database.TransactionTypeIncome
		case "r":
			m = m.transactionsRefresh()
		case "n":
			m, cmd = m.switchPage(newTransactionPage)
		}
	}

	cmds := []tea.Cmd{cmd}
	m.state.transactions.expenseTable, cmd = m.state.transactions.expenseTable.Update(msg)
	cmds = append(cmds, cmd)
	m.state.transactions.incomeTable, cmd = m.state.transactions.incomeTable.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func transactionsInit(db *sql.DB, month database.Month) transactionsState {
	columns := []table.Column{
		{Title: "Date", Width: 15},
		{Title: "Amount", Width: 10},
		{Title: "Description", Width: 15},
		{Title: "Category", Width: 15},
	}

	expenseRows := []table.Row{}
	incomeRows := []table.Row{}

	expenses, income := getTransactions(db, month)

	for _, t := range expenses {
		expenseRows = append(expenseRows, []string{util.FormatDate(t.Date), util.FormatAmountWithDollarSign(t.Amount), t.Description, string(t.Category)})
	}

	for _, t := range income {
		incomeRows = append(incomeRows, []string{util.FormatDate(t.Date), util.FormatAmountWithDollarSign(t.Amount), t.Description, string(t.Category)})
	}

	expenseTable := table.New(
		table.WithColumns(columns),
		table.WithRows(expenseRows),
		table.WithFocused(true),
		table.WithHeight(50),
	)

	incomeTable := table.New(
		table.WithColumns(columns),
		table.WithRows(incomeRows),
		table.WithFocused(false),
		table.WithHeight(50),
	)

	return transactionsState{expenseTable: expenseTable, incomeTable: incomeTable, focusedTable: database.TransactionTypeExpense}
}

func getTransactions(db *sql.DB, month database.Month) ([]database.Transaction, []database.Transaction) {
	ctx := context.Background()

	q := database.New(db)

	expenses, err := q.GetTransactionsByMonthIDAndType(ctx, database.GetTransactionsByMonthIDAndTypeParams{ID: month.ID, TransactionType: database.TransactionTypeExpense})
	if err != nil {
		expenses = []database.Transaction{}
	}

	income, err := q.GetTransactionsByMonthIDAndType(ctx, database.GetTransactionsByMonthIDAndTypeParams{ID: month.ID, TransactionType: database.TransactionTypeIncome})
	if err != nil {
		income = []database.Transaction{}
	}

	sort.Slice(expenses, func(i, j int) bool {
		return expenses[i].Date.After(expenses[j].Date)
	})

	sort.Slice(income, func(i, j int) bool {
		return income[i].Date.After(income[j].Date)
	})

	return expenses, income
}
