package string_functions

// burada alias kullanabiliriz mesela string için s kullanımı
import (
	"fmt"
	s "strings"
)

// go case sensitive bir dildir
func Demo1() {
	isim := "Eylem"
	fmt.Println(s.Count(isim, "E"))    //ismin içindde kaç tane büyük e var demek
	fmt.Println(s.Contains(isim, "E")) //ismin içinde e var mı yok mu  demek
	fmt.Println(s.Index(isim, "E"))    //aranan ifadenin ismin içinde kaçıncı indexte olduğunu gösterir.Eğer aranan ifadeyi bulamazsa -1 değerini döndürür

	sonuc := s.Index(isim, "a")
	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))

}
