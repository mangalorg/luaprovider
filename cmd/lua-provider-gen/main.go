package main

import (
	"flag"
	"github.com/mangalorg/luaprovider"
	"log"
	"os"
	"path/filepath"
)

func main() {
	name := flag.String("name", "provider", "name of the created provider")
	dir := flag.String("dir", ".", "output directory")
	flag.Parse()

	err := os.WriteFile(
		filepath.Join(*dir, "sdk.lua"),
		[]byte(luaprovider.LuaDoc()),
		0655,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(
		filepath.Join(*dir, *name+".lua"),
		[]byte(luaprovider.LuaTemplate()),
		0655,
	)
	if err != nil {
		log.Fatal(err)
	}
}
