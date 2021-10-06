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

	/*
		Eğitmenin Kodu

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
	*/

	// kullanıcı kaç tahminde buldu sonucu onu yazacaz

	/* Benim Kodum :)
	aklimdakiSayi := 34
	var tahminEdilenSayi int
	fmt.Println("1 ile 100 arasında tuttuğum sayıyı tahmin et!")

	for i := 1; i <= 50; i++ {
		fmt.Scanln(&tahminEdilenSayi)
		if aklimdakiSayi < tahminEdilenSayi {
			fmt.Println("Daha küçük bir sayı gir")
		}
		if aklimdakiSayi > tahminEdilenSayi {
			fmt.Println("Daha büyük bir sayı gir")
		}
		if aklimdakiSayi == tahminEdilenSayi {

			if i <= 3 {
				fmt.Printf(" Süper %v. tahminde buldunuz.", i)
			}

			if i >= 4 && i <= 6 {
				fmt.Printf(" Güzel %v. tahminde buldunuz.", i)
			}

			if i > 6 {
				fmt.Printf(" Fena değil %v. tahminde buldunuz.", i)
			}
			break
		}

	}
	if aklimdakiSayi != tahminEdilenSayi {
		fmt.Println("Oyunu Kaybettin :((")
	}*/

	//Eğitmenin kodu daha güzel :))))
	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	tahminSayisi := 0
	fmt.Println("Bir sayı tahmin edin.")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1
	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha büyük bir sayi giriniz.")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük bir sayi giriniz.")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
	}

	basariDurumu := ""
	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Süper"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel"
	} else {
		basariDurumu = "Fena Değil"
	}
	fmt.Printf("Bravo Bildiniz. %v tahmin :%v", tahminSayisi, basariDurumu)

}
