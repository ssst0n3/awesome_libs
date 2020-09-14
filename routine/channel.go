package routine

type Result struct {
	Key   string
	Value interface{}
}

func Channel2Map(c chan Result, m map[string]Result) {
	for {
		result := <-c
		m[result.Key] = result
	}
}

func Run(singleThreadFlag *bool, c chan Result, m map[string]Result) {
	if !*singleThreadFlag {
		*singleThreadFlag = true
		go Channel2Map(c, m)
	}
}
