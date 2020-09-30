# awesome libs

- [ ] cipher
- [x] error: print and trace
- [x] reflect: some useful reflect gadgets
- [x] slice
    - func In(item interface{}, slice interface{}) bool
- [x] python like format

## format
```go
import "github.com/ssst0n3/awesome_libs"

func Format()  {
	awesome_libs.Format("Hello {.name}", awesome_libs.Dict{
		"name": "awesome",
	})
}
```