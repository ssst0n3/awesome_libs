package example

import "github.com/ssst0n3/awesome_libs"

func Format()  {
	awesome_libs.Format("Hello {.name}", awesome_libs.Dict{
		"name": "awesome",
	})
}