# 9_String - Advance Function - Pointer Method - Struct and Interface
1. String : Go menyediakan package strings, isinya banyak fungsi untuk keperluan pengolahan data string. Chapter ini berisi pembahasan mengenai beberapa fungsi yang ada di dalam package tersebut.
    Terdapat beberapa fungsi yaitu : 
    1.1 Fungsi strings.Contains()
Dipakai untuk deteksi apakah string (parameter kedua) merupakan bagian dari string lain (parameter pertama). Nilai kembaliannya berupa bool.

package main

import "fmt"
import "strings"

func main() {
    var isExists = strings.Contains("john wick", "wick")
    fmt.Println(isExists)
}
Variabel isExists akan bernilai true, karena string "wick" merupakan bagian dari "john wick"
    1.2 Fungsi strings.HasPrefix()
Digunakan untuk deteksi apakah sebuah string (parameter pertama) diawali string tertentu (parameter kedua).

var isPrefix1 = strings.HasPrefix("john wick", "jo")
fmt.Println(isPrefix1) // true

var isPrefix2 = strings.HasPrefix("john wick", "wi")
fmt.Println(isPrefix2) // false
    1.3 Fungsi strings.HasSuffix()
Digunakan untuk deteksi apakah sebuah string (parameter pertama) diakhiri string tertentu (parameter kedua).

var isSuffix1 = strings.HasSuffix("john wick", "ic")
fmt.Println(isSuffix1) // false

var isSuffix2 = strings.HasSuffix("john wick", "ck")
fmt.Println(isSuffix2) // true
    1.4 Fungsi strings.Count()
Memiliki kegunaan untuk menghitung jumlah karakter tertentu (parameter kedua) dari sebuah string (parameter pertama). Nilai kembalian fungsi ini adalah jumlah karakternya.

var howMany = strings.Count("ethan hunt", "t")
fmt.Println(howMany) // 2
Nilai yang dikembalikan 2, karena pada string "ethan hunt" terdapat dua buah karakter "t"
    1.5 Fungsi strings.Index()
Digunakan untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2
String "ha" berada pada posisi ke 2 dalam string "ethan hunt" (indeks dimulai dari 0). Jika diketemukan dua substring, maka yang diambil adalah yang pertama, contoh:

var index2 = strings.Index("ethan hunt", "n")
fmt.Println(index2) // 4

Dan masih banyak lagi fungsi yang tersedia.

2. Pointer : Pointer adalah 
reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai. Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori di mana nilai 4 disimpan, bukan nilai 4 itu sendiri.
Variabel-variabel yang memiliki reference atau alamat pointer yang sama, saling berhubungan satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang referensi-nya sama) yaitu nilainya ikut berubah.
 
 3. Method : Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek.
Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level private (level akses nantinya akan dibahas lebih detail pada chapter selanjutnya). Dan juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik.

Penerapan Method
Cara menerapkan method sedikit berbeda dibanding penggunaan fungsi. Ketika deklarasi, ditentukan juga siapa pemilik method tersebut. Contohnya bisa dilihat pada kode berikut:

package main

import "fmt"
import "strings"

type student struct {
    name  string
    grade int
}

func (s student) sayHello() {
    fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}
Cara deklarasi method sama seperti fungsi, hanya saja perlu ditambahkan deklarasi variabel objek di sela-sela keyword func dan nama fungsi. Struct yang digunakan akan menjadi pemilik method.

func (s student) sayHello() maksudnya adalah fungsi sayHello dideklarasikan sebagai method milik struct student. Pada contoh di atas struct student memiliki dua buah method, yaitu sayHello() dan getNameAt().

4. Interface : Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.
Interface merupakan tipe data. Nilai objek bertipe interface zero value-nya adalah nil. Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek konkret yang memiliki definisi method minimal sama dengan yang ada di interface-nya.

Penerapan Interface
Yang pertama perlu dilakukan untuk menerapkan interface adalah menyiapkan interface beserta definisi method nya. Keyword type dan interface digunakan untuk pendefinisian interface.

package main

import "fmt"
import "math"

type hitung interface {
    luas() float64
    keliling() float64
}
Pada kode di atas, interface hitung memiliki 2 definisi method, luas() dan keliling(). Interface ini nantinya digunakan sebagai tipe data pada variabel, di mana variabel tersebut akan menampung objek bangun datar hasil dari struct yang akan kita buat.

Dengan memanfaatkan interface hitung, perhitungan luas dan keliling bangun datar bisa dilakukan, tanpa perlu tahu jenis bangun datarnya sendiri itu apa.

Siapkan struct bangun datar lingkaran, struct ini memiliki method yang beberapa di antaranya terdefinisi di interface hitung.

5. Error Handling : Error handling berfungsi untuk menangani error yang terjadi sehingga software yang mengalami error tidak langsung crash dan software menjadi lebih andal (reliable). Penanganan error di Golang lebih simpel.
