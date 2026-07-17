func removeElement(nums []int, val int) int {
    var w int
	for r := 0; r < len(nums); r++ {
		if nums[r] != val {
			nums[w] = nums[r]
			w++
		}
	}
	return w
}
