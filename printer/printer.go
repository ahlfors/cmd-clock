package printer

import (
	"errors"
	"os"
	"time"

	"github.com/lonord/cmd-clock/weather"
	"golang.org/x/crypto/ssh/terminal"
)

// Print print clock and weather
func Print(date time.Time, wea *weather.Weather) error {
	clearScreen()
	w, h, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}
	if w >= largeWidth && h >= largeHeight {
		printLarge(date, wea, w, h)
		return nil
	}
	return errors.New("terminal window size too small")
}
