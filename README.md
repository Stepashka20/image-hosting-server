# Image hosting backend

This is a backend service for hosting images, built using Golang and the Gin framework.

## Getting Started
1. Clone the repository:
```
git clone https://github.com/Stepashka20/image-hosting-server
```
2. Install dependencies:
```
go get
```
3. Fill .env file
```
cp .env.example .env
```
4. Set up a MongoDB database: create `imagecloud` db and `image_groups`, `images` collections
5. Start the server:
```
go run main.go
```
## Frontend
The frontend for this project can be found [here](https://github.com/Stepashka20/image-hosting)

## Built With

- [Golang](https://golang.org/)
- [Gin](https://gin-gonic.com/)

## License
This project is licensed under the MIT License.
