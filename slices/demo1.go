package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	isimler[0] = "Eylem"
	isimler[1] = "Çiçek"
	isimler[2] = "Gargi"

	isimler = append(isimler, "Büşra")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
