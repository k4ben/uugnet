package art

import "strings"

var title = `
█ █ █ █ █▀▀          █
█ █ █ █ █ ▄ █▀█ █▀█ ▀█▀
█ █ █ █ █ █ █ █ █▀▀  █
▀▀▀ ▀▀▀ ▀▀▀ ▀ ▀ ▀▀▀  ▀
`

var tux = `
    .--.
   |o_o |
   |:_/ |
  //   \ \
 (|     | )
/'\_   _/'\
\___)=(___/
`

func clean(s string) string {
	return strings.Trim(s, "\n")
}

var (
	Title = clean(title)
	Tux   = clean(tux)
)
