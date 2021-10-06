package conditionals

import "fmt"

func Demo3() {
	//üç adet int değşkeni tanımlayınız.
	//ekrana en büyük olanı yazdırılsın

	/* Benim Yöntemim
	sayi1 := 20
	sayi2 := 25
	sayi3 := 11

	if sayi1 > sayi2 && sayi1 > sayi3 {
		fmt.Println("en büyük sayi:", sayi1)
	} else if sayi2 > sayi1 && sayi2 > sayi3 {
		fmt.Println("en büyük sayi:", sayi2)
	} else {
		fmt.Println("en büyük sayi:", sayi3)
	}
	*/

	// Eğitmenin görevi
	var sayi1, sayi2, sayi3 int = 1005, 60, 200
	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}
	//burada else if bloğu yaparsak sonuç hatalı olur çünkü ilk if doğru olduğu için hep o çalışır.
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("En büyük sayi:%v", enBuyuk)
}
