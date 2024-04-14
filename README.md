# carbon-service

`carbon-service` is a lightweight API written in Go, designed to facilitate life cycle assessment (LCA) and carbon calculations for the building industry. It provides endpoints for various operations related to building materials, assemblies, and constructions, allowing users to perform carbon footprint calculations and analyze environmental impacts.

## Features

- RESTful API: Easily interact with the service using standard HTTP requests.
  -Gorm and PostgreSQL: Utilizes Gorm as an ORM library and PostgreSQL as the database backend.
- Dockerized Environment: Ships with Docker and Docker Compose configurations for easy deployment and development.
- Test Suite: Includes comprehensive test suites to ensure reliability and stability.

## Getting Started

To get started with the Carbon Service, follow these steps:

Prerequisites

1. Go (version 1.22.1 or higher)
2. Docker
3. Docker Compose

### Installation

Clone the repository to your local machine:

```
git clone https://github.com/karim-daw/carbon-service.git
```

Navigate to the project directory:

```
cd carbon-service
```

Build the Docker containers:

```
docker-compose build
```

Start the Docker containers:

```
docker-compose up
```

## Usage

The API will be accessible at http://localhost:80.
Use tools like cURL, Postman, or your preferred HTTP client to interact with the API endpoints.

## Environment Variables

The .env file contains environment variables used by the application. Customize it according to your requirements.

## Tests

Run the test suite using:

```
docker-compose run carbon-service go test ./tests
```

# Contributing

We welcome contributions from the community! If you find any bugs, have feature requests, or want to contribute improvements, please open an issue or submit a pull request.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments
This project was inspired by the need for sustainable solutions in the building industry.
Special thanks to the contributors who helped make this project possible.
