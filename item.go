package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell"
)

type FileItem struct {
	Parent string
	Info   os.FileInfo
}

func (f *FileItem) Name() string {
	return f.Info.Name()
}

func (f *FileItem) Path() string {
	return f.Parent + "/" + f.Info.Name()
}

func (f *FileItem) FullPath() string {
	path := f.Path()

	if f.IsSymlink() {
		if resolve, err := filepath.EvalSymlinks(path); err == nil {
			return resolve
		}
	}

	return path
}

func (f *FileItem) IsDir() bool {
	if f.IsSymlink() {
		if path, err := filepath.EvalSymlinks(f.Path()); err == nil {
			if info, err := os.Stat(path); err == nil {
				return info.IsDir()
			}
		}
	}

	return f.Info.IsDir()
}

func (f *FileItem) IsFile() bool {
	return !f.Info.IsDir()
}

func (f *FileItem) IsSymlink() bool {
	return (f.Info.Mode() & os.ModeSymlink) != 0
}

func (f *FileItem) Style(sel *FileItem) tcell.Style {
	style := window.DefaultStyle

	if f.IsDir() {
		style = style.Foreground(tcell.ColorNavy)
		if f == sel {
			style = style.Background(tcell.ColorWhite)
		}
	} else {
		if f == sel {
			style = style.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
		}
	}

	return style
}

func (f *FileItem) Format(width int, detail bool) string {
	var sigil string

	switch mode := f.Info.Mode(); {
	case mode&os.ModeSymlink != 0:
		sigil = "@"
	case mode.IsDir():
		sigil = "/"
	case mode&os.ModeSocket != 0:
		sigil = "="
	case mode&os.ModeNamedPipe != 0:
		sigil = "|"
	case mode&0111 != 0: // executable
		sigil = "*"
	}

	name := f.Info.Name()

	if detail && !f.IsDir() {
		return ljust(trunc(name+sigil, width-5), width-5) +
			fmt.Sprintf("%5s", human(f.Info.Size()))
	}

	return ljust(trunc(name+sigil, width), width)
}
