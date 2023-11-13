# simple-api-go
Simple user API developed with Golang to training.

# :point_right: about
This API provides some endpoints to contemplate the methods PUT, POST, DELETE, and GET in order to handle a user list in memory, without a database or any robust implementations.

The objective of this simple project is to know, practice, and train a little progrmaming with the Golang.

# 💻 technologies
- [Gin Web Framework](https://gin-gonic.com/) ;
- [Go Lang version go1.21.3](https://go.dev/) .

# how to run
- Clone this project;
- On the main folder open a terminal an then run:

            go run main.go

- The server will run on `http://localhost:3000/` (Windows) or 0.0.0.0:3000 (Linux).

# endpoints
The available endpoints are:
- GET (simple welcome JSON): http://localhost:3000/
- PUT, DELETE and GET (by id): http://localhost:3000/user/id
- GET (list): http://localhost:3000/users
- POST (with the following JSON): http://localhost:3000/user

            {
                "id": "1",
                "name": "John Doe",
                "nickname": "Johnny",
                "mail": "john@doe.net",
                "cellphone": 55811112222
            }
