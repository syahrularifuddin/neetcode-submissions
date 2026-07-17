func countStudents(students []int, sandwiches []int) int {
    var i int

	for i <= len(students)-1 {
		if len(sandwiches) == 0 {
			break
		}
		if students[i] == sandwiches[0] {
			students = requeue(students, i)
			dequeue(&students)
			pop(&sandwiches)
			i = 0
			continue
		}
		i++
	}
	return len(students)
}

func requeue(students []int, i int) []int {
	if i == 0 {
		return students
	}
	return append(students[i:], students[:i]...)
}

func dequeue(students *[]int) {
	if len(*students) == 0 {
		return
	}
	*students = (*students)[1:]
}

func pop(sandwich *[]int) {
	if len(*sandwich) == 0 {
		return
	}
	*sandwich = (*sandwich)[1:]
}