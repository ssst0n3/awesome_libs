# awesome libs

- [ ] cipher
- [x] error: print and trace
- [x] reflect: some useful reflect gadgets
- [x] slice
    - func In(item interface{}, slice interface{}) bool
- [x] python like format

## cipher
There's a CommonCipher already inited for you. It's key read from SecretDir/awesome_libs_cipher_common.

```go
import "github.com/ssst0n3/awesome_libs/cipher"

func any()  {
    ...
    enc, err := cipher.CommonCipher.Encrypt(pt)
}
```

## python like format
```go
import "github.com/ssst0n3/awesome_libs"

func Format()  {
	awesome_libs.Format("Hello {.name}", awesome_libs.Dict{
		"name": "awesome",
	})
}
```