package main

// https://github.com/alecthomas/chroma A general purpose syntax highlighter in pure Go

import (
	"io/ioutil"
	"os"

	"github.com/gdamore/tcell"
)

var window *Window

func main() {
	CheckTerminal()

	window = NewWindow()

	view := NewFileView()

	pwd, _ := os.Getwd()
	view.CdFull(pwd)

LOOP:
	for {
		sel := view.Active().Selection()
		if sel != nil {
			emitStr(window.Screen, 1, 0, window.DefaultStyle, sel.FullPath())
		}
		view.Draw()
		window.Show()

		ev := window.Screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			window.OnResize()

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
				window.Destroy()
				break LOOP
			} else if ev.Key() == tcell.KeyDown || ev.Rune() == 'j' {
				view.Active().Cursor(1)
			} else if ev.Key() == tcell.KeyUp || ev.Rune() == 'k' {
				view.Active().Cursor(-1)
			} else if ev.Key() == tcell.KeyRight || ev.Rune() == 'l' {
				view.Enter()
			} else if ev.Key() == tcell.KeyLeft || ev.Rune() == 'h' {
				view.Leave()
			} else if ev.Rune() == '.' {
				view.HiddenFiles = !view.HiddenFiles
				view.CdFull(view.Active().Root)
			}
		}

		window.Clear()
	}

	currentPath := view.Active().Root
	saveDirFile := getSaveDirFile()
	ioutil.WriteFile(saveDirFile, []byte(currentPath), 0644)
}
