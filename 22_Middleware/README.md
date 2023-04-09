# 22_Middleware
Middleware adalah perangkat lunak komputer yang memberikan layanan untuk menghubungkan bagian-bagian berbeda dari sebuah aplikasi dengan sistem operasi. Middleware umumnya digunakan dalan sistem terdistribusi untuk memudahkan pengembang perangkat lunak dalam melakukan komunikasi input/output.

Middleware berfungsi sebagai penerjemah informasi sehingga setiap aplikasi mendapatkan format data yang dapat mereka proses. Tujuan utama software middleware adalah untuk membantu memecahkan interkoneksi beberapa aplikasi dan masalah interoperabilitas.
Pada umumnya, middleware adalah authentication, untuk mengecek apakah seseorang yang mengakses suatu web sudah log-in dan memiliki hak akses atau belum. Web hanya akan bisa berjalan diatas Web Server. Contohnya pada laravel, authentication terletak di middleware. Tujuan utama dari middleware adalah agar controller fokus dalam menyelesaikan logika alur bisnis dari suatu flow aplikasi tanpa harus untuk melakukan hal-hal di luar itu seperti validasi input untuk setiap flow. Jadi ketika data yang masuk ke controller sudah dalam keadaan bersih dan sudah siap untuk diproses sesuai dengan alur bisnis dari aplikasi tersebut, dengan adanya middleware, kode-kode yang telah dibuat akan lebih reusable, maintainable dan readable.

Apa fungsi middleware laravel?
Middleware menyediakan mekanisme penyaringan HTTP request yang masuk ke aplikasi anda atau dengan kata lain setiap kali ada request yang masuk maka akan difilter oleh Middleware. Sebagai contoh, Laravel telah menyediakan sebuah middleware untuk menverifikasi setiap user yang melakukan otentikasi pada aplikasi anda.

Middleware khususnya di dunia web programming dapat diartikan sebagai suatu software layer yang berada di antara router dengan controller. Karena posisi dari middleware berada di antara router dengan controller, fungsi dari middleware rata-rata memiliki fungsi generik.
Middleware dapat digunakan untuk beberapa tujuan seperti:

Memberikan fasilitas untuk programmer supaya dapat mendistribusikan obyek akan dipakai pada beberapa bagian proses yang berbeda.
Sebagai interkoneksi ke beberapa aplikasi dan masalah interoperabilitas. Middleware menjadi sangat dibutuhkan untuk bermigrasi dari aplikasi mainframe ke aplikasi client/server dan juga untuk menyediakan komunikasi antar platform yang berbeda.
Middleware dapar memberikan manfaat tertentu seperti penjelasan dibawah ini:

Pada sistem yang terdistribusi maka dapat dijalankan 2 buah platform atau aplikasi secara bersamaan.
Dapat melakukan komunikasi pada aplikasi yang berjalan di platform berbeda.
Adanya transparansi pada seluruh jaringan sehingga dapat menyediakan interaksi dengan layanan atau aplikasi lainnya.

 
