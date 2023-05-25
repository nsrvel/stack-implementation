![Example Image](https://example.com/image.jpg){width=300px}

Dalam contoh di atas, kami membuat kelas `Stack` yang menggunakan list `data` untuk menyimpan elemen-elemen stack. Terdapat beberapa metode yang diimplementasikan pada kelas `Stack`:

-   `push`: Digunakan untuk menambahkan elemen ke stack. Elemen baru ditambahkan ke akhir list `data`.
-   `pop`: Menghapus dan mengembalikan elemen teratas dari stack. Jika stack kosong, akan memunculkan pengecualian (exception).
-   `peek`: Mengembalikan elemen teratas dari stack tanpa menghapusnya. Jika stack kosong, akan memunculkan pengecualian (exception).
-   `is_empty`: Memeriksa apakah stack kosong atau tidak.
-   `size`: Mengembalikan jumlah elemen dalam stack.

Di dalam contoh penggunaan, kami menguji operasi stack dengan melakukan beberapa operasi sebagai berikut:

1. Menambahkan beberapa elemen ke stack menggunakan metode `push`.
2. Menghapus elemen teratas dari stack menggunakan metode `pop`.
3. Melihat elemen teratas dari stack menggunakan metode `peek`.
4. Memeriksa apakah stack kosong menggunakan metode `is_empty`.
