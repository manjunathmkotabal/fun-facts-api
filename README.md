# Fun Facts API

## Overview
The Fun Facts API allows users to sign up, log in, and query trivia facts. The API supports role-based access control (RBAC) to differentiate between admin and regular users.

## Features
- User Signup and Login
- JWT Authentication
- Role-Based Access Control (RBAC)
- Query Trivia by Category (Admin Only)
- Query Random Trivia (Admin and User)

## Endpoints

### Public Endpoints

#### Signup
- **URL**: `/signup`
- **Method**: `POST`
- **Description**: Register a new user.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string",
    "role": "string" // "admin" or "user"
  }
  ```
- **Response**:
  ```json
  {
    "message": "User created successfully"
  }
  ```

#### Login
- **URL**: `/login`
- **Method**: `POST`
- **Description**: Authenticate a user and get a JWT token.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Response**:
  ```json
  {
    "token": "string"
  }
  ```

#### Refresh Token
- **URL**: `/refresh`
- **Method**: `POST`
- **Description**: Refresh the JWT token.
- **Headers**:
  - `Authorization: Bearer <token>`
- **Response**:
  ```json
  {
    "token": "string"
  }
  ```

### Admin Endpoints

#### Query Trivia by Category
- **URL**: `/admin/query-by-category/:category`
- **Method**: `GET`
- **Description**: Query trivia by category.
- **Headers**:
  - `Authorization: Bearer <token>`
- **Response**:
  ```json
  {
    "data": [
      {
        "ID": "integer",
        "Category": "string",
        "Content": "string"
      }
    ]
  }
  ```

#### Query Random Trivia
- **URL**: `/admin/query-trivia`
- **Method**: `GET`
- **Description**: Query random trivia.
- **Headers**:
  - `Authorization: Bearer <token>`
- **Response**:
  ```json
  {
    "data": [
      {
        "ID": "integer",
        "Category": "string",
        "Content": "string"
      }
    ]
  }
  ```

### User Endpoints

#### Query Random Trivia
- **URL**: `/user/query-trivia`
- **Method**: `GET`
- **Description**: Query random trivia.
- **Headers**:
  - `Authorization: Bearer <token>`
- **Response**:
  ```json
  {
    "data": [
      {
        "ID": "integer",
        "Category": "string",
        "Content": "string"
      }
    ]
  }
  ```

## Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/manjunath.kotabal/fun-facts-api.git
   cd fun-facts-api
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Set up the database:
   ```sh
   go run main.go
   ```

4. Run the application:
   ```sh
   go run main.go
   ```

## Environment Variables

Create a `.env` file in the root directory and add the following environment variables:

```
JWT_SECRET=my_secret_key
```

## Testing

Use the following `curl` commands to test the API endpoints:

### Signup
```sh
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" -d '{
  "username": "adminuser",
  "password": "adminpassword",
  "role": "admin"
}'
```

### Login
```sh
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{
  "username": "adminuser",
  "password": "adminpassword"
}'
```

### Query Trivia by Category (Admin)
```sh
TOKEN=<your_token>
curl -X GET http://localhost:8080/admin/query-by-category/science -H "Authorization: Bearer $TOKEN"
```

### Query Random Trivia (Admin)
```sh
TOKEN=<your_token>
curl -X GET http://localhost:8080/admin/query-trivia -H "Authorization: Bearer $TOKEN"
```

### Query Random Trivia (User)
```sh
TOKEN=<your_token>
curl -X GET http://localhost:8080/user/query-trivia -H "Authorization: Bearer $TOKEN"
```