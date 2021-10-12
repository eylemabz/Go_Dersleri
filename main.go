package main

import "golesson/interfaces"

func main() {
	//go goroutines.CiftSayilar()
	//conditionals.Demo1()
	//conditionals.Demo2()
	//conditionals.Demo3()
	//loops.Demo1()
	//loops.Demo2()
	//loops.Demo3()
	//loops.Demo4()
	//loops.Demo5()
	//arrays.Demo1()
	//arrays.Demo2()
	//arrays.Demo3()
	//slices.Demo1()
	//slices.Demo2()

	/*	functions.SelamVer("Eylem")
		var sonuc = functions.Topla(60, 50)
		fmt.Println(sonuc)*/

	/*
		sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)
		fmt.Println("toplam:", sonuc1)
		fmt.Println("çıkarım:", sonuc2)
		fmt.Println("çarpım:", sonuc3)
		//fmt.Println("bölüm:", sonuc4)
	*/

	/*
		fmt.Println(functions.ToplaVaryadic(1, 2, 6, 9, 4, 8))
		fmt.Println(functions.ToplaVaryadic(2, 4))
		fmt.Println(functions.ToplaVaryadic())

		sayilar := []int{4, 6, 5, 7}
		fmt.Println(functions.ToplaVaryadic(sayilar...))
	*/

	//maps.Demo1()
	//for_range.Demo2()
	//for_range.Demo3()

	/*
		sayi := 20
		pointers.Demo1(&sayi)
		fmt.Println("maindeki sayi:", sayi)
	*/

	/*
		sayilar := []int{1, 2, 3}
		pointers.Demo2(sayilar)
		fmt.Println("maindeli sayı:", sayilar[0])

	*/

	//structs.Demo1()
	//structs.Demo2()

	/* go goroutines.CiftSayilar()
	go goroutines.TekSayilar()
	time.Sleep(5 * time.Second)
	fmt.Println("main bitti")

	*/

	/*

		ciftSayiCn := make(chan int)
		tekSayiCn := make(chan int)
		go channels.CiftSayilar(ciftSayiCn)
		go channels.TekSayilar(tekSayiCn)

		ciftsayiToplam, teksayiToplam := <-ciftSayiCn, <-tekSayiCn
		carpim := ciftsayiToplam * teksayiToplam
		fmt.Println("carpım: ", carpim)

	*/

	//	interfaces.Demo1()
	interfaces.Demo2()
}
