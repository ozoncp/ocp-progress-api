package main

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/ozoncp/ocp-progress-api/internal/utils"
)

func main() {

	fmt.Printf("Hello, my name is Dima Larin. I`ll work on progress-api %v \n", emoji.WavingHand.Tone(emoji.Dark))

	myTests()
	//var l int = 16
	//var k int = 33
	//fmt.Println(math.Ceil(float64(l) / float64(k)))
	//fmt.Println(float64(l) / float64(k))

}

func myTests() {
	newArray := []int{1, 2, 3, 4, 5, 6}

	rez := utils.SplitSlice(newArray, 4)
	fmt.Print("rez = ", rez)
	fmt.Print("\n\n")

	dick := make(map[string]string, 5)
	dick["one"] = "1"
	dick["one"] = "2"
	fmt.Print(dick)

	fmt.Print("\n\n")
	var v *int
	//fmt.Println(*v)
	fmt.Println(v) //<nil>
	v = new(int)
	fmt.Println(*v) //
	fmt.Println(v)  //0xc00004c088
}
