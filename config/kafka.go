package config

import (
	"log"

	"github.com/segmentio/kafka-go"
)

// ConnectBroker mencoba menghubungkan aplikasi ke Kafka broker.
// Fungsi ini digunakan untuk memverifikasi apakah aplikasi bisa terhubung ke broker Kafka.
func ConnectBroker(brokerAddress string) {
	
	// Dial ke broker Kafka untuk memverifikasi koneksi
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke Kafka broker di alamat %s: %v", brokerAddress, err)
	}
	defer conn.Close() // Pastikan koneksi ditutup setelah selesai

	// Log jika koneksi berhasil
	log.Printf("Berhasil terhubung ke Kafka broker di alamat: %s", brokerAddress)
}


// InitWriter digunakan untuk membuat Kafka writer yang bertugas mengirimkan pesan pada topic.
func InitKafkaWriter(broker, topic string) *kafka.Writer {

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return writer

}
