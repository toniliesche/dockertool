package menu

import (
	"sort"
	"strings"
)

type MenuEntry struct {
	Label    string
	Page     PageInterface
	Args     []string
	IntArgs  []int
	Shortcut string
	Divider  bool
}

type MenuEntryList []*MenuEntry

func (l *MenuEntryList) Sort() {
	list := *l
	sort.Slice(*l, func(i, j int) bool {
		return strings.ToLower(list[i].Label) < strings.ToLower(list[j].Label)
	})
}
