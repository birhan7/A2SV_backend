# Task Management API Documentation

This document is detailed information on the Task Management REST API implemented using Go and the Gin Framework. The API supports basic CRUD operations for managing tasks. And also unlike the previous task which uses in memory storage, mongoDB is integrated with it.
**Base URL**
http://localhost:9000

**Endpoints**

**1. Create Task**
Method: POST
Path: /tasks
Description: Creates a new task from the json object sent inside the request body.
Request Body:

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "due_date": "YYYY-MM-DDTHH:MM:SSZ",
  "status": "pending|completed|not-done"
}
```

Response:
Status code: 201

```json
{ "message": "Task created" }
```

400 Bad Request: Invalid request body or validation error (e.g., empty title, invalid status)

```json
{ "message": "Task Must have title and status." }
```

**2. Get a Task by ID**
Method: GET
Path: /tasks/:id
Description: Retrieves a task using the id from the request query string..
Response:
200 OK: Task found.

```json
{
  "id": "1",
  "title": "Task 1",
  "description": "First task",
  "due_date": "2025-07-18T00:59:13.7319603-07:00",
  "status": "Pending"
}
```

404 Not Found: Task not found.

```json
{ "message": "Task not found" }
```

**3. Get All Tasks**
Method: GET
Path: /tasks
Description: Retrieves all tasks from the in memory storage used.
Response:
200 OK: List of tasks.

```json
{
  "tasks": [
    {
      "id": "1",
      "title": "Task 1",
      "description": "First task",
      "due_date": "2025-07-18T00:59:13.7319603-07:00",
      "status": "Pending"
    },
    {
      "id": "2",
      "title": "Task 2",
      "description": "Second task",
      "due_date": "2025-07-19T00:59:13.7319603-07:00",
      "status": "In Progress"
    }
  ]
}
```

**4. Update a Task**
Method: PUT
Path: /tasks/:id
Description: Updates an existing task using the id from the query string and values from the provided request body.
Request Body:

```json
{
  "title": "string",
  "description": "string",
  "due_date": "YYYY-MM-DDTHH:MM:SSZ",
  "status": "pending|completed|not-done"
}
```

Response:
200 OK: Task updated successfully.

```json
{ "message": "Task updated" }
```

400 Bad Request: Invalid request body or validation error.

```json
{ "message": "task title cannot be empty" }
```

404 Not Found: Task not found.

```json
{ "message": "Task not found" }
```

**5. Delete a Task**
Method: DELETE
Path: /tasks/:id
Description: Deletes a task by its ID from the query string.
Response:
200 OK

```json
{ "message": "Task removed" }
```

404 Not Found: Task not found.

```json
{ "message": "Task not found" }
```

**Error Handling**
All errors are returned in JSON format with a message field describing the issue.
Common HTTP status codes:
200 OK: Successful request and Successful deletion.
201 Created: Resource created.
400 Bad Request: Invalid input or validation error.
404 Not Found: Resource not found.
**Example Usage**
Create a task:
curl -X POST http://localhost:8080/tasks \
 -H "Content-Type: application/json" \
 -d '{"id":"1","title":"Sample Task","description":"Test task","due_date":"2025-07-18T12:00:00Z","status":"pending"}'
Get a task:
curl http://localhost:9000/tasks/1
