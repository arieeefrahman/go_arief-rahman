### Problem 1 - Analysis

Berikut adalah kekurangan yang terdapat pada code problem 1.
- Tipe data username dan password yang di-*declare* dengan tipe data integer. Sebaiknya tipe data kedua variabel tersebut dibuat dalam bentuk string menjadi `username string` dan `password string`.
- Menyimpan username dan password dalam sebuah variabel tidak aman. Sebaiknya simpan kedua data tersebut di dalam sebuah tabel database, lalu *fetch* dan *assign* ke variabel baru di file golang.
- Penulisan nama variabel yang *lowercase*. Jika ingin menulis kata berikutnya pada penamaan variabel, sebaiknya tulis dengan huruf kapital. Function `userservice(), getallusers(), getuserbyid()` sebaiknya ditulis `userService(), getAllUsers(), getUserById()`.
- Nama variabel `user` pada `type user struct {...}` menurut saya kurang deskriptif. Sebaiknya diganti dengan `type account struct {...}` karena variabel didalamnya terdapat data-data akun saja.
- Tidak ada pointer pada saat membuat method `getAllUsers(), getUserById()`. Sebaiknya ditambahkan pointer sehingga menjadi `func (u *userService)`.
- 