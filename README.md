# Kafka2MySQL

Kafka2MySQL is a Go-based project that demonstrates how to consume messages from an Apache Kafka topic and store them in a MySQL database. It uses several powerful libraries including Watermill for message processing, GORM for ORM (Object-Relational Mapping), and Gin for HTTP server routing.

## Table of Contents

- [Kafka2MySQL](#kafka2mysql)
    - [Table of Contents](#table-of-contents)
    - [Features](#features)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Project Structure](#project-structure)
    - [Libraries Used](#libraries-used)
    - [Contributing](#contributing)
    - [License](#license)

## Features

- Consume messages from a Kafka topic.
- Store consumed messages in a MySQL database.
- Expose a REST API to publish messages to the Kafka topic.
- Automatic schema migration for the MySQL database.

## Prerequisites

- Go 1.20+
- Docker and Docker Compose
- MySQL
- Apache Kafka

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/kafka2mysql.git
    cd kafka2mysql
    ```

2. **Install dependencies**:
    ```sh
    go mod tidy
    ```

3. **Set up MySQL and Kafka using Docker Compose**:
   Ensure you have Docker and Docker Compose installed. Run the following command to start MySQL and Kafka:
    ```sh
    docker-compose up -d
    ```

## Usage

1. **Run the application**:
    ```sh
    go run main.go
    ```

2. **Publish a message to Kafka**:
   Use an HTTP GET request to publish a message to Kafka.
    ```sh
    curl "http://localhost:8080/api/v1/kafka2mysql?message=HelloKafka"
    ```

3. **Check MySQL**:
   The message should be stored in the MySQL database.

## Project Structure

```
kafka2mysql/
├── configs/
│   └── config.go            # Configuration loading logic
├── providers/
│   ├── gin.go               # Gin server setup
│   ├── mysql.go             # MySQL setup
│   └── watermill.go         # Watermill router setup
├── subscriber/
│   └── kafka_subscriber.go  # Kafka subscriber implementation
├── publisher/
│   └── kafka_publisher.go   # Kafka publisher implementation
├── main.go                  # Main entry point
└── go.mod                   # Go modules file
```


## Libraries Used

- [Watermill](https://github.com/ThreeDotsLabs/watermill): A library for building event-driven applications.
- [GORM](https://gorm.io/): An ORM library for Go.
- [Gin](https://github.com/gin-gonic/gin): A web framework for Go.
- [Go-JSON](https://github.com/goccy/go-json): A JSON library for Go, used for marshalling and unmarshalling JSON.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss what you would like to change.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

