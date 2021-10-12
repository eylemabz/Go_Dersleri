package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Eylem"
	fmt.Println(s.HasPrefix(isim, "Ey")) //prefix=ön ek
	fmt.Println(s.HasSuffix(isim, "Ey")) //suffix=son ek
	fmt.Println(s.Index(isim, "lem"))    // lem de l harfi kaçıncı indexteyse onu döndürüyor
	harfler := []string{"e", "y", "l", "e", "m"}
	fmt.Println(s.Join(harfler, ""))

	sonuc := s.Join(harfler, "-")
	fmt.Println(s.Replace(sonuc, "*", "-", -1)) //-1 yerine 3 verisek en fazla 3 tane değiştir demek

	fmt.Println(s.Split(sonuc, "  "))
	fmt.Println(s.Repeat(sonuc, 5))
}
