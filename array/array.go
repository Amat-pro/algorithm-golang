package array

import (
	"sort"
)

// 1.
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

// 2.
// question: container with most water
// description: 给你 n 个非负整数 a1，a2，…，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
// (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。    说明：你不能倾斜容器。
func maxArea(height []int) int {
	// step one: validated
	maxMulti := 0
	left, right := 0, len(height)-1
	for left < right {
		width := right - left
		height := min(height[left], height[right])
		maxMulti = max(maxMulti, width*height)
	}
	return maxMulti
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 3.
// question: man of two sorted arrays
// description: 寻找两个正序数组的中位数 -> 给定两个大小分别为m和n的正序（从小到大）数组nums1和nums2，找出并返回这两个正序数组的中位数
// method one: 归并排序 -> 找中位数  complexity(TIME): O(m+n)
func findMedianOOfSortedArrays1(nums1 []int, nums2 []int) float64 {
	return 0
}

// method two:  分治策略 -> complexity(TIME): O(log(M+N))
func findMedianOfSortedArrays2(nums1 []int, nums2 []int) float64 {
	// step one: validated
	// make len(nums1) <= len(nums2)

	// query kth mum -> it's index is kth-1
	// find the kth1 in nums1 and find the kth2 in nums2
	// if kth1 < kth2 then remove the elements before kth1 which in nums1 and keep going on finding (k-kth1)th

	// find kth or kth, k+1th according to odd-even
	len := len(nums1) + len(nums2)
	if len%2 != 0 {
		return float64(findKth(nums1, nums2, len/2))
	}
	kth := len / 2
	kth1 := (len + 1) / 2
	return float64((findKth(nums1, nums2, kth) + findKth(nums1, nums2, kth1)) / 2)
}

func findKth(nums1, nums2 []int, k int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		n1, n2 = n2, n1
		nums1, nums2 = nums2, nums1 // 为避免数组长度的分类讨论，先做预处理
	}

	if n1 == 0 {
		return nums2[k-1] // bingo
	}

	if k == 1 {
		return min(nums1[0], nums2[0]) // bingo
	}

	k1 := min(k/2, n1) // 避免越界
	k2 := k - k1       // 不能理想的 k/2, k/2 划分

	// fmt.Println(nums1, k1-1, nums2, k2-1)

	switch {
	case nums1[k1-1] < nums2[k2-1]:
		return findKth(nums1[k1:], nums2, k2) // 彻底舍弃区域 1
	case nums1[k1-1] > nums2[k2-1]:
		return findKth(nums1, nums2[k2:], k1) // 彻底舍弃区域 3
	default:
		return nums1[k1-1] // bingo
	}
}

// 4.
// question: three sum
// description:  给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
// 且不重复的三元组
// note: 答案中不可以包含重复的三元组
// 处理结果与元素原顺序无关，可排序预处理，方便去重
// 使用双指针遍历后部分剩余数组
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // [-4 -1 -1 0 1 2] // 预排序有 2 个好处：去重 & 指导双指针的下一步方向
	n := len(nums)
	var res [][]int
	for i, num := range nums {
		if num > 0 {
			break // 优化，再往后三个正数和不可能为 0
		}

		// 第一层遍历数向前去重
		if i > 0 && nums[i] == nums[i-1] { // 因为双指针从 i 之后取，不能使用 nums[i] == nums[i+1] 向后去重
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := num + nums[l] + nums[r]
			switch {
			case sum > 0:
				r--
			case sum < 0:
				l++
			default:
				res = append(res, []int{num, nums[l], nums[r]})
				// 第二层候选数向后去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}

// 5.
// question: three sum closest
// description: 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个
// 数的和。假定每组输入只存在唯一答案。
// 同样也是指针法
// 和 15 一样，排序预处理能知道双指针移动的方向，记录最小 abs
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	// todo Bit Operation
	minAbs := 1<<31 - 1
	minSum := 0

	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] { // 优化：可选的去重
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := num + nums[l] + nums[r]
			if abs(target-sum) < minAbs {
				minAbs = abs(target - sum)
				minSum = sum
			}
			switch {
			case sum < target:
				l++
			case sum > target:
				r--
			default:
				return target
			}
		}
	}
	return minSum
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// 6.
// question: four sum
// description: 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d
// 的值与 target 相等？找出所有满足条件且不重复的四元组
// note: 答案中不可以包含重复的四元组
//
// N 数之和的本质是，在有序数组内，寻找 N 个数的和恰好是 S
// 解决办法还是 3sum 3sum_closest 的双指针法，不过需要外部 N-2 层循环，内部双指针循环即可
// 注意双指针在遍历时外部所有循环要去重，指针移动时也要去重
//
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

	for i := 0; i < n-1; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 去重1
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] { // 去重2    // 注意条件：j>i+1 与 i>0 相同都是为了排除第一个相同数
				continue
			}
			head, tail := j+1, n-1
			for head < tail {
				sum := nums[i] + nums[j] + nums[head] + nums[tail]
				switch {
				case sum < target: // 向后走
					head++
				case sum > target: // 向前走
					tail--
				case sum == target: // 向前向后走
					res = append(res, []int{nums[i], nums[j], nums[head], nums[tail]})
					// 去重3：注意 for 循环条件的判断，避开死循环
					for head < tail && nums[head] == nums[head+1] {
						head++
					}
					for head < tail && nums[tail] == nums[tail-1] {
						tail--
					}
					head++
					tail--
				}
			}
		}
	}
	return res
}

