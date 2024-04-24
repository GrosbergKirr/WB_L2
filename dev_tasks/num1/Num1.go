package num1

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// CLI runs the go-ntp command line app and returns its exit status.
func NTP(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}

	if err := app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type appEnv struct {
	host string
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("ntp", flag.ContinueOnError)
	fl.StringVar(&app.host, "host", "pool.ntp.org", "ntp host")

	fmt.Println(app.host)
	if err := fl.Parse(args); err != nil {
		return err
	}

	return nil
}

func (app *appEnv) run() error {
	t, err := ntp.Time(app.host)
	fmt.Println(t)
	if err != nil {
		return err
	}
	fmt.Println(t.UTC().Format(time.UnixDate))
	return nil
}
