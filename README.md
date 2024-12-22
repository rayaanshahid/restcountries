# REST Countries Go Project

This project is a console application that fetches data from the [REST Countries API](https://restcountries.com/) and processes it to provide two main functionalities:

1. A list of all countries in Europe with the following attributes:
   - Name
   - Currency
   - Ability to sort results based on name or currency in ascending or descending order.
   
2. A list of all currencies in the world, along with information on which countries each currency is used in.

## Technologies Used

### 1. **Go (Golang)**
   - The main programming language used for building this project. Go is a statically typed, compiled language known for its simplicity, concurrency support, and performance. It is ideal for building scalable applications, especially those involving networking and APIs.

   - **Version**: Go 1.22.4

## Features

1. Fetches all countries in Europe and filters them based on name and currency.
2. Provides sorting functionality for country names and currencies in ascending or descending order.
3. Lists all currencies in the world with information on which countries use each currency.
4. Utilizes Go’s built-in support for concurrent HTTP requests and JSON parsing.


## Getting Started

### Prerequisites
- Install Go: [Go Downloads](https://go.dev/dl/)
- Clone the repository: `git clone https://github.com/rayaanshahid/restcountries.git`

### Installing Dependencies

Navigate to the project directory and run the following command to fetch dependencies:

### Running the Application

To run the application, execute the following command:

go run main.go

### Running Tests

To run the tests, use the Go test command:

go test ./…