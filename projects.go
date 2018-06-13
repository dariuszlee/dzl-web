package main

import ()

func getProjects(amount int) []project {
	b1 := []project{}
	for i := 0; i < amount; i++ {
		b := project{"First Post", "This is a longer description, etc, etc", i}
		b1 = append(b1, b)
	}
	return b1
}
