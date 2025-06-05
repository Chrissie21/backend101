# Expense Tracker API

A robust backend for an expense tracking application built with Go, Gin, GORM, PostgreSQL, and JWT-based authentication. This API allows users to register, log in, manage transactions, and calculate their financial balance securely.

## Features

-   **User Authentication**:
    
    -   Register with name, email, and password (hashed with bcrypt).
    -   Login to receive a JWT token for secure access.
    -   Protected routes using JWT middleware to ensure only authenticated users access their data.
-   **Transaction Management**:
    
    -   Create, read, update, and delete transactions (income or expense).
    -   Transactions are tied to the authenticated user, ensuring data privacy.
    -   Input validation to enforce correct data formats (e.g., positive amounts, valid transaction types).
-   **Financial Insights**:
    
    -   Calculate total income, total expenses, and net balance.
    -   Indicate whether the user is in a positive or negative financial zone.
-   **Tech Stack**:
    
    -   **Go**: Backend programming language.
    -   **Gin**: Web framework for routing and HTTP handling.
    -   **GORM**: ORM for PostgreSQL database interactions.
    -   **PostgreSQL**: Database for storing users and transactions.
    -   **JWT**: Secure authentication with token-based access.
    -   **Viper & godotenv**: Configuration management with .env files.
    -   **go-playground/validator**: Input validation for robust data integrity.
    -   **Redis**: (Optional, included for future caching or session management).
    -   **Swagger**: (Planned for API documentation, included in dependencies).

## Project Structure

```
backend101/
├── cmd/                    # Application entry point
│   └── main.go
├── config/                 # Configuration loading (Viper, godotenv)
│   └── config.go
├── controllers/            # API endpoint logic
│   ├── auth_controller.go
│   ├── transaction_controller.go
│   └── user_controller.go
├── database/               # Database connection and setup
│   └── postgres.go
├── middleware/             # Middleware for JWT authentication
│   └── auth_middleware.go
├── models/                 # Data models and structs
│   ├── auth.go
│   ├── transaction.go
│   └── user.go
├── routes/                 # API route definitions
│   ├── auth_routes.go
│   ├── transaction_routes.go
│   └── user_routes.go
├── services/               # Business logic (e.g., password hashing, JWT generation)
│   ├── auth_service.go
│   └── jwt_service.go
├── utils/                  # Utility functions (e.g., validation)
│   └── validator.go
├── .env                    # Environment variables
├── go.mod                  # Go module dependencies
└── README.md               # Project documentation

```

## Prerequisites

-   **Go**: Version 1.18 or higher.
-   **PostgreSQL**: Running locally or on a server (e.g., version 13 or higher).
-   **Redis**: (Optional, for future features).
-   A code editor (e.g., VS Code).
-   Tools like Postman or Insomnia for testing API endpoints.

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/backend101.git
cd backend101

```

Replace `yourusername` with your GitHub username.

### 2. Initialize Go Module

```bash
go mod init github.com/yourusername/backend101

```

### 3. Install Dependencies

Install the required Go packages:

```bash
go get github.com/gin-gonic/gin \
       github.com/golang-jwt/jwt/v5 \
       github.com/joho/godotenv \
       github.com/spf13/viper \
       github.com/sirupsen/logrus \
       github.com/go-playground/validator/v10 \
       github.com/swaggo/swag/cmd/swag \
       github.com/swaggo/gin-swagger \
       github.com/swaggo/files \
       gorm.io/gorm \
       gorm.io/driver/postgres \
       github.com/redis/go-redis/v9 \
       golang.org/x/crypto/bcrypt

```

### 4. Create the `.env` File

At the project root, create a `.env` file with the following content:

```env
PORT=8080
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=expense_tracker
DB_PORT=5432
JWT_SECRET=your_super_secret_key
JWT_EXPIRE_HOURS=24
REDIS_ADDR=localhost:6379
ENV=development

```

Replace `yourpassword` and `your_super_secret_key` with secure values.

### 5. Set Up PostgreSQL

Ensure PostgreSQL is running. Create a database named `expense_tracker`:

```bash
createdb expense_tracker

```

Alternatively, use `psql` or a GUI like PgAdmin to create the database.

### 6. Run the Application

Start the server:

```bash
go run cmd/main.go

