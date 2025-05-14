# Apollo Server

> [!CAUTION]
> This is a heavily work in progress api
> Things may change

Apollo server is a HTTP server for a custom music library tracking / recommendation app.

## Usage

This server is designed to be used by either the [Apollo CLI](https://github.com/Wraith29/apollo-cli)
or the [Mobile App](https://github.com/Wraith29/apollo-mobile).

At the time of writing both of these implements are still WIP. The CLI is mostly functional, however may be missing important functionality.

## Workflow

The workflow of someone wanting to use this purely from the API would be:

> Register a user `POST <apollo-url>/auth/register` (This will return an auth token)
> Add artist(s) `POST <apollo-url>/artist`

```json
{
  "artistName": "My Favourite Artist"
}
```

> Get a recommendation `GET <apollo-url>/album/recommendation`
> Rate that recommendation `PUT <apollo-url>/album/rating`

```json
{
  "albumId": "my-album-id", // This is returned by the Recommendation endpoint. You will need to remember this / store it somewhere
  "rating": 3 // This can be any number between 1 - 5 inclusive. 1 is a negative rating, 3 is neutral, and 5 is positive, with stages between them.
}
```

## Running

I recommend running the app through docker.

It will require creating a `.env` file on the root with the following settings:

```
APOLLO_POSTGRES_USERNAME=<my_postgres_username>
APOLLO_POSTGRES_PASSWORD=<my_postgres_password>
APOLLO_DB_HOST=database
APOLLO_DB_HOST_PORT=1301
APOLLO_DB_CONN_PORT=5432
APOLLO_SECRET_KEY=<my_secret_key>
```

- `APOLLO_DB_HOST` is the name of the service within the `docker-compose.yml`.

- `APOLLO_DB_HOST_PORT` is the port which will be exposed by docker to connect to the database.

- `APOLLO_DB_CONN_PORT` is the port which Postgres is running on within the container.

Once this file is created, you can simply run:

```sh
docker compose up -d
```

To get the API & Database up and running.

You can then access the API on `http://localhost:1300` (Or whatever you changed the API port to in the `docker-compose.yml`)

And the database is accessible on port `1301` through PGAdmin4
