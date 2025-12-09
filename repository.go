package main

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetTopCustomerPerMonth(t *testing.T) {
	// 1. Setup Mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewReportRepository(db)

	// 2. Siapkan Data Dummy yang diharapkan kembali
	rows := sqlmock.NewRows([]string{"bulan", "name", "total"}).
		AddRow("2025-12", "Andre", 2).
		AddRow("2025-11", "Budi", 1)

	// 3. Query yang diharapkan (Sama persis dengan di repository.go)
	query := `
		SELECT 
			TO_CHAR(start_time, 'YYYY-MM') as bulan,
			c.name, 
			COUNT(r.id) as total
		FROM ride r
		JOIN customer c ON r.customer_id = c.id
		GROUP BY bulan, c.name
		ORDER BY bulan DESC, total DESC
	`

	// PENTING: Gunakan regexp.QuoteMeta agar simbol () dibaca sebagai text biasa
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// 4. Jalankan Fungsi
	results, err := repo.GetTopCustomerPerMonth()

	// 5. Validasi (Assert)
	if err != nil {
		t.Errorf("error was not expected: %s", err)
		return
	}

	// Cek apakah hasil kosong (Mencegah Panic index out of range)
	if len(results) == 0 {
		t.Errorf("expected results, got empty list")
		return
	}

	// Cek Data Pertama
	if results[0].CustomerName != "Andre" {
		t.Errorf("expected CustomerName Andre, got %s", results[0].CustomerName)
	}
	if results[0].TotalOrder != 2 {
		t.Errorf("expected TotalOrder 2, got %d", results[0].TotalOrder)
	}

	// Pastikan semua ekspektasi terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTopLocations(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewReportRepository(db)

	rows := sqlmock.NewRows([]string{"pickup_location", "total"}).
		AddRow("Grogol", 10).
		AddRow("ITPLN", 5)

	query := `
		SELECT pickup_location, COUNT(*) as total 
		FROM ride 
		GROUP BY pickup_location 
		ORDER BY total DESC
	`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	stats, err := repo.GetTopLocations()

	if err != nil {
		t.Errorf("error was not expected: %s", err)
		return
	}

	if len(stats) == 0 {
		t.Errorf("expected results, got empty list")
		return
	}

	if stats[0].Location != "Grogol" {
		t.Errorf("expected top location to be Grogol, got %s", stats[0].Location)
	}
	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
