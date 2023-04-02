# 21_ORM and Code Structure (MVC)

ORM atau kependekan dari Object Relational Mapping adalah suatu metode pemrograman yang di gunakan untuk mengkonfersi data dari Object Oriented Programming (OOP) ke lingkungan database relational. Metode ORM memudahkan kita dalam pembuatan aplikasi yang menggunakan database relasional.

Kekurangan dan Kelebihan
Menggunakan ORM menghemat banyak waktu karena:

Anda menulis model data Anda hanya di satu tempat, dan lebih mudah untuk memperbarui, memelihara, dan menggunakan kembali kode.
Banyak hal dilakukan secara otomatis, mulai dari penanganan basis data hingga I18N.
Ini memaksa Anda untuk menulis kode MVC, yang pada akhirnya membuat kode Anda sedikit lebih bersih.
Anda tidak perlu menulis SQL dengan format yang buruk (sebagian besar pemrogram Web benar-benar payah dalam hal itu, karena SQL diperlakukan seperti bahasa "sub", padahal sebenarnya bahasa itu sangat kuat dan kompleks).
sanitasi; menggunakan pernyataan atau transaksi yang disiapkan semudah memanggil metode.
Menggunakan perpustakaan ORM lebih fleksibel karena:

Ini cocok dengan cara pengkodean alami Anda (itu bahasa Anda!).
Ini mengabstraksi sistem DB, sehingga Anda dapat mengubahnya kapan pun Anda mau.
Model terikat lemah dengan aplikasi lainnya, sehingga Anda dapat mengubahnya atau menggunakannya di tempat lain.
Ini memungkinkan Anda menggunakan kebaikan OOP seperti pewarisan data tanpa pusing.
Tapi ORM bisa menyebalkan:

Anda harus mempelajarinya, dan pustaka ORM bukanlah alat yang ringan;
Anda harus mengaturnya. Permasalahan yang sama.
Performanya OK untuk kueri biasa, tetapi master SQL akan selalu melakukan yang lebih baik dengan SQL-nya sendiri untuk proyek besar.
Ini mengabstraksi DB. Meskipun tidak apa-apa jika Anda tahu apa yang terjadi di balik layar, ini adalah jebakan bagi pemrogram baru yang dapat menulis pernyataan yang sangat rakus, seperti pukulan berat dalam perulangan for.

Model View Controller atau yang dapat disingkat MVC adalah sebuah pola arsitektur dalam membuat sebuah aplikasi dengan cara memisahkan kode menjadi tiga bagian yang terdiri dari:

Model
Bagian yang bertugas untuk menyiapkan, mengatur, memanipulasi, dan mengorganisasikan data yang ada di database.

View
Bagian yang bertugas untuk menampilkan informasi dalam bentuk Graphical User Interface (GUI).

Controller
Bagian yang bertugas untuk menghubungkan serta mengatur model dan view agar dapat saling terhubung.

Manfaat dari MVC
Ada beragam manfaat ketika kamu menerapkan MVC ini dalam pembuatan aplikasi kamu. Berikut ini adalah manfaatnya.

Proses pengembangan aplikasi menjadi lebih efisien
Penggunaan MVC dapat mempercepat pengembangan aplikasi karena kode dapat dikerjakan oleh beberapa developer. Contohnya dalam kasus pengembangan aplikasi web, bagian model dan controller dapat dikerjakan oleh back-end developer sedangkan bagian view dapat dilakukan oleh front-end developer.

Penulisan kode menjadi lebih rapi
Karena dibagi menjadi tiga bagian, maka penulisan kode akan jadi lebih rapi dan memudahkan developer lain untuk mengembangkan kode tersebut.

Dapat melakukan testing dengan lebih mudah
Untuk memastikan seluruh aplikasi bekerja sesuai dengan rencana maka langkah testing atau uji coba wajib dilakukan. Dengan menggunakan model view controller ini, maka proses uji coba dapat dilakukan pada setiap bagian.

Perbaikan bug atau error lebih cepat untuk diselesaikan
Penggunaan MVC dapat memudahkan developer untuk memperbaiki error atau bug yang terjadi. Developer dapat fokus untuk menemukan dan memperbaiki masalah yang terjadi karena kode dituliskan pada bagian-bagian terpisah.

Mempermudah pemeliharaan
Konsep MVC ini dapat mempermudah pemeliharaan aplikasi, karena script atau kode yang lebih rapi dan terstruktur sehingga mempermudah developer dalam proses pemeliharaan aplikasi.
