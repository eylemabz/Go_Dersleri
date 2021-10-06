package loops

import "fmt"

func Demo4() {
	//1. aşamada kullanıcıcan bir sayı girmesini isteyecez
	//kullanıcının girdiği sayının asal olup olmadığını bulacaz
	//Ben yazamadım kodu:( Eğitmenin kodu
	var sayi int
	fmt.Println("Bir sayı giriniz: ")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}
	if asalMi == true {
		fmt.Println("Asaldır")
	} else {
		fmt.Println("Asal Değildir")

	}

}
