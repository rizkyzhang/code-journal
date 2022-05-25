# Commit

## Commit log

- Untuk melihat riwayat/history commit, jalankan command `git log`.
- Agar setiap commit ditampilkan secara singkat dalam satu baris, jalankan command `git log --oneline`.
- Untuk melihat riwayat/history commit beserta hubungan antara commit (branch), jalankan command `git log --oneline --graph`.
- Untuk melihat detail suatu commit secara detail, jalankan command `git log hash`.

## Compare commit

- Untuk membandingkan satu commit dengan commit lainnya, jalankan command `git diff older-hash newer-hash`.
- Jika kita ingin menggunakan visual studio code untuk melihat perbedaan antar commit, kita bisa gunakan perintah `git diff older-hash newer-hash`.

## Rename file

- Sebenarnya rename file merupakan operasi gabungan antara hapus file, lalu menambah file baru dengan isi yang sama.
- Git bisa otomatis mendeteksi jika terjadi perubahan nama file, karena isi file sebagian besar masih sama.

## Reset commit

- Reset commit merupakan mekanisme dimana kita menggeser HEAD pointer ke posisi commit yang kita mau, artinya commit selanjutnya akan dilakukan pada posisi HEAD baru
- Untuk melakukan reset commit, kita bisa gunakan perintah `git reset <mode> hash`
- Ada beberapa mode pengaturan melakukan reset commit:
  - soft, memindahkan HEAD pointer, namun tidak melakukan perubahan apapun di staging area dan working directory.
  - mixed (default), memindahkan HEAD pointer, mengubah staging area menjadi sama seperti dengan repository, namun tidak mengubah apapun di working directory.
  - hard, memindahkan HEAD pointer, dan mengubah staging area dan working directory sehingga sama dengan repository. Ini **mode yang paling berbahaya** karena semua commit yang telah direset akan hilang dan tidak bisa dikembalikan seperti dua mode sebelumnya.
- Jika kita melakukan reset, namun kita belum membuat commit baru, kita masih bisa kembali maju lagi ke commit yang paling baru. Namun jika kita membuat commit baru, secara otomatis commit lama akan ditimpa oleh commit baru.

## Ammend commit

- Kadang saat sudah melakukan commit, mungkin ada beberapa hal yang terlupakan. Biasanya kita akan lakukan reset soft ke commit sebelumnya, lalu tambahkan perubahan yang terlupakan, lalu kita lakukan commit ulang. Hal tersebut bisa dilakukan tanpa manual melakukan reset, jalankan command `git commit --amend`. Perlu diingat, amend akan mengubah hash commit karena data perubahan yang dicommit bertambah.
