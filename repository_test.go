package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetTopCustomerPerMonth(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewReportRepository(db)

	rows := sqlmock.NewRows([]string{"bulan", "name", "total"}).
		AddRow("2023-01", "Alice", 5).
		AddRow("2023-02", "Bob", 3)

	mock.ExpectQuery("SELECT TO_CHAR(start_time, 'YYYY-MM') as bulan, c.name, COUNT\\(r.id\\) as total FROM ride r JOIN customer c ON r.customer_id = c.id GROUP BY bulan, c.name ORDER BY bulan DESC, total DESC").
		WillReturnRows(rows)

	customers, err := repo.GetTopCustomerPerMonth()

	assert.NoError(t, err)
	assert.Len(t, customers, 2)
	assert.Equal(t, "Alice", customers[0].CustomerName)
	assert.Equal(t, 5, customers[0].TotalOrder)

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
		AddRow("ITPLN", 5).
		AddRow("Ciledug", 2)

	mock.ExpectQuery("SELECT pickup_location, COUNT\\(\\*\\) as total FROM ride GROUP BY pickup_location ORDER BY total DESC").
		WillReturnRows(rows)

	stats, err := repo.GetTopLocations()

	assert.NoError(t, err)
	assert.Len(t, stats, 3)
	assert.Equal(t, "Grogol", stats[0].Location)
	assert.Equal(t, 10, stats[0].TotalOrder)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetHourlyStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewReportRepository(db)

	rows := sqlmock.NewRows([]string{"jam", "total"}).
		AddRow(8, 15).
		AddRow(9, 20).
		AddRow(10, 5)

	mock.ExpectQuery("SELECT EXTRACT\\(HOUR FROM start_time\\) as jam, COUNT\\(\\*\\) as total FROM ride GROUP BY jam ORDER BY total DESC").
		WillReturnRows(rows)

	hours, err := repo.GetHourlyStats()

	assert.NoError(t, err)
	assert.Len(t, hours, 3)
	assert.Equal(t, 8, hours[0].Hour)
	assert.Equal(t, 15, hours[0].TotalOrder)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}