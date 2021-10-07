package slices

import "fmt"

func Demo2() {
	sehirler := []string{"ankara", "izmir", "istanbul"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)
	sehirler = append(sehirler, "adana") //append ile yepyeni bir slice oluyor onnu şehirlere aktardık kopya kısmına eklenmiyor
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[1:4]) //1. indisten itibaren 4 e kadar al demek (4 dahil değil)

}
