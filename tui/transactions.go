package tui

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
)

type transactionsState struct {
	expenseTable table.Model
	incomeTable  table.Model
	focusedTable database.TransactionType
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
	case onSwitchPageMsg:
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
		case "b":
			m, cmd = m.switchPage(budgetsPage, nil)
		case "n":
			m, cmd = m.switchPage(editorPage, nil)
		case "m":
			m, cmd = m.switchPage(monthsPage, nil)
		case "enter":
			m, cmd = m.switchPage(editorPage, rowToTransaction(m.selectedTransactionRow(), m.state.transactions.focusedTable))
		case "backspace":
			m.err = m.deleteTransaction()
			m = m.transactionsRefresh()
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
		{Title: "ID", Width: 0},
		{Title: "Date", Width: 15},
		{Title: "Amount", Width: 10},
		{Title: "Description", Width: 15},
		{Title: "Category", Width: 15},
	}

	expenses, income := getTransactions(db, month)
	expenseRows := formatTransactionsToRows(expenses)
	incomeRows := formatTransactionsToRows(income)

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
		table.WithHeight(40),
		table.WithKeyMap(keyMap),
	)

	incomeTable := table.New(
		table.WithColumns(columns),
		table.WithRows(incomeRows),
		table.WithFocused(false),
		table.WithHeight(40),
		table.WithKeyMap(keyMap),
	)

	return transactionsState{expenseTable: expenseTable, incomeTable: incomeTable, focusedTable: database.TransactionTypeExpense}
}

func formatTransactionsToRows(transactions []database.Transaction) []table.Row {
	rows := []table.Row{}

	for _, t := range transactions {
		rows = append(rows, []string{strconv.Itoa(t.ID), util.FormatDate(t.Date), util.FormatAmountWithDollarSign(t.Amount), t.Description, string(t.Category)})
	}

	return rows
}

func (m model) transactionsRefresh() model {
	expenses, income := getTransactions(m.db, m.month)

	expenseRows := formatTransactionsToRows(expenses)
	incomeRows := formatTransactionsToRows(income)

	m.state.transactions.expenseTable.SetRows(expenseRows)
	m.state.transactions.incomeTable.SetRows(incomeRows)

	return m
}

func (m model) selectedTransactionRow() table.Row {
	var table table.Model
	if m.state.transactions.focusedTable == database.TransactionTypeExpense {
		table = m.state.transactions.expenseTable
	} else {
		table = m.state.transactions.incomeTable
	}
	return table.SelectedRow()
}

func rowToTransaction(row table.Row, transactionType database.TransactionType) database.Transaction {
	id, _ := strconv.Atoi(row[0])
	date, _ := util.ParseDate(row[1])
	amount, _ := strconv.ParseFloat(strings.TrimPrefix(row[2], "$"), 64)
	return database.Transaction{
		ID:              id,
		Date:            date,
		Amount:          amount,
		Description:     row[3],
		Category:        database.Category(row[4]),
		TransactionType: transactionType,
	}
}

func (m model) deleteTransaction() error {
	ctx := context.Background()
	q := database.New(m.db)

	id, _ := strconv.Atoi(m.selectedTransactionRow()[0])
	if err := q.DeleteTransaction(ctx, id); err != nil {
		return fmt.Errorf("Error deleting transaction %d", id)
	}

	return nil

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
