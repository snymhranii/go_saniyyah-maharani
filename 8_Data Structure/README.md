# 8_Data Structure
1. Array : Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel. Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatan, menjadikan elemen/data yang disimpan di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan. Default nilai tiap elemen array pada awalnya tergantung dari tipe datanya. Jika int maka tiap element zero value-nya adalah 0, jika bool maka false, dan seterusnya. Setiap elemen array memiliki indeks berupa angka yang merepresentasikan posisi urutan elemen tersebut. Indeks array dimulai dari 0.

Contoh penerapan array:

var names [4]string
names[0] = "trafalgar"
names[1] = "d"
names[2] = "water"
names[3] = "law"

2. Slice : Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.

Pembuatan slice mirip seperti pembuatan array, bedanya tidak perlu mendefinisikan jumlah elemen ketika awal deklarasi. Pengaksesan nilai elemen-nya juga sama. Contoh pembuatan slice:

var fruits = []string{"apple", "grape", "banana", "melon"}
fmt.Println(fruits[0]) // "apple"
Salah satu perbedaan slice dan array bisa diketahui pada saat deklarasi variabel-nya, jika jumlah elemen tidak dituliskan, maka variabel tersebut adalah slice.

var fruitsA = []string{"apple", "grape"}      // slice
var fruitsB = [2]string{"banana", "melon"}    // array
var fruitsC = [...]string{"papaya", "grape"}  // array
 
 3. Map : Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair. Untuk setiap data (atau value) yang disimpan, disiapkan juga key-nya. Key harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.

Kalau dilihat, map mirip seperti slice, hanya saja indeks yang digunakan untuk pengaksesan bisa ditentukan sendiri tipe-nya (indeks tersebut adalah key).

Cara menggunakan map cukup dengan menuliskan keyword map diikuti tipe data key dan value-nya. Agar lebih mudah dipahami, silakan perhatikan contoh di bawah ini.

var chicken map[string]int
chicken = map[string]int{}

chicken["januari"] = 50
chicken["februari"] = 40

fmt.Println("januari", chicken["januari"]) // januari 50
fmt.Println("mei",     chicken["mei"])     // mei 0
Variabel chicken dideklarasikan sebagai map, dengan tipe data key adalah string dan value-nya int. Dari kode tersebut bisa dilihat bagaimana cara penggunaan keyword map.

Kode map[string]int maknanya adalah, tipe data map dengan key bertipe string dan value bertipe int.

Default nilai variabel map adalah nil. Oleh karena itu perlu dilakukan inisialisasi nilai default di awal, caranya cukup dengan tambahkan kurung kurawal pada akhir tipe, contoh seperti pada kode di atas: map[string]int{}.

Cara menge-set nilai pada sebuah map adalah dengan menuliskan variabel-nya, kemudian disisipkan key pada kurung siku variabel (mirip seperti cara pengaksesan elemen slice), lalu isi nilainya. Contohnya seperti chicken["februari"] = 40. Sedangkan cara pengambilan value adalah cukup dengan menyisipkan key pada kurung siku variabel.

Pengisian data pada map bersifat overwrite, ketika variabel sudah memiliki item dengan key yang sama, maka value lama akan ditimpa dengan value baru.

4. Function : Function adalah sekelompok pernyataan yang bersama-sama melakukan tugas. Setiap program Go memiliki setidaknya satu fungsi, yaitu main (). Anda dapat membagi kode Anda menjadi fungsi-fungsi terpisah. Bagaimana Anda membagi kode Anda di antara fungsi-fungsi yang berbeda tergantung pada Anda, tetapi secara logis, pembagiannya harus sedemikian rupa sehingga masing-masing fungsi melakukan tugas tertentu.

