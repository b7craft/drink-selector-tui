package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	selected int
	choices  []string

	// モーダル
	showModal    bool
	modalMessage string

	// ターミナルサイズ
	width  int
	height int
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// モーダル表示中は、モーダルを閉じる処理のみを行う
		if m.showModal {
			switch msg.String() {
			case "esc", "q":
				m.closeModal()
				return m, nil
			case "ctrl+c":
				m.closeModal()
				return m, tea.Quit
			default:
				// do nothing
			}
		}

		switch msg.String() {
		case "up", "k":
			m.selected--
		case "down", "j":
			m.selected++
		case "enter":
			m.showModal = true
			m.modalMessage = fmt.Sprintf("You selected: %s! Enjoy your drink!", m.choices[m.selected])
		case "q", "ctrl+c":
			return m, tea.Quit
		default:
			// do nothing
		}
	case tea.WindowSizeMsg:
		// ターミナルサイズが変更されたらモデルに保存
		m.width = msg.Width
		m.height = msg.Height

	}
	// 選択範囲を境界内に収める
	if m.selected < 0 {
		m.selected = 0
	}
	if m.selected >= len(m.choices) {
		m.selected = len(m.choices) - 1
	}
	return m, nil
}

func (m *model) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m *model) View() string {

	// ヘッダー部分
	header := "Select your favorite drink.\n\nAvailable options:\n"

	// メニューリスト部分
	var menuBuilder strings.Builder
	for i, choice := range m.choices {
		if i == m.selected {
			menuBuilder.WriteString(fmt.Sprintf(" -> %s\n", choice))
		} else {
			menuBuilder.WriteString(fmt.Sprintf("    %s\n", choice))
		}
	}
	menu := menuBuilder.String()

	// 選択された項目表示部分
	selectedDrink := m.choices[m.selected]
	answer := fmt.Sprintf("You selected: %s\n", selectedDrink)

	footer := "\nUse the up/down keys to navigate and select with Enter. Press 'q' to quit"

	mainView := header + menu + answer + footer

	// モーダル表示
	if m.showModal {
		var modalContent strings.Builder
		modalContent.WriteString("\n" + m.modalMessage + "\n\nPress ESC to close.\n")
		return modalContent.String()
	}

	return mainView
}

// モーダルを閉じる専用メソッド
func (m *model) closeModal() {
	m.showModal = false
	m.modalMessage = ""
}

func main() {
	initialModel := &model{
		selected: 0,
		choices: []string{
			"Té de burbujas", "Café", "Cerveza", "Zumo de naranja", "Agua",
		},
		showModal: false,
	}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}
