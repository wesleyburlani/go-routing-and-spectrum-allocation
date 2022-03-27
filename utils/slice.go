package utils

import "strconv"

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func IntSliceToStringSlice(slice []int) []string {
	var str []string
	for _, i := range slice {
		str = append(str, strconv.Itoa(i))
	}
	return str
}
