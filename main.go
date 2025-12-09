package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Pastikan connection string ini sesuai dengan settingan Postgre di laptopmu
	connStr := "host=127.0.0.1 port=5432 user=postgres password=password dbname=ojek_online sslmode=disable"
	
	// Perbaikan: Gunakan sql.Open, BUKAN db.Connect
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Cek koneksi
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal konek ke DB:", err)
	}

	repo := NewReportRepository(db)

	fmt.Println("=== REPORT OJEK ONLINE ===")

	// 1. Customer Paling Sering Order Tiap Bulan
	fmt.Println("\n[Fitur 2] Top Customer per Bulan:")
	topCustomers, err := repo.GetTopCustomerPerMonth()
	if err != nil {
		log.Println("Error:", err)
	} else {
		for _, tc := range topCustomers {
			fmt.Printf("Bulan: %s | Customer: %s | Total Order: %d\n", tc.Month, tc.CustomerName, tc.TotalOrder)
		}
	}

	// 2. Lokasi Terbanyak
	fmt.Println("\n[Fitur 3] Lokasi Pickup Teramai:")
	locations, err := repo.GetTopLocations()
	if err != nil {
		log.Println("Error:", err)
	} else {
		for i, l := range locations {
			fmt.Printf("%d. %s (%d order)\n", i+1, l.Location, l.TotalOrder)
		}
	}

	// 3. Waktu Ramai dan Sepi
	fmt.Println("\n[Fitur 4] Analisa Waktu Order (Jam):")
	hours, err := repo.GetHourlyStats()
	if err != nil {
		log.Println("Error:", err)
	} else if len(hours) > 0 {
		fmt.Printf("Waktu Paling RAMAI: Jam %02d:00 (%d order)\n", hours[0].Hour, hours[0].TotalOrder)
		
		lastIdx := len(hours) - 1
		fmt.Printf("Waktu Paling SEPI : Jam %02d:00 (%d order)\n", hours[lastIdx].Hour, hours[lastIdx].TotalOrder)
	}
}