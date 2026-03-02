package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
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
	stateFlags
	stateResult
)

type model struct {
	input     textinput.Model
	list      list.Model
	flagInput textinput.Model

	state    state
	text     string
	result   string
	err      error
	info     string
	quitting bool

	selected  processors.Processor
	flags     []processors.Flag
	flagIndex int
	entered   []processors.Flag
}

func initialModel(input string) model {
	ti := textinput.New()
	ti.Placeholder = "Type or paste text here, then press Enter…"
	ti.Focus()
	ti.Width = 80

	delegate := list.NewDefaultDelegate()
	l := list.New(processors.List, delegate, 80, 20)
	l.Title = "Select a transformation"

	fi := textinput.New()
	fi.Width = 80

	m := model{
		input:     ti,
		list:      l,
		flagInput: fi,
		state:     stateInput,
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
		m.flagInput.Width = msg.Width - 4
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			switch m.state {
			case stateResult, stateInput:
				m.quitting = true
				return m, tea.Quit
			case stateList:
				m.state = stateInput
				m.err = nil
				m.info = ""
				return m, nil
			case stateFlags:
				m.state = stateList
				m.err = nil
				m.info = ""
				return m, nil
			}

		case "enter":
			switch m.state {
			case stateInput:
				m.text = m.input.Value()
				if m.text == "" {
					return m, nil
				}
				m.state = stateList
				m.err = nil
				m.info = ""
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
				m.selected = p
				m.flags = p.Flags()
				m.entered = nil
				m.flagIndex = 0
				m.err = nil
				m.info = ""

				if len(m.flags) == 0 {
					m.executeSelected()
					return m, nil
				}

				m.prepareFlagInput()
				m.state = stateFlags
				return m, nil

			case stateFlags:
				current := m.flags[m.flagIndex]
				flag, err := parseFlagInput(current, m.flagInput.Value())
				if err != nil {
					m.err = err
					return m, nil
				}
				m.err = nil
				m.info = ""
				m.entered = append(m.entered, flag)
				if m.flagIndex == len(m.flags)-1 {
					m.executeSelected()
					return m, nil
				}
				m.flagIndex++
				m.prepareFlagInput()
				return m, nil

			case stateResult:
				m.quitting = true
				return m, tea.Quit
			}

		case "esc":
			switch m.state {
			case stateList:
				m.state = stateInput
				m.err = nil
				m.info = ""
				return m, nil
			case stateFlags:
				m.err = nil
				m.info = ""
				if m.flagIndex == 0 {
					m.state = stateList
					return m, nil
				}
				if len(m.entered) > 0 {
					m.entered = m.entered[:len(m.entered)-1]
				}
				m.flagIndex--
				m.prepareFlagInput()
				return m, nil
			case stateResult:
				m.state = stateList
				m.err = nil
				m.info = ""
				return m, nil
			}

		case "y":
			if m.state == stateResult && m.err == nil {
				if err := clipboard.WriteAll(m.result); err != nil {
					m.info = fmt.Sprintf("Clipboard error: %v", err)
				} else {
					m.info = "Copied result to clipboard"
				}
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
	case stateFlags:
		m.flagInput, cmd = m.flagInput.Update(msg)
	}
	return m, cmd
}

func (m *model) prepareFlagInput() {
	if m.flagIndex >= len(m.flags) {
		return
	}
	flag := m.flags[m.flagIndex]
	ti := textinput.New()
	ti.Focus()
	ti.Width = m.input.Width
	if flag.Sensitive {
		ti.EchoMode = textinput.EchoPassword
		ti.EchoCharacter = '*'
	}
	defaultText := flagDefaultAsString(flag)
	required := "optional"
	if flag.Required {
		required = "required"
	}
	if defaultText != "" {
		ti.Placeholder = fmt.Sprintf("%s (%s, default=%s)", flag.HelpLabel(), required, defaultText)
	} else {
		ti.Placeholder = fmt.Sprintf("%s (%s)", flag.HelpLabel(), required)
	}
	m.flagInput = ti
}

func (m *model) executeSelected() {
	result, err := m.selected.Transform([]byte(m.text), m.entered...)
	m.result = result
	m.err = err
	m.state = stateResult
}

func parseFlagInput(flag processors.Flag, raw string) (processors.Flag, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		if flag.Required {
			switch v := flag.Value.(type) {
			case string:
				if v == "" {
					return processors.Flag{}, fmt.Errorf("--%s is required", flag.Name)
				}
			default:
				if v == nil {
					return processors.Flag{}, fmt.Errorf("--%s is required", flag.Name)
				}
			}
		}
		return processors.Flag{Short: flag.Short, Value: flag.Value}, nil
	}

	out := processors.Flag{Short: flag.Short}
	switch flag.Type {
	case processors.FlagString:
		out.Value = raw
	case processors.FlagInt:
		v, err := strconv.Atoi(trimmed)
		if err != nil {
			return processors.Flag{}, fmt.Errorf("--%s expects an integer", flag.Name)
		}
		out.Value = v
	case processors.FlagUint:
		v, err := strconv.ParseUint(trimmed, 10, 64)
		if err != nil {
			return processors.Flag{}, fmt.Errorf("--%s expects an unsigned integer", flag.Name)
		}
		out.Value = uint(v)
	case processors.FlagBool:
		v, err := parseBool(trimmed)
		if err != nil {
			return processors.Flag{}, fmt.Errorf("--%s expects a boolean (true/false)", flag.Name)
		}
		out.Value = v
	default:
		out.Value = raw
	}

	if flag.Required {
		if s, ok := out.Value.(string); ok && strings.TrimSpace(s) == "" {
			return processors.Flag{}, fmt.Errorf("--%s is required", flag.Name)
		}
	}

	return out, nil
}

