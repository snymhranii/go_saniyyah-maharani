# 10_Recursive - Number Theory - Sorting - Searching
1. Recursive : Recursive function adalah sebuah function yang memanggil/mengeksekusi dirinya sendiri. Recursive function bisa dikatakan salah satu yang bisa kita gunakan untuk melakukan perulangan. Ketika menulis kode aplikasi, terkadang ada kasus dimana akan lebih mudah jika dilakukan dengan recursive function. Contoh penggunaan sederhana recursive function adalah ketika melakukan operasi factorial.

Saat membuat recursive function, kita harus memastikan function tersebut dapat berhenti. Jika tidak, maka akan terkena error stack overflow atau melebihi limit stack karena function terus memanggil dirinya sendiri. Oleh karena itu, biasanya recursive function tidak langsung memanggil dirinya sendiri tetapi bergantung pada kondisi tertentu.
 
2. Sorting : Sorting merupakan salah satu algoritma pengurutan yang dimana melakukan perbandingkan elemen array dengan elemen array dengan memiliki nilai tertinggi dalam array. Jika elemen array pembanding memiliki nilai lebih tinggi, maka elemen array tersebut akan tukar posisi dengan elemen array yang memiliki nilai lebih rendah dan elemen array yang memiliki nilai lebih rendah akan menjadi elemen pembanding sampai posisi yang tepat ditemukan. Untuk mengetahui proses pengurutan ada dibagian kesimpulan.

Pada contoh insertion sort kali ini kita mencoba untuk mengurutkan 10 bilangan yang secara acak, dengan angka yang akan diurutkan sebagai berikut: 6, 9, 21, 14, 3, 52, 107, 99, 5, 1

3. Algoritma Sequential Search adalah suatu teknik pencarian beruntun (sequential search) atau yang juga disebut dengan pencarian lurus (linear search), nilai yang dicari dibandingkan dengan nilai dari setiap elemen array data, mulai dari elemen pertama sampai nilai yang dicari ditemukan atau sampai elemen terakhir, dimana data-data tidak perlu diurutkan terlebih dahulu. Pencarian berurutan menggunakan prinsip-prinsip sebagai berikut : data yang ada dibandingkan satu per satu secara berurutan dengan yang dicari hingga data tersebut ditemukan atau tidak ditemukan.

Semisal kita ambil contoh,
Kita mempunyai 5 data didalam array [1,2,4,3,5]
Disini kita akan mencari data dengan nilainya 5.

maka akan mengecek dari awal sampai data ketemu, jadi akan di periksa satu persatu ,
dengan cara membandingkan seperti ini

apakah 5 ada di index 0 ? 1 False
apakah 5 ada di index 1 ? 2 False
apakah 5 ada di index 2 ? 4 False
apakah 5 ada di index 3 ? 3 False
apakah 5 ada di index 4 ? 5 True

Semisal datanya ada di tengah maka akan berhenti ketika kondisi menjadi true, dan tidak akan di ulangi pencarian lagi, contoh kita mencari data 5 dari array [1,2,5,3,4]

apakah 5 ada di index 0 ? 1 False
apakah 5 ada di index 1 ? 2 False
apakah 5 ada di index 2 ? 5 True
(pencarian data akan berhenti)
