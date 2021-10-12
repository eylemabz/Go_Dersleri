package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {

	aklimdakisayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz")
	}
	if tahmin > aklimdakisayi {
		return "daha küçük bir sayı giriniz", nil
	}
	if tahmin < aklimdakisayi {
		return "daha büyük bir sayı giriniz", nil
	}

	return "Bildiniz", nil
}
func Demo2() {
	mesaj, hata := TahminEt(111)
	fmt.Println(mesaj)
	fmt.Println(hata)
	//fmt.Println(TahminEt(101))
	//fmt.Println(TahminEt(-10))
}
