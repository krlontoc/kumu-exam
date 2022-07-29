# Kumu Coding Challenge

## Description
This app is built in Golang. Uses Iris for router handling and also for endpoint unit testing. The app's main functionality is to get github user's basic information using the provided username. Informations were fetched from GIT by using http request to hit their public API `https://api.github.com/users/{username}`.

## How to run the app?
1. Build the docker image using this command.
```
docker build -t kumu-exam
```
2. After a successful build, use this command to run the docker image.
```
docker run -d --rm -p 1007:1007 kumu-exam
```
## Testing
For local testing, use command `go test -v` to run the test.

## Endpoints
- **Get Bearer Toke** : gen access token, token will be use for the other routes.  
Method: `POST`  
URL: `/auth/token`  
cURL:
```
curl --location --request POST 'localhost:1007/auth/token'
```
Response:
```
{
    "data": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTkwNzcwMTYsInZsZCI6dHJ1ZX0.Ak2N_7q_UaW86GwKFX02zx-FRbquXqbjCkQP50pA3yN95VP3s_6oOFxPOZzvUT9MW7vuWFi8rJ0HDE-Dn9rEES2UqCq_Gej5iJVfw3NVsQFkYHG6D3orY5QR528dVe-KdvZo19sXTFL_Az45fjlDExqP0ul-BI0KVZfMkHd2628QAVzCIM3nVXerOn7anYz8BTfs_dI3PNiBaIe7cMcLh-q6BW4cX-4rk5jR4pq7yQN1IUpqStr4yxtVhKGaUPQQ39E0W3FVEB1VNEBuyF36q3T9Nldsx1qGxWgKCmDwgYZXKMkq6BhD3XhK0xKfNCOy0HZNUZMLxQB25raGmh7KfA",
    "status": 200
}
```
- **Get Git Users Info** : get GIT users info, blank info of the user will automatically be omitted
Method: `GET`  
URL: `/api/v1/git-users`  
cURL:
```
curl --location --request GET 'localhost:1007/api/v1/git-users' \
--header 'Authorization: Bearer <token here>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "users": [
        "some-user-1",
        "krlontoc",
        "some-user-2",
        "",
        "krlontoc"
    ]
}'
```
Payload:
```
{
    "users": [
        "some-user-1",
        "krlontoc",
        "some-user-2",
        "",
        "krlontoc"
    ]
}
```
Response:
```
{
    "data": [
        {
            "name": "Kurt Russel Lontoc",
            "login": "krlontoc",
            "public_repos": 1
        },
        {
            "name": "some-user-1",
            "login": "some-user-1",
            "message": "Not Found"
        },
        {
            "name": "some-user-2",
            "login": "some-user-2",
            "message": "Not Found"
        }
    ],
    "status": 200
}
```
## Links
The app is also hosted at heroku, access the endpoints using this `https://kumu-exam.herokuapp.com`.  
Import postman collection using this `https://www.getpostman.com/collections/fdff8f00bb68f534d345`.  
View my resume by clicking [here](https://drive.google.com/file/d/1NI8yFlbc9Xp0HsasL071PmLNx1SjxruQ/view?usp=sharing).