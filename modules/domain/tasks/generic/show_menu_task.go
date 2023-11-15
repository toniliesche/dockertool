package generic

import (
	"bufio"
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/domain/tasks/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"os"
	"strconv"
	"strings"
)

type ShowMenuTask struct {
	base.Task
	menuEntries    models.EntryList
	specialEntries models.EntryList
	opts           []string
}

func (t *ShowMenuTask) Run() error {
	reader := bufio.NewReader(os.Stdin)

	var info string
	if len(t.opts) > 0 {
		info = fmt.Sprintf("%s :", t.opts[0])
	} else {
		info = "Select an Action :"
	}

	fmt.Println(info)
	fmt.Println(console.StringPad("=", len(info)))

	for i, entry := range t.menuEntries {
		fmt.Printf("%6d) %s\n", i+1, entry.Label)
		if entry.Divider {
			fmt.Println()
		}
	}

	if len(t.specialEntries) > 0 {
		fmt.Println()

		for _, entry := range t.specialEntries {
			fmt.Printf("%6s) %s\n", entry.Shortcut, entry.Label)
		}
	}

	fmt.Println()
	fmt.Println("     q) Exit")
	fmt.Println()

	var returnEntry *models.Entry
out:
	for {
		console.PrintHeadline("Enter your choice :")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if v, err := strconv.Atoi(text); nil == err {
			if v > 0 && v <= len(t.menuEntries) {
				returnEntry = t.menuEntries[v-1]
				break
			}
		} else {
			if "q" == text || "Q" == text {
				break
			}

			for _, entry := range t.specialEntries {
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

	t.Result = returnEntry

	return nil
}

func CreateShowMenuTask(entries models.EntryList, specialEntries models.EntryList) (*ShowMenuTask, error) {
	return &ShowMenuTask{menuEntries: entries, specialEntries: specialEntries}, nil
}
