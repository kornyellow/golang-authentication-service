# 📡 API Documentation

Base URL: `http://localhost:8080`

## 🔐 Authentication
This API uses **JWT (JSON Web Token)** for authentication.
- **Login** to receive a token.
- **Pass the token** in the `Authorization` header for protected routes.

---

## 1. Register User (IT 02-2)
Create a new user account. Password will be hashed using Bcrypt.

- **URL:** `/api/register`
- **Method:** `POST`
- **Content-Type:** `application/json`

### Request Body
```json
{
  "username": "johndoe",
  "password": "password123",
  "confirm_password": "password123"
}
```

### Response (201 Created)
```json
{
  "message": "User created successfully"
}
```

### Response (400 Bad Request)
```json
{
  "error": "Key: 'RegisterRequest.ConfirmPassword' Error:Field validation for 'ConfirmPassword' failed on the 'eqfield' tag"
}
```

---

## 2. Login (IT 02-1)
Authenticate user and receive a JWT Token.
- URL: `/api/login`
- Method: `POST`
- Content-Type: `application/json`

### Request Body
```json
{
  "username": "johndoe",
  "password": "password123"
}
```

### Response (200 OK)
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "username": "johndoe"
}
```

### Response (401 Unauthorized)
```json
{
  "error": "invalid username or password"
}
```

---

## 3. Get User Profile (IT 02-3)
[Protected Route] Retrieve current user information. Validates JWT signature.
- URL: `/api/profile`
- Method: `GET`
- Headers:
    - `Authorization`: `Bearer <your_token_here>`

### Response (200 OK)
```json
{
  "message": "Welcome to the secret zone!",
  "user": "johndoe"
}
```

### Response (401 Unauthorized)
```json
{
  "error": "Invalid or expired token"
}
```