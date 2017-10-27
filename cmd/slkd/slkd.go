// slkd is daemon process. It gets channel history and shows
// unseen messages. That simple.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yarikbratashchuk/slk/internal/api"
	"github.com/yarikbratashchuk/slk/internal/config"
	"github.com/yarikbratashchuk/slk/internal/log"
	"github.com/yarikbratashchuk/slk/internal/message"
	"github.com/yarikbratashchuk/slk/internal/print"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-sigs
		os.Exit(0)
	}()

	for {
		conf, err := config.Read()
		if err != nil {
			log.Fatal(err)
		}

		hist, err := api.GetChannelHistory(conf)
		if err != nil {
			log.Fatalf("\nslkd: %s", err)
		}

		diff := message.TsFilterNewer(conf.ChannelTs[conf.Channel], hist)

		diff = message.ReverseOrder(diff)
		message.RemoveURefs(diff)
		message.FormatLines(hist)

		print.ListenChat(conf.Username, conf.Users, diff)

		if len(hist) == 0 {
			continue
		}

		conf.ChannelTs[conf.Channel] = hist[0].Ts
		if err := config.Write(conf); err != nil {
			fmt.Printf("\nslkd: %s", err)
		}

		time.Sleep(3 * time.Second)
	}
}
