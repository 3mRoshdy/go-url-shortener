# Go-URL-Shortener 
A simple url shortener using golang that serves on localhost:8080

## Running the project
For running the server (without compliling the go project)
```
go run main.go
```
You should receive this message in terminal `URL Shortener is running on :8080`

For running the server (with compiling)
```
go build -o go-url-shortener main.go
```

and then execute the executable file
```
./go-url-shortener
```

## API
Creating a short version of a url
```
POST 'localhost:8080/shorten'

body: {
  "Url": "https://www.mywebsite.com"
}

response: {
    "Url": "https://www.mywebsite.com",
    "ShortUrl": "LfUbHL"
}
```

Redirection to original url using the short version
```
GET 'localhost:8080/short/:short_url'
```


## Packages

  * Shortener
    * shortener package for definition and logic of short urls
  * Generator:
    * generator package for generating random hash for url
