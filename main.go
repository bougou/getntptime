package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/beevik/ntp"
)

func main() {
	var ntpdserver = flag.String("server", "127.0.0.1", "ntpd server")
	flag.Parse()

	t, err := ntp.Time(*ntpdserver)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(t.UnixNano())
}
