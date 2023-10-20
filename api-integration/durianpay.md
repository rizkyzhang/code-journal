## Durianpay API Integration

- Daftar akun durianpay
- Copy API key Sandbox, masuk ke postman, ke tab authorization, pilih Basic Auth, masukkan API key ke username dan passwordnya kosong. Lakukan request ke endpoint durianpay manapun, ke tab Headers, cari key Authorization dengan value `Basic xxxxxxxxxx`, Basic auth ini akan digunakan untuk melakukan request, untuk production sama juga caranya.
- Menghubungi pihak durianpay untuk membuat grup support whatsapp dan memberikan data diri atau perusahaan untuk mempercepat proses untuk mendapatkan API Key production
- Buat endpoint create order `router.post('api/payments/gateway/durianpay/orders',gateway.createOrder)`
- Buat endpoint get order by id untuk mendapatkan detail dari sebuah order
  `router.get('api/payments/gateway/durianpay/orders/:orderId', checkUser, gateway.getOrderById)`
- Buat kolom durianpayOrderId dan durianpayOrderToken di database
- Masukkan script `<script type="text/javascript" src="https://js.durianpay.id/0.1.23/durianpay.min.js"></script>` di dalam head element
- Buat objek payload untuk post request create order
- Cek apakah order sudah memiliki durianpayOrderId
  - Jika sudah, maka get request ke endpoint get order by id untuk mendapatkan access token
  - Jika belum, maka post request ke endpoint create order, lalu simpan durianpayOrderId dan durianpayOrderToken di database
- Buat objek init config untuk init durianpay, didalamnya juga berisi access token, gunakan access token yang telah disimpan di database atau dari objek payment yang dikembalikan oleh create order
- Lakukan init durianpay

```js
const dpay = window.Durianpay.init(initConfig);
dpay.checkout();
```

- Buat webhook untuk event payment completed, selain cara ini bisa juga menggunakan method onsucces di initConfig `router.post('/durianpay/webhook/payment', gateway.postDurianpayWebhookPayment)`. Di dalam sini kita akan verifikasi request untuk mengecek apakah benar dari pihak durianpay dengan membandingkan verfication_signature, apabila sesuai ubah order menjadi paid.
