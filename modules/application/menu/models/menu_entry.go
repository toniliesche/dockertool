package models

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"sort"
	"strings"
)

type Entry struct {
	Label    string
	Page     interfaces.PageInterface
	Args     []string
	IntArgs  []int
	Shortcut string
	Divider  bool
}

type EntryList []*Entry

func (l *EntryList) Sort() {
	list := *l
	sort.Slice(*l, func(i, j int) bool {
		return strings.ToLower(list[i].Label) < strings.ToLower(list[j].Label)
	})
}
