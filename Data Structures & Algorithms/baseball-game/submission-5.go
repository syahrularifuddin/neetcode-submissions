func calPoints(operations []string) int {
	var stack []int
	var n int
	for _, v := range operations {
		n = len(stack)
		switch v {
			case "+":
				stack = append(stack, stack[n-1]+stack[n-2])
			case "C":
				stack = stack[:n-1]
			case "D":
				stack = append(stack, 2*stack[n-1])
			default: 
				x, _ := strconv.Atoi(v)
				stack = append(stack, x)
		}
		fmt.Println(stack)
	}
	var sum int
	for i := range stack {
		sum += stack[i]
	}
	return sum
}
