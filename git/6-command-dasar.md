# Command dasar

## Menambahkan file

- Secara default, file yang ditambahkan masih berstatus _untracked_ yang artinya belum ditrack/dilacak perubahan filenya oleh Git.
- Untuk memindahkan file ke staging area, jalankan command `git add nama-file`.
- Untuk memindahkan/commit file ke repository, jalankan command `git commit -m "pesan"`.

## Menghapus file

- File yang dihapus berstatus _deleted_.

## Memodifikasi file

- File yang dimodifikasi berstatus _modified_.
- Untuk melihat perubahan pada file, jalankan command `git diff`.
