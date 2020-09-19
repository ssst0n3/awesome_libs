package log

import "fmt"

func Format(isRequest bool,  method string, rawUrl string, headers map[string][]string, body string, statusCode string) string {
	headersSplit := ""
	if isRequest {
		return fmt.Sprintf("%s %s\n%s\n%s", method, rawUrl, headersSplit, body)
	} else {
		return fmt.Sprintf("%s %s\n%s\n%s", method, rawUrl, headersSplit, body)
	}
}
