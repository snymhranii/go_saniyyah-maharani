# 18_System Design
Design system merupakan kumpulan dari komponen yang digunakan secara berulang dalam suatu dapat menjaga standar kualitas dan konsistensi dari design seperti yang telah kita bahas sebelumnya. Para designer dapat langsung mengambil asset dari komponen yang ada dalam design system ini sehingga hasil design menjadi lebih konsisten. Tak hanya itu, design system ini juga memiliki peranan penting lainnya, yaitu ia dapat membuat proses design menjadi lebih efisien karena kita dapat langsung mengambil asset yang sudah ada dan tidak perlu membuat berulang-ulang kali, dengan begitu kolaborasi tim yang memiliki banyak anggota juga dapat menjadi lebih mudah.
Dalam pembuatan design system ini, biasanya akan meliputi beberapa hal yang antara lain adalah:

Tujuan dan Nilai yang ingin di capai
Tentunya dalam membuat suatu design, kita memiliki tujuan dan ada nilai yang ingin kita capai. Jadi, dalam design system ini juga penting untuk menempatkan tujuan yang jelas agar satu tim mengetahui apa sebenarnya tujuan mereka. Dan untuk nilai-nilai yang ingin di capai juga harus diselaraskan agar pada saat pembuatannya tidak melenceng dari nilai yang ingin di capai ini.

Design Principle
Perlu juga di tentukan prinsip design agar dapat memandu tim untuk menciptakan design yang sesuai dan sama antar anggota tim.

Brand Identity dan Language
Beberapa hal yang harus dispesifikasikan di dalam sini, seperti warna, font, spasi, bentuk, icon, ilustrasi, foto, animasi, suara, dan lain sebagainya.

Component dan Patterns
Perlu juga di tentukan komponen dan juga pattern yang akan dipakai dalam design. Dengan demikian, akan terciptanya keselarasan pada saat mendesign.
Fundamental System Design
Berikut ini beberapa pertimbangan utama yang perlu diingat saat mendesain sistem.

1.    Mengidentifikasi requirements (persyaratan)
Langkah pertama dalam merancang sistem adalah mengidentifikasi kebutuhan. Ini termasuk memahami masalah yang coba dipecahkan oleh sistem, pengguna yang akan berinteraksi dengan sistem, dan kendala di mana sistem harus beroperasi. Sangat penting untuk memiliki pemahaman yang jelas tentang masalah untuk memastikan bahwa sistem dirancang untuk mengatasinya secara efektif.

2.    Skalabilitas
Skalabilitas adalah salah satu pertimbangan terpenting dalam desain sistem. Sistem yang dapat menangani peningkatan jumlah pengguna dan transaksi tanpa penurunan kinerja yang signifikan dianggap dapat diskalakan. Sangat penting untuk merancang sistem yang dapat menangani beban yang meningkat tanpa memerlukan perubahan signifikan pada arsitektur yang mendasarinya.

3.    High availability 
High availability mengacu pada kemampuan sistem untuk provide uninterrupted service kepada penggunanya bahkan jika terjadi kegagalan. Ini sangat penting dalam sistem yang perlu tersedia 24/7, seperti pengecer online atau sistem keuangan. Untuk mencapai high availability, penting untuk merancang sistem yang dapat menangani kegagalan pada tingkat yang berbeda, seperti kegagalan perangkat keras, kegagalan jaringan, dan bug perangkat lunak.

4.    Kinerja
Kinerja adalah pertimbangan penting lainnya dalam desain sistem. Suatu sistem yang dapat menangani jumlah transaksi atau permintaan yang tinggi per detik dianggap memiliki kinerja yang baik. Penting untuk merancang sistem yang dapat menangani beban yang diharapkan secara tepat waktu untuk memastikan pengalaman pengguna yang baik.

5.    Keamanan
Keamanan adalah aspek penting dari desain sistem. Sistem harus dirancang untuk melindungi terhadap akses tidak sah dan pelanggaran data. Ini termasuk menerapkan protokol komunikasi yang aman, kontrol akses, dan enkripsi. Sangat penting untuk merancang sistem yang dapat mencegah akses tidak sah dan melindungi data sensitif.
 

Keuntungannya Menggunakan System Design
1.    Dapat merencanakan desain dengan konsisten
Komponen standar yang digunakan secara konsisten dan berulang kali membuat aplikasi yang lebih dapat diprediksi dan lebih mudah dipahami. Komponen memungkinkan desainer untuk tidak terlalu membuang waktu pada gaya dan lebih banyak waktu untuk mengembangkan pengalaman pengguna yang lebih baik.

2.    Membuat Prototipe dengan lebih cepat
Bekerja dengan sistem desain memungkinkan kita merancang aplikasi dengan cepat seperti membuat mainan Lego karena kita tidak perlu mendesain dari awal, cukup drag and drop.

3.    Kolaboratif dan berbagi informasi
Tingkatkan kerja sama dengan tim dan bangun semangat kerja sama yang baik ketika sistem perencanaan dilakukan bersama. jadi bekerja sama dengan desainer lain untuk membuat dan mengembangkan sistem desain itu sendiri.
 

Komponen System Design
1.    Tempat khusus untuk dokumen
Komponen utama tentu saja adalah tempat digital khusus. Perusahaan biasanya menggunakan ruang penyimpanan online seperti Google Drive atau Dropbox. Beberapa perusahaan membuka akses untuk umum. Namun, ada juga yang hanya memberikan akses kepada karyawan produksinya.

2.    Komponen desain yang dapat digunakan kembali
Komponen penting berikutnya dari sistem desain adalah karya desain itu sendiri, misalnya dokumen desain berupa file Adobe XD, figma, gambar, dan lain-lain. Elemen-elemen ini nantinya dapat diambil alih dan dimodifikasi sesuai kebutuhan saat mendesain produk.  

3.    Komponen development yang bisa digunakan kembali 
Selain elemen desain, desain sistem juga memiliki elemen pemrograman. Misalnya kode yang berbeda dalam bahasa yang berbeda. Kode juga dapat berbeda untuk produk aplikasi, produk situs web, atau produk lainnya.  
 

Types System Design
1.    Logical design
Desain logis berkaitan dengan representasi abstrak dari aliran data, input, dan output dari sistem. Ini menggambarkan input (sumber), output (tujuan), database (penyimpanan data), prosedur (aliran data) semuanya dalam format yang memenuhi persyaratan pengguna. Sambil menyiapkan desain logis dari suatu sistem, analis sistem menentukan kebutuhan pengguna pada tingkat detail yang secara virtual menentukan arus informasi masuk dan keluar dari sistem dan sumber data yang diperlukan. Diagram aliran data, pemodelan diagram E-R digunakan.

2.    Physical design
Desain fisik berkaitan dengan proses input dan output aktual dari sistem. Ini berfokus pada bagaimana data dimasukkan ke dalam sistem, diverifikasi, diproses, dan ditampilkan sebagai output. Ini menghasilkan sistem kerja dengan mendefinisikan spesifikasi desain yang menentukan dengan tepat apa yang dilakukan sistem kandidat. Ini berkaitan dengan desain antarmuka pengguna, desain proses, dan desain data.

3.    Architectural design
Ini juga dikenal sebagai desain tingkat tinggi yang berfokus pada desain arsitektur sistem. Ini menggambarkan struktur dan perilaku sistem. Ini mendefinisikan struktur dan hubungan antara berbagai modul proses pengembangan sistem.

4.    Detail design
Ini mengikuti desain Arsitektur dan berfokus pada pengembangan setiap modul.
