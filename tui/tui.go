package tui

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

var (
	errorStyle = lipgloss.NewStyle().Foreground(red)
)

type page int

const (
	editorPage page = iota
	transactionsPage
	budgetsPage
)

type state struct {
	transactions transactionsState
	budgets      budgetsState
	editor       editorState
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
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	switch m.page {
	case editorPage:
		m, cmd = m.editorUpdate(msg)
	case transactionsPage:
		m, cmd = m.transactionsUpdate(msg)
	case budgetsPage:
		m, cmd = m.budgetsUpdate(msg)
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
	case editorPage:
		content.WriteString(m.editorView())
	case transactionsPage:
		content.WriteString(m.transactionsView())
	case budgetsPage:
		content.WriteString(m.budgetsView())
	}

	if m.err != nil {
		content.WriteString("\n")
		content.WriteString(errorStyle.Render(m.err.Error()))
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

	databaseURL := os.Getenv("DATABASE_URL")

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
			transactions: transactionsInit(db, month),
			budgets:      budgetsInit(db, month),
			editor:       editorInit(),
		},
		err: nil,
	}, nil
}

type onSwitchPageMsg struct {
	data any
}

func (m model) switchPage(p page, data any) (model, tea.Cmd) {
	m.page = p
	m.err = nil
	return m, func() tea.Msg { return onSwitchPageMsg{data} }
}
