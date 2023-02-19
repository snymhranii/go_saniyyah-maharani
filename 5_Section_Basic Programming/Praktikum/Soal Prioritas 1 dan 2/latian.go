package main

import "fmt"

func main() {
	// <== simbol ini di pake buat komentar, dan tidak di tampilkan di output
	//deklarasi variable dan tipe data beserta nilainya
	/*
		variable luas itu masuknya tipe data float sebab di pake untuk hasil inputan
		sisi di pake untuk menampung nilai dari rumus trapesium itu sendiri dimana teradapt sisi di antatra bangunnya
		begitu juga dengan variable tinggi
	*/
	var luas, sisi1, sisi2, tinggi float32
	//code proses
	/*
		#penggunaan Printf itu buat nampilin string (kata,kalimat) tapi tanpa baris baru seperti Println (print line)
		nah fmt sendiri itu package supaya kita bisa make perintah Printf,Println,atau Scan
		#nah penggunaan Scanf itu buat nampung nilai yang di input dari Printf,intinya ngescan apa yang kita ketik pas progam jalan
		di satu sisi dalam penggunaan Printf,Println , WAJIB di awali kata fmt.
	*/
	fmt.Println("=====LUAS TRAPESIUM=====")
	fmt.Println("INPUT NILAI : ")
	fmt.Printf("masukan nilai dari alas a => ")
	fmt.Scanf("%f\n", &sisi1)
	fmt.Printf("masukan nilai dari alas b => ")
	fmt.Scanf("%f\n", &sisi2)
	fmt.Printf("masukan nilai dari tinggi trapesium => ")
	fmt.Scanf("%f\n", &tinggi)
	//statement atau pernyataan
	//di sini deklarasiin rumus dari trapesium
	//usahakan pendeklarasian itu sebelum baris kode selanjutnya yang bakal nampilin nilai (output)
	//dan sesudah kita masukin nilai dari setiap variable yang ada (alas a ,alas b,tinggi)
	//jadi nnti nilainya masuk ,dan bisa kehidtun , proses ini yang di namakan SEQUENTIAL ALGOTIMA
	luas = 0.5 * (sisi1 + sisi2) * tinggi
	/*
		#penggunaan quoted
		# %f itu untuk tipe data jenis float
		# %d itu untuk tipe data jenis integer
		# %s itu untuk tipe data jenis string
		# %t itu untuk menentukan sebuah jenis tipe data
		# \n itu untuk newline (pindah baris)
		# %v itu untuk menampung nilai apapun yang di masukan user baik tipe data string,float,int,bool dsb
	*/
	//output setelah statement dan input proses user
	fmt.Println("=====<HASIL OPERASI>=====")
	fmt.Printf("luas dari trapesium\ndengan sisi a = %v dan sisi b = %v serta tinggi = %v \nadalah %v cm\n ", sisi1, sisi2, tinggi, luas)

}
