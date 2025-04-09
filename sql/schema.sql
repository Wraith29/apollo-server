CREATE DATABASE apollo WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';

\connect apollo

CREATE TABLE public.genre (
    id text PRIMARY KEY,
    name text NOT NULL,
    rating integer NOT NULL DEFAULT 0
);

CREATE TABLE public.artist (
    id text PRIMARY KEY,
    name text NOT NULL,
    rating integer NOT NULL DEFAULT 0,
    updated_on date NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE public.artist_genre (
    artist_id text NOT NULL REFERENCES "artist" ("id") ON DELETE CASCADE,
    genre_id text NOT NULL REFERENCES "genre" ("id") ON DELETE CASCADE,

    PRIMARY KEY ("artist_id", "genre_id")
);

CREATE TABLE public.album (
    id text PRIMARY KEY,
    name text NOT NULL,
    release_date date NOT NULL,
    artist_id text NOT NULL REFERENCES "artist" ("id") ON DELETE CASCADE,
    rating integer NOT NULL DEFAULT 0
);

CREATE TABLE public.album_genre (
    album_id text NOT NULL REFERENCES "album" ("id") ON DELETE CASCADE,
    genre_id text NOT NULL REFERENCES "genre" ("id") ON DELETE CASCADE,

    PRIMARY KEY ("album_id", "genre_id")
);

CREATE TABLE public.user (
    name text PRIMARY KEY,
    password text NOT NULL
);

CREATE TABLE public.user_artist (
    user_id text NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    artist_id text NOT NULL REFERENCES "artist" ("id") ON DELETE CASCADE,
    rating integer NOT NULL DEFAULT 0,
    added_on date NOT NULL DEFAULT CURRENT_DATE,

    PRIMARY KEY ("user_id", "artist_id")
);

CREATE TABLE public.user_album (
    user_id text NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    album_id text NOT NULL REFERENCES "album" ("id") ON DELETE CASCADE,
    rating integer NOT NULL DEFAULT 0,
    recommended boolean NOT NULL DEFAULT false,
    added_on date NOT NULL DEFAULT CURRENT_DATE,
    updated_on date NOT NULL DEFAULT CURRENT_DATE,

    PRIMARY KEY ("user_id", "album_id")
);

CREATE TABLE public.user_genre (
    user_id text NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    genre_id text NOT NULL REFERENCES "genre" ("id") ON DELETE CASCADE,
    rating integer NOT NULL DEFAULT 0,
    added_on date NOT NULL DEFAULT CURRENT_DATE,

    PRIMARY KEY ("user_id", "genre_id")
);

CREATE TABLE public.recommendation (
    id serial PRIMARY KEY,
    user_id text NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    album_id text NOT NULL REFERENCES "album" ("id") ON DELETE CASCADE,
    recommended_date date NOT NULL DEFAULT CURRENT_DATE
);
