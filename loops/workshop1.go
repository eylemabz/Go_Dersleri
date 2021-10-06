package loops

import "fmt"

func Demo3() {
	//bir sayı tut
	//kullanıcıdan sayı al
	//sayıyı tahmin edelim

	/* Benim Kodum
	aklimdakiSayi := 34
	var tahminEdilenSayi int
	fmt.Println("1 ile 100 arasında tuttuğum sayıyı tahmin et!  3 hakkın var.")

	for i := 1; i <= 3; i++ {
		fmt.Scanln(&tahminEdilenSayi)
		if aklimdakiSayi < tahminEdilenSayi {
			fmt.Println("Tuttuğum sayı daha küçük")
		} else if aklimdakiSayi > tahminEdilenSayi {
			fmt.Println("Tuttuğum sayı daha büyük")
		} else if aklimdakiSayi == tahminEdilenSayi {
			fmt.Println("Tebrikler Doğru!!")
			break
		}
	}
	if aklimdakiSayi != tahminEdilenSayi {
		fmt.Println("Oyunu Kaybettin :((")
	}
	*/
	//Eğitmenin Kodu

	aklimdakiSayi := 80
	tahminEdilenSayi := 0

	fmt.Println("Bir sayı tahmin edin.")
	fmt.Scanln(&tahminEdilenSayi)
	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha büyük bir sayi giriniz.")
			fmt.Scanln(&tahminEdilenSayi)
		}
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük bir sayi giriniz.")
			fmt.Scanln(&tahminEdilenSayi)
		}
	}
	if aklimdakiSayi == tahminEdilenSayi {
		fmt.Println("Bravo Bildiniz.")
	}
}
