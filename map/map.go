// // An Empty map
// map[Key_Type]Value_Type{}

// // Map with key-value pair
// map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

package main

import "fmt"

func main() {
	var abc map[int]int

	abc = map[int]int{
		11: 10,
		12: 20,
		13: 30,
	}

	for x, j := range abc {
		//abc[x] = abc[x] * 10

		if abc[x] == 20 {
			fmt.Print("\n", x, j)
		}

	}

	abc[14] = 40
	abc[15] = 50

	number, ok := abc[15]
	fmt.Println("\nKey present or not:", ok)
	fmt.Println("Value:", number)
	fmt.Print("\n", abc)

	xyz := map[int]string{
		1: "gaurav",
		2: "bavita",
		3: "Aradhya",
	}

	xyz1 := xyz
	delete(xyz, 3)

	fmt.Print("\n", xyz)
	fmt.Print("\n", xyz1)

	ggg := make(map[int]string)
	fmt.Print("\nvalue of map made by make", ggg)

	ggg[1] = "aadha"
	ggg[2] = "bobo"
	fmt.Print("\nvalue of map made by make after giving data", ggg)

}
