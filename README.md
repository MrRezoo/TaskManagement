# Task Management API

This project is a simple Task Management API built with Go, Fiber, and PostgreSQL. It allows users to create, read, update, and delete tasks.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Configuration](#configuration)
- [Running Tests](#running-tests)
- [Docker Support](#docker-support)
- [Contributing](#contributing)
- [License](#license)

## Features

- Create, read, update, and delete tasks
- Health check endpoint

## Prerequisites

- Go 1.18+
- PostgreSQL
- Docker (optional)

## Installation

### Clone the repository

```sh
git clone https://github.com/MrRezoo/task-management-api.git
cd task-management-api
```

### Server Setup

1. **Install Go dependencies**

    ```sh
    go mod tidy
    ```

2. **Run the server**

    ```sh
    go run cmd/main.go
    ```

## Usage

### Running the Server

```sh
go run cmd/main.go
```

## API Endpoints

### Health Check

- **URL:** `/api/health`
- **Method:** `GET`
- **Responses:**
    - `200 OK`

### Create Task

- **URL:** `/api/tasks`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "title": "Task Title",
    "description": "Task Description",
    "status": "Pending"
  }
  ```
- **Responses:**
    - `201 Created`
    - `400 Bad Request`

### Get All Tasks

- **URL:** `/api/tasks`
- **Method:** `GET`
- **Responses:**
    - `200 OK`

### Get Task by ID

- **URL:** `/api/tasks/:id`
- **Method:** `GET`
- **Responses:**
    - `200 OK`
    - `404 Not Found`

### Update Task

- **URL:** `/api/tasks/:id`
- **Method:** `PATCH`
- **Request Body:**
  ```json
  {
    "title": "Updated Title",
    "description": "Updated Description",
    "status": "In Progress"
  }
  ```
- **Responses:**
    - `200 OK`
    - `400 Bad Request`
    - `404 Not Found`

### Delete Task

- **URL:** `/api/tasks/:id`
- **Method:** `DELETE`
- **Responses:**
    - `204 No Content`
    - `404 Not Found`

## Configuration

Configuration files are located in the `config` directory. Modify the `.yaml` files according to your environment.

- `config-development.yaml`
- `config-docker.yaml`

## Running Tests

To run the tests, use the following command:

```sh
go test ./...
```

## Docker Support

To run the server using Docker, use the provided `docker-compose` file.

### Build and Run

```sh
docker-compose up --build
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

This structure and documentation should provide a comprehensive guide to setting up and running your Task Management API project. Adjust paths, package names, and other details as needed.