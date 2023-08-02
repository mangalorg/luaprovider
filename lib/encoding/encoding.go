package encoding

import (
	luadoc "github.com/mangalorg/gopher-luadoc"
	"github.com/mangalorg/luaprovider/lib/encoding/base64"
	"github.com/mangalorg/luaprovider/lib/encoding/json"
	lua "github.com/yuin/gopher-lua"
)

func Lib(L *lua.LState) *luadoc.Lib {
	return &luadoc.Lib{
		Name:        "encoding",
		Description: "",
		Libs: []*luadoc.Lib{
			base64.Lib(L),
			json.Lib(),
		},
	}
}
