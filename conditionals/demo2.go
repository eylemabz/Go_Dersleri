package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paranız hazılanıyor")
		fmt.Println("Dikkat hesapta paranız kalmadı.")
	} else {
		fmt.Println("Bitti.")
	}
}
