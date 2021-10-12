package defer_statement

import "fmt"

func Demo2(sayi int) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		fmt.Println("Çift Sayı Çalıştı")
		return "Çift Sayıdır"
	}
	if sayi%2 != 0 {

		return "Sayı Tektir"
	}

	return "Belli Değil"
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
