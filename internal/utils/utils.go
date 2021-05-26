// package utils is for necessary functions
package utils

import (
	"errors"
	"math"
	"sync"

	"github.com/ozoncp/ocp-progress-api/core/progress"
)

func SplitSlice(input []int, n int) ([][]int, error) {

	if n <= 0 || len(input) == 0 || input == nil {
		return nil, errors.New("not correct input parameters")
	}

	// если считать в целых числах еще надо будет использовать операцию %
	newSliceLen := int(math.Ceil(float64(len(input)) / float64(n))) // вынес в отдельную переменную гораздо лушче стало
	mainSlice := make([][]int, newSliceLen)

	for i := 0; i < newSliceLen; i++ {
		length := n
		if (i*n + n) > len(input) {
			length = len(input) - i*n
		}
		mainSlice[i] = input[i*n : i*n+length]
	}
	return mainSlice, nil
}

//SplitSlice  using for separate slice for new slices
//this function came due to my interest just for me
func SplitSliceAsynchDeepCopy(input []int, n int) ([][]int, error) {

	if n <= 0 || len(input) == 0 || input == nil {
		return nil, errors.New("not correct input parameters")
	}

	var wg sync.WaitGroup
	// если считать в целых числах еще надо будет использовать операцию %
	newSliceLen := int(math.Ceil(float64(len(input)) / float64(n))) // вынес в отдельную переменную гораздо лушче стало
	mainSlice := make([][]int, newSliceLen)

	for i := 0; i < newSliceLen; i++ {
		wg.Add(1)
		go addSliceToMainSliceDeepCopy(&wg, &mainSlice, i, &input, i*n, n) // в комит в тот раз не попало go :((( сейчас должный быть операции асинхронны
	}
	wg.Wait()
	return mainSlice, nil
}

func addSliceToMainSliceDeepCopy(wg *sync.WaitGroup, mainSlice *[][]int, pos int, input *[]int, from int, length int) {

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

//SplitToBulks split slice of Pogress struct to slice of bulks
func SplitToBulks(users []progress.Pogress, n int) ([][]progress.Pogress, error) {
	if n <= 0 || len(users) == 0 || users == nil {
		return nil, errors.New("not correct input parameters")
	}

	newSliceLen := int(math.Ceil(float64(len(users)) / float64(n)))
	mainSlice := make([][]progress.Pogress, newSliceLen)

	for i := 0; i < newSliceLen; i++ {
		length := n
		if (i*n + n) > len(users) {
			length = len(users) - i*n
		}
		mainSlice[i] = users[i*n : i*n+length]
	}
	return mainSlice, nil
}

//CreatMapFromSlise creat map from slice of Pogress struct
func CreatMapFromSlise(users []progress.Pogress) (map[uint64]progress.Pogress, error) {
	if len(users) == 0 || users == nil {
		return nil, errors.New("not correct input parameters")
	}

	resalt := make(map[uint64]progress.Pogress, len(users))

	for _, progesValue := range users {
		if _, foundKey := resalt[progesValue.Id]; foundKey {
			// как я понял ни разработчик ни программа не должны паниковать
			//panic(fmt.Sprintf("duplicate value %d", progesValue.UserId))

			// по этому возвращаю nil и ошибку
			return nil, errors.New("duplicate value " + string(progesValue.Id))
		}
		resalt[progesValue.Id] = progesValue
	}
	return resalt, nil
}