func parseBool(s string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "1", "true", "t", "yes", "y", "on":
		return true, nil
	case "0", "false", "f", "no", "n", "off":
		return false, nil
	default:
		return false, fmt.Errorf("invalid bool")
	}
}

func flagDefaultAsString(flag processors.Flag) string {
	if flag.Value == nil {
		return ""
	}
	switch v := flag.Value.(type) {
	case string:
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", v)
	}
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

	case stateFlags:
		flag := m.flags[m.flagIndex]
		b.WriteString(dimStyle.Render("Input: ") + truncate(m.text, 60) + "\n")
		b.WriteString(dimStyle.Render(fmt.Sprintf("Processor: %s", m.selected.Name())) + "\n")
		b.WriteString(dimStyle.Render(fmt.Sprintf("Flag %d/%d", m.flagIndex+1, len(m.flags))) + "\n\n")
		b.WriteString(fmt.Sprintf("%s\n", flag.Desc))
		b.WriteString(m.flagInput.View())
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("Press Enter to continue • Esc to go back"))
		if m.err != nil {
			b.WriteString("\n" + errorStyle.Render("Error: ") + m.err.Error())
		}

	case stateResult:
		b.WriteString(dimStyle.Render("Input: ") + truncate(m.text, 60) + "\n\n")
		if m.err != nil {
			b.WriteString(errorStyle.Render("Error: ") + m.err.Error() + "\n")
		} else {
			b.WriteString(resultStyle.Render("Result:") + "\n" + m.result + "\n")
		}
		if m.info != "" {
			b.WriteString("\n" + dimStyle.Render(m.info))
		}
		b.WriteString("\n")
		b.WriteString(dimStyle.Render("Press y to copy • Enter or q to quit • Esc to go back"))
	}

	return b.String()
}

func truncate(s string, max int) string {
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	if max <= 3 {
		return string(r[:max])
	}
	return string(r[:max-3]) + "..."
}

type UI struct {
	input string
}

func New(input string) *UI {
	return &UI{input: input}
}

func (u *UI) Render() error {
	p := tea.NewProgram(initialModel(u.input), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running TUI: %w", err)
	}
	return nil
}
