package array

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println(" twoSum test is beginning... ")
	m.Run()
}

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if len(result) != 2 {
		t.Fail()
	}
}

func BenchmarkTwoSum(b *testing.B) {

}
