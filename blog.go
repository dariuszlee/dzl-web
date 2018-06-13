package main

import ()

func getBlogs(amount int) []blog {
	b1 := []blog{}
	for i := 0; i < amount; i++ {
		b := blog{"First Post", "This is a longer description, etc, etc", i}
		b1 = append(b1, b)
	}
	return b1
}
