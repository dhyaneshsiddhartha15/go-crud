# API Documentation with Swagger

## How Swagger Works in This Project

### In Compass (Complex way):
- Uses `.proto` files (gRPC)
- Auto-generates `swagger.json` from protobuf
- Serves Swagger UI

### In Your Project (Simple way):
- Uses **Swaggo** tool
- Generates `swagger.json` from **code comments**
- No protobuf needed!

## Setup Steps

### 1. Install Swag Tool
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 2. Add Dependencies
```bash
go get -u github.com/swaggo/http-swagger
go get -u github.com/swaggo/files
```

### 3. Add Comments to Your Code

Example in `main.go`:
```go
// @title           CRUD-GO Blog API
// @version         1.0
// @description     Simple REST API for blog posts
// @host            localhost:8080
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

Example in handlers:
```go
// Register godoc
// @Summary Register new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User data"
// @Success 201 {object} UserResponse
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
    // code
}
```

### 4. Generate Docs
```bash
swag init -g cmd/api/main.go
```

This creates `docs/` folder with swagger files.

### 5. View Swagger UI
```
http://localhost:8080/swagger/index.html
```

## Result

You get beautiful interactive API documentation just like Compass!
