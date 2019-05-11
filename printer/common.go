package printer

import "os"

func printString(s string) {
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\033[?25l")
}

func printSpace(count int) {
	for i := 0; i < count; i++ {
		printString(" ")
	}
}

func printReturn(count int) {
	for i := 0; i < count; i++ {
		printString("\n")
	}
}

func clearScreen() {
	printString("\033c")
}

func charLen(s string) int {
	return len([]rune(s))
}
