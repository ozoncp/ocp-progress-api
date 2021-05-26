package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/enescakir/emoji"
	"github.com/ozoncp/ocp-progress-api/internal/utils"
)

const mainConfigName string = "config/config.yaml"

func main() {

	fmt.Printf("Hello, my name is Dima Larin. I`ll work on progress-api %v \n", emoji.WavingHand.Tone(emoji.Dark))

	//myTests()
	openAndCloseFile(5)
}

func openAndCloseFile(count int) error {

	var f error
	for i := 0; i < count; i++ {

		f = func() error {
			file, err := os.Open(mainConfigName)
			if err != nil {
				fmt.Println(err)
				return err
			}

			defer func() {
				fmt.Println("File closed")
				file.Close()
			}()

			in := bufio.NewReader(file)
			str, err := in.ReadString('\n')

			if errors.Is(err, io.EOF) {
				buf := make([]byte, in.Buffered())
				read, err := io.ReadFull(in, buf)
				if err != nil {
					fmt.Println(err)
					return err
				}
				str += string(buf[:read])
			} else if err != nil {
				fmt.Println("Error of read string from file ", err)
				return err
			}

			fmt.Println(str)
			time.Sleep(2 * time.Second)

			return nil
		}()
		if f != nil {
			break
		}
	}
	return f
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
