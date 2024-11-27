## Goyda-Return Simple Task Manager

Goyda-Return is a simple task management API built with Go, Gin, and GORM. It allows users to create, update, retrieve, and delete tasks. This project is designed to demonstrate best practices in building a RESTful API using Go.

---

## Features

- **Task Management**: Create, read, update, and delete tasks.
- **Task Types**: Supports categorization of tasks into _personal_, _work_, and _other_.
- **Validation**: Ensures proper validation of request payloads.
- **Database Integration**: Built with GORM for seamless database operations.
- **Extensible Architecture**: Follows a modular structure to easily add new features.

---

## Prerequisites

Make sure you have the following installed on your system:

- [Go](https://golang.org/) (Version 1.20+ recommended)
- [PostgreSQL](https://www.postgresql.org/) or your preferred database
- [Git](https://git-scm.com/)

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/goyda-return.git
   cd goyda-return
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure the database:
    - Update the database configuration in `config/config.go` to match your database credentials.

4. Run database migrations:
   ```bash
   go run migrations/migrate.go
   ```

5. Start the server:
   ```bash
   go run main.go
   ```

---

## API Endpoints

### Task Routes

| Method | Endpoint        | Description             |
|--------|-----------------|-------------------------|
| `GET`  | `/tasks`        | Retrieve all tasks     |
| `GET`  | `/tasks/:id`    | Retrieve a task by ID  |
| `POST` | `/tasks`        | Create a new task      |
| `PUT`  | `/tasks/:id`    | Update a task          |
| `DELETE` | `/tasks/:id` | Delete a task          |

### Task Payloads

#### Create Task

```json
{
  "name": "Buy groceries",
  "type": "personal",
  "done": false
}
```

#### Update Task

```json
{
  "name": "Buy groceries and cook",
  "type": "work",
  "done": true
}
```

---

## Project Structure

```
goyda-return/
│
├── config/             # Configuration files
├── internal/           # Internal modules
│   ├── modules/        # Task module (controller, service, model, etc.)
│   ├── dto/            # Data Transfer Objects
│   ├── ent/            # Entity definitions
├── migrations/         # Database migrations
├── main.go             # Entry point
├── README.md           # Project documentation
```

---

## Contributing

Contributions are welcome! Please fork the repository, make your changes, and submit a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.