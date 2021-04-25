package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/beevik/ntp"
)

func main() {
	var ntpdserver = flag.String("server", "0.pool.ntp.org", "ntpd server")
	var format = flag.String("format", "datetime", "output format, one of: unixnsec|unixmsec|unixsec|datetime")
	flag.Parse()

	t, err := ntp.Time(*ntpdserver)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch *format {
	case "unixnsec":
		fmt.Println(t.UnixNano())
	case "unixmsec":
		fmt.Println(t.UnixNano() / 1000000)
	case "unixsec":
		fmt.Println(t.Unix())
	case "datetime":
		fmt.Println(t)
	default:
		fmt.Println(t)
	}
}
