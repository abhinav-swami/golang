package main

import (
	"fmt"
)

// import different packages , no commas

func main() {

	// var x int
	// if value is not ASSIGNED it will have value as zero
	// var x int = 5
	// or
	x := 5
	var sum int = x + x
	if x > 2 {
		fmt.Println("greater then two")
	} else if x < 2 {
		fmt.Println("less then two")
	} else {
		fmt.Println("lesasafafafa")
	}
	fmt.Println(sum)
	// var a [5]int
	// a[3] = 7
	//length of array is fixed initial values are zero
	//or
	// a := [5]int{1, 2, 3, 4, 5}
	//length is fixed add empty block to put slice sort of method
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 6)
	fmt.Println(a)

	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12
	fmt.Println(vertices)
	// return map[dodecagon:12 square:3 triangle:2]
	fmt.Println(vertices["square"])
	//return 3
	//delete something from the map
	delete(vertices, "square")

	// for i := 0; i < 50000; i++ {
	// 	fmt.Println(i)
	// }

	//for loop can also work as while -
	// z := 2
	// for z < 1000 {
	// 	fmt.Println(z)
	// 	z++
	// }

	//array
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index : ", index, "value", value)
	}

	//calling other function
	result := sumat(5, 7)
	fmt.Println(result)

	//
	// result, err := sqrt(16)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }
	// fmt.Println(sqrt(5))

	user1 := logindata{name: "abhinav", totp: 1}
	fmt.Println(user1)
	fmt.Println(user1.totp)

	//print memory address using &
	fmt.Println(&user1.totp)

	//POINTERS: use * to orignally modify value from that memory referance point

	fmt.Println("alalalalal")
	fmt.Println("alalalalal")
}

// if want to return something give type of that
//name and type of argument

func sumat(
	a int,
	b int) int {
	return a + b
}

// function can have multiptle return values
// func sqrt(i float64) (float64, error) {
// 	if i < 0 {
// 		return 0, errors.New("undefined imaginary no.")
// 	}
// 	//need to return two values because we defined it
// 	return math.Sqrt(i), nil

// }

//str type
type logindata struct {
	name string
	totp int
}
