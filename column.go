package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

const min_col_width = 8
const max_col_width = 20
const max_active_col_width = 28

type FileColumn struct {
	Root   string
	Items  []*FileItem
	cursor int
	Active bool
}

func NewFileColumn(root string) *FileColumn {
	v := &FileColumn{
		Root:  root,
		Items: []*FileItem{},
	}

	infos, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err)
	}

	files := []*FileItem{}
	dirs := []*FileItem{}

	if root == "/" {
		root = ""
	}

	for _, info := range infos {
		item := &FileItem{
			Parent: root,
			Info:   info,
		}
		if item.IsDir() {
			dirs = append(dirs, item)
		} else {
			files = append(files, item)
		}
	}

	for _, item := range dirs {
		v.Items = append(v.Items, item)
	}

	for _, item := range files {
		v.Items = append(v.Items, item)
	}

	return v
}

func (col *FileColumn) Width() int {
	var max_name_length int
	for _, item := range col.Items {
		nameLen := len(item.Name()) + 1
		if max_name_length < nameLen {
			max_name_length = nameLen
		}
	}

	a := math.Max(min_col_width, float64(max_name_length))

	if col.Active {
		return int(math.Min(a, max_active_col_width)) + 5
	}

	return int(math.Min(a, max_col_width))
}

func (col *FileColumn) Cursor(offset int) {
	col.cursor = int(math.Min(math.Max(float64(col.cursor+offset), 0), float64(len(col.Items)-1)))
}

func (col *FileColumn) Selection() *FileItem {
	if len(col.Items) == 0 {
		return nil
	}
	return col.Items[col.cursor]
}

func (col *FileColumn) Select(name string) bool {
	for n, item := range col.Items {
		if item.Name() == name {
			col.cursor = n
			return true
		}
	}
	return false
}

func (col *FileColumn) Draw(xpos, ypos, max_y int) {
	sel := col.Selection()
	width := col.Width()

	skiplines := int(math.Max(0, float64(col.cursor-max_y+1)))

	for j, item := range col.Items {
		if j < skiplines {
			continue
		}
		if j-skiplines > max_y {
			break
		}

		text := item.Format(width, col.Active)

		style := item.Style(sel)
		emitStr(window.Screen, xpos, ypos, style, text)
		ypos += 1
	}
}
