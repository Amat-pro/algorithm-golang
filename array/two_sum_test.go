package array

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println(" my test is beginning... ")
	// m.run() -> trigger test execute
	m.Run()
}

func TestTwoSum1(t *testing.T) {
	// t.run(name, func)
}

func BenchmarkTwoSum(b *testing.B) {

}
