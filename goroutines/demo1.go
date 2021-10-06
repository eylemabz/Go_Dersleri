package goroutines

import "log"

func CiftSayilar() {
	for i := -0; i <= 10; i += 2 {
		log.Println("Çift sayı:", i)
	}
}
func TekSayilar() {
	for i := 1; i <= 10; i += 2 {
		log.Println("Tek sayı:", i)
	}
}
