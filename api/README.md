# Devbook (WIP)

**Devbook** is a simple social network backend implemented in Golang. The goal is to provide a minimal, educational example of user management and authentication using Go, with clean code and modular structure.

## Features (Current)
- User registration
- User login (with password verification)
- User search by name or nickname
- User update and deletion
- Basic password hashing
- RESTful API endpoints

## Important Notes
- **Security:** Password handling is basic and for demonstration only. Do not use this implementation for production. Use strong password hashing and validation in real applications.
- **Database:** Avoid using the `root` user for your database. Create a dedicated, limited-privilege user for your application.

## How to Use

### Prerequisites
- Go 1.18+
- MySQL database (or compatible)

### Setup
1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd devbook/api
   ```
2. Configure your database connection in `src/config/config.go`.
3. Run the SQL script in `sql/sql.sql` to create the required tables.
4. Build and run the API:
   ```bash
   go run main.go
   ```

### API Endpoints
- `POST /login` — User login
- `POST /users` — Register new user
- `GET /users?search=...` — Search users by name or nickname
- `PUT /users/:id` — Update user
- `DELETE /users/:id` — Delete user

## Project Structure
```
go.mod
main.go
sql/
  sql.sql
src/
  config/
  controllers/
  database/
  model/
  repository/
  response/
  router/
  security/
```

## License
MIT (see LICENSE)

---
This project is under active development. Contributions and feedback are welcome!