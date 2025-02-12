# Golang JWT Authentication with Redis

This project demonstrates how to build a **JWT-based authentication system** using Golang. The application utilizes various modern technologies and tools to create a secure and efficient authentication mechanism.

### Tech Stack:

- **[Fiber](https://gofiber.io/):** A fast, lightweight, and efficient web framework for Golang.
- **[GORM](https://gorm.io/):** An ORM library for interacting with a PostgreSQL database.
- **[JWT](https://jwt.io/):** JSON Web Token for secure user authentication and session management.
- **[Redis](https://redis.io/):** An in-memory key-value store used for caching and managing tokens.
- **[PostgreSQL](https://www.postgresql.org/):** A powerful, open-source relational database for storing user data.
- **[Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt):** A secure hashing algorithm for password encryption.

### External Components:

- **Redis:** Used as a cache to store tokens for quick verification and logout.
- **PostgreSQL:** The primary database for persisting user credentials and data.
