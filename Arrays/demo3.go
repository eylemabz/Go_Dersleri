package arrays

import "fmt"

func Demo3() {
	var sayilar [2][3]int

	sayilar[0][0] = 1
	sayilar[0][1] = 2
	sayilar[0][2] = 3
	sayilar[1][0] = 4
	sayilar[1][1] = 5
	sayilar[1][2] = 6

	fmt.Println(sayilar[1][1])

	for i := 0; i < 2; i++ {
		for k := 0; k < 3; k++ {
			fmt.Print(sayilar[i][k])

			fmt.Println(" ")
		}
		fmt.Println()
	}
}
