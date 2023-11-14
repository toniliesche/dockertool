package menu

import (
	"bufio"
	"fmt"
	"github.com/toniliesche/dockertool/modules/console"
	"os"
	"strconv"
	"strings"
)

type Menu struct {
}

func (m *Menu) RunMenu(menuEntries []*MenuEntry, specialEntries []*MenuEntry, opts ...string) *MenuEntry {
	reader := bufio.NewReader(os.Stdin)

	var info string
	if len(opts) > 0 {
		info = fmt.Sprintf("%s :", opts[0])
	} else {
		info = "Select an Action :"
	}

	fmt.Println(info)
	fmt.Println(console.StringPad("=", len(info)))

	for i, entry := range menuEntries {
		fmt.Printf("%6d) %s\n", i+1, entry.Label)
		if entry.Divider {
			fmt.Println()
		}
	}

	if len(specialEntries) > 0 {
		fmt.Println()

		for _, entry := range specialEntries {
			fmt.Printf("%6s) %s\n", entry.Shortcut, entry.Label)
		}
	}

	fmt.Println()
	fmt.Println("     q) Exit")
	fmt.Println()

	var returnEntry *MenuEntry
out:
	for {
		console.PrintHeadline("Enter your choice :")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if v, err := strconv.Atoi(text); err == nil {
			if v > 0 && v <= len(menuEntries) {
				returnEntry = menuEntries[v-1]
				break
			}
		} else {
			if "q" == text || "Q" == text {
				break
			}

			for _, entry := range specialEntries {
				if entry.Shortcut == text {
					returnEntry = entry
					break out
				}
			}
		}

		console.PrintError("Invalid option")
	}

	if nil != returnEntry && nil != returnEntry.Args && nil != returnEntry.Page {
		returnEntry.Page.SetArguments(returnEntry.Args)
	}

	return returnEntry
}
