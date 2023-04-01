# 20_Intro Echo Golang
Echo adalah framework bahasa golang untuk pengembangan aplikasi web. Framework ini cukup terkenal di komunitas. Echo merupakan framework besar, di dalamnya terdapat banyak sekali dependensi.

Salah satu dependensi yang ada di dalamnya adalah router, dan pada chapter ini kita akan mempelajarinya.

Dari banyak routing library yang sudah penulis gunakan, hampir seluruhnya mempunyai kemiripan dalam hal penggunaannya, cukup panggil fungsi/method yang dipilih (biasanya namanya sama dengan HTTP Method), lalu sisipkan rute pada parameter pertama dan handler pada parameter kedua.

Berikut contoh sederhana penggunaan echo framework.

r := echo.New()
r.GET("/", handler)
r.Start(":9000")
Sebuah objek router r dicetak lewat echo.New(). Lalu lewat objek router tersebut, dilakukan registrasi rute untuk / dengan method GET dan handler adalah closure handler. Terakhir, dari objek router di-start-lah sebuah web server pada port 9000.

Echo router mengadopsi konsep radix tree, membuat performa lookup nya begitu cepat. Tak juga itu, pemanfaatan sync pool membuat penggunaan memory lebih hemat, dan aman dari GC overhead.

selain ctx.String() ada banyak method sejenis lainnya, berikut selengkapnya.

• Method .String()
Digunakan untuk render plain text sebagai output (isi response header Content-Type adalah text/plain). Method ini tugasnya sama dengan method .Write() milik objek http.ResponseWriter.

r.GET("/index", func(ctx echo.Context) error {
    data := "Hello from /index"
    return ctx.String(http.StatusOK, data)
})
• Method .HTML()
Digunakan untuk render html sebagai output. Isi response header Content-Type adalah text/html.

r.GET("/html", func(ctx echo.Context) error {
    data := "Hello from /html"
    return ctx.HTML(http.StatusOK, data)
})
• Method .Redirect()
Digunakan untuk redirect, pengganti http.Redirect().

r.GET("/index", func(ctx echo.Context) error {
    return ctx.Redirect(http.StatusTemporaryRedirect, "/")
})
• Method .JSON()
Digunakan untuk render data JSON sebagai output. Isi response header Content-Type adalah application/json.

r.GET("/json", func(ctx echo.Context) error {
    data := M{"Message": "Hello", "Counter": 2}
    return ctx.JSON(http.StatusOK, data)
})

 
