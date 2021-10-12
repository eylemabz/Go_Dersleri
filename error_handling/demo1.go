package error_handling

import (
	"fmt"
	"os"
)

//type assertion
func Demo1() {
	f, err := os.Open("demo11.txt") // eğer varsa dosyayı dönüyor yoksa err dönüyor
	//dosya bulunursa err nil değerine eşit oluyor nil=null ay şey go da nil var

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı:", pErr.Path)
			return
		} else {
			fmt.Println("Dosya Bulunamadı")
			return
		}

	} else {
		fmt.Println(f.Name())
	}

}
