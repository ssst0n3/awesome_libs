package slice

func In(item interface{}, slice []interface{}) bool {
	result := false
	for _, v := range slice {
		if item == v {
			result = true
		}
	}
	return result
}

func InInt(item int, slice []int) bool {
	result := false
	for _, v := range slice {
		if item == v {
			result = true
		}
	}
	return result
}