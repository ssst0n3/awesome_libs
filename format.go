package awesome_libs

import (
	"bytes"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"text/template"
)

var t = template.New("awesome_libs").Delims("{", "}")

func Format(tpl string, arg Dict, delimiters ...string) string {
	var msg bytes.Buffer

	if len(delimiters) == 2 {
		left := delimiters[0]
		right := delimiters[1]
		t = template.New("awesome_libs").Delims(left, right)
	}
	t, err := t.Parse(tpl)
	awesome_error.CheckFatal(err)

	awesome_error.CheckFatal(t.Execute(&msg, arg))
	return msg.String()
}
