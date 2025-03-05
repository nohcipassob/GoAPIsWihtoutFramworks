Go API Server with Middleware

A simple HTTP API server built using Go's standard library. It includes middleware for authentication and request logging.

Features

Minimalistic HTTP server using net/http

Middleware support for request logging and authentication

Simple route handling with dynamic user IDs

Installation

Clone the repository:

git clone
cd

Run the server:

go run .

API Endpoints

POST /users/{userId}

Description: Returns the user ID provided in the URL.

Example Request:

curl -X POST http://localhost:8080/users/123 -H "Authorization: Bearer token"

Response:

User ID: 123

ANY /users/{userId}

Description: Catch-all route for handling other HTTP methods.

Response:

"CATCH ALL METHOD"

Middleware

Request Logger Middleware

Logs HTTP method and request URI.

Auth Middleware

Checks for a valid Authorization header with Bearer token. If invalid, returns 401 Unauthorized.
