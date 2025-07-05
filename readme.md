# 🍹 Drink Selector TUI

This is a simple yet elegant terminal-based UI (TUI) application built with [Bubble Tea](https://github.com/charmbracelet/bubbletea), allowing users to select their favorite drink from a list using their keyboard.

✨ In the future, it will be enhanced using [Lip Gloss](https://github.com/charmbracelet/lipgloss) for stylish, colorful UI components.

---

## 🖼️ Preview

```

Select your favorite drink.

Available options:
Té de burbujas
-> Café
Cerveza
Zumo de naranja
Agua

You selected: Café

Use the up/down keys to navigate and select with Enter. Press 'q' to quit

```

After pressing Enter:

```

You selected: Café! Enjoy your drink!

Press ESC to close.

````

---

## 🛠️ Features

- Keyboard navigation (`↑/↓` or `k/j`)
- Modal-style confirmation after selection
- Terminal resize awareness
- Clean and extendable structure
- Ready to be styled with `lipgloss`

---

## 🎮 Controls

| Key        | Action                      |
|------------|-----------------------------|
| ↑ / k      | Move selection up           |
| ↓ / j      | Move selection down         |
| Enter      | Show modal confirmation     |
| ESC / q    | Close modal / quit          |
| Ctrl + C   | Force quit                  |

---

## 🚀 Getting Started

### Prerequisites

- Go 1.24 or later

### Installation

```bash
git clone https://github.com/b7craft/drink-selector-tui.git
cd drink-selector-tui
go mod tidy
go run .
````

---

## 📦 Dependencies

* [Bubble Tea](https://github.com/charmbracelet/bubbletea) — framework for TUI
* (Soon) [Lip Gloss](https://github.com/charmbracelet/lipgloss) — for styling and theming

---

## 📁 Project Structure

```
.
├── main.go       # Main application logic
└── README.md     # Project description
```

---

## 💡 Future Enhancements

* 🌈 Style the list, modal, and footer using `lipgloss`
* 🐹 Add animated transitions using [bubbles](https://github.com/charmbracelet/bubbles)
* 🌐 Localize drink names and messages
* 🎨 Theme toggle (dark/light mode)

---

## 📜 License

MIT License

---

## 🙌 Acknowledgments

Special thanks to [Charmbracelet](https://github.com/charmbracelet) for creating the delightful TUI ecosystem.

---

*Made with Go and good taste. Cheers! 🥂*