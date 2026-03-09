package main

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	ollamaModels []string
	cursor       int
	status       string
	err          error
	textInput    textinput.Model
	quitting     bool
}

func initialModel() Model {

	ti := textinput.New()
	ti.Focus()
	ti.SetWidth(80)
	ti.CharLimit = 1000

	return Model{
		status:       "",
		quitting:     false,
		ollamaModels: []string{},
		cursor:       0,
		textInput:    ti,
	}
}

func (m Model) Init() tea.Cmd {

	return textinput.Blink
}

func (m Model) View() tea.View {
	var c *tea.Cursor
	if !m.textInput.VirtualCursor() {
		c = m.textInput.Cursor()
		c.Y += lipgloss.Height(m.headerView())
	}

	str := lipgloss.JoinVertical(lipgloss.Top, m.headerView(), m.textInput.View(), m.footerView())
	if m.quitting {
		str += "\n"
	}

	v := tea.NewView(str)
	v.Cursor = c
	return v
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:

		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) headerView() string { return "ChattaTUI\n" }
func (m Model) footerView() string { return "\n(esc or Ctrl-c to quit)" }
