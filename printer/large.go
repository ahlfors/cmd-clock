package printer

import (
	"time"
)

const (
	largeWidth  = 30
	largeHeight = 10
)

func printLarge(date time.Time, width, height int) {
	spaceTop := (height - 7) / 2
	printReturn(spaceTop)
	printTimeLarge(date, width)
	printReturn(1)
	printDateLarge(date, width)
	printReturn(1)
}

func printTimeLarge(date time.Time, width int) {
	spaceLeft := (width - 35) / 2
	for i := 0; i < 7; i++ {
		printSpace(spaceLeft)
		hr, min, _ := date.Clock()
		hr1 := hr / 10
		hr2 := hr % 10
		min1 := min / 10
		min2 := min % 10
		printString(largeSub(largeDigital[i], hr1))
		printString(largeSub(largeDigital[i], hr2))
		printString(largeColon[i])
		printString(largeSub(largeDigital[i], min1))
		printString(largeSub(largeDigital[i], min2))
		printReturn(1)
	}
}

func printDateLarge(date time.Time, width int) {
	d := date.Format("2006-01-02 Mon")
	spaceLeft := (width + 1 - len(d)) / 2
	printSpace(spaceLeft)
	printString(d)
	printReturn(1)
}

func largeSub(l string, n int) string {
	n = n % 10
	r := []rune(l)
	return string(r[n*8 : (n+1)*8])
}

var largeDigital = []string{
	" ██████    ██    ██████  ██████  ██  ██  ██████  ██████  ██████  ██████  ██████ ",
	" ██  ██    ██        ██      ██  ██  ██  ██      ██          ██  ██  ██  ██  ██ ",
	" ██  ██    ██        ██      ██  ██  ██  ██      ██          ██  ██  ██  ██  ██ ",
	" ██  ██    ██    ██████  ██████  ██████  ██████  ██████      ██  ██████  ██████ ",
	" ██  ██    ██    ██          ██      ██      ██  ██  ██      ██  ██  ██      ██ ",
	" ██  ██    ██    ██          ██      ██      ██  ██  ██      ██  ██  ██      ██ ",
	" ██████    ██    ██████  ██████      ██  ██████  ██████      ██  ██████  ██████ ",
}

var largeColon = []string{
	"    ",
	"    ",
	" ██ ",
	"    ",
	" ██ ",
	"    ",
	"    ",
}
