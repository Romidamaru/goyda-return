## Goyda-Return Simple Task Manager

Goyda-Return is a simple task management API built with Go, Gin, and GORM. It allows users to create, update, retrieve, and delete tasks. This project demonstrates best practices in building a RESTful API using Go, with a modular architecture and user-task integration.

---

## Features

- **Task Management**: Create, read, update, and delete tasks.
- **Task Types**: Categorize tasks as _personal_, _work_, or _other_.
- **User Integration**: Each task is associated with a user.
- **Validation**: Ensures proper validation of request payloads.
- **Database Integration**: Uses GORM for seamless database interactions.
- **Modular Architecture**: Built with a modular structure to easily add new features.

---

## Prerequisites

Ensure you have the following installed:

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
   - Update the database configuration in `internal/config/config.go` to match your database credentials.

4. Start the server:
   ```bash
   go run main.go
   ```

---

## API Endpoints

### Task Routes

| Method   | Endpoint        | Description               |
|----------|-----------------|---------------------------|
| `GET`    | `/tasks`        | Retrieve all tasks        |
| `GET`    | `/tasks/:id`    | Retrieve a task by ID     |
| `POST`   | `/tasks`        | Create a new task         |
| `PUT`    | `/tasks/:id`    | Update a task             |
| `DELETE` | `/tasks/:id`    | Delete a task             |

### User Routes

| Method   | Endpoint        | Description               |
|----------|-----------------|---------------------------|
| `POST`   | `/users`        | Create a new user         |

### Task Payloads

#### Create Task

```json
{
  "name": "Buy groceries",
  "type": "personal",
  "done": false,
  "user_id": 1
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

### User Payloads

#### Create User

```json
{
  "username": "johndoe",
  "email": "johndoe@example.com",
  "password": "securepassword"
}
```

---

## Project Structure

```
goyda-return/
│
├── internal/                     # Internal modules
│   ├── config/                   # Configuration files
│   │   └── config.go             # Database and app configuration
│   ├── modules/                  # Modules for tasks and users
│   │   ├── core/                 # Core utilities
│   │   ├── tasks/                # Task module (controller, service, model, etc.)
│   │   │   ├── ctrl/             # Task controllers
│   │   │   ├── dto/              # Task DTOs
│   │   │   ├── ent/              # Task entities (GORM models)
│   │   │   ├── svc/              # Task services
│   │   │   └── router.task.go    # Task routes
│   │   ├── users/                # User module (controller, service, model, etc.)
│   │   │   ├── ctrl/             # User controllers
│   │   │   ├── dto/              # User DTOs
│   │   │   ├── ent/              # User entities (GORM models)
│   │   │   ├── svc/              # User services
│   │   │   └── router.user.go    # User routes
│   ├── utils/                    # Utility files
│   │   └── validator.go          # Validation logic
├── main.go                       # Entry point for the application
├── README.md                     # Project documentation
```

---

## Contributing

Contributions are welcome! Please fork the repository, make your changes, and submit a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

This updated README should now align with your project's structure, reflecting the correct paths and modular setup. Let me know if you'd like to make any further adjustments!