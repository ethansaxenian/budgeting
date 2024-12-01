package tui

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethansaxenian/budgeting/database"
)

type item database.Month

func (i item) Title() string       { return fmt.Sprintf("%s %d", i.Month, i.Year) }
func (i item) Description() string { return fmt.Sprintf("%d", i.Year) }
func (i item) FilterValue() string { return i.Title() }

type monthsState struct {
	list list.Model
}

func (m model) monthsView() string {
	return m.state.months.list.View()
}

func (m model) monthsUpdate(msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case onSwitchPageMsg:
		m = m.monthsRefresh()

	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			m = m.monthsRefresh()
		case "esc":
			if !m.hasMonthsFilterOpen() {
				m, cmd = m.switchPage(m.lastPage, nil)
			}
		case "enter":
			if !m.hasMonthsFilterOpen() {
				m, cmd = m.selectMonth(m.state.months.list.SelectedItem())
			}
		}
	}

	cmds := []tea.Cmd{cmd}
	m.state.months.list, cmd = m.state.months.list.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func monthsInit(db *sql.DB) monthsState {
	months := getMonthItems(db)

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.SetSpacing(0)

	keyMap := list.KeyMap{
		// Browsing.
		CursorUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		PrevPage: key.NewBinding(
			key.WithKeys("left", "h", "pgup", "b", "u"),
			key.WithHelp("←/h/pgup", "prev page"),
		),
		NextPage: key.NewBinding(
			key.WithKeys("right", "l", "pgdown", "f", "d"),
			key.WithHelp("→/l/pgdn", "next page"),
		),
		GoToStart: key.NewBinding(
			key.WithKeys("home", "g"),
			key.WithHelp("g/home", "go to start"),
		),
		GoToEnd: key.NewBinding(
			key.WithKeys("end", "G"),
			key.WithHelp("G/end", "go to end"),
		),
		Filter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "filter"),
		),
		ClearFilter: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "clear filter"),
		),

		// Filtering.
		CancelWhileFiltering: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel"),
		),
		AcceptWhileFiltering: key.NewBinding(
			key.WithKeys("enter", "tab", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
			key.WithHelp("enter", "apply filter"),
		),
	}

	l := list.New(months, delegate, 20, 40)
	l.KeyMap = keyMap
	l.Title = "Select a Month"
	l.SetShowHelp(true)
	return monthsState{list: l}
}

func (m model) hasMonthsFilterOpen() bool {
	return m.state.months.list.IsFiltered() || m.state.months.list.SettingFilter()
}

func (m model) selectMonth(selected list.Item) (model, tea.Cmd) {
	i, ok := selected.(item)
	if !ok {
		return m, nil
	}

	m.month = database.Month(i)
	m, cmd := m.switchPage(m.lastPage, nil)
	return m, cmd
}

func (m model) monthsRefresh() model {
	months := getMonthItems(m.db)

	m.state.months.list.SetItems(months)
	m.state.months.list.ResetFilter()

	return m
}

func getMonthItems(db *sql.DB) []list.Item {
	ctx := context.Background()

	q := database.New(db)

	months, err := q.GetAllMonths(ctx)
	if err != nil {
		months = []database.Month{}
	}

	items := []list.Item{}
	for _, m := range months {
		items = append(items, item(m))
	}

	return items
}
