# API Documentation – User Management

Base URL: `http://localhost:8080`

### Create User

**Endpoint:** `/users`  
**Method:** `POST`  
**Description:** Creates a new user.

## Request Body
```
{
  "name": "Alice",
  "email": "alice@example.com"
}
```
## Response Body
```
{
	"id": "5580f306-bd36-4c7c-a6e5-d99aa7263f5d",
	"name": "Alice",
	"email": "alice@example.com"
}
```

### Get User

**Endpoint:** `/users`  
**Method:** `GET`  
**Description:** Retrieves all users.

## Response Body
```
[
	{
		"id": "5580f306-bd36-4c7c-a6e5-d99aa7263f5d",
		"name": "Alice",
		"email": "alice@example.com"
	}
]
```

### Update User

**Endpoint:** `/users/update`  
**Method:** `PUT`  
**Description:** Updates a user by id.

## Request Body
```
{
	"id": "5580f306-bd36-4c7c-a6e5-d99aa7263f5d",
	"name": "Alice Whitman",
	"email": "alice@example.com"
}
```
## Response Body
```
{
	"id": "5580f306-bd36-4c7c-a6e5-d99aa7263f5d",
	"name": "Alice whitman",
	"email": "alice@example.com"
}
```

### Delete User

**Endpoint:** `/users/delete`  
**Method:** `DELETE`  
**Description:** Deletes a user by id.

## Request Body
```
{
	"id": "5580f306-bd36-4c7c-a6e5-d99aa7263f5d",
}
```
## Response Body
```
No body returned for response
```