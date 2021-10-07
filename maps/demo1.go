package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"
	// silme işlemi delete
	fmt.Println(sozluk["book"])
	fmt.Println("eleman sayısı:", len(sozluk))
	fmt.Println("Sözlük:", sozluk)
	delete(sozluk, "book")
	fmt.Println("eleman sayısı:", len(sozluk))
	fmt.Println("Sözlük:", sozluk)
	//arama işlemi search
	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu :", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "pen": "kalem"} // başka bir yazım şekli
	fmt.Println(sozluk2)
}
