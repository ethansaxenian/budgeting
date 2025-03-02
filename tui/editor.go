package tui

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethansaxenian/budgeting/database"
	"github.com/ethansaxenian/budgeting/util"
)

const (
	date = iota
	amount
	description
	category
	transactionType
	submit
)

var inputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))

type editorState struct {
	inputs  []textinput.Model
	focused int
	currID  *int
}

func (m model) editorView() string {
	inputs := m.state.editor.inputs
	var header = "New Transaction:"
	if m.state.editor.currID != nil {
		header = fmt.Sprintf("Editing Transaction %d:", *m.state.editor.currID)
	}
	content := fmt.Sprintf(
		`%s
%s
%s

%s
%s

%s
%s

%s
%s

%s
%s

`,
		header,
		inputStyle.Width(11).Render("Date"),
		inputs[date].View(),
		inputStyle.Width(8).Render("Amount"),
		inputs[amount].View(),
		inputStyle.Width(50).Render("Description"),
		inputs[description].View(),
		inputStyle.Width(15).Render("Category"),
		inputs[category].View(),
		inputStyle.Width(7).Render("Type"),
		inputs[transactionType].View(),
	)

	return content
}

func (m model) editorUpdate(msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case onSwitchPageMsg:
		switch msg.data.(type) {
		case database.Transaction:
			m = m.loadTransaction(msg.data.(database.Transaction))
		default:
			m = m.onEditorSwitch()
		}
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyShiftTab, tea.KeyCtrlK:
			m = m.editorPrevInput()
		case tea.KeyTab, tea.KeyCtrlJ:
			m = m.editorNextInput()
		case tea.KeyEnter:
			m.err = m.editorSubmit()
			if m.err == nil {
				m, cmd = m.switchPage(m.lastPage, nil)
			}
		case tea.KeyEsc:
			m, cmd = m.switchPage(m.lastPage, nil)
		}

		for i := range m.state.editor.inputs {
			m.state.editor.inputs[i].Blur()
		}
		m.state.editor.inputs[m.state.editor.focused].Focus()
	}

	cmds := []tea.Cmd{cmd}
	for i := range m.state.editor.inputs {
		m.state.editor.inputs[i], cmd = m.state.editor.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) editorNextInput() model {
	m.state.editor.focused = (m.state.editor.focused + 1) % len(m.state.editor.inputs)
	return m
}

func (m model) editorPrevInput() model {
	m.state.editor.focused--
	if m.state.editor.focused < 0 {
		m.state.editor.focused = len(m.state.editor.inputs) - 1
	}

	return m
}

func (m model) onEditorSwitch() model {
	m.state.editor.inputs[date].SetValue(util.FormatDate(time.Now()))
	m.state.editor.inputs[amount].SetValue("")
	m.state.editor.inputs[description].SetValue("")
	m.state.editor.inputs[category].SetValue(string(database.CategoryTransportation))
	m.state.editor.inputs[transactionType].SetValue(string(database.TransactionTypeExpense))
	m.state.editor.currID = nil
	return m
}

func (m model) loadTransaction(transaction database.Transaction) model {
	m.state.editor.inputs[date].SetValue(util.FormatDate(transaction.Date))
	m.state.editor.inputs[amount].SetValue(util.FormatAmount(transaction.Amount))
	m.state.editor.inputs[description].SetValue(transaction.Description)
	m.state.editor.inputs[category].SetValue(string(transaction.Category))
	m.state.editor.inputs[transactionType].SetValue(string(transaction.TransactionType))
	m.state.editor.currID = &transaction.ID
	return m
}

