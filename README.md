# User Management API

This project provides a RESTful API for managing user data, including fetching, creating, reading, updating, and deleting users. The API uses Gin (a Go web framework) and GORM for ORM with a MySQL database.

---

## Features

- Fetch data from an external API (JSONPlaceholder) and save it to the database.
- Perform CRUD (Create, Read, Update, Delete) operations on user data.
- Built using Go, Gin, and GORM.

---

## Technologies Used

- **Language**: Go
- **Framework**: Gin (Go Web Framework)
- **Database**: MySQL
- **ORM**: GORM
- **External Data Source**: JSONPlaceholder

---

## API Endpoints

### Fetch and Save Users
- **URL**: `/api/fetch-and-save-users`
- **Method**: `GET`
- **Description**: Fetches data from JSONPlaceholder and saves it to the database.

### Get All Users
- **URL**: `/api/users`
- **Method**: `GET`
- **Description**: Retrieves all users from the database.
- **Sukses**:  
  ```json
  {
   "users": [
    {
      "id": 1,
      "userId": 1,
      "title": "Example title",
      "body": "Example body"
    }
   ]
  }

### Get User by ID
- **URL**: `/api/user/:id`
- **Method**: `GET`
- **Description**: Retrieves a user by their ID.
- **Sukses**:  
  ```json
    {
      "id": 1,
      "userId": 1,
      "title": "Example title",
      "body": "Example body"
    }


### Create a New User
- **URL**: `/api/user`
- **Method**: `POST`
- **Description**: Creates a new user in the database.
- **Sukses**:  
  ```json
    {
      "userId": 1,
      "title": "Example title",
      "body": "Example body"
    }

### Update a User
- **URL**: `/api/user/:id`
- **Method**: `PUT`
- **Description**: Updates an existing user by ID.
- **Sukses**:  
  ```json
    {
      "userId": 1,
      "title": "Example title",
      "body": "Example body"
    }

### Delete a User
- **URL**: `/api/user/:id`
- **Method**: `DELETE`
- **Description**: Deletes a user by their ID.
- **Sukses**:  
  ```json
    {
      "message": "User deleted successfully"
    }

---

**Link GPT**
https://chatgpt.com/share/674d892e-7234-8004-8614-3dc6d1871e9f
