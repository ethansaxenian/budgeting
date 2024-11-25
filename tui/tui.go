package tui

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type page int

const (
	newTransactionPage page = iota
	transactionsPage
	budgetsPage
)

type state struct {
	transactions   transactionsState
	newTransaction newTransactionState
}

type model struct {
	db    *sql.DB
	month database.Month
	page  page
	state state
	err   error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case error:
		m.err = msg
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		// case "esc":
		// 	if m.transactionsTable.Focused() {
		// 		m.transactionsTable.Blur()
		// 	} else {
		// 		m.transactionsTable.Focus()
		// 	}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	switch m.page {
	case newTransactionPage:
		m, cmd = m.newTransactionUpdate(msg)
	case transactionsPage:
		m, cmd = m.transactionsUpdate(msg)
	}

	return m, cmd
}

func (m model) View() string {
	content := strings.Builder{}
	content.WriteString(m.month.Month.String())
	content.WriteString(" ")
	content.WriteString(strconv.Itoa(m.month.Year))
	content.WriteString("\n\n")

	switch m.page {
	case newTransactionPage:
		content.WriteString(m.newTransactionView())
	case transactionsPage:
		content.WriteString(m.transactionsView())
	}

	// var renderedTabs []string
	//
	// for _, t := range m.Tabs {
	// 	renderedTabs = append(renderedTabs, t.header())
	// }

	// content.WriteString("\n\n")
	// content.WriteString(strings.Join(renderedTabs, "\t"))
	// content.WriteString("\n\n")
	// content.WriteString(m.page.getContent())
	return content.String()
}

func NewModel() (model, error) {
	ctx := context.Background()

	databaseURL := os.Getenv("DB_URL")

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return model{}, err
	}
	q := database.New(db)

	currYear, currMonth, _ := util.CurrentDate()

	month, err := q.GetMonthByMonthAndYear(ctx, database.GetMonthByMonthAndYearParams{Month: currMonth, Year: currYear})
	if err != nil {
		return model{}, err
	}
	return model{
		db:    db,
		month: month,
		page:  transactionsPage,
		state: state{
			transactions:   transactionsInit(db, month),
			newTransaction: newTransactionInit(),
		},
		err: nil,
	}, nil
}

type switchPageMsg struct{}

func (m model) switchPage(p page) (model, tea.Cmd) {
	m.page = p
	return m, func() tea.Msg { return switchPageMsg{} }
}
