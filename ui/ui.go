package ui

import (
	"github.com/rivo/tview"
)

// type UI interface {
//
// }
type UI struct {
	version  string
	Itmes    []string // skal endres
	TODOList *tview.List
	TODODATA *tview.TextView
}

func NewUI(version string) *UI {
	return &UI{version: version}
}
