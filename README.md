# LRU Cache Application

This project is an implementation of an LRU (Least Recently Used) cache system with a React frontend and a GoLang backend. The application allows users to interact with the cache by setting key-value pairs with expiration times, retrieving values by key, and deleting entries.

## Features

- **LRU Cache**: Stores key-value pairs with expiration times.
- **API Endpoints**: RESTful API for GET, SET, and DELETE operations.
- **React UI**: User interface for interacting with the cache.
- **Concurrency**: Concurrent cache access in the backend.
- **WebSocket (Optional)**: Real-time updates to the UI (good to have).

## Tech Stack

- **Backend**: GoLang
- **Frontend**: React, Axios
- **Communication**: HTTP, WebSocket (Optional)


## Getting Started

### Prerequisites

- **Node.js**: Ensure you have Node.js installed (required for the frontend).
- **Go**: Ensure you have Go installed (required for the backend).

### Backend Setup

1. **Clone the repository:**

   git clone https://github.com/christy-mariam/lru-cache-backend.git
   cd lru-cache-backend

2. **Install dependencies:**

   The Go backend does not have external dependencies beyond the standard library, so you can skip this step.

3. **Run the backend server:**

   go run main.go

   The backend server will start on `http://localhost:8080`.

### Frontend Setup

1. **Navigate to the frontend directory:**

   cd lru-cache-frontend

2. **Install dependencies:**

   npm install

3. **Run the frontend application:**

   npm start

   The frontend will be accessible at `http://localhost:3000`.

### Note

Ensure that both the backend and frontend servers are running concurrently to enable full functionality.

## API Endpoints

Here are the available API endpoints for the LRU cache:

- **GET `/api/cache?key={key}`**: Retrieve the value for a given key.

  **Example:**

  curl -X GET "http://localhost:8080/api/cache?key=test"

- **POST `/api/cache`**: Set a new key-value pair with an expiration time.

  **Payload:**

  {
    "key": "test",
    "value": "123",
    "expiration": 5
  }

  **Example:**

  curl -X POST "http://localhost:8080/api/cache" -H "Content-Type: application/json" -d '{"key":"test","value":"123","expiration":5}'

- **DELETE `/api/cache?key={key}`**: Delete the key-value pair for a given key.

  **Example:**

  curl -X DELETE "http://localhost:8080/api/cache?key=test"

## WebSocket Integration (Optional)

For real-time updates, you can implement WebSocket support. The backend and frontend need additional code for WebSocket communication, which can be added later.

## Project Structure

### Backend

lru-cache-backend/
|-- main.go
|-- cache/
| |-- cache.go
| |-- lru.go
|-- handlers/
| |-- handlers.go
|-- go.mod

### Frontend

lru-cache-frontend/
|-- src/
| |-- App.js
| |-- CacheInput.js
| |-- CacheDisplay.js
| |-- CacheItem.js
|-- package.json

