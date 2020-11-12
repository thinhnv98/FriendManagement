# S3_FriendManagement_ThinhNguyen

##Build and run App
```
docker-compose build
docker-compose up
```
##APIs

###Create an email
```http request
POST /user
```

- Request body
```json
{
    "Email":"abc@example.com"
}
```

- Response body:
```json
{
    "Success": true
}
```

###Create friend connection
```http request
POST /friend
```

- Request body:
```json
{ 
    "friends": [
        "andy@example.com",
        "john@example.com"
    ]
}
```

- Response body:
```json
{ 
    "success": "true"
}
```

### Get friend list for an email address
```http request
GET /friend/friends
```

- Request body:
```json
{ 
    "email": "andy@example.com"
}
```

- Response body:
```json
{ 
    "success": "true",
    "friends": [
        "john@example.com"
    ],
    "count" : 1
}
```

### Get common friend list between two email addresses
```http request
GET /friend/common-friends
```

- Request body:
```json
{ 
    "friends": [
        "andy@example.com",
        "john@example.com"
    ]
}
```

- Response body:
```json
{ 
    "success": "true",
    "friends": [
        "common@example.com"
    ],
    "count" : 1
}
```

### Subscribe to update from an email address
```http request
POST /subscription
```

- Request body:
```json
{
  "requestor": "lisa@example.com",
  "target": "john@example.com"
}
```

- Response body:
```json
{ 
    "success": "true"
}
```


### Block update from an email address
```http request
POST /block
```

- Request body:
```json
{
  "requestor": "andy@example.com",
  "target": "john@example.com"
}
```

- Response body:
```json
{ 
    "success": "true"
}
```


### Retrieve all email addresses which can receive update from an email address
```http request
GET /friend/emails-receive-update
```

- Request body:
```json
{
  "sender": "john@example.com",
  "text": "Hello World! kate@example.com"
}
```

- Response body:
```json
{ 
    "success": "true",
    "recipients": [
        "lisa@example.com",
        "kate@example.com"
    ]
}
```

## Project architecture
- Workflow: Request => Handlers => Services => Repositories => Database

- Three layers model:
    + Handlers: Get request from httpRequest, decode, validate, call services, write httpResponse
    + Services: Handle business logic, call repositories
    + Repositories: Data access layer 
     