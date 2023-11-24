package printer

import (
	"errors"
	"os"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// Print print clock
func Print(date time.Time) error {
	clearScreen()
	w, h, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}
	if w >= largeWidth && h >= largeHeight {
		printLarge(date, w, h)
		return nil
	}
	return errors.New("terminal window size too small")
}
