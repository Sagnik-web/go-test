package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// InputVal()
	// PointerVal()
	ArrVal()
}

func InputVal() {
	render := bufio.NewReader(os.Stdin)

	fmt.Println("--------SB Input-----------")
	inputData, _ := render.ReadString('\n')

	fmt.Println("My Name Is:", inputData)
}

func PointerVal() {

	var data int = 10

	pointerValue := &data

	fmt.Println("Memory Location ", pointerValue)
	fmt.Println("Value ", *pointerValue+15)
	fmt.Println("After Add Memory Location ", pointerValue)

}

func ArrVal() {
	var val [5]int

	val[0] = 10
	val[1] = 20
	val[2] = 30

	fmt.Println(len(val))

	var newArr = [5]int{12, 14, 15}
	fmt.Println(newArr)

}
