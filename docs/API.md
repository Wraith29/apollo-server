# Api

## User / Authentication

### Creating a new user

Request

```
ENDPOINT: /auth/register
METHOD: POST
BODY: {
  "username": "my_username",
  "password": "my_password"
}
```


Responses

> | Code | Body                            | Description                                |
> | ---- | ------------------------------- | ------------------------------------------ |
> | 200  | `{"authToken": "a.jwt.token"}`  | The account has successfully been created. |
> | 400  | `{"err": "an error message"}`   | The request was malformed or invalid.      |
> | 409  | `{"err": "an error message"}`   | The username requested is already taken.   |
> | 500  | `{"err": "an error message"}`   | Something went wrong on the server side.   |

### Logging into your account

Request

```
ENDPOINT: /auth/login
METHOD: POST
BODY: {
  "username": "my_username",
  "password": "my_password"
}
```

Responses

> | Code | Body                            | Description                                    |
> | ---- | ------------------------------- | ---------------------------------------------- |
> | 200  | `{"authToken": "a.jwt.token"}`  | The account login was successful.              |
> | 400  | `{"err": "an error message"}`   | The request was malformed or invalid.          |
> | 401  | `{"err": "an error message"}`   | Something went wrong when trying to authorize. |
> | 500  | `{"err": "an error message"}`   | Something went wrong on the server side.       |

### Refreshing your Token

Request
```
ENDPOINT: /auth/refresh
METHOD: GET
HEADERS: {
  Authorization: your.jwt.token
}
```

Responses

> | Code | Body                            | Description                                    |
> | ---- | ------------------------------- | ---------------------------------------------- |
> | 200  | `{"authToken": "a.jwt.token"}`  | The account login was successful.              |
> | 401  | `{"err": "an error message"}`   | Something went wrong when trying to authorize. |
> | 500  | `{"err": "an error message"}`   | Something went wrong on the server side.       |
