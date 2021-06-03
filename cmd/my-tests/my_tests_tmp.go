package main

import (
	"fmt"

	"github.com/ozoncp/ocp-progress-api/internal/utils"
)

func main() {
	myTests()
	testWithInterface()
}

func myTests() {
	newArray := []int{1, 2, 3, 4, 5, 6}

	rez, _ := utils.SplitSlice(newArray, 4)
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

type Sl interface {
	String() string
}

type User struct {
	Name string
}

func (u *User) String() string {
	if u != nil {
		return "LOLO  = " + u.Name
	}
	return ""
}

func foo(s Sl) {
	fmt.Println(s.String())
}

func testWithInterface() {
	user := &User{"Anna"}
	foo(user)
	fmt.Println(user)

	user1 := &User{"Rita"}

	//var stringer Sl

	//stringer = user1

	user1.Name = "Masha"

	fmt.Println(user1)

}
