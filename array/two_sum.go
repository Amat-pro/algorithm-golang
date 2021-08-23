package array

// question: two sum
// description: 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标
func twoSum(nums []int, target int) []int {
	// step one: validated
	numsIndex := make(map[int]int, len(nums))
	for i, num := range nums {
		pair := target - num
		if j, ok := numsIndex[pair]; ok {
			return []int{i, j}
		}
		numsIndex[num] = i
	}
	return nil
}
