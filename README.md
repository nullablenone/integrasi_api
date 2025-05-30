# Integrasi API

Proyek ini adalah backend service berbasis Go yang dirancang untuk mengintegrasikan berbagai API eksternal secara modular dan efisien. Dengan struktur folder yang rapi, pengembangan dan pemeliharaan kode menjadi lebih mudah dan terorganisir.

## Struktur Folder

```
integrasi_api/
├── config/                  # Konfigurasi aplikasi
│   ├── database.go          # Koneksi database
│   ├── env.go               # Pengaturan environment
│   ├── kafka.go             # Koneksi Kafka
│   └── redis.go             # Koneksi Redis
├── constants/               # Konstanta global
│   └── constants.go
├── internal/                # Logika bisnis dan integrasi
│   ├── domain/              # Domain khusus aplikasi
│   │   └── user/            # Modul user
│   │       ├── handler.go   # Handler untuk HTTP request
│   │       ├── model.go     # Definisi model data
│   │       ├── repository.go # Operasi database
│   │       └── service.go   # Logika bisnis
│   ├── integration/         # Modul integrasi pihak ketiga
│   │   └── jsonplaceholder/
│   │       ├── client.go    # Koneksi API eksternal
│   │       ├── dto.go       # Data Transfer Object
│   │       └── service.go   # Logika pengolahan data eksternal
│   └── kafka/producer/      # Modul producer Kafka
│       └── service.go
├── routes/                  # Definisi rute HTTP
│   └── routes.go
├── utils/                   # Fungsi utilitas umum
│   └── redis.go
├── .env                     # Variabel lingkungan
├── docker-compose.yml       # Konfigurasi Docker Compose
├── go.mod                   # File modul Go
├── go.sum                   # Checksum dependensi Go
└── main.go                  # Entry point aplikasi
```

## Fitur

* **Modular Architecture**: Struktur kode yang terpisah untuk domain, integrasi, dan utilitas.
* **Integrasi API Pihak Ketiga**: Mendukung API seperti JSONPlaceholder.
* **Message Broker**: Producer Kafka untuk pengiriman pesan.
* **Database dan Caching**: Dukungan untuk database MySQL dan Redis sebagai caching layer.
* **Dockerized Deployment**: Kemudahan dalam pengaturan dan deployment dengan Docker.

## Tech Stack

* **Bahasa Pemrograman**: Go (Golang)
* **Message Broker**: Kafka
* **Database**: MySQL
* **Caching**: Redis
* **Containerization**: Docker & Docker Compose

## Setup

1. **Clone Repository**

   ```bash
   git clone https://github.com/nullablenone/integrasi_api.git
   cd integrasi_api
   ```

2. **Siapkan File `.env`**

   Buat file `.env` di root proyek dan tambahkan variabel lingkungan berikut sesuai dengan kebutuhan aplikasi Anda:

   ```env
   # MySQL
   DB_USER=root
   DB_PASS=
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=integrasi_api

   # API External
   EXTERNAL_API_URL=https://jsonplaceholder.typicode.com

   # Redis
   REDIS_PORT=6379

   # Zookeeper
   ZOOKEEPER_CLIENT_PORT=2181
   ZOOKEEPER_TICK_TIME=2000

   # Kafka
   KAFKA_BROKER_ID=1
   KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
   KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
   KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092
   KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1
   ```

3. **Install Dependensi Go**

   Jalankan perintah berikut untuk memastikan semua dependensi Go terpasang:

   ```bash
   go mod tidy
   ```

4. **Jalankan dengan Docker Compose**

   Pastikan Docker dan Docker Compose telah terinstal, lalu jalankan:

   ```bash
   docker-compose up --build
   ```

   Aplikasi akan tersedia di `http://localhost:8080` (sesuai konfigurasi port).

2. **Menambahkan Modul Baru**

   * Tambahkan domain baru di folder `internal/domain/`.
   * Jika perlu, tambahkan modul integrasi di `internal/integration/`.

3. **Menjalankan Uji Coba**

   Jalankan unit test dengan:

   ```bash
   go test ./...
   ```
