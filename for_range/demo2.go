package for_range

import "fmt"

func Demo2() {
	//1- 10 arası tek olan sayıları toplayan program
	//benim kodum
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	for _, sayi := range sayilar {

		if sayi%2 == 1 {
			toplam = toplam + sayi

		}
	}
	fmt.Println(toplam)
}
