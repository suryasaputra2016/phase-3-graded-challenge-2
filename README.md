# Graded Challenge 2 - P3

Graded Challenge ini dibuat guna mengevaluasi pembelajaran pada Hacktiv8 Program Fulltime Golang khususnya pada pembelajaran grpc dan implementasi auth dengan Tema :

`REST API menggunakan gRPC dengan sistem autentikasi untuk Sistem Manajemen Tugas.`

## Assignment Objective
Graded Challenge 2 ini dibuat guna mengevaluasi pemahaman gRPC sebagai berikut:
- Mampu memahami konsep gRPC
- Mampu membuat service dengan gRPC
- Mampu melakukan pengujian dengan unit testing
- Mampu melakukan deployment ke Google Cloud dengan menggunakan docker image

## Assignment Directions

 - Buat proto file untuk definisi layanan gRPC dengan pesan (message) yang sesuai.
 - Implementasikan server gRPC untuk layanan yang dibuat dalam Golang.
 - Tambahkan auth menggunakan token JWT.
 - Dokumentasikan keseluruhan service dengan Swagger
 - Lakukan pengujian menggunakan unit testing minimal 3 function
 - Setelah aplikasi running maka lakukan deployment dengan Google Cloud menggunakan docker image

## Database Schema:
Database bebas (sql or no sql)
### Tabel Users:
 - ID (UUID): Identifikasi unik pengguna.
 - Username (String): Nama pengguna.
 - Password (String): Kata sandi yang di-hash.

### Tabel Tasks:

 - ID (UUID): Identifikasi unik tugas.
 - Title (String): Judul tugas.
 - Description (Text): Deskripsi tugas.
 - DueDate (Timestamp): Tanggal jatuh tempo tugas.
 - UserID (UUID): ID pengguna yang membuat tugas.


 #### Sebagai tambahan dari requirement yang sudah diberikan sebelumnya, Student juga diharapkan untuk memahami dan menerapkan konsep-konsep berikut:
- Cloud Deployment using GCP
Student diharapkan untuk mengimplementasikan Cloud Deployment menggunakan Google Cloud Platform (GCP).
Pastikan aplikasi Anda dapat diakses secara publik setelah deployment.
Sediakan dokumentasi sederhana mengenai langkah-langkah deployment yang Anda lakukan.
- Job Scheduling
Implementasikan konsep job scheduling untuk beberapa proses yang memerlukannya, seperti proses pembaharuan data atau pembersihan data yang tidak diperlukan.
- Unit Test
Buat unit test untuk memastikan bahwa setiap fungsi atau method dalam aplikasi Anda bekerja dengan semestinya.
- Docker
Kontainerisasi aplikasi Anda menggunakan Docker.
Pastikan Anda menyediakan Dockerfile dan dokumentasi singkat tentang bagaimana menjalankan aplikasi Anda menggunakan Docker.


## Database Requirements

![Alt text](db.png)

Buatlah database dengan skema mengikuti sample struct yang ada pada gambar diatas!


## Expected Result 
### Image hanya sample
 - File Proto.
   <img width="842" alt="image" src="https://github.com/H8-FTGO-P3/FTGO-P3-V1-GC2/assets/134502566/0dc710bc-a417-406d-98b3-20e329ae6404">

- Endpoint gRPC menggunakan url Google Cloud, contoh : http://url-google-cloud.com/
  <img width="627" alt="image" src="https://github.com/H8-FTGO-P3/FTGO-P3-V1-GC2/assets/134502566/5571ab9f-a203-44b4-ae3c-6f93bf681b62">

- Dokumentasi API Swagger.
  <img width="983" alt="image" src="https://github.com/H8-FTGO-P3/FTGO-P3-V1-GC2/assets/134502566/156ecb65-857d-426e-87f2-a48b27f66417">
