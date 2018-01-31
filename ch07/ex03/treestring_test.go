package treestring

import "testing"

func TestString(t *testing.T) {
	excepted := "{0 1 2 3 4 5 6 7 8 9 10}"
	data := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{5, 4, 3, 2, 1, 10, 9, 8, 7, 6, 0},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	}

	for _, array := range data {
		var tree *tree
		for _, d := range array {
			tree = add(tree, d)
		}

		if tree.String() != excepted {
			t.Errorf("tree shoud be %s, tree is %s\n", excepted, tree.String())
		}
	}
}
