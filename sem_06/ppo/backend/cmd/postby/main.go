package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/app/http"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/app/tech"
)

func main() {
	uiFlag := flag.String("ui", "t", "config ui")
	cfgFile := flag.String("cfg", "./configs/config.yaml", "config file name")

	flag.Parse()

	if *uiFlag == "t" {
		a := tech.New()

		err := a.Init(*cfgFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "init: %s", err)
			os.Exit(1)
		}

		a.Run(context.Background())
	} else if *uiFlag == "w" {
		a := http.New()

		err := a.Init(*cfgFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "init: %s", err)
			os.Exit(1)
		}

		err = a.Run(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "run: %s", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintf(os.Stderr, "parse: flag error")
		os.Exit(1)
	}
}
