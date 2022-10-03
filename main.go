package main

import (
	"fmt"
	"os"

	"github.com/Caiosev/weather-BubbleTea/metaweather"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	t := textinput.NewModel()
	t.Focus()
	initialModel := model{
		textInput: t,
		typing:    true,
	}
	err := tea.NewProgram(initialModel, tea.WithAltScreen()).Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{}
}

type model struct {
	textInput textinput.Model

	typing   bool
	loading  bool
	err      error
	location metaweather.Location
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	if m.typing {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	if m.typing {
		return fmt.Sprintf("Enter location:\n%s", m.textInput.View())
	}
	return "Press Ctrl+C to exit"
}
