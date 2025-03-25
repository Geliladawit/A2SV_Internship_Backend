# Task Management API

This API provides endpoints for managing tasks. It uses MongoDB for persistent data storage.

## Prerequisites

*   **MongoDB:** A MongoDB instance must be running and accessible.
*   **Go:** Make sure you have Go installed (version 1.16 or later).
*   **Postman:** Recommended for testing the API endpoints.

## Configuration

The API requires a connection string to your MongoDB instance. This connection string should be set as an environment variable named `MONGODB_URI`.

**Setting the `MONGODB_URI` environment variable:**

*   **Linux/macOS:**

    ```bash
    export MONGODB_URI="mongodb+srv://<username>:<password>@<cluster_name>.mongodb.net/?retryWrites=true&w=majority"
    ```

*   **Windows (PowerShell):**

    ```powershell
    $env:MONGODB_URI="mongodb+srv://<username>:<password>@<cluster_name>.mongodb.net/?retryWrites=true&w=majority"
    ```

    Replace `<username>`, `<password>`, and `<cluster_name>` with your MongoDB credentials.
    **Important**: Remember to source your `.env` file to load the MONGODB_URI environment variable if using the `godotenv` library.

## Endpoints

### GET /tasks

*   **Description:** Retrieves a list of all tasks.
*   **Request:**
    *   Method: GET
    *   Headers: `Content-Type: application/json` (optional)
*   **Response:**
    *   Status: 200 OK
    *   Body:

        ```json
        [
            {
                "id": "654321abcdef01234567890",
                "title": "Grocery Shopping",
                "description": "Buy groceries for the week",
                "dueDate": "2024-01-20T18:00:00Z",
                "status": "To Do"
            },
            {
                "id": "987654zyxwvu9876543210",
                "title": "Laundry",
                "description": "Wash and dry clothes",
                "dueDate": "2024-01-21T12:00:00Z",
                "status": "In Progress"
            }
        ]
        ```

### GET /tasks/:id

*   **Description:** Retrieves a specific task by ID.
*   **Request:**
    *   Method: GET
    *   URL: `/tasks/654321abcdef01234567890` (replace `654321abcdef01234567890` with the task ID)
    *   Headers: `Content-Type: application/json` (optional)
*   **Response (Success):**
    *   Status: 200 OK
    *   Body:

        ```json
        {
            "id": "654321abcdef01234567890",
            "title": "Grocery Shopping",
            "description": "Buy groceries for the week",
            "dueDate": "2024-01-20T18:00:00Z",
            "status": "To Do"
        }
        ```

*   **Response (Not Found):**
    *   Status: 404 Not Found
    *   Body:

        ```json
        {
            "error": "task not found"
        }
        ```

### POST /tasks

*   **Description:** Creates a new task.
*   **Request:**
    *   Method: POST
    *   URL: `/tasks`
    *   Headers: `Content-Type: application/json`
    *   Body:

        ```json
        {
            "title": "Pay Bills",
            "description": "Pay all outstanding bills",
            "dueDate": "2024-01-25T17:00:00Z",
            "status": "To Do"
        }
        ```

*   **Response (Success):**
    *   Status: 201 Created
    *   Body:

        ```json
        {
            "id": "fedcba9876543210fedcba98",
            "title": "Pay Bills",
            "description": "Pay all outstanding bills",
            "dueDate": "2024-01-25T17:00:00Z",
            "status": "To Do"
        }
        ```

*   **Response (Bad Request):**
    *   Status: 400 Bad Request
    *   Body:

        ```json
        {
            "error": "Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag"
        }
        ```

### PUT /tasks/:id

*   **Description:** Updates an existing task.
*   **Request:**
    *   Method: PUT
    *   URL: `/tasks/654321abcdef01234567890` (replace `654321abcdef01234567890` with the task ID)
    *   Headers: `Content-Type: application/json`
    *   Body:

        ```json
        {
            "title": "Grocery Shopping",
            "description": "Buy groceries for the week (urgent)",
            "dueDate": "2024-01-20T18:00:00Z",
            "status": "In Progress"
        }
        ```

*   **Response (Success):**
    *   Status: 200 OK
    *   Body:

        ```json
        {
            "id": "654321abcdef01234567890",
            "title": "Grocery Shopping",
            "description": "Buy groceries for the week (urgent)",
            "dueDate": "2024-01-20T18:00:00Z",
            "status": "In Progress"
        }
        ```

*   **Response (Not Found):**
    *   Status: 404 Not Found
    *   Body:

        ```json
        {
            "error": "task not found"
        }
        ```

### DELETE /tasks/:id

*   **Description:** Deletes a task.
*   **Request:**
    *   Method: DELETE
    *   URL: `/tasks/654321abcdef01234567890` (replace `654321abcdef01234567890` with the task ID)
    *   Headers: `Content-Type: application/json` (optional)
*   **Response (Success):**
    *   Status: 204 No Content  (Indicates successful deletion)
*   **Response (Not Found):**
    *   Status: 404 Not Found
    *   Body:

        ```json
        {
            "error": "task not found"
        }
        ```

**Important Notes:**

*   The `id` field in all JSON responses is a string representation of a MongoDB `ObjectID`.
*   Ensure that the `MONGODB_URI` environment variable is set correctly before running the application.
*   Error handling is implemented to provide informative error messages in case of failures.
