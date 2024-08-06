package main

func repeatChar(c string, count int) string {
	var ans string
	for i := 0; i < count; i++ {
		ans += c
	}
	return ans
}
