package console

import "fmt"

func PrintHeadline(msg string) {
	fmt.Printf("%s%s%s\n\n", HeadlineColor, msg, RegularColor)
}

func PrintOK(msg string) {
	fmt.Printf("%s%s%s\n\n", OKColor, msg, RegularColor)
}

func PrintError(msg string) {
	fmt.Println()
	fmt.Printf("%s%s%s\n", ErrorColor, msg, RegularColor)
	fmt.Println()
}
