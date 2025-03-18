--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE iacnaylor;
ALTER ROLE iacnaylor WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN NOREPLICATION BYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:HRlvR1miJNwT5qK2tofWfA==$+RBb/TyQq33fVRmncCdqgIRT6zn5w/RzPgfEBtIxYX0=:/PlQQXxsUg83xOD3SGOls4bu/xcIwezw8oqdMZdtxtg=';

--
-- User Configurations
--








--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.4 (Debian 17.4-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- PostgreSQL database dump complete
--

--
-- Database "apollo" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.4 (Debian 17.4-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: apollo; Type: DATABASE; Schema: -; Owner: iacnaylor
--

CREATE DATABASE apollo WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE apollo OWNER TO iacnaylor;

\connect apollo

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: albums; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.albums (
    album_id text NOT NULL,
    name text NOT NULL,
    rating integer NOT NULL,
    artist_id text NOT NULL
);


ALTER TABLE public.albums OWNER TO iacnaylor;

--
-- Name: artists; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.artists (
    artist_id text NOT NULL,
    name text NOT NULL,
    rating integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.artists OWNER TO iacnaylor;

--
-- Name: genres; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.genres (
    genre_id text NOT NULL,
    name text NOT NULL,
    rating integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.genres OWNER TO iacnaylor;

--
-- Name: recommendations; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.recommendations (
    id integer NOT NULL,
    user_id text NOT NULL,
    album_id text NOT NULL,
    date date DEFAULT CURRENT_DATE NOT NULL
);


ALTER TABLE public.recommendations OWNER TO iacnaylor;

--
-- Name: recommendations_id_seq; Type: SEQUENCE; Schema: public; Owner: iacnaylor
--

ALTER TABLE public.recommendations ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.recommendations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: user_albums; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.user_albums (
    user_id text NOT NULL,
    album_id text NOT NULL,
    rating integer DEFAULT 0 NOT NULL,
    notes text
);


ALTER TABLE public.user_albums OWNER TO iacnaylor;

--
-- Name: user_artists; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.user_artists (
    user_id text NOT NULL,
    artist_id text NOT NULL,
    rating integer DEFAULT 0 NOT NULL,
    notes text
);


ALTER TABLE public.user_artists OWNER TO iacnaylor;

--
-- Name: user_genres; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.user_genres (
    user_id text NOT NULL,
    genre_id text NOT NULL,
    rating integer DEFAULT 0 NOT NULL,
    notes text
);


ALTER TABLE public.user_genres OWNER TO iacnaylor;

--
-- Name: users; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.users (
    user_id text NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.users OWNER TO iacnaylor;

--
-- Data for Name: albums; Type: TABLE DATA; Schema: public; Owner: iacnaylor
--

COPY public.albums (album_id, name, rating, artist_id) FROM stdin;
\.


--
-- Data for Name: artists; Type: TABLE DATA; Schema: public; Owner: iacnaylor
--

COPY public.artists (artist_id, name, rating) FROM stdin;
\.


--
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: iacnaylor
--

COPY public.genres (genre_id, name, rating) FROM stdin;
\.


--
-- Data for Name: recommendations; Type: TABLE DATA; Schema: public; Owner: iacnaylor
--

COPY public.recommendations (id, user_id, album_id, date) FROM stdin;
\.


--
-- Data for Name: user_albums; Type: TABLE DATA; Schema: public; Owner: iacnaylor
--

COPY public.user_albums (user_id, album_id, rating, notes) FROM stdin;
\.


--
-- Data for Name: user_artists; Type: TABLE DATA; Schema: public; Owner: iacnaylor
--

COPY public.user_artists (user_id, artist_id, rating, notes) FROM stdin;
\.


--
-- Data for Name: user_genres; Type: TABLE DATA; Schema: public; Owner: iacnaylor
--

COPY public.user_genres (user_id, genre_id, rating, notes) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: iacnaylor
--

COPY public.users (user_id, name) FROM stdin;
\.


--
-- Name: recommendations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: iacnaylor
--

SELECT pg_catalog.setval('public.recommendations_id_seq', 1, false);


--
-- Name: users PK_user_id; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT "PK_user_id" PRIMARY KEY (user_id);


--
-- Name: albums albums_pkey; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_pkey PRIMARY KEY (album_id);


--
-- Name: artists artists_pkey; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.artists
    ADD CONSTRAINT artists_pkey PRIMARY KEY (artist_id);


--
-- Name: genres genres_pkey; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (genre_id);


--
-- Name: recommendations recommendations_pkey; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.recommendations
    ADD CONSTRAINT recommendations_pkey PRIMARY KEY (id);


--
-- Name: user_albums user_albums_pkey; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_albums
    ADD CONSTRAINT user_albums_pkey PRIMARY KEY (user_id, album_id);


--
-- Name: user_artists user_artists_pkey; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_artists
    ADD CONSTRAINT user_artists_pkey PRIMARY KEY (user_id, artist_id);


--
-- Name: user_genres user_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_genres
    ADD CONSTRAINT user_genres_pkey PRIMARY KEY (user_id, genre_id);


--
-- Name: user_albums FK_album; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_albums
    ADD CONSTRAINT "FK_album" FOREIGN KEY (album_id) REFERENCES public.albums(album_id) ON DELETE CASCADE;


--
-- Name: albums FK_album_artist; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT "FK_album_artist" FOREIGN KEY (artist_id) REFERENCES public.artists(artist_id) ON DELETE CASCADE NOT VALID;


--
-- Name: recommendations FK_album_recommendation; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.recommendations
    ADD CONSTRAINT "FK_album_recommendation" FOREIGN KEY (album_id) REFERENCES public.albums(album_id) ON DELETE CASCADE NOT VALID;


--
-- Name: user_artists FK_artist; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_artists
    ADD CONSTRAINT "FK_artist" FOREIGN KEY (artist_id) REFERENCES public.artists(artist_id) ON DELETE CASCADE NOT VALID;


--
-- Name: user_genres FK_genres; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_genres
    ADD CONSTRAINT "FK_genres" FOREIGN KEY (genre_id) REFERENCES public.genres(genre_id) ON DELETE CASCADE;


--
-- Name: user_artists FK_user; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_artists
    ADD CONSTRAINT "FK_user" FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


--
-- Name: user_albums FK_user; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_albums
    ADD CONSTRAINT "FK_user" FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


--
-- Name: user_genres FK_user; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_genres
    ADD CONSTRAINT "FK_user" FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


--
-- Name: recommendations FK_user_recommendation; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.recommendations
    ADD CONSTRAINT "FK_user_recommendation" FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE NOT VALID;


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

\connect postgres

--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.4 (Debian 17.4-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: albums; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.albums (
    album_id text NOT NULL,
    name text NOT NULL,
    rating integer NOT NULL,
    artist_id text NOT NULL
);


ALTER TABLE public.albums OWNER TO postgres;

--
-- Data for Name: albums; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.albums (album_id, name, rating, artist_id) FROM stdin;
\.


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

