package main

import ()

func getBlogs() []blog {
	b := blog{"First Post", "This is a longer description, etc, etc"}
	b1 := []blog{b, b, b}
	return b1
}
