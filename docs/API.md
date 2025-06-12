## Authentication / Authorization

Authenticated requests require that the `Authorization` header be present, with a valid JWT token as the value.

A valid JWT can be acquired through one of the following endpoints: [Register](#creating-a-new-user), [Login](#logging-into-your-account) 
or, if you already have a valid token, you can refresh the expiry with [Refresh](#refreshing-your-token)

### Creating a new user

**Request**

```
ENDPOINT: /auth/register
METHOD: POST
BODY: {
  "username": "my_username",
  "password": "my_password"
}
```

Create a new user account

**Responses**

> | Code | Body                            | Description                                |
> | ---- | ------------------------------- | ------------------------------------------ |
> | 200  | `{"authToken": "a.jwt.token"}`  | The account has successfully been created. |
> | 400  | `{"err": "an error message"}`   | The request was malformed or invalid.      |
> | 409  | `{"err": "an error message"}`   | The username requested is already taken.   |
> | 500  | `{"err": "an error message"}`   | Something went wrong on the server side.   |

### Logging into your account

**Request**

```
ENDPOINT: /auth/login
METHOD: POST
BODY: {
  "username": "my_username",
  "password": "my_password"
}
```

Login to your existing account

**Responses**

> | Code | Body                            | Description                                    |
> | ---- | ------------------------------- | ---------------------------------------------- |
> | 200  | `{"authToken": "a.jwt.token"}`  | The account login was successful.              |
> | 400  | `{"err": "an error message"}`   | The request was malformed or invalid.          |
> | 401  | `{"err": "an error message"}`   | Something went wrong when trying to authorize. |
> | 500  | `{"err": "an error message"}`   | Something went wrong on the server side.       |

### Refreshing your Token

**Request**

```
ENDPOINT: /auth/refresh
AUTHENTICATED: Yes
METHOD: GET
```

Get a new token, the main purpose of this is to allow an Apollo client to refresh the users token 
without them having to login again

**Responses**

> | Code | Body                            | Description                                    |
> | ---- | ------------------------------- | ---------------------------------------------- |
> | 200  | `{"authToken": "a.jwt.token"}`  | The account login was successful.              |
> | 401  | `{"err": "an error message"}`   | Something went wrong when trying to authorize. |
> | 500  | `{"err": "an error message"}`   | Something went wrong on the server side.       |

## Artist

### Adding a new Artist

**Request**

```
ENDPOINT: /artist
AUTHENTICATED: Yes
METHOD: POST
BODY: {
  "artistName": "My Artist Name"
}
```

Apollo will go to [MusicBrainz](https://musicbrainz.org) and search for the given artist name,
and then using [Levenshtein Distance](https://en.wikipedia.org/wiki/Levenshtein_distance),
find the closest match, and add them to your library.

In future versions, there may be some alternatives for this endpoint added which will allow you to:

 - Pass in an artists [MusicBrainz ID](https://musicbrainz.org/doc/MusicBrainz_Identifier) to add them to your library

 - Have the endpoint return to you the list of artists found, and allow you to select a specific one to add.

**Responses**

> | Code | Body                            | Description                                    |
> | ---- | ------------------------------- | ---------------------------------------------- |
> | 202  | N/A                             | The artist was succesfully received.           |
> | 400  | `{"err": "an error message"}`   | The request body was invalid.                  |
> | 401  | `{"err": "an error message"}`   | Something went wrong when trying to authorize. |
> | 500  | `{"err": "an error message"}`   | Something went wrong on the server side.       |

### Updating an Artists data

**Request**

```
ENDPOINT: /artist/update
AUTHENTICATED: Yes
METHOD: POST
BODY: {
  "artistId": "string"
}
```

Apollo will find the artist in the database, and ensure that their Albums are up to date,
and any genre tags they have are accurate.

**Responses**


> | Code | Body                            | Description                                    |
> | ---- | ------------------------------- | ---------------------------------------------- |
> | 202  | N/A                             | The artist was succesfully received.           |
> | 400  | `{"err": "an error message"}`   | The request body was invalid.                  |
> | 401  | `{"err": "an error message"}`   | Something went wrong when trying to authorize. |
> | 500  | `{"err": "an error message"}`   | Something went wrong on the server side.       |

## Albums

### Getting a Recommendation

**Request**

```
ENDPOINT: /album/recommendation
AUTHENTICATED: Yes
METHOD: GET
QUERY: {
  "genres": "string,string",
  "includeListened": "bool"
}
```

Get an album recommended from your personal collection.

With no query passed in, this will find any album you have 
