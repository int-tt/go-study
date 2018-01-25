package main

import "fmt"

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	zero := []int{}
	fmt.Println("max:", max(vals[0], vals[1:]...))
	_, err := maxWithError(zero...)
	if err != nil {
		fmt.Println("maxWithError:", err.Error())
	}
	fmt.Println("max:", min(vals[0], vals[1:]...))

	_, err = minWithError(zero...)
	if err != nil {
		fmt.Println("minWithError:", err.Error())
	}

}
func max(first int, vals ...int) int {
	max := first
	for _, v := range vals {
		if max < v {
			max = v
		}
	}
	return max
}
func maxWithError(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("vals is zero slice")
	}
	max := 0
	for _, v := range vals {
		if max < v {
			max = v
		}
	}
	return max, nil
}

func min(first int, vals ...int) int {
	min := first
	for _, v := range vals {
		if min > v {
			min = v
		}
	}
	return min
}

func minWithError(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("vals is zero slice")
	}
	min := 0
	for _, v := range vals {
		if min > v {
			min = v
		}
	}
	return min, nil
}
