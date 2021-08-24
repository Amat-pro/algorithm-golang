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

func BenchmarkTwoSum(b *testing.B) {

}

func TestMaxArea(t *testing.T) {
	maxArea1 := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	maxArea2 := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if maxArea1 != 49 || maxArea2 != 49 {
		t.Fail()
	}
}
