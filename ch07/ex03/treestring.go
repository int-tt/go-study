package treestring

import "strconv"

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	str := "{" + t.string() + "}"
	return str
}
func (t *tree) string() string {
	str := ""
	if t == nil {
		return str
	}

	//一番左端(小さい値)まで進む
	if t.left != nil {
		str = t.left.string() + " "
	}
	s := strconv.Itoa(t.value)
	str += s
	if t.right != nil {
		str = str + " " + t.right.string()
	}
	return str

}
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[0:], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)

	} else {
		t.right = add(t.right, value)
	}
	return t
}
