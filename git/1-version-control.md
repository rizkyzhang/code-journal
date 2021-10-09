# Version Control

- Sebelum version control diciptakan, cara untuk backup file seperti word cukup sederhana. Jika ada revisi, kita akan simpan di file baru dengan nama document-2, lalu jika ada revisi lagi kita akan simpan dengan nama document-3, dan begitu seterusnya. Tujuannya adalah agar kita bisa membandingkan dan mengetahui perubahan file lama dengan file revisi serta dapat dengan mudah kembali ke versi lama jika diperlukan.

- Version control = sistem yang merekam perubahan pada file dari waktu ke waktu (mirip dengan fitur save point dalam game).

- Dengan version control, jika seorang programmer melakukan kesalahan dalam kodenya, dia dapat dengan mudah kembali ke versi sebelumnya.

- Version control dapat digunakan untuk berbagai file.

- Version control terbagi menjadi 3:

1. Local Version Control
2. Centralized Version Control
3. Distributed Version Control

## Local Version Control

![Local Version Control Diagram](https://www.researchgate.net/profile/Amith-Raj-Megalapete-Puttaraju-2/publication/341667446/figure/fig1/AS:895632688562177@1590546881514/Local-version-control.ppm)

- Local Version Control = version control yang menyimpan riwayat revisi data di local computer (offline).

- Karena offline, kelemahan dari Local Version Control adalah jika komputer rusak, maka semua data revisi akan hilang. Selain itu akan susah untuk berkolaborasi secara online karena file tidak bisa diakses secara online.

## Centralized Version Control

- Centralized Version Control = version control yang menyimpan riwayat revisi data di server (online).

- Kelemahannya adalah jika sedang offline/server down, pengguna tidak bisa melakukan perubahan dan melihat riwayat revisi data.

- Contoh Centralized Version Control adalah Subversion yang cukup terkenal 10 tahun yang lalu.

![Centralized Version Control Diagram](https://git-scm.com/book/en/v2/images/centralized.png)

## Disributed Version Control

- Fitur dari Disributed Version Control adalah pengguna dapat menyimpan seluruh riwayat revisi data di local computer sehingga jika server down pengguna masih dapat melakukan perubahan dan melihat riwayat revisi data.

- Contoh Disributed Version Control adalah Git.

![Disributed Version Control Diagram](https://reader021.docslide.net/reader021/html5/20170731/55cf8f32550346703b99d8ca/bg4.png)
