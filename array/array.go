package array

import "sort"

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
// description: 给你 n 个非负整数 a1，a2，…，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai)
// 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。    说明：你不能倾斜容器。
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
