package utils_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-progress-api/internal/utils"
)

type funcsToSplits func([]int, int) ([][]int, error)

func testSlpitSlice(f funcsToSplits, t *testing.T) {

	type TestCase struct {
		NumberOfTest uint
		InputSlice   []int
		N            int
		OutPut       [][]int
	}

	testCases := []TestCase{
		{1, []int{1, 2, 3, 4, 5, 6}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{2, []int{1, 2, 3, 4, 5, 6}, 4, [][]int{{1, 2, 3, 4}, {5, 6}}},
		{3, []int{1, 2, 3, 4, 5, 6}, 7, [][]int{{1, 2, 3, 4, 5, 6}}},
		{4, nil, 2, nil},
		{5, []int{1, 2, 3}, 0, nil},
	}

	for _, testCase := range testCases {
		res, _ := f(testCase.InputSlice, testCase.N)
		if !reflect.DeepEqual(res, testCase.OutPut) {
			fmt.Println("TC number = ", testCase.NumberOfTest, " Fail result: ", res, " Correct anser = ", testCase.OutPut)
			t.Error("fail ")
			return
		}
		fmt.Println("Good result: ", res)

	}
}

func TestSplitSlice(t *testing.T) {
	testSlpitSlice(utils.SplitSlice, t)
}

func TestSplitSliceAsynchDeepCopy(t *testing.T) {
	testSlpitSlice(utils.SplitSliceAsynchDeepCopy, t)
}

func TestReverseKeyValue(t *testing.T) {
	var m = map[string]int{}
	m["a"] = 1
	m["b"] = 2

	correctAns := map[int]string{1: "a", 2: "b"}

	res := utils.ReverseKeyValue(m)
	if !reflect.DeepEqual(res, correctAns) {
		fmt.Println("Fail result: ", res)
		t.Error("fail")
		return
	}
	fmt.Println("Good result: ", res)
}

func TestFilterSlice(t *testing.T) {
	var sl = []string{"a", "b", "c"}
	var bl = []string{"a", "c"}

	res := utils.FilterSlice(sl, bl)
	if !reflect.DeepEqual(res, []string{"b"}) {
		fmt.Println("Fail result: ", res)
		t.Error("fail")
		return
	}
	fmt.Println("Good result: ", res)
}
