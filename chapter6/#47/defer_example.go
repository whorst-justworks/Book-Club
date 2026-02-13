func main() {
	i := 0
	j := 0
	defer func(i int) {
		fmt.Println(i, j)
	}(i)
	i++
	j++
}

// Output:
// 0 1
