package array

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println(" array test is beginning... ")
	m.Run()
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if len(result) != 2 {
		t.Fail()
	}
}

// example: testing.B
func BenchmarkTwoSum(b *testing.B) {

}

func TestMaxArea(t *testing.T) {
	maxArea1 := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	maxArea2 := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if maxArea1 != 49 || maxArea2 != 49 {
		t.Fail()
	}
}

// do test findMedianOfSortedArray2
func TestFindMedianOfSortedArrays2(t *testing.T) {

}

// todo test threeSum
func TestThreeSum(t *testing.T) {

}

// todo test threeSumClosest
func TestThreeSumClosest(t *testing.T) {

}

// todo test nurSum
func TestFourSum(t *testing.T) {

}

// todo test removeDuplicates1
func TestRemoveDuplicates1(t *testing.T) {

}

// todo test removeDuplicates2
func TestRemoveDuplicates2(t *testing.T) {

}

// todo test removeElement1
func TestRemoveElement1(t *testing.T) {

}

// todo test removeElement2
func TestRemoveElement2(t *testing.T) {

}

// todo test nextPermutation
func TestNextPermutation(t *testing.T) {

}

// todo test searchInRotatedSortedArray
func TestSearchInRotatedSortedArray(t *testing.T) {

}
