# 16_Join - Union - Agregasi - Subquery - Function (DBMS)
union : merupakan metode overlay dimana apabila batas luar data grafis yang akan dilakukan tumpang susun tidak sama, maka batas luar yang baru adalah gabungan antara kedua data grafis tersebut
Fungsi UNION pada SQL digunakan untuk menggabungkan dua tabel dalam bentuk baris baru ke bawah dimana field yang di-SELECT antara tabel satu dan lainnya adalah harus sama. Atau sederhananya yaitu untuk menempatkan baris dari kueri satu sama lain dan nilainya distinct/unik.

JOIN merupakan salah satu operasi dalam SQL yang digunakan untuk menggabungkan dua atau lebih tabel yang berbeda menjadi satu basis data. Operasi join ini dapat dilakukan jika tabel-tabel yang akan digabungkan itu memiliki key kolom yang sama.

Agregasi data adalah jenis proses penambangan data dan informasi di mana data dicari, dikumpulkan, dan disajikan dalam format yang dirangkum berdasarkan laporan untuk mencapai tujuan atau proses bisnis tertentu dan / atau melakukan analisis manusia.

Pengumpulan data dapat dilakukan secara manual atau melalui perangkat lunak khusus.
Agregasi data adalah komponen dari solusi intelijen bisnis (BI). Petugas agregasi data atau basis data pencarian perangkat lunak menemukan data permintaan pencarian yang relevan dan menyajikan temuan data dalam format ringkasan yang bermakna dan berguna bagi pengguna akhir atau aplikasi.

Subquery atau Inner query atau Nested query adalah query dalam query SQL lain dan tertanam dalam klausa WHERE. Sebuah subquery digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil. Subqueries dapat digunakan dengan SELECT, INSERT, UPDATE, dan DELETE statements bersama dengan operator seperti =, <,>,> =, <=, IN, BETWEEN dll.

Ada beberapa aturan yang subqueries harus mengikuti:

Subqueries harus tertutup dalam tanda kurung.
Sebuah subquery hanya dapat memiliki satu kolom pada klausa SELECT, kecuali beberapa kolom yang di query utama untuk subquery untuk membandingkan kolom yang dipilih.
ORDER BY tidak dapat digunakan dalam subquery, meskipun permintaan utama dapat menggunakan ORDER BY.GROUP BY dapat digunakan untuk melakukan fungsi yang sama seperti ORDER BY dalam subquery.
Subqueries yang kembali lebih dari satu baris hanya dapat digunakan dengan beberapa value operator, seperti operator IN.
Daftar SELECT tidak bisa menyertakan referensi ke nilai-nilai yang mengevaluasi ke BLOB, ARRAY, CLOB, atau NCLOB.
Sebuah subquery tidak dapat segera tertutup dalam fungsi set.
Operator BETWEEN tidak dapat digunakan dengan subquery;Namun, BETWEEN dapat digunakan dalam subquery.