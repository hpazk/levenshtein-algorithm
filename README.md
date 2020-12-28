# Task 02 - Talent Factory Batch 1

## Intro

### Algoritma Levenshtein Distance

Algoritma Levenshtein distance atau algoritma Edit distance adalah algoritma yang mengukur kesamaan dua string, yaitu membandingkan text-target (t) dengan text-source (s) untuk mengetahui berapa nilai perbedaan dari kedua string tersebut.\
\
Secara matematika dapat didefinisikan sebagai berikut: \
\
`Lev t, s = (|t|, |s|)`\
\
dimana t adalah text-target dan s adalah source-text\
\
> Adapun yang dimaksud dengan edit distance adalah pengeditan pada karakter tunggal pada kata yang dibandingkan seperti insertion (penyisipan), deletion (penghapusan), atau subtitution (penggantian)

#### Contoh

sebagai contoh perhitungan kata "computer" dengan "komputer"

* t = komputer
* s = computer

maka `Levenshtein(t, s) = 1`

### Aplikasi

Algoritma Levenshtein distance diantaranya dapat diaplikasikan untuk:

* String Matching
* Spelling Checking

## Studi Kasus

Mencari judul buku klasik berdasarkan keyword yang dilansir dari halaman: [https://www.penguin.co.uk/articles/2018/100-must-read-classic-books.html](https://www.penguin.co.uk/articles/2018/100-must-read-classic-books.html) dengan metode Fuzzy Search

### Langkah-Langkah

1. Seeding data dengan cara scraping judul buku dari halaman web: [https://www.penguin.co.uk/articles/2018/100-must-read-classic-books.html](https://www.penguin.co.uk/articles/2018/100-must-read-classic-books.html) menggunakan goquery, kemudian disimpan pada sebuah variable dalam bentuk array

2. Masukan keyword yang ingin dicari

3. Melakukan String Matching berdasarkan tingkat kemiripan antara keyword yang ingin dicari dengan kata pada list judul buku

4. Mengembalikan hasil dengan kondisi

    * Jika keyword yang dimasukan tidak ada yang sesuai dengan source string dari list judul buku, maka tampilkan pesan *"Book not found, search instead for (`keyword`)*

    * Jika keyword yang dimasukan memiliki kemiripan dengan source string dari list judul buku, maka tampilkan pesan *Did you mean: (`levenshtein(keyword, source`))*
    
    * Jika ada error maka tampilkan pesan error