// 7.
// question: remove duplicates from sorted arrays
// description: 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度
// note: 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成

// 针对有序数组，双指针法是十分常见且有用的
func removeDuplicates1(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums)-1 {
		if nums[fast] != nums[fast+1] { // 相邻的数不相等
			slow++
			fast++
			nums[slow] = nums[fast] // 将最新的新数存储到慢指针的位置
			continue
		}
		fast++
	}
	return slow + 1
}

// 充分利用数组有序的已知条件
func removeDuplicates2(nums []int) int {
	n := len(nums)
	l, r := 0, 1
	for r < n {
		if nums[l] < nums[r] { // 比我大就放到我的下一个
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
		r++
	}
	return l + 1
}

// 8.
// question: remove element
// description: 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度
// note: 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素

// 遍历数组一如既往地联想到双指针法
func removeElement1(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast] // 快指针在前，慢指针在后，遇到不等的元素就放到慢指针的位置
			slow++
		}
	}
	return slow
}

// 类似 26 的解法
func removeElement2(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			if slow != fast {
				nums[slow] = nums[fast]
			}
			slow++
			fast++
			continue
		}
		fast++
	}
	return slow
}

// 9.
// question: next permutation
// description: 实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）
// note: 必须 原地 修改，只允许使用额外常数空间

// 数组规律题
// 从后往前找第一个下降点 i，再从后往前找它的 ceil 值，交换
// 再将 [i+1:] 之后的数据从降序反转为升序，为最小序列
func nextPermutation(nums []int) {
	// 处理降序的 case
	desc := true
	n := len(nums)
	for i := range nums[:n-1] {
		if nums[i] < nums[i+1] {
			desc = false
		}
	}
	if desc {
		// todo write func of reverse
		// reverse(nums)
		return
	}

	// 从后向前找第一个下降的点
	var i int
	for i = n - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			i-- // 找到 2
			break
		}
	}

	// 从后向前，找向上最接近的值
	for j := n - 1; j > i; j-- {
		if nums[j] > nums[i] {
			nums[j], nums[i] = nums[i], nums[j] // 交换 2 和 3     // [1 3 7 4 2 1]
			break
		}
	}
	// todo write func of reverse
	// reverse(nums[i+1:]) // 反转 4 2 1    // [1 3 1 2 4 7]
}

// 10.
// question: search in rotated sorted array
// description: 整数数组 nums 按升序排列，数组中的值 互不相同
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ...,
// nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为
// [4,5,6,7,0,1,2] 。
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// note: 整数数组 nums 按升序排列，数组中的值 互不相同
// 类二分搜索  log(n)
// 最左边数 < 中间数则左侧有序，
// 最右边数 > 中间数则右侧有序
// 在缩小搜索区域时，一直只在确定的有序区域内查找
func searchInRotatedSortedArray(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		switch {
		case nums[mid] == target: // bingo
			return mid
		case nums[l] <= nums[mid]: // 左侧有序
			if nums[l] <= target && target < nums[mid] { // 保证 target 一定在有序的左侧内
				r = mid - 1
			} else {
				l = mid + 1
			}
		case nums[mid] <= nums[r]: // 右侧有序
			if nums[mid] < target && target <= nums[r] { // 保证 target 一定在有序的右侧内
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
