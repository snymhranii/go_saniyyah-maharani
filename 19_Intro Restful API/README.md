# 19_Intro Restful API
API RESTful adalah antarmuka yang digunakan oleh dua sistem komputer untuk bertukar informasi secara aman melalui internet. Sebagian besar aplikasi bisnis harus berkomunikasi dengan aplikasi internal dan pihak ketiga lainnya untuk melakukan berbagai tugas. Misalnya, untuk menghasilkan slip gaji bulanan, sistem akun internal Anda harus berbagi data dengan sistem perbankan pelanggan Anda untuk mengotomatiskan tagihan dan berkomunikasi dengan aplikasi absensi internal. API RESTful mendukung pertukaran informasi ini karena mengikuti standar komunikasi perangkat lunak yang aman, andal, dan efisien.

Antarmuka Program Aplikasi (API) menentukan aturan yang harus Anda ikuti untuk berkomunikasi dengan sistem perangkat lunak lain. Developer mengekspos dan membuat API sehingga aplikasi lain dapat berkomunikasi dengan aplikasinya secara terprogram. Contohnya, aplikasi absensi mengekspos API yang meminta nama lengkap pegawai dan rentang tanggal. Saat menerima informasi ini, API memproses absensi pegawai secara internal dan mengembalikan jumlah jam kerja dalam rentang tanggal tersebut.

Anda dapat menganggap sebuah API web sebagai gateway antara klien dan sumber daya pada web.

Klien
Klien adalah pengguna yang ingin mengakses informasi dari web. Klien dapat berupa orang atau sistem perangkat lunak yang menggunakan API. Contohnya, developer dapat menulis program yang mengakses data cuaca dari sistem data. Atau Anda dapat mengakses data yang sama dari peramban Anda saat mengunjungi situs web cuaca secara langsung.

Sumber daya
Sumber daya adalah informasi yang disediakan aplikasi yang berbeda untuk klien mereka. Sumber daya dapat berupa citra, video, teks, angka, atau tipe data apa pun. Mesin yang memberikan sumber daya kepada klien disebut juga sebagai server. Organisasi menggunakan API untuk berbagi sumber daya dan menyediakan layanan web sembari memelihara keamanan, kendali, dan autentikasi. Selain itu, API membantu menentukan klien yang mendapatkan akses ke sumber daya internal tertentu.

Representational State Transfer (REST) adalah arsitektur perangkat lunak yang memberlakukan syarat mengenai cara API bekerja. REST pada awalnya dibuat sebagai panduan untuk mengelola komunikasi pada jaringan kompleks seperti internet. Anda dapat menggunakan arsitektur berbasis REST untuk mendukung komunkasi berperforma tinggi dan andal sesuai skala. Anda dapat dengan mudah menerapkan dan memodifikasinya, membawa visibilitas dan portabilitas lintas platform ke semua sistem API.

Developer API dapat merancang API menggunakan beberapa arsitektur yang berbeda. API yang mengikuti gaya arsitektur REST disebut sebagai API REST. Layanan web yang menerapkan arsitektur REST disebut sebagai layanan web RESTful. Istilah API RESTful umumnya merujuk pada API web RESTful. Namun, Anda dapat menggunakan istilah API REST dan API RESTful secara bergantian.

Berikut adalah beberapa prinsip gaya arsitektur REST:

Antarmuka seragam
Antarmuka seragam adalah desain fundamental dari semua layanan web RESTful. Ini mengindikasikan bahwa server mentransfer informasi dalam format standar. Sumber daya terformat disebut representasi dalam REST. Format ini dapat berbeda dari representasi internal sumber daya pada aplikasi server. Contohnya, server dapat menyimpan data sebagai teks tetapi mengirimkannya dalam format representasi HTML.

Antarmuka seragam memberlakukan empat hambatan arsitektur:

Permintaan harus mengidentifikasi sumber daya. Permintaan melakukannya dengan menggunakan pengidentifikasi sumber daya seragam.
Klien memiliki cukup informasi dalam representasi sumber daya untuk memodifikasi atau menghapus sumber daya jika diinginkan. Server memenuhi syarat ini dengan mengirimkan metadata yang menjelaskan sumber daya lebih lanjut.
Klien menerima informasi mengenai cara untuk memproses representasi lebih lanjut. Server mencapainya dengan mengirimkan pesan deskriptif mandiri yang berisi metadata mengenai cara klien dapat menggunakannya sebaik mungkin.
Klien menerima informasi mengenai semua sumber daya terkait lainnya yang dibutuhkan untuk menyelesaikan tugas. Server mencapainya dengan mengirim hyperlink di dalam representasi sehingga klien dapat menemukan lebih banyak sumber daya secara dinamis.
Statelessness
Dalam arsitektur REST, statelessness mengacu pada metode komunikasi tempat server menyelesaikan setiap permintaan klien secara independen dari semua permintaan sebelumnya. Klien dapat meminta sumber daya dalam perintah apa pun, dan setiap permintaan berada dalam status stateless atau terisolasi dari permintaan lain. Hambatan rancangan API REST ini menyiratkan bahwa server dapat sepenuhnya memahami dan memenuhi permintaan setiap waktu. 

