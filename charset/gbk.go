package charset

/*
borrowed from https://studygolang.com/articles/24814
*/
func DetectGBK(content []byte) (isGBK bool) {
	for i := 0; i < len(content); {
		if content[i] <= 0x7f {
			// compatible with ascii
			i++
			continue
		} else {
			// gbk
			if 0x81 <= content[i] && content[i] <= 0xfe && len(content) > i+1 &&
				0x40 <= content[i+1] && content[i+1] <= 0xfe && content[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return
			}
		}
	}
	return true
}
