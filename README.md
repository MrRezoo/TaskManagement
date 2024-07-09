# Task Management API

This repository contains the Task Management API which provides endpoints to manage tasks including creating, retrieving, and deleting tasks.

## Table of Contents

- [Task Management API](#task-management-api)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Installation](#installation)
  - [API Endpoints](#api-endpoints)
    - [Health Check](#health-check)
    - [Retrieve All Tasks](#retrieve-all-tasks)
    - [Create a New Task](#create-a-new-task)
    - [Retrieve a Specific Task](#retrieve-a-specific-task)
    - [Delete a Specific Task](#delete-a-specific-task)
  - [Examples](#examples)
    - [Health Check Example](#health-check-example)
    - [Retrieve All Tasks Example](#retrieve-all-tasks-example)
    - [Create a New Task Example](#create-a-new-task-example)
    - [Retrieve a Specific Task Example](#retrieve-a-specific-task-example)
    - [Delete a Specific Task Example](#delete-a-specific-task-example)
  - [License](#license)

## Overview

The Task Management API provides a simple interface to manage tasks. The API allows you to perform CRUD (Create, Read, Update, Delete) operations on tasks.

## Installation

To set up the Task Management API locally:

1. Clone the repository:
   ```sh
   git clone https://github.com/MrRezoo/TaskManagement.git
   cd TaskManagement
   ```

2. Install the dependencies:
   ```sh
   npm install
   ```

3. Start the server:
   ```sh
   npm start
   ```

The API will be available at `http://localhost:3000`.

## API Endpoints

### Health Check

**GET** `/api/health/`

Checks the health status of the API.

#### Responses
- `200 OK`: Health status of the API.

### Retrieve All Tasks

**GET** `/api/tasks/`

Retrieves a list of all tasks.

#### Responses
- `200 OK`: A list of tasks.

### Create a New Task

**POST** `/api/tasks/`

Creates a new task.

#### Request Body
- `title` (string): Title of the task.
- `description` (string): Description of the task.
- `status` (string): Status of the task (`pending`, `completed`, `in_progress`).

#### Responses
- `201 Created`: Task created successfully.
- `400 Bad Request`: Invalid input.

### Retrieve a Specific Task

**GET** `/api/tasks/{id}/`

Retrieves a specific task by ID.

#### Path Parameters
- `id` (string): ID of the task.

#### Responses
- `200 OK`: Task details.
- `404 Not Found`: Task not found.

### Delete a Specific Task

**DELETE** `/api/tasks/{id}/`

Deletes a specific task by ID.

#### Path Parameters
- `id` (string): ID of the task.

#### Responses
- `204 No Content`: Task deleted successfully.

## Examples

### Health Check Example

**Request**
```sh
curl -X GET http://localhost:3000/api/health/
```

**Response**
```json
{
  "message": "",
  "success": true
}
```

### Retrieve All Tasks Example

**Request**
```sh
curl -X GET http://localhost:3000/api/tasks/
```

**Response**
```json
{
  "result": [
    {
      "id": "3ea40a96-045a-45ae-b526-8e7f515d335e",
      "title": "One",
      "description": "des",
      "status": "pending",
      "created_at": "2024-07-09T19:08:55.994283+03:30",
      "updated_at": "2024-07-09T19:08:55.994283+03:30"
    }
  ],
  "success": true,
  "code": 200,
  "validation_errors": null,
  "error": null
}
```

### Create a New Task Example

**Request**
```sh
curl -X POST http://localhost:3000/api/tasks/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "One",
    "description": "des",
    "status": "pending"
  }'
```

**Response**
```json
{
  "result": {
    "id": "3ea40a96-045a-45ae-b526-8e7f515d335e",
    "title": "One",
    "description": "des",
    "status": "pending",
    "created_at": "2024-07-09T19:08:55.994283+03:30",
    "updated_at": "2024-07-09T19:08:55.994283+03:30"
  },
  "success": true,
  "code": 201,
  "validation_errors": null,
  "error": null
}
```

### Retrieve a Specific Task Example

**Request**
```sh
curl -X GET http://localhost:3000/api/tasks/3ea40a96-045a-45ae-b526-8e7f515d335e/
```

**Response**
```json
{
  "result": {
    "id": "3ea40a96-045a-45ae-b526-8e7f515d335e",
    "title": "One",
    "description": "des",
    "status": "pending",
    "created_at": "2024-07-09T19:08:55.994283+03:30",
    "updated_at": "2024-07-09T19:08:55.994283+03:30"
  },
  "success": true,
  "code": 200,
  "validation_errors": null,
  "error": null
}
```

### Delete a Specific Task Example

**Request**
```sh
curl -X DELETE http://localhost:3000/api/tasks/3ea40a96-045a-45ae-b526-8e7f515d335e/
```

**Response**
```
204 No Content
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

This README includes detailed explanations for each API endpoint along with examples of how to use them. Adjust the server URL (`http://localhost:3000`) and any other specific details as necessary to match your actual setup.