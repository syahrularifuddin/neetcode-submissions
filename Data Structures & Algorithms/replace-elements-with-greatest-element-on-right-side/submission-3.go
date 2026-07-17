func replaceElements(arr []int) []int {
	max := -1
	var res = make([]int, len(arr))
	for i:=len(arr)-1; i>=0; i--{
		res[i] = max
		if arr[i] > max {
			max = arr[i]
		}
	}
	return res
}