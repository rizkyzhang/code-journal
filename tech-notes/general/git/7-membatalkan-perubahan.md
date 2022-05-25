# Membatalkan perubahan

## Working directory

- Membatalkan penambahan file  
  `git clean -f`

- Membatalkan modifikasi file  
  `git restore nama-file`

- Membatalkan penghapusan file  
  `git restore nama-file`

## Staging area

- Command masih sama seperti working directory, namun karena sudah berada di staging area, tidak bisa membatalkan perubahan secara langsung, harus menjalankan command `git restore --staged nama-file` untuk memindahkan file kembali ke working directory, baru command di atas bisa dijalankan.

## Repository

- Jika file sudah dicommit, maka perubahan tidak bisa dibatalkan lagi. Yang bisa dilakukan hanyalah _rollback commit_ atau _revert commit_.
