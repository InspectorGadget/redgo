# RedGo (Gin Gonic + Redis)

## Description
RedGo is a simple web application built using the Gin Gonic framework and Redis. It serves as a demonstration of how to use these technologies together to create a basic web application with caching capabilities.

## Endpoints
- `GET /`: A simple GET endpoint that returns a welcome message.
- `GET /get/:key`: A GET endpoint that retrieves a value from Redis. If the value is not found, it will say so.
- `POST /set`: A POST endpoint that sets a key-value pair in Redis. The key and value are passed as JSON in the request body.
