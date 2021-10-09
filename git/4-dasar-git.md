# Dasar Git

## Repository

- Repository merupakan istilah projek di Git.
- Repository dapat dibuat dari folder kosong atau folder yang sudah berisi file.
- Selain membuat repository sendiri, kita dapat menggunakan repository yang sudah ada di server Git dengan cara clone dari server Git.

## Cara membuat repository

- Untuk membuat repository, kita harus masuk ke folder yang dipilih, lalu jalan command `git init`.
- Setelah menjalankan command, akan muncul folder .git di folder tersebut. Folder tersebut berisi database Git.

## The three states

- Git memiliki tiga states/fase yaitu: modified, staged dan committed.
- Modified artinya kita telah memodifikasi (mengubah, menghapus, mengedit) file, namun belum disimpan secara permanen di repository.
- Staged artinya file yang telah kita modifikasi akan disimpan secara permanen di repository.
- Committed artinya file telah disimpan secara permanen di repository.

## The three section

- Ketiga fase tersebut masing-masing terjadi di tempat yang berbeda yaitu: working directory, staging area dan repository.
- Working directory merupakan tempat terjadinya fase modified. Jika kita menjalankan command `git add`, maka file akan masuk ke staging area.
- Staging area merupakan tempat terjadinya fase staged. Jika kita menjalankan command `git commit`, maka file akan masuk ke repository.
- Repository merupakan tempat terjadinya fase committed.

## The three tree

![Three tree diagram](https://miro.medium.com/max/686/1*diRLm1S5hkVoh5qeArND0Q.png)
