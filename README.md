# go-url-encounter
A simple URL encounter.

## Dependencies

- Docker
- Go

## Running


```bash
git clone https://github.com/lu-css/go-url-encounter # Clone the repository.

sudo docker-compose up -d # Start postgres in docker.

go run . # Finally run the program.
```

The program will run on the port `3000`.!

## Usage

- Send a POST to `localhost8080` with an `url` field with your URL.

A POST example.
```bash
curl -X POST http://localhost:8080/ -d "{\"url\": \"https://github.com/lu-css/\"}" -H "Content-Type: application/json"
```

- To access your shortened URL, use `localhost:8080/<your-url-code>`, and you will be automatically redirected for your URL.

## How it Works

Randonly generates URL-Safe characters, the characters are `abcdefghijklmnopqrstuvwxyz0123456789-._~()'!*:@,;`, this is the code access, currently is 6 characters long. So it can represent 49^6 diferent URLs, a total of `13.841.287.201` URLs.

