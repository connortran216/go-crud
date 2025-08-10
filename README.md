# Go CRUD - My First RESTful API with Golang

A clean and modern RESTful API built with Go, featuring a service-oriented architecture and complete CRUD operations for blog posts.



## 📋 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/posts` | Create a new post |
| GET | `/posts` | Get all posts |
| GET | `/posts/:id` | Get post by ID |
| PUT | `/posts/:id` | Update entire post |
| PATCH | `/posts/:id` | Partial update post |
| DELETE | `/posts/:id` | Delete post |

## 🏗️ Project Structure

```
go-crud/
├── controllers/           # HTTP request handlers (legacy)
│   └── controlPost.go    
├── services/             # Business logic layer
│   └── post_service.go   
├── viewsets/             # Generic API viewsets
│   ├── base.go          # Generic CRUD operations
│   └── post_viewset.go  # Post-specific viewset
├── models/              # Data models
│   └── postModel.go     
├── initializers/        # Application initialization
│   ├── initPostgres.go  # Database connection
│   └── loadEnvVariables.go # Environment loader
├── migration/           # Database migration
│   └── migration.go     
├── examples/            # Usage examples
│   └── user_example.go  
├── main.go             # Application entry point
├── go.mod              # Go module file
└── go.sum              # Dependency checksums
```

### Architecture Layers

1. **Models** (`models/`): Define data structures and database schemas
2. **Services** (`services/`): Contain business logic and data validation
3. **ViewSets** (`viewsets/`): Handle HTTP requests/responses with generic CRUD operations
4. **Controllers** (`controllers/`): Legacy layer, being replaced by ViewSets
5. **Initializers** (`initializers/`): Handle app startup and configuration

## 🛠️ Technologies Used

- **Go 1.24.5** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM library for Go
- **PostgreSQL** - Database
- **godotenv** - Environment variable management

## ⚡ Getting Started

### Prerequisites

- Go 1.24+ installed
- PostgreSQL database running
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd go-crud
   ```

2. **Set up environment variables**
   ```bash
   # Create .env file in root directory
   touch .env
   ```
   
   Add your database configuration:
   ```env
   DB_DSN=postgres://username:password@localhost:5432/database_name?sslmode=disable
   ```

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Run database migration**
   ```bash
   go run migration/migration.go
   ```

5. **Start the server**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`

## 📖 How to Generate go.mod and go.sum

### Creating go.mod

The `go.mod` file is the heart of Go modules. Here's how it's generated:

1. **Initialize a new module**:
   ```bash
   go mod init go-crud
   ```
   This creates a `go.mod` file with the module name.

2. **Add dependencies**:
   When you import packages in your Go files, use:
   ```bash
   go mod tidy
   ```
   This automatically adds required dependencies to `go.mod`.

3. **Manual dependency addition**:
   ```bash
   go get github.com/gin-gonic/gin
   go get gorm.io/gorm
   go get gorm.io/driver/postgres
   ```

### Understanding go.sum

The `go.sum` file contains cryptographic checksums of module dependencies:

- **Automatically generated** when you run `go mod tidy` or `go build`
- **Ensures integrity** of downloaded modules
- **Version verification** - prevents tampering
- **Should be committed** to version control

**Key commands**:
```bash
go mod tidy        # Add missing and remove unused modules
go mod verify      # Verify dependencies match go.sum
go mod download    # Download modules to local cache
```


## 🎯 Key Features Explained

### Generic ViewSet Pattern
The project uses a generic ViewSet pattern that allows for reusable CRUD operations:

- **BaseViewSet**: Provides generic CRUD operations for any model
- **Type Safety**: Uses Go generics for compile-time type checking
- **Standardized Responses**: Consistent API response format
- **Easy Extension**: Simple to add custom endpoints

### Service Layer
Clean separation of business logic:

- **Validation**: Input validation and business rules
- **Error Handling**: Proper error messages and HTTP status codes
- **Database Operations**: Abstracted database interactions
- **Reusability**: Services can be used across different controllers

### Database Integration
- **GORM ORM**: Powerful and developer-friendly ORM
- **Auto-migration**: Automatic database schema updates
- **Connection Pooling**: Efficient database connection management
- **Environment Configuration**: Database settings from environment variables



## 🔮 Future Enhancements

- [ ] Add authentication and authorization
- [ ] Implement pagination for list endpoints
- [ ] Add request validation middleware
- [ ] Include API documentation with Swagger
- [ ] Add unit and integration tests
- [ ] Implement logging middleware
- [ ] Add rate limiting
- [ ] Docker containerization

## 📝 License

This project is for learning purposes.

---
*Built with ❤️ using Go and modern software architecture principles*
