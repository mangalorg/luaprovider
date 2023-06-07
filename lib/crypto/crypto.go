package crypto

import (
	luadoc "github.com/mangalorg/luaprovider/doc"
	"github.com/mangalorg/luaprovider/lib/crypto/aes"
	"github.com/mangalorg/luaprovider/lib/crypto/md5"
	"github.com/mangalorg/luaprovider/lib/crypto/sha1"
	"github.com/mangalorg/luaprovider/lib/crypto/sha256"
	"github.com/mangalorg/luaprovider/lib/crypto/sha512"
	lua "github.com/yuin/gopher-lua"
)

func Lib(L *lua.LState) *luadoc.Lib {
	return &luadoc.Lib{
		Name:        "crypto",
		Description: "Various cryptographic functions.",
		Libs: []*luadoc.Lib{
			aes.Lib(),
			md5.Lib(),
			sha1.Lib(),
			sha256.Lib(),
			sha512.Lib(),
		},
	}
}
