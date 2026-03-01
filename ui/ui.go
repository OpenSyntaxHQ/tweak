package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

var (
	titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7C3AED")).Bold(true)
	resultStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#10B981")).Bold(true)
	errorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#EF4444")).Bold(true)
	dimStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280"))
)

type state int

const (
	stateInput state = iota
	stateList
	stateResult
)

type model struct {
	input    textinput.Model
	list     list.Model
	state    state
	text     string
	result   string
	err      error
	quitting bool
}

func initialModel(input string) model {
	ti := textinput.New()
	ti.Placeholder = "Type or paste text here, then press Enter…"
	ti.Focus()
	ti.Width = 80

	delegate := list.NewDefaultDelegate()
	l := list.New(processors.List, delegate, 80, 20)
	l.Title = "Select a transformation"

	m := model{
		input: ti,
		list:  l,
		state: stateInput,
	}
	if input != "" {
		m.text = input
		m.state = stateList
	}
	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-4)
		m.input.Width = msg.Width - 4
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.state == stateResult {
				m.quitting = true
				return m, tea.Quit
			}
			if m.state == stateList {
				m.state = stateInput
				return m, nil
			}
			m.quitting = true
			return m, tea.Quit

		case "enter":
			switch m.state {
			case stateInput:
				m.text = m.input.Value()
				if m.text == "" {
					return m, nil
				}
				m.state = stateList
				return m, nil

			case stateList:
				selected := m.list.SelectedItem()
				if selected == nil {
					return m, nil
				}
				p, ok := selected.(processors.Processor)
				if !ok {
					m.err = fmt.Errorf("invalid processor selection")
					m.state = stateResult
					return m, nil
				}
				result, err := p.Transform([]byte(m.text))
				m.result = result
				m.err = err
				m.state = stateResult
				return m, nil

			case stateResult:
				m.quitting = true
				return m, tea.Quit
			}

		case "esc":
			switch m.state {
			case stateList:
				m.state = stateInput
				return m, nil
			case stateResult:
				m.state = stateList
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	switch m.state {
	case stateInput:
		m.input, cmd = m.input.Update(msg)
	case stateList:
		m.list, cmd = m.list.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render("✦ tweak") + dimStyle.Render(" — interactive text transformer") + "\n\n")

	switch m.state {
	case stateInput:
		b.WriteString("Enter your text:\n\n")
		b.WriteString(m.input.View())
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("Press Enter to continue • Ctrl+C to quit"))

	case stateList:
		b.WriteString(dimStyle.Render("Input: ") + truncate(m.text, 60) + "\n\n")
		b.WriteString(m.list.View())

	case stateResult:
		b.WriteString(dimStyle.Render("Input: ") + truncate(m.text, 60) + "\n\n")
		if m.err != nil {
			b.WriteString(errorStyle.Render("Error: ") + m.err.Error() + "\n")
		} else {
			b.WriteString(resultStyle.Render("Result:") + "\n" + m.result + "\n")
		}
		b.WriteString("\n")
		b.WriteString(dimStyle.Render("Press Enter or q to quit • Esc to go back"))
	}

	return b.String()
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

type UI struct {
	input string
}

func New(input string) *UI {
	return &UI{input: input}
}

func (u *UI) Render() {
	p := tea.NewProgram(initialModel(u.input), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
		os.Exit(1)
	}
}
