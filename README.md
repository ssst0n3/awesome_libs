# awesome libs

- [ ] cipher
- [x] error: print and trace
- [x] reflect: some useful reflect gadgets
- [x] slice
    - func In(item interface{}, slice interface{}) bool
- [x] python like format

## error
### use awesome_libs/log.Logger as Logger
```go
import "github.com/ssst0n3/awesome_libs/awesome_error"

func example()  {
    awesome_error.CheckError(errors.New("test"))
}
```

### use self-defined logger
```go
import (
    "github.com/ssst0n3/awesome_libs/awesome_error"
    "github.com/ssst0n3/awesome_libs/awesome_error/error_logger"
)

func example()  {
    l := logger.InitLogger("my-logger", os.Stdout)
    error_logger.Init(l)
    awesome_error.CheckError(errors.New("test"))
}
```

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