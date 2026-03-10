package main

import (
	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"strings"
)

type Model struct {
	ollamaModels []string
	cursor       int
	status       string
	err          error
	viewport     viewport.Model
	textarea     textarea.Model
	quitting     bool
	thinking     bool
	messages     []string
	senderStyle  lipgloss.Style
}

func initialModel() Model {

	ta := textarea.New()
	ta.Placeholder = "Ask away!"
	ta.Focus()
	ta.SetVirtualCursor(false)
	ta.SetWidth(30)
	ta.CharLimit = 1000
	ta.SetHeight(3)
	ta.ShowLineNumbers = false

	s := ta.Styles()
	s.Focused.CursorLine = lipgloss.NewStyle()
	ta.SetStyles(s)

	vp := viewport.New(viewport.WithHeight(30), viewport.WithHeight(5))
	vp.SetContent(`Let's get started!`)
	vp.KeyMap.Left.SetEnabled(false)
	vp.KeyMap.Right.SetEnabled(false)

	return Model{
		status:       "",
		quitting:     false,
		ollamaModels: []string{},
		messages:     []string{},
		// cursor:       0,
		textarea:    ta,
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) View() tea.View {
	viewportView := m.viewport.View()

	v := tea.NewView(viewportView + "/n" + m.textarea.View())
	c := m.textarea.Cursor()
	if c != nil {
		c.Y += lipgloss.Height(m.headerView())
	}
	v.Cursor = c
	// v.AltScreen = true
	return v
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.SetWidth(msg.Width)
		m.textarea.SetWidth(msg.Width)
		m.viewport.SetHeight(msg.Height - m.textarea.Height())

		if len(m.messages) > 0 {
			m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
		}
		m.viewport.GotoBottom()
	case tea.KeyPressMsg:

		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
			m.textarea.Reset()
			m.viewport.GotoBottom()
			return m, nil

		default:
			var cmd tea.Cmd
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		}

	case cursor.BlinkMsg:
		var cmd tea.Cmd
		m.textarea, cmd = m.textarea.Update(msg)
		return m, cmd

	}

	// m.textarea, cmd = m.textarea.Update(msg)
	return m, nil
}

func (m Model) headerView() string { return "ChattaTUI\n" }
func (m Model) footerView() string { return "\n(esc or Ctrl-c to quit)" }
