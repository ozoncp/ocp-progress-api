// package utils is for necessary functions
package utils

import (
	"math"
	"sync"
)

//SplitSlice  using for separate slice for new slices
func SplitSlice(input []int, n int) [][]int {

	if tmp := [][]int{{}}; n <= 0 || len(input) == 0 || input == nil {
		return tmp // почему нельзя return [][]int или return [[]] ??????
	}

	var wg sync.WaitGroup
	mainSlice := make([][]int, int(math.Ceil(float64(len(input))/float64(n))))

	for i := 0; i < int(math.Ceil(float64(len(input))/float64(n))); i++ {
		wg.Add(1)
		addSliceToMainSlice(&wg, &mainSlice, i, &input, i*n, n)
	}
	wg.Wait()
	return mainSlice
}

func addSliceToMainSlice(wg *sync.WaitGroup, mainSlice *[][]int, pos int, input *[]int, from int, length int) {

	defer func() { wg.Done() }()

	if (from + length) > len(*input) {
		length = len(*input) - from
	}

	(*mainSlice)[pos] = make([]int, length)

	copy((*mainSlice)[pos], (*input)[from:from+length])
}

//ReverseKeyValue using for creation new reverse map
func ReverseKeyValue(inputMap map[string]int) map[int]string {

	outPutMap := make(map[int]string, len(inputMap))

	for key, value := range inputMap {
		outPutMap[value] = key
	}
	return outPutMap
}

//FilterSlice filter for slice
func FilterSlice(inputSlice []string, blackList []string) []string {

	var outputSlice []string
	tmpSet := make(map[string]struct{}, len(blackList))

	for _, blackItem := range blackList {
		tmpSet[blackItem] = struct{}{}
	}

	for _, val := range inputSlice {

		if _, exist := tmpSet[val]; !exist {
			outputSlice = append(outputSlice, val)
		}

	}

	return outputSlice
}
