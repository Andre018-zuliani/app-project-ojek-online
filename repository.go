package main

import (
	"database/sql"
)

// Struct untuk menampung hasil query
type MonthlyTopCustomer struct {
	Month        string
	CustomerName string
	TotalOrder   int
}

type LocationStats struct {
	Location   string
	TotalOrder int
}

type HourlyStats struct {
	Hour       int
	TotalOrder int
}

// Repository Interface
type ReportRepository interface {
	GetTopCustomerPerMonth() ([]MonthlyTopCustomer, error)
	GetTopLocations() ([]LocationStats, error)
	GetHourlyStats() ([]HourlyStats, error)
}

type reportRepo struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepo{db: db}
}

// Fitur 2: Top Customer Per Bulan
func (r *reportRepo) GetTopCustomerPerMonth() ([]MonthlyTopCustomer, error) {
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
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []MonthlyTopCustomer
	processedMonths := make(map[string]bool)

	for rows.Next() {
		var d MonthlyTopCustomer
		if err := rows.Scan(&d.Month, &d.CustomerName, &d.TotalOrder); err != nil {
			return nil, err
		}
		
		if !processedMonths[d.Month] {
			results = append(results, d)
			processedMonths[d.Month] = true
		}
	}
	return results, nil
}

// Fitur 3: Lokasi Terbanyak
func (r *reportRepo) GetTopLocations() ([]LocationStats, error) {
	query := `
		SELECT pickup_location, COUNT(*) as total 
		FROM ride 
		GROUP BY pickup_location 
		ORDER BY total DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []LocationStats
	for rows.Next() {
		var l LocationStats
		if err := rows.Scan(&l.Location, &l.TotalOrder); err != nil {
			return nil, err
		}
		results = append(results, l)
	}
	return results, nil
}

// Fitur 4: Waktu Ramai
func (r *reportRepo) GetHourlyStats() ([]HourlyStats, error) {
	query := `
		SELECT EXTRACT(HOUR FROM start_time) as jam, COUNT(*) as total 
		FROM ride 
		GROUP BY jam 
		ORDER BY total DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []HourlyStats
	for rows.Next() {
		var h HourlyStats
		if err := rows.Scan(&h.Hour, &h.TotalOrder); err != nil {
			return nil, err
		}
		results = append(results, h)
	}
	return results, nil
}