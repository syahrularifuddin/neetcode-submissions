func findMaxConsecutiveOnes(nums []int) int {
    var cons, max int
    for i, _ := range nums {
        if nums[i] != 1 {
            cons = 0
            continue
        }
        cons++
        if cons > max {
            max = cons
        }
    }
    return max
}
