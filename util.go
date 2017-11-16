package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"
)

func getSaveDirFile() string {
	u, _ := user.Current()
	var path string
	if _, err := os.Stat(fmt.Sprintf("/tmp/%s", u.Username)); err == nil {
		path = fmt.Sprintf("/tmp/%s/.colfmdir", u.Username)
	} else {
		path = fmt.Sprintf("%s/.colfmdir", u.HomeDir)
	}
	return path
}

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func LessFile(name string) {
	cmd := exec.Command("less", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func CheckTerminal() {
	if !terminal.IsTerminal(int(os.Stdout.Fd())) {
		panic("no tty")
	}
}

func trunc(str string, width int) string {
	if len(str) > width {
		a := str[0:(2 * width / 3)]
		b := str[(len(str) - (width / 3)):]
		return fmt.Sprintf("%s…%s", a, b)
	} else {
		return str
	}
}

func ljust(str string, width int) string {
	tmpl := fmt.Sprintf("%%-%ds", width)
	return fmt.Sprintf(tmpl, str)
}

func human(size int64) string {
	units := []string{"B", "K", "M", "G", "T", "P", "E", "Z", "Y"}

	for {
		if size < 1024 {
			break
		}
		units = units[1:]
		size /= 1024.0
	}

	return fmt.Sprintf("%d%s", size, units[0])
}
