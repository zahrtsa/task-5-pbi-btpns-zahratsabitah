# Api using Golang

The task of creating a simple API using Golang, with the feature that users cand add profile photo if they are logged in, using "project-api" for location project initialization.

## Features 🧑‍💻

- Users can add a profile photo
- The system can identify users(log in/sign up)
- Only users who logged in/signed up can delete/add profile photos
- Users can delete images that have been posted
- Different users can't delete/change photos that have been crated by other users

## Install module/package

```go
go get -u github.com/asaskevich/govalidator
go get -u github.com/gin-gonic/gin
go get -u github.com/go-sql-driver/mysql
go get -u github.com/golang-jwt/jwt/v5 
go get -u github.com/jinzhu/gorm 
go get -u github.com/joho/godotenv 
go get -u golang.org/x/crypto 
go get -u gorm.io/gorm 
```

## Structure 📋

```php
project-api 
 ┣ controllers  // Folder of all method 
 ┃ ┣ photoController.go
 ┃ ┗ userController.go
 ┣ database
 ┃ ┣ configEnv.go // Load env variable
 ┃ ┗ connectDB.go // Connect database & migrate
 ┣ middleware
 ┃ ┗ requireAuth.go // Validate user & claims token
 ┣ models  // All models
 ┃ ┣ photo.go 
 ┃ ┗ user.go
 ┣ router
 ┃ ┗ route.go // Call all method from folder controllers & set the router
 ┣ .env // File to save .env variable
 ┣ .env.example // Example of .env
 ┣ go.mod // All of module
 ┣ go.sum
 ┣ main.go // Main project, run this app for run project
 ┣ project-api.exe // I'm using compiledaemon that's why this file was created
 ┣ project-api.exe~
 ┗ readme.md
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

#### Database Config

```bash
'DB_HOST'
'DB_DRIVER'
'DB_USER'
'DB_PASSWORD'
'DB_NAME'
'DB_PORT'
```

#### Gin http running on port

```bash
'PORT'
```

#### JWT Token

```bash
'SECRET'
```

## API Reference

#### Create User
```
 POST /users/register 
```
```json
{
  "username": "myusername",
  "email": "myemail@gmail.com",
  "password": "mypassword"
}
```

#### Login

```
  POST /users/login
```

```json
{
  "email": "myemail@gmail.com",
  "password": "mypassword"
}
```

#### Validate user

```
  GET /users/validate
```

```json
{
  "User who logged is ": "2"
}
```

#### Update user

```
  PUT /users/update/{id}
```

```json
{
  "username": "new.myusername",
  "email": "new.myemail@gmail.com",
  "password": "new.mypassword"
}
```

#### Delete user

```
  DELETE /users/delete/{id}
```

#### Show All User

```
  GET /users/showAll
```

```json
{
  "message": [
    {
      "ID": 2,
      "CreatedAt": "2023-09-27T13:39:39+07:00",
      "UpdatedAt": "2023-09-27T13:39:39+07:00",
      "DeletedAt": null,
      "username": "myusername",
      "email": "myemail@gmail.com",
      "password": "$2a$10$Q2IQ2ddNMtrVOTJHsocmzOo4peFcpVOJ6y9LnNWxrfu0Pu001FP3u",
      "photo": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "title": "",
        "caption": "",
        "photo_url": "",
        "users_id": 0
      }
    },
    {
      "ID": 3,
      "CreatedAt": "2023-09-27T13:40:25+07:00",
      "UpdatedAt": "2023-09-27T13:40:25+07:00",
      "DeletedAt": null,
      "username": "myusername2",
      "email": "myemail2@gmail.com",
      "password": "$2a$10$P3iaJkgbduYR7s.btXH9QOMPJLXuXr/Zno0gn5RkL7bCZ8bECbSNW",
      "photo": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "title": "",
        "caption": "",
        "photo_url": "",
        "users_id": 0
      }
    }
  ]
}
```

---

#### Post photos

```
  POST /photo/post
```

```json
{
  "title": "myphoto",
  "caption": "#thisismyphoto",
  "photo_url": "http://example.com"
}
```

#### Update photos

```
  POST /photo/update/{id}
```

```json
{
  "title": "newmyphoto",
  "caption": "#newpostphotos",
  "photo_url": "http://example.com"
}
```

#### Delete photos

```
  DELETE /photo/delete/{id}
```

#### Show all photos

```
  GET /photo/show
```

```json
{
  "message": [
    {
      "ID": 1,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "DeletedAt": null,
      "title": "newmyphoto",
      "caption": "#newpostphotos",
      "photo_url": "http://example.com",
      "users_id": 2
    },
    {
      "ID": 2,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "DeletedAt": null,
      "title": "myphoto2",
      "caption": "postphotos2",
      "photo_url": "http://example2.com",
      "users_id": 3
    }
  ]
}
```

## Tools 🛠

[![My Skills](https://skillicons.dev/icons?i=go,git,mysql,postman,vscode)](https://skillicons.dev)
