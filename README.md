# Go CMS Project: My Learning Journey

## Introduction
This project represents my journey in learning the Go programming language by building a simple Content Management System (CMS) API. My goal was to understand Go's core concepts, web development patterns, and best practices by applying them in a real-world project.

## Project Setup

### Prerequisites
- Go 1.23.3 or later
- PostgreSQL database
- Vercel account (for deployment)

### Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd cms-project
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   - Create a `.env` file in the root directory.
   - Add the following variables:
     ```
     DATABASE_URL=postgres://username:password@localhost:5432/dbname
     PORT=8080
     ```

4. Run the project locally:
   ```bash
   go run main.go
   ```

5. Access the API documentation at `http://localhost:8080/swagger/index.html`.

### Deployment
- Deploy the project on Vercel using the provided `vercel.json` configuration.

## What I Learned
- **Go Basics:** Syntax, data types, functions, error handling, and package management.
- **Project Structure:** Organizing code using Go modules, packages, and a clean directory structure (`internal/`, `router/`, etc.).
- **Web Development:**
  - Using the [Gorilla Mux](https://github.com/gorilla/mux) router for HTTP routing and middleware.
  - Implementing RESTful API endpoints for users and navigation links.
  - Handling CORS and HTTP methods.
- **Database Integration:**
  - Connecting to a PostgreSQL database using [sqlx](https://github.com/jmoiron/sqlx).
  - Managing environment variables securely with [godotenv](https://github.com/joho/godotenv).
- **API Documentation:**
  - Generating and serving Swagger documentation with [swaggo](https://github.com/swaggo/swag) and [http-swagger](https://github.com/swaggo/http-swagger).
- **Deployment:**
  - Preparing the project for deployment on Vercel using `vercel.json`.

## Project Structure
- `main.go`: Entry point, server setup, and initialization.
- `internal/`: Business logic, database, and feature modules (users, navLinks).
- `router/`: HTTP routing and middleware.
- `docs/`: Auto-generated Swagger documentation.

## Architectural Decisions

- **Modular Design:** The project is structured into modules (`internal/`, `router/`, etc.) to promote separation of concerns and maintainability.
- **Use of Gorilla Mux:** Chosen for its flexibility and ease of use in handling HTTP routing and middleware.
- **Database Abstraction:** Using `sqlx` for database operations to simplify SQL queries and improve code readability.
- **Environment Configuration:** Utilizing `godotenv` for managing environment variables, ensuring secure and flexible configuration.
- **API Documentation:** Implementing Swagger for auto-generating API documentation, making it easier for developers to understand and use the API.
- **Deployment Strategy:** Using Vercel for deployment, leveraging its serverless architecture for scalability and ease of use.

## Key Takeaways
- Go's simplicity and performance make it ideal for backend APIs.
- Strong typing and clear error handling improve code reliability.
- The Go ecosystem provides robust libraries for web, database, and documentation needs.
- Structuring a Go project with clear separation of concerns is crucial for maintainability.

## Next Steps
- Add authentication and authorization.
- Expand API features (e.g., more CMS entities).
- Write unit and integration tests.
- Explore Go's concurrency features for performance improvements.

## Acknowledgments
- [Go by Example](https://gobyexample.com/)
- [The Go Programming Language Tour](https://tour.golang.org/)
- [Gorilla Mux Documentation](https://github.com/gorilla/mux)
- [Swaggo Documentation](https://github.com/swaggo/swag)

---

This project is a work in progress and a reflection of my ongoing learning in Go. Feedback and suggestions are welcome! 