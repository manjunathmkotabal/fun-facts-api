# Fun Facts API

This is an Open-Source API for Random Fun Facts & Trivia. It provides users with random trivia or fun facts on demand. The API features categorized trivia (science, history, tech, etc.), a daily random fact generator, and API access for chatbots or apps.

## Features

- Categorized trivia (science, history, tech, etc.)
- Daily random fact generator
- API access for chatbots or apps
- User authentication with JWT
- Rate limiting

## Prerequisites

- Go 1.17 or later
- SQLite

## Setup

1. **Clone the repository**:

```sh
git clone https://github.com/manjunathmkotabal/fun-facts-api.git
cd fun-facts-api
```

2. **Install dependencies**:

```sh
go mod tidy
```

3. **Create a `.env` file**:

Create a `.env` file in the root of the project directory with any necessary environment variables. For this example, you can leave it empty or add any required variables.

```sh
touch .env
```

4. **Run the application**:

You can use `reflex` to automatically restart the application whenever you make changes to your code.

```sh
reflex -c reflex.conf
```

Alternatively, you can run the application manually:

```sh
go run main.go
```

## API Endpoints

### Public Endpoints

- **Signup**: Create a new user

```sh
POST /signup
```

Request body:

```json
{
  "username": "user",
  "password": "password"
}
```

- **Login**: Authenticate a user and get a JWT token

```sh
POST /login
```

Request body:

```json
{
  "username": "user",
  "password": "password"
}
```

- **Refresh Token**: Refresh the JWT token

```sh
POST /refresh
```

Headers:

```sh
Authorization: <your-token>
```

### Protected Endpoints

- **Get Random Trivia**: Get a random trivia fact

```sh
GET /trivia/random
```

Headers:

```sh
Authorization: <your-token>
```

- **Get Trivia by Category**: Get trivia facts by category

```sh
GET /trivia/category/:category
```

Headers:

```sh
Authorization: <your-token>
```

## Project Structure

```
/Users/manjunath.kotabal/go/fun-facts-api
├── config
│   └── config.go
├── controllers
│   ├── auth_controller.go
│   ├── trivia_controller.go
│   └── user_controller.go
├── middleware
│   ├── auth_middleware.go
│   └── rate_limiter.go
├── models
│   └── trivia.go
├── routes
│   └── routes.go
├── services
│   └── trivia_service.go
├── .gitignore
├── go.mod
├── main.go
├── reflex.conf
└── README.md
```

## License

This project is licensed under the MIT License.
````

This `README.md` file provides an overview of the project, setup instructions, API endpoints, and the project structure. It should help users understand how to set up and run the project.