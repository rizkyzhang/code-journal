## Previous snapshot/commit

- Git juga memiliki fitur seperti mesin waktu, dimana kita bisa kembali pada snapshot sebelumnya.
  Jalankan command `git checkout hash`. Jika ingin kembali ke HEAD terakhir sebelum menjalankan command, jalankan command `git checkout nama-branch`.

## Revert commit

- Git memiliki fitur revert commit, yaitu fitur untuk membatalkan commit yang sudah kita lakukan dengan cara membuat commit baru yang membatalkan commit sebelumnya. Untuk melakukan revert commit, jalankan command `git revert hash`.

## Git blame

- Git memiliki fitur yang bernama blame, ini digunakan untuk mencari tahu, siapa yang menambah perubahan pada file dan juga untuk mengetahui commit nya, jalankan command `git blame nama-file`.
