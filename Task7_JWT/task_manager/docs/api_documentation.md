## Authentication and Authorization

This API uses JSON Web Tokens (JWT) for authentication and authorization.  This ensures that only authorized users can access protected resources.

### User Registration

To create a new user account, send a `POST` request to the `/register` endpoint.

**Request:**


POST /register
Content-Type: application/json

```json
{
  "username": "your_username",
  "password": "your_password"
}

Response (Success - 201 Created):

{
  "message": "User registered successfully"
}


Response (Error - 400 Bad Request):

Indicates an issue with the request body (e.g., missing fields, invalid data types). The response will contain an error message.

{
  "error": "Invalid request body: username is required"
}


Response (Error - 500 Internal Server Error):

Indicates a server-side error during user creation.

{
  "error": "Failed to create user"
}
User Login

To authenticate an existing user and obtain a JWT token, send a POST request to the /login endpoint.

Request:

POST /login
Content-Type: application/json

{
  "username": "your_username",
  "password": "your_password"
}


Response (Success - 200 OK):

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJyb2xlIjoidXNlciIsImV4cCI6MTcwMzkwNTA1NH0.aJ09-88N-Y2aN0c3NnI8q28-WqXqQ03t_a7iHhJ7x8"
}


The token field contains the JWT token. Store this token securely on the client-side (e.g., in local storage or a cookie).

Response (Error - 401 Unauthorized):

Indicates invalid username or password.

{
  "error": "Invalid credentials"
}


Response (Error - 500 Internal Server Error):

Indicates a server-side error during authentication.

{
  "error": "Failed to retrieve user"
}


To access protected API endpoints, you must include the JWT token in the Authorization header of your HTTP requests. Use the Bearer authentication scheme.

Example Request:

GET /tasks
Authorization: Bearer <your_jwt_token>
Content-Type: application/json


Replace <your_jwt_token> with the actual JWT token you received during login.

Without a valid JWT token in the Authorization header, protected endpoints will return a 401 Unauthorized error.

User Roles and Authorization

This API uses role-based access control (RBAC) to restrict access to certain endpoints based on user roles. The following roles are defined:

admin: Has full access to all API endpoints.

user: Has limited access to API endpoints. Can typically read data but cannot create, update, or delete.

Endpoint Permissions:

Endpoint	Method	Role(s) Required	Description
/tasks	POST	admin	Creates a new task. Only administrators can create tasks.
/tasks	GET	user, admin	Retrieves a list of all tasks. All authenticated users (including administrators) can access this endpoint.
/tasks/:id	GET	user, admin	Retrieves a specific task by ID. All authenticated users can access this endpoint.
/tasks/:id	PUT	admin	Updates an existing task. Only administrators can update tasks.
/tasks/:id	DELETE	admin	Deletes an existing task. Only administrators can delete tasks.
/promote/:id	PUT	admin	Promotes a standard user to Admin. Only admins can promote users.

Authorization Errors:

If you attempt to access an endpoint without the required role, you will receive a 403 Forbidden error.

{
  "error": "Insufficient permissions"
}


Important Notes:

Keep your JWT token secret. Do not share it with anyone.

Store your JWT token securely on the client-side.

Always use HTTPS to protect your data in transit.


