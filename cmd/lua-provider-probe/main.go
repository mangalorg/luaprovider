package main

import (
	"encoding/json"
	"fmt"
	"github.com/mangalorg/luaprovider"
	"log"
	"os"
)

func printUsage() {
	_, _ = fmt.Fprintln(os.Stderr, `Usage:
lua-provider-probe <path>
lua-provider-probe [-h]elp `)
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	arg := os.Args[1]

	switch arg {
	case "-help", "-h":
		printUsage()
		return
	default:
		file, err := os.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
			return
		}

		loader, err := luaprovider.NewLoader(file, luaprovider.Options{})
		if err != nil {
			log.Fatal(err)
			return
		}

		info, err := json.MarshalIndent(loader.Info(), "", "  ")
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(string(info))
	}
}
