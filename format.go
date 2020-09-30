package awesome_libs

import (
	"bytes"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"text/template"
)

func Format(tpl string, arg Dict) string {
	var msg bytes.Buffer

	t, err := template.New("awesome").Parse(tpl)
	awesome_error.CheckFatal(err)

	awesome_error.CheckFatal(t.Execute(&msg, arg))
	return msg.String()
}
