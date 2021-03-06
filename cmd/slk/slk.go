// The slk command provides inteface to slack api.
package main

import (
	"fmt"
	"os"

	"github.com/yarikbratashchuk/slk/internal/cli"

	_ "github.com/yarikbratashchuk/slk/internal/cli/setup"

	_ "github.com/yarikbratashchuk/slk/internal/cli/ignore"
	_ "github.com/yarikbratashchuk/slk/internal/cli/listen"

	_ "github.com/yarikbratashchuk/slk/internal/cli/read"
	_ "github.com/yarikbratashchuk/slk/internal/cli/write"

	_ "github.com/yarikbratashchuk/slk/internal/cli/on"
	_ "github.com/yarikbratashchuk/slk/internal/cli/to"
)

func main() {
	if len(os.Args) < 2 {
		usage(os.Args[0])
	}

	commands := cli.InitCommands()

	if len(os.Args) >= 2 && os.Args[1] == "help" {
		if len(os.Args) > 2 {
			if commands[os.Args[2]] != nil {
				commands[os.Args[2]]().Usage()
			}
		}
		usage(os.Args[0])
	}

	if c, ok := commands[os.Args[1]]; ok {
		c().Run()
		return
	}
	usage(os.Args[0])
}

func usage(command string) {
	fmt.Printf(`Usage: %s <command> <options>

Minimalistic slack cli. It has just a few simple commands.
It's NOT a complete replacement for slack client, but convenient
messaging cli.

Commands:
  setup  - set up and change slk configuration

  listen - start listen channel for messages
  ignore - stop listening channel for messages

  read   - get last 10 messages
  write  - write message

  to     - switch between channels
  on     - name of the current channel 
  
Options:
  command specific, use "slk help <command>" for details
`, command)
	os.Exit(0)
}
