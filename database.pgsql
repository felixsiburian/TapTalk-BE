--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
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
-- Name: daily_entries; Type: TABLE; Schema: public; Owner: TapTalk
--

CREATE TABLE public.daily_entries (
    id bigint NOT NULL,
    user_id bigint,
    created_date text,
    updated_date text
);


ALTER TABLE public.daily_entries OWNER TO "TapTalk";

--
-- Name: daily_entries_id_seq; Type: SEQUENCE; Schema: public; Owner: TapTalk
--

CREATE SEQUENCE public.daily_entries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.daily_entries_id_seq OWNER TO "TapTalk";

--
-- Name: daily_entries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: TapTalk
--

ALTER SEQUENCE public.daily_entries_id_seq OWNED BY public.daily_entries.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: TapTalk
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    fullname text,
    birthday timestamp with time zone,
    email text,
    username text,
    password text,
    created_date text
);


ALTER TABLE public.users OWNER TO "TapTalk";

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: TapTalk
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO "TapTalk";

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: TapTalk
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: daily_entries id; Type: DEFAULT; Schema: public; Owner: TapTalk
--

ALTER TABLE ONLY public.daily_entries ALTER COLUMN id SET DEFAULT nextval('public.daily_entries_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: TapTalk
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: daily_entries; Type: TABLE DATA; Schema: public; Owner: TapTalk
--

COPY public.daily_entries (id, user_id, created_date, updated_date) FROM stdin;
1	1	2021-January-16	2021-January-16 01:38:09
2	2	2021-January-16	2021-January-16 01:39:09
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: TapTalk
--

COPY public.users (id, fullname, birthday, email, username, password, created_date) FROM stdin;
1	Test User	1997-09-22 19:42:31+07	test@gmail.com	testuser	$2a$10$TLrEbAikl/8rgalf9Pd6euaOdM8jsrnNo.FrSaP8LiVpbF6LEQHfi	2021-January-15
2	Test User 1	1997-09-22 19:42:31+07	test1@gmail.com	testuser1	$2a$10$YhXo0Gx8mhNv4TO.W63mQeBIpYpzpkeZVbgnGFLODyIhilvGBxVgq	2021-August-15
\.


--
-- Name: daily_entries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: TapTalk
--

SELECT pg_catalog.setval('public.daily_entries_id_seq', 2, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: TapTalk
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- Name: daily_entries daily_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: TapTalk
--

ALTER TABLE ONLY public.daily_entries
    ADD CONSTRAINT daily_entries_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: TapTalk
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

