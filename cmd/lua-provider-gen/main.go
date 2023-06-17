package main

import (
	"flag"
	"github.com/mangalorg/luaprovider"
	"log"
	"os"
	"path/filepath"
)

const (
	filenameSDK       = "sdk.lua"
	filenameProvider  = "provider.lua"
	filenameLuarcJSON = ".luarc.json"
)

func main() {
	dir := flag.String("dir", ".", "output directory")

	generateSDK := flag.Bool("sdk", false, "generate sdk.lua file for language server")
	generateProvider := flag.Bool("provider", false, "generate provider.lua template")
	generateLuarc := flag.Bool("luarc", false, "generate .luarc.json file for language server configuration")

	flag.Parse()

	if !*generateSDK && !*generateProvider && !*generateLuarc {
		flag.Usage()
		return
	}

	if *generateSDK {
		err := os.WriteFile(
			filepath.Join(*dir, filenameSDK),
			[]byte(luaprovider.LuaDoc()),
			0655,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *generateProvider {
		err := os.WriteFile(
			filepath.Join(*dir, filenameProvider),
			[]byte(luaprovider.LuaTemplate()),
			0655,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *generateLuarc {
		err := os.WriteFile(
			filepath.Join(*dir, filenameLuarcJSON),
			[]byte(luaprovider.LuarcJSON()),
			0655,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
