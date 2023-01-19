<img src="instruction/chatapp_logo.png" alt="isolated" width="150"/>

# A Web Socket Chat Application - Backend

After running the backend code, please go to https://github.com/Briantam0422/chatapp-frontend to run frontend code.

## Backend

Run the project in the localhost

```jsx
cd cmd
cd main
go run main.go
```

## Database

Database migration has not finished yet. I will update it as soon as possible.

table

- users
    - id int
    - username varchar(255)
    - password varchar(255)
    - token varchar(500)
    - created_at timestamp
    - updated_at timestamp
    - deleted_at timestamp
- chat_records
  - id int
  - message TEXT
  - created_by varchar(255)
  - created_at timestamp
  - updated_at timestamp
  - deleted_at timestamp
- chat_rooms
  - id int
  - name varchar(255)
  - created_by varchar(255)
  - created_at timestamp
  - updated_at timestamp
  - deleted_at timestamp
- schema_migrations
  - version bigint
  - dirty tinyint(1)

### APIs

- Login
- Register
    - Create a new user
- isAuth
    - Check user is authorised
- chat/new
    - For create a new room
- chat/close
    - Close the web socket connection

 

### Middlewares

- CORS
- Authentication
    - Check token
    - Check use authentication
    - Auto login

## Login

![Untitled](instruction/login.png)

## Registration

![Untitled](instruction/register.png)

## Chat Room and Real Time Chat

![Untitled](instruction/chat.png)

1. Create a chat room
2. Share room id to your friends
    1. Copy the room id and send to your friends
3. Join room by id
    1. Paste room id in the “ROOM ID” input
    2. Click connect
    3. Now you can start sending messages
4. Live chat
    1. real time message
5. Multiple user in a room
    1. A chat room supports multiple user in a room