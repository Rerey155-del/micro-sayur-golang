# Micro Sayur

Micro Sayur adalah aplikasi e-commerce sayur berbasis microservices dengan frontend `Nuxt 4` dan backend `Go`. Proyek ini memisahkan fitur utama ke beberapa service agar alur autentikasi, produk, dan pesanan dapat dikelola secara modular.

## Fitur

- Autentikasi user dengan role `Customer` dan `Super Admin`
- Manajemen produk oleh admin
- Pembelian produk oleh customer
- Riwayat pesanan customer
- Monitoring dan update status pesanan dari dashboard admin
- Docker Compose untuk menjalankan seluruh service

## Stack

- Frontend: `Nuxt 4`, `Vue 3`
- Backend:
  - `user-service`: `Go` + `Echo`
  - `product-service`: `Go` + `Fiber`
  - `order-service`: `Go` + `Fiber`
- Database: `MySQL 8`
- ORM: `GORM`
- Containerization: `Docker`, `Docker Compose`

## Arsitektur

Project ini memakai pendekatan `microservices`:

- `user-service`
  Menangani registrasi, login, dan pengambilan data user.
- `product-service`
  Menangani data produk, stok, harga, dan upload gambar produk.
- `order-service`
  Menangani pembuatan order, riwayat order customer, dan update status order.
- `nuxt-frontend`
  Menyediakan UI untuk customer dan admin.

## Struktur Folder

```text
micro-sayur/
├─ user-service/
├─ product-service/
├─ order-service/
├─ payment-service/
├─ notification-service/
├─ nuxt-frontend/
├─ docker-compose.yml
└─ init.sql
```

## Port Service

| Service | Port | Keterangan |
|---|---:|---|
| `frontend` | `3000` | Nuxt frontend |
| `product-service` | `8081` | API produk |
| `order-service` | `8082` | API order |
| `user-service` | `8085` | API user dan auth |
| `mysql` | `3307` | MySQL host port |

## Database

Saat container database pertama kali dijalankan, file [init.sql](/c:/xampp/htdocs/micro-sayur/init.sql) akan membuat database:

- `sayur-user-service`
- `sayur-product-service`
- `sayur-order-service`

Schema order juga tersedia di:

- [order-service/database/schema.sql](/c:/xampp/htdocs/micro-sayur/order-service/database/schema.sql)

## Endpoint Utama

### User Service

- `POST /signin`
- `POST /signup`
- `GET /users`

### Product Service

- `GET /api/v1/products`
- `POST /api/v1/products`
- `PUT /api/v1/products/:id`
- `DELETE /api/v1/products/:id`

### Order Service

- `GET /api/v1/orders`
- `GET /api/v1/orders/user/:userId`
- `POST /api/v1/orders`
- `PATCH /api/v1/orders/:id/status`

## Cara Menjalankan Dengan Docker

Pastikan `Docker` dan `Docker Compose` sudah terpasang, lalu dari root project jalankan:

```bash
docker-compose up --build
```

Setelah semua service aktif:

- Frontend: `http://localhost:3000`
- User service: `http://localhost:8085`
- Product service: `http://localhost:8081`
- Order service: `http://localhost:8082`
- MySQL: `localhost:3307`

## Cara Menjalankan Manual

### 1. Jalankan database MySQL

Pastikan database berikut sudah tersedia:

- `sayur-user-service`
- `sayur-product-service`
- `sayur-order-service`

### 2. Jalankan backend service

Contoh:

```bash
cd user-service
go run .
```

```bash
cd product-service
go run .
```

```bash
cd order-service
go run .
```

### 3. Jalankan frontend

```bash
cd nuxt-frontend
npm install
npm run dev
```

## Flow Aplikasi

### Customer

1. User register atau login
2. Customer membuka halaman produk
3. Customer membuat pesanan dari halaman produk
4. Customer melihat riwayat pesanan di halaman pesanan

### Admin

1. Admin login
2. Admin mengelola produk
3. Admin memantau daftar pesanan
4. Admin mengubah status pesanan

