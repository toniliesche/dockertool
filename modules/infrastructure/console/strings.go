package console

import "fmt"

func StringPad(char string, length int) string {
	ret := ""

	for i := 0; i < length; i++ {
		ret = ret + char
	}

	return ret
}

func BoolToYesNo(input bool) string {
	if input {
		return "yes"
	}

	return "no"
}

func BoolToYesNoColored(input bool) string {
	if input {
		return fmt.Sprintf("%syes%s", OKColor, RegularColor)
	}

	return fmt.Sprintf("%sno%s", ErrorColor, RegularColor)
}

func YesNoToBool(input string) bool {
	return "yes" == input
}
