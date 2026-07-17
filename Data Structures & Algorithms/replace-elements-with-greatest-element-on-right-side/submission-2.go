func replaceElements(arr []int) []int {
	for i:=0;i<len(arr);i++{
		if i+1 == len(arr) {
			arr[i] = -1
			break
		}
		var max int
		for j:=i+1; j<len(arr);j++{
			if arr[j] >= max {
				max = arr[j]
			}
		}
		arr[i] = max
	}
	return arr
}