```

The server will run on `http://localhost:8080` (or the port specified in `.env`).

### 7. Test the API

Use Postman, Insomnia, or `curl` to test the endpoints. Example:

```bash
curl http://localhost:8080/

```

Expected response:

```json
{
  "message": "Welcome to Expense Tracker API 🚀"
}

```

## API Endpoints

### Authentication

-   **POST /api/auth/register**
    
    -   Register a new user.
    -   Request body:
        
        ```json
        {
          "name": "Somebody Someone",
          "email": "somebody@some_e-mail.com",
          "password": "mypassword123"
        }
        
        ```
        
    -   Response: `201 Created` with `{ "message": "User registered successfully" }`.
-   **POST /api/auth/login**
    
    -   Log in to receive a JWT token.
    -   Request body:
        
        ```json
        {
          "email": "someone@example.com",
          "password": "mypassword123"
        }
        
        ```
        
    -   Response: `200 OK` with `{ "token": "eyJhbGciOiJIUzI..." }`.

### User

-   **GET /api/user/me** (Protected)
    -   Get the authenticated user's ID.
    -   Headers: `Authorization: Bearer <your_token>`
    -   Response: `200 OK` with `{ "message": "You are authenticated", "user_id": 1 }`.

### Transactions

-   **POST /api/transactions** (Protected)
    
    -   Create a new transaction.
    -   Headers: `Authorization: Bearer <your_token>`
    -   Request body:
        
        ```json
        {
          "amount": 150.50,
          "category": "Food",
          "description": "Lunch at cafe",
          "type": "expense"
        }
        
        ```
        
    -   Response: `200 OK` with the created transaction.
-   **GET /api/transactions** (Protected)
    
    -   Retrieve all transactions for the authenticated user.
    -   Headers: `Authorization: Bearer <your_token>`
    -   Response: `200 OK` with an array of transactions.
-   **PUT /api/transactions/:id** (Protected)
    
    -   Update a transaction (owned by the authenticated user).
    -   Headers: `Authorization: Bearer <your_token>`
    -   Request body:
        
        ```json
        {
          "amount": 220.0,
          "category": "Transport",
          "description": "Fuel refill",
          "type": "expense",
          "date": "2025-05-17T00:00:00Z"
        }
        
        ```
        
    -   Response: `200 OK` with the updated transaction.
-   **DELETE /api/transactions/:id** (Protected)
    
    -   Delete a transaction (owned by the authenticated user).
    -   Headers: `Authorization: Bearer <your_token>`
    -   Response: `200 OK` with `{ "message": "Transaction deleted" }`.
-   **GET /api/transactions/balance** (Protected)
    
    -   Calculate total income, expenses, balance, and financial status.
    -   Headers: `Authorization: Bearer <your_token>`
    -   Response: `200 OK` with:
        
        ```json
        {
          "income_total": 1200.00,
          "expense_total": 950.00,
          "balance": 250.00,
          "financial_zone": "positive"
        }
        
        ```
        

## Input Validation

The API uses `go-playground/validator` to enforce:

-   Positive transaction amounts (`amount > 0`).
-   Required fields for categories, descriptions, and transaction types.
-   Transaction types limited to `income` or `expense`.
-   Minimum/maximum length for certain fields (e.g., category: 2–30 characters).

Example error response for invalid input:

```json
{
  "validation_errors": {
    "Amount": "gt",
    "Category": "required",
    "Type": "oneof",
    "Description": "required",
    "Date": "required"
  }
}

```

## Future Enhancements

-   **Monthly Breakdown**: Filter transactions by month/year.
-   **Category Insights**: Generate spending summaries by category.
-   **Swagger Docs**: Auto-generate API documentation using `swaggo`. (Sort of Implemented already.)
-   **Redis Integration**: Add caching for frequent queries (e.g., balance).
-   **AI/ML Features**: Analyze spending habits and provide saving tips.

## Contributing

1.  Fork the repository.
2.  Create a feature branch (`git checkout -b feature-name`).
3.  Commit your changes (`git commit -m 'Add feature'`).
4.  Push to the branch (`git push origin feature-name`).
5.  Create a pull request.


