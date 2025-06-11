# Api

## User / Authentication

### Creating a new user

**POST** `/auth/register`

Parameters

> | Name     | Type   | Description                                    |
> | -------- | ------ | ---------------------------------------------- |
> | username | string | The new user's desired name                    |
> | password | string | The (unencrypted) password for the new account |

Responses

> | Code | Body                           | Description                                |
> | ---- | ------------------------------ | ------------------------------------------ |
> | 200  | `{"authToken": "a.jwt.token"}` | The account has successfully been created. |
