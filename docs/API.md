# Api

## User / Authentication

### Creating a new user

<details id="post_auth_register">

<summary>**POST** `/auth/register`</summary>

<details id="post_auth_register-parameters">

<summary>Parameters</summary>

> | Name     | Type   | Description                                    |
> | -------- | ------ | ---------------------------------------------- |
> | username | string | The new user's desired name                    |
> | password | string | The (unencrypted) password for the new account |

</details>

<details id="post_auth_register-responses">
<summary>Responses</summary>

> | Code | Body                           | Description                                |
> | ---- | ------------------------------ | ------------------------------------------ |
> | 200  | `{"authToken": "a.jwt.token"}` | The account has successfully been created. |

</details>
</details>
