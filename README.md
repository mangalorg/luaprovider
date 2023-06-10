# Lua Provider

This is a generic provider for [libmangal](https://github.com/mangalorg/libmangal)
that uses [Lua](https://www.lua.org/) scripts to create *subproviders*.

It uses [native go implementation of the Lua interpreter](https://github.com/yuin/gopher-lua)
and provides a set of libraries that can be used in the scripts.

## Features

- Built-in Lua VM without CGO
- Batteries-included library
- Ships with [CLI tool](./cmd/lua-provider-gen) to generate script template
- Luadoc generation which enables autocompletion for you IDE
- Script template generation

## Scripts

Scripts must contain these global functions:

```lua
function SearchMangas(query)
  return {}
end

function MangaVolumes(manga)
 return {} 
end

function VolumeChapters(volume)
 return {} 
end

function ChapterPages(chapter)
  return {}
end
```

The scripts can load sdk with

```lua
local sdk = require("sdk")
```

Which provides these packages:

- [crypto](./lib/crypto)
  - [md5](./lib/crypto/md5)
  - [sha1](./lib/crypto/sha1)
  - [sha256](./lib/crypto/sha256)
  - [sha512](./lib/crypto/sha512)
  - [aes](./lib/crypto/aes)
- [encoding](./lib/encoding)
  - [base64](./lib/encoding/base64)
  - [json](./lib/encoding/json)
- [html](./lib/html)
- [http](./lib/http)
- [js](./lib/js) - a javascript virtual machine
- [regexp](./lib/regexp)
- [time](./lib/time)
- [strings](./lib/strings)
- [levenshtein](./lib/levenshtein)
- [util](./lib/util)