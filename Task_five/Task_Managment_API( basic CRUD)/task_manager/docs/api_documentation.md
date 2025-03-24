# Task Management API

This API provides endpoints for managing tasks.

## Endpoints

### GET /tasks

*   **Description:**  Retrieves a list of all tasks.
*   **Request:**
    *   Method: GET
    *   Headers: `Content-Type: application/json` (optional)
*   **Response:**
    *   Status: 200 OK
    *   Body:
        ```json
        [
            {
                "id": 1,
                "title": "Grocery Shopping",
                "description": "Buy groceries for the week",
                "dueDate": "2024-01-20T18:00:00Z",
                "status": "To Do"
            },
            {
                "id": 2,
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
    *   URL: `/tasks/1` (replace `1` with the task ID)
    *   Headers: `Content-Type: application/json` (optional)
*   **Response (Success):**
    *   Status: 200 OK
    *   Body:
        ```json
        {
            "id": 1,
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
            "id": 3,
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
    *   URL: `/tasks/1` (replace `1` with the task ID)
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
            "id": 1,
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
    *   URL: `/tasks/1` (replace `1` with the task ID)
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
