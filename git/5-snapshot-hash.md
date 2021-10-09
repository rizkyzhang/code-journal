## Snapshot (Commit) & Hash

![Snapshot Diagram](https://jaimeiniesta.github.io/learn.github.com/images/snapshots.png)

- Semua perubahan file yang direkam dinamakan snapshot.
- Snapshot berisi semua perubahan dari semua file yang kita commit.
- Setiap snapshot memiliki hash/identitas unik.
- Hash merupakan checksum yang menghitung perubahan yang terjadi.
- Hash dibutuhkan untuk menjaga integritas data, sehingga snapshot yang ada tidak bisa diubah lagi.
- Contoh hash Git : 30534aeabde65981732c6c469b7583483d379b00

## HEAD

![HEAD diagram](https://static.javatpoint.com/tutorial/git/images/git-head.png)

- HEAD merupakan pointer untuk menuju ke snapshot yang paling akhir.
- Karena sangat menyulitkan jika harus menuliskan hash value, maka untuk menuju ke snapshot paling baru, kita bisa gunakan HEAD.
