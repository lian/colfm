package main

import (
	"fmt"
	"os"
	"strings"
)

type FileView struct {
	Columns     []*FileColumn
	HiddenFiles bool
}

func NewFileView() *FileView {
	v := &FileView{
		Columns: []*FileColumn{},
	}
	return v
}

func (v *FileView) CdFull(dir string) {
	v.Columns = []*FileColumn{}

	for _, name := range strings.Split(dir, "/") {
		if name == "" {
			v.Cd("/")
		} else {
			v.Active().Select(name)
			v.Enter()
		}
	}
}

func (v *FileView) Cd(dir string) {
	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		fmt.Println("not a directory", dir)
		return
	}

	if len(v.Columns) != 0 {
		v.Columns[len(v.Columns)-1].Active = false
	}

	column := NewFileColumn(dir, v.HiddenFiles)
	column.Active = true
	v.Columns = append(v.Columns, column)
}

func (v *FileView) Active() *FileColumn {
	return v.Columns[len(v.Columns)-1]
}

func (v *FileView) Next() {
	v.Active().Cursor(1)
}

func (v *FileView) Prev() {
	v.Active().Cursor(-1)
}

func (v *FileView) Enter() {
	item := v.Active().Selection()
	if item == nil {
		return
	}
	if item.IsDir() {
		v.Cd(item.FullPath())
	} else {
		window.Screen.Clear()
		window.Screen.Sync()
		LessFile(item.FullPath())
		window.Screen.Sync()
	}
}

func (v *FileView) Leave() {
	if len(v.Columns) == 1 {
		return
	}
	v.Columns = v.Columns[:len(v.Columns)-1]
	v.Columns[len(v.Columns)-1].Active = true
}

func (v *FileView) Skipcols() int {
	var cols, total int
	for i := len(v.Columns) - 1; i >= 0; i-- {
		total += v.Columns[i].Width() + 1
		if total > window.Width {
			break
		}
		cols += 1
	}
	return len(v.Columns) - cols
}

func (v *FileView) Draw() {
	skip := v.Skipcols()

	xpos := 1
	ypos := 1
	max_y := window.Height - 2

	for n, col := range v.Columns {
		if n < skip {
			continue
		}
		col.Draw(xpos, ypos, max_y)
		xpos += col.Width() + 1
	}
}
