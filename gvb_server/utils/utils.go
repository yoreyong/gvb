package utils

func InList(key string, list []string) (index int, ok bool) {
	for index, s := range list {
		if key == s {
			return index, true
		}
	}
	return -1, false
}
