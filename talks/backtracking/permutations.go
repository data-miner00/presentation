package main

import "fmt"

// permute prints every node of the search tree it visits, indented by depth.
func permute(path []byte, used []bool, letters string) {
	if len(path) > 0 {
		fmt.Printf("%*s%s\n", 2*(len(path)-1), "", path)
	}
	if len(path) == len(letters) {
		return // complete: a leaf of the tree
	}
	for i := range letters {
		if used[i] {
			continue // each letter at most once
		}
		used[i] = true // choose
		path = append(path, letters[i])
		permute(path, used, letters) // explore
		path = path[:len(path)-1]    // unchoose
		used[i] = false
	}
}

func main() {
	letters := "abc"
	permute(nil, make([]bool, len(letters)), letters)
}
