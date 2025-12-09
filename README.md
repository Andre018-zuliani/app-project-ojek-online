# Ojek Online Report

This project is a Go application that generates reports for an online ride-hailing service. It connects to a PostgreSQL database to retrieve and display various statistics, including top customers, top pickup locations, and hourly order statistics.

## Project Structure

```
ojek-online-report
├── go.mod
├── go.sum
├── main.go
├── repository.go
├── repository_test.go
├── pkg
│   └── db
│       └── connection.go
└── README.md
```

## Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd ojek-online-report
   ```

2. **Initialize Go Modules**
   Ensure you have Go installed, then run:
   ```bash
   go mod tidy
   ```

3. **Database Setup**
   - Ensure you have PostgreSQL installed and running.
   - Create a database named `ojek_online`.
   - Import the necessary data into the database from your SQL backup.

4. **Configuration**
   - Update the database connection string in `main.go` to match your PostgreSQL credentials.

## Usage

To run the application, execute the following command:

```bash
go run main.go
```

This will display the reports for top customers, top pickup locations, and hourly statistics.

## Testing

To run the unit tests, use the following command:

```bash
go test -v
```

This will execute the tests defined in `repository_test.go`, ensuring that the repository methods work as expected.

## Dependencies

- `github.com/lib/pq`: PostgreSQL driver for Go.
- `github.com/DATA-DOG/go-sqlmock`: Mocking library for testing database interactions.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.