<div align="center">
  <img width="150px" alt="a logo depicting a moon" src="https://github.com/mangalorg/luaprovider/assets/62389790/219a58d4-1ca9-484f-a34b-389a2486fed0">
  <h1>Lua Provider</h1>
</div>

> **Warning**
> 
> This is a beta software. The API is not stable and may change at any time.

This is a generic provider for [libmangal](https://github.com/mangalorg/libmangal)
that uses [Lua](https://www.lua.org/) scripts to create *subproviders*.

It uses [native go implementation of the Lua interpreter](https://github.com/yuin/gopher-lua)
and provides a set of libraries that can be used in the scripts.

Take a look at [the official lua scripts repository](https://github.com/mangalorg/saturno)

## Features

- Built-in Lua VM without CGO with [gopher-lua](https://github.com/yuin/gopher-lua)
- Batteries-included library
- Ships with [CLI helper tools](./cmd) for generating templates & probing.
- Luadoc generation which enables autocompletion for you IDE
- Script template generation

> **Note**
> 
> It is recommended to use [lua-language-server](https://github.com/LuaLS/lua-language-server)
> to get nice completions for your IDE
> 
> [VSCode extension](https://marketplace.visualstudio.com/items?itemName=sumneko.lua)

## Scripts

### Overview

See [examples of scripts](https://github.com/mangalorg/saturno/tree/main/luas)

Scripts must look like this:

```lua
---> name: Script name
---> description: Script description
---> version: v0.1.0
---> website: https://example.com

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

Notice the four required global functions
- `SearchMangas` - searches for mangas based on the given query.
- `MangaVolumes` - gets manga volumes. Each manga must have at least 1 volume.
- `VolumeChapters` - gets chapters of the given volume.
- `ChapterPages` - gets pages of the given chapter.

Also, each script **must** contain the following lines:

```lua
---> name: Script name
---> description: Script description
---> version: v0.1.0
---> website: https://example.com
```

Basically, `--->` indicates that this line contains script information field in YAML format.

- `name` - Name of the script
- `description` - Description of the script
- `version` - Script version. It must be a valid [semver](https://semver.org/)

---

The scripts can load sdk with

```lua
local sdk = require("sdk")
```

Which provides these packages:

<details>
<summary>Packages</summary>

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

Headless browser coming soon...
</details>

### Developing

Install [lua-provider-gen](./cmd/lua-provider-gen) helper tool

Then use it to generate a new workflow

```bash
lua-provider-gen -sdk -provider
```

This command will create the following files:

- `sdk.lua` - gives IDE autocompletion
- `provider.lua` - a provider script template