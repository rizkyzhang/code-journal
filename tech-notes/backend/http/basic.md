# Basic

- HTTP merupakan singkatan dari Hypertext Transfer Protocol.
- HTTP merupakan sebuah protokol yang digunakan dalam transmisi files seperti HTML, CSS, JS, gambar, audio, video dan lain-lain.
- HTTP pada awalnya digunakan sebagai perantara komunikasi antara web browser dan web server, sekarang memiliki banyak fungsi lain.
- Client akan mengirim HTTP request untuk meminta atau mengirim informasi ke server, server akan membalasnya dengan mengirim balik HTTP response.

## Stateless

- HTTP merupakan protokol yang stateless, yang artinya tiap HTTP request bersifat independen atau tidak terikat dengan HTTP request sebelumnya dan setelahnya.
- Hal ini dilakukan agar HTTP request tidak harus dilakukan secara berurut.

## HTTPS

- Secara default HTTP tidak aman, sehingga dibuatlah HTTPS sebagai HTTP dengan fitur enkripsi.
- HTTPS menggunakan SSL (Secure Sockets Layer) untuk melakukan enkripsi HTTP request dan HTTP response.

## IP

- IP merupakan singkatan dari Internet Protocol
- IP digunakan sebagai identitas komputer di jaringan

## URL

- URL merupakan singkatan dari Universal Resource Locator
- URL merupakan alamat dari sebuah resource di web

## DNS

- DNS merupakan singkatan dari Domain Name Server
- DNS merupakan server yang memetakan antara nama domain di URL dengan IP address.
- Web browser -> domain name -> DNS -> IP address
- DNS bisa dianalogikan sebagai "alias" dari IP address yang susah diingat.

## HTTP Method

- HTTP Method merupakan kategori HTTP request yang dilakukan.
- GET method digunakan untuk meminta data.
- HEAD mirip dengan GET, tetapi membutuhkan body saat melakukan request.
- POST digunakan untuk mengirim data.
- PUT digunakan untuk mengganti semua data yang terdapat di server.
- DELETE digunakan untuk menghapus data.
- PATCH digunakan untuk mengubah sebagian data yang terdapat di server.
- OPTIONS digunakan untuk mendeskripsikan opsi komunikasi yang tersedia.
- TRACE digunakan untuk debugging. TRACE akan mengembalikan semua informasi yang dikirim oleh client.

## Anatomi URL

- Scheme = protokol yang digunakan oleh Client.
- Authority terdiri dari domain name dan port. Secara default, port HTTP 80, port HTTPS 443.
- Path to resource
- Parameter (key=value). Dua parameter dipisahkan oleh &.
- Anchor = bagian tertentu di sebuah web page.

## Anatomi HTTP Header

- Host = authority pada URL
- Content-type = tipe data pada body
- User-agent = informasi seperti browser dan OS
- Accept = tipe data yang diterima oleh client
- Authorization = credential untuk autentikasi

## HTTP Status

- HTTP status merupakan kode HTTP response mengindikasikan status dari HTTP request yang diterima server.
- HTTP status terbagi menjadi 5 kategori:
  - Informational response (100-199)
  - Successful response (200-299)
  - Redirect (300-399)
  - Client error (400-499) (body tidak valid, autentikasi gagal)
  - Server error (500-599)