Sistem berlapis
Dalam arsitektur sistem berlapis, klien dapat terhubung ke perantara resmi lain antara klien dan server, dan masih akan menerima respons dari server tersebut. Server juga dapat meneruskan permintaan ke server lain. Anda dapat merancang layanan web RESTful untuk berjalan pada beberapa server dengan beberapa lapisan seperti keamanan, aplikasi, dan logika bisnis, bekerja sama untuk memenuhi permintaan klien. Lapisan ini tetap terlihat oleh klien.

Ketersinggahan (Cacheability)
Layanan web RESTful mendukung pembuatan cache, proses penyimpanan beberapa respons pada klien atau pada perantara untuk meningkatkan waktu respons server. Contohnya, bayangkan Anda mengunjungi situs web yang memiliki citra header dan footer umum pada setiap halaman. Setiap kali Anda mengunjungi halaman situs web baru, server harus mengirim ulang citra yang sama. Untuk menghindari hal ini, klien membuat cache atau menyimpan citra setelah respons pertama lalu menggunakan citra ini langsung dari cache tersebut. Layanan web RESTful mengontrol pembuatan cache dengan menggunakan respons API yang menentukannya apakah dapat dibuat cache atau tidak.

Kode sesuai permintaan
Dalam gaya arsitektur REST, server dapat sementara memperluas atau menyesuaikan fungsionalitas klien dengan mentransfer kode pemrograman perangkat lunak ke klien. Contohnya, saat Anda mengisi formulir pendaftaran pada situs web, peramban Anda dengan segera menyoroti kesalahan yang Anda buat, seperti nomor telepon yang salah. Ini dapat terjadi karena kode yang dikirimkan oleh server.

API RESTful memiliki manfaat berikut:

Skalabilitas
Sistem yang menerapkan API REST dapat menskalakan secara efisien karena REST mengoptimalkan interaksi klien-server. Statelessness menghapus beban server karena server tidak perlu mempertahankan informasi permintaan klien di masa lalu. Pembuatan cache yang dikelola dengan baik secara parsial atau keseluruhan menghilangkan beberapa interaksi klien-server. Semua fitur ini mendukung skalabilitas tanpa menyebabkan kemacetan komunikasi yang mengurangi performa.

Fleksibilitas
Layanan web RESTful mendukung pemisahan total klien-server. Layanan web RESTful menyederhanakan dan memisahkan berbagai komponen server sehingga masing-masing bagian dapat berkembang secara mandiri. Perubahan platform atau teknologi pada aplikasi server tidak memengaruhi aplikasi klien. Kemampuan untuk melapisi fungsi aplikasi semakin meningkatkan fleksibilitas lebih jauh. Contohnya, developer dapat membuat perubahan pada lapisan basis data tanpa menulis ulang logika aplikasi.

Independensi
API REST independen terhadap teknologi yang digunakan. Anda dapat menulis baik aplikasi klien dan server dalam berbagai bahasa pemrograman tanpa memengaruhi desain API. Anda juga dapat mengubah teknologi mendasar di kedua sisi tanpa memengaruhi komunikasi.

Bagaimana cara kerja API RESTful?
Fungsi dasar API RESTful sama dengan penjelajahan internet. Klien menghubungi server dengan menggunakan API saat meminta sumber daya. Developer API menjelaskan kepada klien cara untuk menggunakan API REST dalam dokumentasi API aplikasi server. Ini adalah langkah umum untuk semua panggilan API REST:

Klien mengirimkan permintaan ke server. Klien mengikuti dokumentasi API untuk memformat permintaan dalam format yang dipahami oleh server.
Server mengautentikasi klien dan mengonfirmasi bahwa klien memiliki hak untuk membuat permintaan.
Server menerima permintaan dan memproses secara internal.
Server mengembalikan respons kepada klien. Respons berisi informasi yang memberitahu klien jika permintaannya berhasil. Respons juga termasuk informasi apa saja yang diminta klien.
Permintaan API REST dan detail respons sedikit berbeda tergantung pada cara developer API merancang API.
