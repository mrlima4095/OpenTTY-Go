package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type AppState struct {
	version      string
	Username     string
	Path         string
	History      []string
	Build        string
	nanoContent  string
	StdinEntry   *widget.Entry
	StdoutView   *widget.MultiLineEntry
	SelectedCmd  string
	MainTitle    string
}

func LoadRMS(key string) string {
	filename := filepath.Join(".", key+".rms")
	data, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

func WriteRMS(key, value string) {
	filename := filepath.Join(".", key+".rms")
	_ = os.WriteFile(filename, []byte(value), 0644)
}

func (state *AppState) AppendOutput(text string) {
	state.StdoutView.SetText(state.StdoutView.Text + "\n" + text)
}

func (state *AppState) ProcessCommand(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return ""
	}
	switch cmd {
	case "help":
		return "Available commands: help, history, clear, nano, exit"
	case "history":
		OpenHistoryViewer(state)
		return "[history viewer opened]"
	case "clear":
		state.StdoutView.SetText("")
		return ""
	case "nano":
		OpenNanoEditor(state)
		return "[nano editor opened]"
	case "exit":
		os.Exit(0)
	}
	return fmt.Sprintf("%s: not found", cmd)
}

func OpenNanoEditor(state *AppState) {
	nanoWin := fyne.CurrentApp().NewWindow("Nano")

	editor := widget.NewMultiLineEntry()
	editor.SetText(LoadRMS("nano"))

	saveBtn := widget.NewButton("Save", func() {
		text := editor.Text
		state.nanoContent = text
		WriteRMS("nano", text)
		nanoWin.Close()
	})

	cancelBtn := widget.NewButton("Cancel", func() {
		nanoWin.Close()
	})

	buttons := container.NewHBox(saveBtn, cancelBtn)
	layout := container.NewBorder(nil, buttons, nil, nil, editor)

	nanoWin.SetContent(layout)
	nanoWin.Resize(fyne.NewSize(500, 400))
	nanoWin.Show()
}

func OpenHistoryViewer(state *AppState) {
	historyWin := fyne.CurrentApp().NewWindow(state.MainTitle)

	list := widget.NewList(
		func() int {
			return len(state.History)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(state.History[i])
		},
	)

	list.OnSelected = func(id int) {
		state.SelectedCmd = state.History[id]
	}

	runBtn := widget.NewButton("Run", func() {
		if state.SelectedCmd != "" {
			output := state.ProcessCommand(state.SelectedCmd)
			if output != "" {
				state.AppendOutput(output)
			}
		}
	})

	editBtn := widget.NewButton("Edit", func() {
		if state.SelectedCmd != "" {
			state.StdinEntry.SetText(state.SelectedCmd)
		}
	})

	closeBtn := widget.NewButton("Close", func() {
		historyWin.Close()
	})

	buttons := container.NewHBox(runBtn, editBtn, closeBtn)
	layout := container.NewBorder(nil, buttons, nil, nil, list)

	historyWin.SetContent(layout)
	historyWin.Resize(fyne.NewSize(500, 400))
	historyWin.Show()
}

func main() {
	a := app.New()
	mainTitle := "OpenTTY Terminal"
	w := a.NewWindow(mainTitle)

	username := LoadRMS("OpenRMS")
	path := "/home/"

	stdin := widget.NewEntry()
	stdout := widget.NewMultiLineEntry()
	stdout.Wrapping = fyne.TextWrapWord
	stdout.SetText("Welcome to OpenTTY 0.6.2\nCopyright (C) 2025 - Mr. Lima\n")
	stdout.SetMinRowsVisible(25)

	state := &AppState{
		Username:     username,
		Path:         path,
		History:      []string{},
		Build:        "2025-1.15-02x06",
		nanoContent:  LoadRMS("nano"),
		StdinEntry:   stdin,
		StdoutView:   stdout,
		SelectedCmd:  "",
		MainTitle:    mainTitle,
	}

	stdin.SetPlaceHolder(fmt.Sprintf("%s %s $ ", state.Username, state.Path))

	stdin.OnSubmitted = func(command string) {
		if strings.TrimSpace(command) != "" {
			state.History = append(state.History, command)
		}
		output := state.ProcessCommand(command)
		if output != "" {
			state.AppendOutput(output)
		}
		stdin.SetText("")
	}

	executeBtn := widget.NewButton("Send", func() {
		stdin.OnSubmitted(stdin.Text)
	})

	helpBtn := widget.NewButton("Help", func() {
		output := state.ProcessCommand("help")
		state.AppendOutput(output)
	})

	nanoBtn := widget.NewButton("Nano", func() {
		output := state.ProcessCommand("nano")
		state.AppendOutput(output)
	})

	clearBtn := widget.NewButton("Clear", func() {
		state.StdoutView.SetText("")
	})

	historyBtn := widget.NewButton("History", func() {
		output := state.ProcessCommand("history")
		state.AppendOutput(output)
	})

	btns := container.NewHBox(executeBtn, helpBtn, nanoBtn, clearBtn, historyBtn)
	term := container.NewBorder(nil, container.NewVBox(stdin, btns), nil, nil, stdout)

	w.SetContent(container.New(layout.NewMaxLayout(), term))
	w.Resize(fyne.NewSize(800, 600))

	w.SetCloseIntercept(func() {
		WriteRMS("nano", state.nanoContent)
		a.Quit()
	})

	w.ShowAndRun()
}
