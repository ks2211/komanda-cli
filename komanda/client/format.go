package client

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
)

func StatusMessage(view *gocui.View, data string) {
	timestamp := time.Now().Format(config.MessageTimestampFormat)

	fmt.Fprintf(view, "-> [%s] * %s\n",
		color.String(color.Yellow, timestamp),
		color.String(color.White, data),
	)
}