func (m model) editorSubmit() error {
	inputs := m.state.editor.inputs

	for _, input := range inputs {
		if input.Err != nil {
			return input.Err
		}
	}

	dateValue, err := util.ParseDate(inputs[date].Value())
	if err != nil {
		return fmt.Errorf("Invalid date")
	}

	amountValue, err := strconv.ParseFloat(inputs[amount].Value(), 64)
	if err != nil {
		return fmt.Errorf("Invalid amount")
	}

	ctx := context.Background()
	q := database.New(m.db)

	switch id := m.state.editor.currID; id {
	case nil:
		if _, err = q.CreateTransaction(ctx, database.CreateTransactionParams{
			Description:     inputs[description].Value(),
			Amount:          amountValue,
			Date:            dateValue,
			Category:        database.Category(inputs[category].Value()),
			TransactionType: database.TransactionType(inputs[transactionType].Value()),
		}); err != nil {
			return fmt.Errorf("Error adding transaction")
		}
	default:
		if _, err = q.UpdateTransaction(ctx, database.UpdateTransactionParams{
			ID:              *m.state.editor.currID,
			Description:     inputs[description].Value(),
			Amount:          amountValue,
			Date:            dateValue,
			Category:        database.Category(inputs[category].Value()),
			TransactionType: database.TransactionType(inputs[transactionType].Value()),
		}); err != nil {
			return fmt.Errorf("Error updating transaction")
		}
	}

	return nil
}

func editorInit() editorState {
	var inputs = make([]textinput.Model, 5)

	keyMap := textinput.DefaultKeyMap
	keyMap.AcceptSuggestion = key.NewBinding(key.WithKeys("ctrl+y"))

	inputs[date] = textinput.New()
	inputs[date].Placeholder = "0000-00-00"
	inputs[date].Focus()
	inputs[date].CharLimit = 10
	inputs[date].Width = 10
	inputs[date].Prompt = ""
	inputs[date].Cursor.SetMode(cursor.CursorBlink)
	inputs[date].SetValue(util.FormatDate(time.Now()))
	inputs[date].Validate = func(s string) error {
		if _, err := util.ParseDate(s); err != nil {
			return fmt.Errorf("Invalid date")
		}
		return nil
	}

	inputs[amount] = textinput.New()
	inputs[amount].Placeholder = "00.00"
	inputs[amount].CharLimit = 8
	inputs[amount].Width = 8
	inputs[amount].Prompt = ""
	inputs[amount].Cursor.SetMode(cursor.CursorBlink)
	inputs[amount].Validate = func(s string) error {
		if _, err := strconv.ParseFloat(s, 64); err != nil {
			return fmt.Errorf("Invalid amount")
		}
		return nil
	}

	inputs[description] = textinput.New()
	inputs[description].Placeholder = "description"
	inputs[description].CharLimit = 50
	inputs[description].Width = 50
	inputs[description].Prompt = ""
	inputs[description].Cursor.SetMode(cursor.CursorBlink)

	inputs[category] = textinput.New()
	inputs[category].Placeholder = "transportation"
	inputs[category].CharLimit = 14
	inputs[category].Width = 14
	inputs[category].Prompt = ""
	inputs[category].Cursor.SetMode(cursor.CursorBlink)
	inputs[category].SetValue(string(database.CategoryTransportation))
	inputs[category].ShowSuggestions = true
	categorySuggestions := []string{}
	for _, cat := range database.AllCategoryValues() {
		categorySuggestions = append(categorySuggestions, string(cat))
	}
	inputs[category].SetSuggestions(categorySuggestions)
	inputs[category].KeyMap = keyMap
	inputs[category].Validate = func(s string) error {
		if !slices.Contains(database.AllCategoryValues(), database.Category(s)) {
			return fmt.Errorf("Invalid category")
		}
		return nil
	}

	inputs[transactionType] = textinput.New()
	inputs[transactionType].Placeholder = "expense"
	inputs[transactionType].CharLimit = 7
	inputs[transactionType].Width = 7
	inputs[transactionType].Prompt = ""
	inputs[transactionType].Cursor.SetMode(cursor.CursorBlink)
	inputs[transactionType].SetValue(string(database.TransactionTypeExpense))
	inputs[transactionType].ShowSuggestions = true
	transactionTypeSuggestions := []string{}
	for _, t := range database.AllTransactionTypeValues() {
		transactionTypeSuggestions = append(transactionTypeSuggestions, string(t))
	}
	inputs[transactionType].SetSuggestions(transactionTypeSuggestions)
	inputs[transactionType].KeyMap = keyMap
	inputs[transactionType].Validate = func(s string) error {
		if !slices.Contains(database.AllTransactionTypeValues(), database.TransactionType(s)) {
			return fmt.Errorf("Invalid transaction type")
		}
		return nil
	}

	return editorState{inputs: inputs, focused: date, currID: nil}
}
