package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// InputVal()
	// PointerVal()
	// ArrVal()
	// Loop()
	// ArrLoop()
	AllData()
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

func Loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("Loop Count: ", i)
	}
}

func ArrLoop() {
	rev := []string{"Ab", "Cd", "Ef", "Gh"}

	for i, j := range rev {
		fmt.Println("Val ", i, " Data ", j)
	}

}

func AllData() {

	myMap := map[int]string{
		10: "SB",
		20: "DB",
		30: "PY",
		40: "GO",
	}

	for key, val := range myMap {
		fmt.Println("Key: ", key, " Value: ", val)
	}

	fmt.Println(myMap[30])
}
