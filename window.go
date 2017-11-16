package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	tcell_encoding "github.com/gdamore/tcell/encoding"
)

type Window struct {
	Screen       tcell.Screen
	DefaultStyle tcell.Style
	Width        int
	Height       int
}

func NewWindow() *Window {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	tcell_encoding.Register()

	w := &Window{
		Screen: screen,
	}

	w.DefaultStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	w.Screen.SetStyle(w.DefaultStyle)
	w.Screen.Clear()

	w.OnResize()

	return w
}

func (w *Window) Show() {
	w.Screen.Show()
}

func (w *Window) Clear() {
	w.Screen.Clear()
}

func (w *Window) Destroy() {
	w.Screen.Fini()
}

func (w *Window) OnResize() {
	w.Width, w.Height = w.Screen.Size()
	w.Screen.Sync()
}