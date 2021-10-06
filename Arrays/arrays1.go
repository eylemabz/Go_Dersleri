package arrays

import "fmt"

func Demo1() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "Adana"
	sehirler[2] = "İstanbul"
	sehirler[3] = "İzmir"
	sehirler[4] = "Diyarbakır"
	fmt.Println(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
