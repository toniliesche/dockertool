package console

import (
	"bufio"
	"fmt"
	"os"
)

func WaitForReturn() {
	fmt.Println("Press return to continue")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
