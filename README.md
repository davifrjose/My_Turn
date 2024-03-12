# My_Turn

## Getting Started

1. If you do not use devcontainer, ensure you have [Go](https://go.dev/dl/) 1.23 or higher and [Task](https://taskfile.dev/installation/) installed on your machine:
[Migrate](https://github.com/golang-migrate/migrate/tree/v4.17.0/cmd/migrate)

    ```bash
    go version && task --version
    ```

2. Create a copy of the `.env.example` file and rename it to `.env`:

    ```bash
    cp .env.example .env
    ```

    Update configuration values as needed.

3. Install all dependencies, run docker compose, create database schema, and run database migrations:

    ```bash
    task
    ```

4. Run the project in development mode:

    ```bash
    task dev
    ```
