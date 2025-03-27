--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

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
-- Name: comments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comments (
    id character varying(100) NOT NULL,
    author_id character varying(100) NOT NULL,
    post_id character varying(100) NOT NULL,
    content text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.comments OWNER TO postgres;

--
-- Name: conversations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.conversations (
    id character varying(100) NOT NULL,
    name character varying(200),
    host_id character varying(100),
    members character varying(500),
    is_group boolean,
    is_delete boolean,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.conversations OWNER TO postgres;

--
-- Name: likes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.likes (
    id character varying(100) NOT NULL,
    author_id character varying(100) NOT NULL,
    post_id character varying(100),
    comment_id character varying(100),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_like_target CHECK ((((post_id IS NOT NULL) AND (comment_id IS NULL)) OR ((post_id IS NULL) AND (comment_id IS NOT NULL))))
);


ALTER TABLE public.likes OWNER TO postgres;

--
-- Name: messages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.messages (
    id character varying(100) NOT NULL,
    author_id character varying(100) NOT NULL,
    conversation_id character varying(100) NOT NULL,
    content text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.messages OWNER TO postgres;

--
-- Name: notifications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.notifications (
    id character varying(100) NOT NULL,
    actor_id character varying(100) NOT NULL,
    target_user_id character varying(100),
    post_id character varying(100),
    comment_id character varying(100),
    action character varying(50) NOT NULL,
    is_read boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.notifications OWNER TO postgres;

--
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    id character varying(100) NOT NULL,
    author_id character varying(100),
    content text NOT NULL,
    attachment text,
    is_private boolean NOT NULL,
    is_hidden boolean NOT NULL,
    status boolean NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id character varying(100) NOT NULL,
    role_name character varying(100) NOT NULL,
    active_status boolean DEFAULT true NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- Name: social_requests; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.social_requests (
    id character varying(100) NOT NULL,
    author_id character varying(100) NOT NULL,
    account_id character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    request_type character varying(50) NOT NULL
);


ALTER TABLE public.social_requests OWNER TO postgres;

--
-- Name: user_securities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_securities (
    user_id character varying(100) NOT NULL,
    access_token text,
    refresh_token text,
    action_token text,
    fail_access integer DEFAULT 0,
    last_fail timestamp without time zone
);


ALTER TABLE public.user_securities OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id character varying(100) NOT NULL,
    role_id character varying(100),
    full_name character varying(255) NOT NULL,
    username character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    date_of_birth date,
    profile_avatar character varying(255),
    bio text,
    friends text,
    followers text,
    followings text,
    block_users text,
    conversations text,
    is_private boolean DEFAULT false,
    is_active boolean DEFAULT true,
    is_activated boolean,
    is_have_to_reset_password boolean,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.comments (id, author_id, post_id, content, created_at, updated_at) FROM stdin;
comment-1	40eebcec-dc43-4668-bc94-ec480f49481d	post-1	Excited to be part of this platform!	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
comment-2	4e8ce455-e59e-4bf6-88ba-defbe325f21c	post-1	The UI design looks great!	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
comment-3	7a105aeb-9fe9-4812-827d-40d496c58d3e	post-2	Looking forward to seeing your project!	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
comment-4	04cabad4-5f5f-4245-9b80-cf49b26dfc49	post-3	Amazing artwork, Jane!	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
comment-5	a00c6143-67e2-42db-8351-d8f05c819e9a	post-3	What tools do you use for your digital art?	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
comment-6	40eebcec-dc43-4668-bc94-ec480f49481d	post-5	Thanks for the update, will review the guidelines.	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
comment-7	4e8ce455-e59e-4bf6-88ba-defbe325f21c	post-6	What a beautiful view! Where is this?	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
comment-8	7a105aeb-9fe9-4812-827d-40d496c58d3e	post-8	Thanks for the heads-up about the maintenance.	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
\.


--
-- Data for Name: conversations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.conversations (id, name, host_id, members, is_group, is_delete, created_at, updated_at) FROM stdin;
conv-1	johndoe|diftlow	\N	04cabad4-5f5f-4245-9b80-cf49b26dfc49|40eebcec-dc43-4668-bc94-ec480f49481	f	f	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
conv-2	diftlow|tester01\n	\N	40eebcec-dc43-4668-bc94-ec480f49481d|4e8ce455-e59e-4bf6-88ba-defbe325f21	f	f	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
conv-3	johndoe|alexj	\N	04cabad4-5f5f-4245-9b80-cf49b26dfc49|7a105aeb-9fe9-4812-827d-40d496c58d3	f	f	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
conv-4	diftlow|janesmith	\N	40eebcec-dc43-4668-bc94-ec480f49481d|a00c6143-67e2-42db-8351-d8f05c819e9	f	f	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
conv-5	tester01|janesmith	\N	4e8ce455-e59e-4bf6-88ba-defbe325f21c|a00c6143-67e2-42db-8351-d8f05c819e9	f	f	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
conv-6	johndoe|tester01	\N	04cabad4-5f5f-4245-9b80-cf49b26dfc49|4e8ce455-e59e-4bf6-88ba-defbe325f21	f	f	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
conv-7	samwilson|namnguyen	\N	b8248036-6b03-4e1e-8f0d-4d86b6ee860f|d1fd2cef-ac30-42f0-a62b-49416418db4	f	f	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
\.


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.likes (id, author_id, post_id, comment_id, created_at) FROM stdin;
like-1	40eebcec-dc43-4668-bc94-ec480f49481d	post-1	\N	2025-03-27 20:30:45.662313
like-2	4e8ce455-e59e-4bf6-88ba-defbe325f21c	post-1	\N	2025-03-27 20:30:45.662313
like-3	7a105aeb-9fe9-4812-827d-40d496c58d3e	post-1	\N	2025-03-27 20:30:45.662313
like-4	a00c6143-67e2-42db-8351-d8f05c819e9a	post-1	\N	2025-03-27 20:30:45.662313
like-5	04cabad4-5f5f-4245-9b80-cf49b26dfc49	post-2	\N	2025-03-27 20:30:45.662313
like-6	4e8ce455-e59e-4bf6-88ba-defbe325f21c	post-2	\N	2025-03-27 20:30:45.662313
like-7	04cabad4-5f5f-4245-9b80-cf49b26dfc49	post-3	\N	2025-03-27 20:30:45.662313
like-8	40eebcec-dc43-4668-bc94-ec480f49481d	post-3	\N	2025-03-27 20:30:45.662313
like-9	7a105aeb-9fe9-4812-827d-40d496c58d3e	post-3	\N	2025-03-27 20:30:45.662313
like-10	04cabad4-5f5f-4245-9b80-cf49b26dfc49	\N	comment-1	2025-03-27 20:30:45.662313
like-11	4e8ce455-e59e-4bf6-88ba-defbe325f21c	\N	comment-1	2025-03-27 20:30:45.662313
like-12	40eebcec-dc43-4668-bc94-ec480f49481d	\N	comment-4	2025-03-27 20:30:45.662313
\.


--
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.messages (id, author_id, conversation_id, content, created_at) FROM stdin;
\.


--
-- Data for Name: notifications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.notifications (id, actor_id, target_user_id, post_id, comment_id, action, is_read, created_at) FROM stdin;
notif-1	40eebcec-dc43-4668-bc94-ec480f49481d	04cabad4-5f5f-4245-9b80-cf49b26dfc49	post-1	\N	like	f	2025-03-27 20:30:45.662313
notif-2	4e8ce455-e59e-4bf6-88ba-defbe325f21c	04cabad4-5f5f-4245-9b80-cf49b26dfc49	post-1	\N	comment	t	2025-03-27 20:30:45.662313
notif-3	7a105aeb-9fe9-4812-827d-40d496c58d3e	40eebcec-dc43-4668-bc94-ec480f49481d	post-2	\N	comment	f	2025-03-27 20:30:45.662313
notif-4	04cabad4-5f5f-4245-9b80-cf49b26dfc49	4e8ce455-e59e-4bf6-88ba-defbe325f21c	post-3	\N	like	t	2025-03-27 20:30:45.662313
notif-5	04cabad4-5f5f-4245-9b80-cf49b26dfc49	4e8ce455-e59e-4bf6-88ba-defbe325f21c	post-3	\N	comment	t	2025-03-27 20:30:45.662313
notif-6	40eebcec-dc43-4668-bc94-ec480f49481d	4e8ce455-e59e-4bf6-88ba-defbe325f21c	\N	\N	follow	f	2025-03-27 20:30:45.662313
notif-7	4e8ce455-e59e-4bf6-88ba-defbe325f21c	40eebcec-dc43-4668-bc94-ec480f49481d	\N	\N	follow	t	2025-03-27 20:30:45.662313
notif-8	a00c6143-67e2-42db-8351-d8f05c819e9a	4e8ce455-e59e-4bf6-88ba-defbe325f21c	post-3	\N	comment	f	2025-03-27 20:30:45.662313
\.


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.posts (id, author_id, content, attachment, is_private, is_hidden, status, created_at, updated_at) FROM stdin;
post-1	04cabad4-5f5f-4245-9b80-cf49b26dfc49	Welcome to our new social network platform!	\N	f	f	t	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
post-2	40eebcec-dc43-4668-bc94-ec480f49481d	Just finished working on a new project. Excited to share more details soon!	project-image.jpg	f	f	t	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
post-3	4e8ce455-e59e-4bf6-88ba-defbe325f21c	Check out my latest digital artwork	artwork1.jpg	f	f	t	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
post-4	4e8ce455-e59e-4bf6-88ba-defbe325f21c	Work in progress - new design concept	design-wip.jpg	f	f	t	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
post-5	7a105aeb-9fe9-4812-827d-40d496c58d3e	Important community guidelines update. Please read and follow.	\N	f	f	t	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
post-6	a00c6143-67e2-42db-8351-d8f05c819e9a	Beautiful sunset from my trip to the mountains	sunset.jpg	f	f	t	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
post-7	40eebcec-dc43-4668-bc94-ec480f49481d	This is a private post only for my friends	\N	t	f	t	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
post-8	04cabad4-5f5f-4245-9b80-cf49b26dfc49	System maintenance scheduled for tomorrow	\N	f	f	t	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id, role_name, active_status, created_at, updated_at) FROM stdin;
1	Admin	t	2025-03-26 22:10:34.860966	2025-03-26 22:10:34.860966
2	User	t	2025-03-26 22:10:34.860966	2025-03-26 22:10:34.860966
\.


--
-- Data for Name: social_requests; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.social_requests (id, author_id, account_id, created_at, request_type) FROM stdin;
req-1	7a105aeb-9fe9-4812-827d-40d496c58d3e	40eebcec-dc43-4668-bc94-ec480f49481d	2025-03-27 20:30:45.662313	add_friend
req-2	a00c6143-67e2-42db-8351-d8f05c819e9a	04cabad4-5f5f-4245-9b80-cf49b26dfc49	2025-03-27 20:30:45.662313	add_friend
req-3	4e8ce455-e59e-4bf6-88ba-defbe325f21c	a00c6143-67e2-42db-8351-d8f05c819e9a	2025-03-27 20:30:45.662313	add_friend
\.


--
-- Data for Name: user_securities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_securities (user_id, access_token, refresh_token, action_token, fail_access, last_fail) FROM stdin;
04cabad4-5f5f-4245-9b80-cf49b26dfc49	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZXhhbXBsZS5jb20iLCJleHBpcmUiOjE3NDMxMzMzOTEsInJvbGUiOiIyIiwidXNlcl9pZCI6IjA0Y2FiYWQ0LTVmNWYtNDI0NS05YjgwLWNmNDliMjZkZmM0OSJ9.ugGO_1ycTtAup-sIf2nA65e2Iur4Bp5PXOj8n0HDOzg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZXhhbXBsZS5jb20iLCJleHBpcmUiOjE3NDM2NTE3OTEsInJvbGUiOiIyIiwidXNlcl9pZCI6IjA0Y2FiYWQ0LTVmNWYtNDI0NS05YjgwLWNmNDliMjZkZmM0OSJ9.2jjCleVQptiQiN4BajVAyW9rPOj-BdE8piLcGTq-ATM	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZXhhbXBsZS5jb20iLCJleHBpcmUiOjE3NDMwNDU5MDgsInJvbGUiOiIyIiwidXNlcl9pZCI6IjA0Y2FiYWQ0LTVmNWYtNDI0NS05YjgwLWNmNDliMjZkZmM0OSJ9.Feiwcq9ieHM3McbTzc-EtLnsmqChp6hw1MgCsoDnCwE	0	1900-01-01 00:00:00
40eebcec-dc43-4668-bc94-ec480f49481d	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpZnRsb3dAZ21haWwuY29tIiwiZXhwaXJlIjoxNzQzMTMzMzI4LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0MGVlYmNlYy1kYzQzLTQ2NjgtYmM5NC1lYzQ4MGY0OTQ4MWQifQ.5IbKh8JS_23KAS8bKHTkPizFipiu2pnCEGXVTFbzWFw	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpZnRsb3dAZ21haWwuY29tIiwiZXhwaXJlIjoxNzQzNjUxNzI4LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0MGVlYmNlYy1kYzQzLTQ2NjgtYmM5NC1lYzQ4MGY0OTQ4MWQifQ.ABTVd8OJAY_JVyJ5uP4K76pU6_40sVqN18noly2k2Fk	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpZnRsb3dAZ21haWwuY29tIiwiZXhwaXJlIjoxNzQzMDE1NDQ1LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0MGVlYmNlYy1kYzQzLTQ2NjgtYmM5NC1lYzQ4MGY0OTQ4MWQifQ.bV5IyGDfWArfRs3PpaI-0KJ4y4pefHX75zKKZ0aSTFo	0	1900-01-01 00:00:00
4e8ce455-e59e-4bf6-88ba-defbe325f21c	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RlcjAxQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzMTMzMjU0LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0ZThjZTQ1NS1lNTllLTRiZjYtODhiYS1kZWZiZTMyNWYyMWMifQ.DtcT5TMldSjcX9NIiLYQDzgxJlHtS2WvUjB88jB6m4A	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RlcjAxQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzNjUxNjU0LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0ZThjZTQ1NS1lNTllLTRiZjYtODhiYS1kZWZiZTMyNWYyMWMifQ.5yds4XUs1FmpV8Ms8PwlL_VlC-8j9cdDEQ7YvEmtY4g	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RlcjAxQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzMDQ3MjExLCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0ZThjZTQ1NS1lNTllLTRiZjYtODhiYS1kZWZiZTMyNWYyMWMifQ.qTQHgrNrDZkukyO29p2mvXZESrP24Ji_cCu4E932pGE	0	1900-01-01 00:00:00
7a105aeb-9fe9-4812-827d-40d496c58d3e	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsZXhqQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzMTMzMjk3LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI3YTEwNWFlYi05ZmU5LTQ4MTItODI3ZC00MGQ0OTZjNThkM2UifQ.x3_X9ROlQt7m3g8hzQHae2o5dzvWv0Q7EOgxFnYlTz8	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsZXhqQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzNjUxNjk3LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI3YTEwNWFlYi05ZmU5LTQ4MTItODI3ZC00MGQ0OTZjNThkM2UifQ.nXWA7jZ57Irn-_mMdXXreBZL1YVznnBdhD9hmHuV5Tw	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsZXhqQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzMDQ1OTg1LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI3YTEwNWFlYi05ZmU5LTQ4MTItODI3ZC00MGQ0OTZjNThkM2UifQ.61oCchWKvZu9XkFLZuFWIll_tv58jHEQykXoBI85JCM	0	1900-01-01 00:00:00
a00c6143-67e2-42db-8351-d8f05c819e9a	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphbmVzbWl0aEBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzEzMjk0NCwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYTAwYzYxNDMtNjdlMi00MmRiLTgzNTEtZDhmMDVjODE5ZTlhIn0.EDWdLNE87gzWyQ_9krxeloe5VMoZqyQ1fSce3DvoSv8	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphbmVzbWl0aEBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzY1MTM0NCwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYTAwYzYxNDMtNjdlMi00MmRiLTgzNTEtZDhmMDVjODE5ZTlhIn0._ii_alwQQ-ea4b-wDclxhEDbKCEQEr9TwW2uSA9yrg4	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphbmVzbWl0aEBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzA0NTk0MSwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYTAwYzYxNDMtNjdlMi00MmRiLTgzNTEtZDhmMDVjODE5ZTlhIn0.FCKzvixt_6PowuazGKuz6B6cScRd80RdMoQrw6rSjSo	0	1900-01-01 00:00:00
b8248036-6b03-4e1e-8f0d-4d86b6ee860f	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXdpbHNvbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzEzMjk1Mywicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYjgyNDgwMzYtNmIwMy00ZTFlLThmMGQtNGQ4NmI2ZWU4NjBmIn0.0bGWCZSeJZEmP0OIQu9UbYqAYFaXx4YWy0Tu-Z1utLM	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXdpbHNvbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzY1MTM1Mywicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYjgyNDgwMzYtNmIwMy00ZTFlLThmMGQtNGQ4NmI2ZWU4NjBmIn0.Q3QfI3dc6VGxftRS1-UEmNT0IZyU4Uj5u0JL_L3_Z5Q	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXdpbHNvbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzA0NTk2MCwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYjgyNDgwMzYtNmIwMy00ZTFlLThmMGQtNGQ4NmI2ZWU4NjBmIn0.uTw9uwWSuI_Euc2Z_oZFRdS8R7zhJd62EW6nAixpcYQ	0	1900-01-01 00:00:00
d1fd2cef-ac30-42f0-a62b-49416418db4d	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbW5ndXllbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzEzMjk2Niwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiZDFmZDJjZWYtYWMzMC00MmYwLWE2MmItNDk0MTY0MThkYjRkIn0.QbToiSj0_1C61x1VLex1cdxUW3EpayEmN5LfJwMrJoU	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbW5ndXllbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzY1MTM2Niwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiZDFmZDJjZWYtYWMzMC00MmYwLWE2MmItNDk0MTY0MThkYjRkIn0.MnzhIzBuQqR0A2uGmtz2-Bihrx_Ix0aEX8Af4vMDBts	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbW5ndXllbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzAwMzg5MSwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiIn0.PeK2JmRqUPH4k9c-1r2XszACCrCwY2-C9b4ET7uTxBg	0	1900-01-01 00:00:00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, role_id, full_name, username, email, password, date_of_birth, profile_avatar, bio, friends, followers, followings, block_users, conversations, is_private, is_active, is_activated, is_have_to_reset_password, created_at, updated_at) FROM stdin;
d1fd2cef-ac30-42f0-a62b-49416418db4d	2	Nam Nguyen	namnguyen	namnguyen@example.com	$2a$10$ZfjPp5sEKAgswQECmoI2rurU5hEf7OIQD3H.pnE4Bm4JjUhXCKI9q	1990-05-15	namnguyen.jpg	Software engineer		7a105aeb-9fe9-4812-827d-40d496c58d3e			conv-7	f	t	t	\N	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
b8248036-6b03-4e1e-8f0d-4d86b6ee860f	2	Sam Wilson	samwilson	samwilson@example.com	$2a$10$jYadsG7f6lnUgq64msvp5OslkC/jS8qgtPLNZGR29syhq/yAwUzEu	1990-05-15	samwilson.jpg	Community manager	7a105aeb-9fe9-4812-827d-40d496c58d3e	7a105aeb-9fe9-4812-827d-40d496c58d3e			conv-7	f	t	t	\N	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
04cabad4-5f5f-4245-9b80-cf49b26dfc49	2	John Doe	johndoe	johndoe@example.com	$2a$10$LE0npKrAFzlubRxYKC3dTOwQVQu/Faymp4UfqrFPu7K0DIeEal9UK	1990-05-15	johndoe.jpg	Software developer and tech enthusiast	40eebcec-dc43-4668-bc94-ec480f49481d|4e8ce455-e59e-4bf6-88ba-defbe325f21c	40eebcec-dc43-4668-bc94-ec480f49481d|4e8ce455-e59e-4bf6-88ba-defbe325f21c|7a105aeb-9fe9-4812-827d-40d496c58d3e	40eebcec-dc43-4668-bc94-ec480f49481d|a00c6143-67e2-42db-8351-d8f05c819e9a		conv-1|conv-3|conv-6	f	t	t	\N	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
40eebcec-dc43-4668-bc94-ec480f49481d	2	diftlow	diftlow	diftlow@gmail.com	$2a$10$9/56Ienbxo0j5Mpkc9nfoOAxPIp2XqDBfgiRfSDtxqtUTYfex7Ehq	1990-05-15	avatar_url.jpg	Tech enthusiast	04cabad4-5f5f-4245-9b80-cf49b26dfc49|4e8ce455-e59e-4bf6-88ba-defbe325f21c	04cabad4-5f5f-4245-9b80-cf49b26dfc49|4e8ce455-e59e-4bf6-88ba-defbe325f21c|a00c6143-67e2-42db-8351-d8f05c819e9a	04cabad4-5f5f-4245-9b80-cf49b26dfc49|4e8ce455-e59e-4bf6-88ba-defbe325f21c|7a105aeb-9fe9-4812-827d-40d496c58d3e		conv-1|conv-2|conv-4	f	t	t	\N	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
4e8ce455-e59e-4bf6-88ba-defbe325f21c	2	Tester One	tester01	tester01@example.com	$2a$10$1XZdqlCyijtLGGICkAbGWuA4PY9NavLIyAXlmhrZhuraf.vOHLzLa	1990-05-15	tester01.jpg	QA tester	04cabad4-5f5f-4245-9b80-cf49b26dfc49|40eebcec-dc43-4668-bc94-ec480f49481d	40eebcec-dc43-4668-bc94-ec480f49481d|7a105aeb-9fe9-4812-827d-40d496c58d3e	04cabad4-5f5f-4245-9b80-cf49b26dfc49|40eebcec-dc43-4668-bc94-ec480f49481d		conv-2|conv-5|conv-6	f	t	t	\N	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
7a105aeb-9fe9-4812-827d-40d496c58d3e	2	Alex Johnson	alexj	alexj@example.com	$2a$10$ynOhFEWfbB2avt3JTnJabeZI4.vmhLo931EkuU/7PWc.nNAQSTAcK	1990-05-15	alexj.jpg	Traveler and photographer	b8248036-6b03-4e1e-8f0d-4d86b6ee860f	40eebcec-dc43-4668-bc94-ec480f49481d|d1fd2cef-ac30-42f0-a62b-49416418db4d	b8248036-6b03-4e1e-8f0d-4d86b6ee860f|d1fd2cef-ac30-42f0-a62b-49416418db4d		conv-3	f	t	t	\N	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
a00c6143-67e2-42db-8351-d8f05c819e9a	2	Jane Smith	janesmith	janesmith@example.com	$2a$10$yV57BBmsOOai5FkOxmoAeOf19b7uvChJO2BZ66SyCcKaxN9bZ6uRK	1990-05-15	janesmith.jpg	Digital artist and designer		04cabad4-5f5f-4245-9b80-cf49b26dfc49	40eebcec-dc43-4668-bc94-ec480f49481d		conv-4|conv-5	f	t	t	\N	2025-03-27 20:30:45.662313	2025-03-27 20:30:45.662313
\.


--
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- Name: conversations conversations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conversations
    ADD CONSTRAINT conversations_pkey PRIMARY KEY (id);


--
-- Name: likes likes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_pkey PRIMARY KEY (id);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: social_requests social_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.social_requests
    ADD CONSTRAINT social_requests_pkey PRIMARY KEY (id);


--
-- Name: user_securities user_securities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_securities
    ADD CONSTRAINT user_securities_pkey PRIMARY KEY (user_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: comments_author_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX comments_author_idx ON public.comments USING btree (author_id);


--
-- Name: comments_post_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX comments_post_idx ON public.comments USING btree (post_id);


--
-- Name: likes_author_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX likes_author_idx ON public.likes USING btree (author_id);


--
-- Name: likes_comment_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX likes_comment_idx ON public.likes USING btree (comment_id);


--
-- Name: likes_post_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX likes_post_idx ON public.likes USING btree (post_id);


--
-- Name: posts_author_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX posts_author_idx ON public.posts USING btree (author_id);


--
-- Name: comments fk_comment_post; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_comment_post FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE CASCADE;


--
-- Name: comments fk_comment_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_comment_user FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: conversations fk_conversation_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conversations
    ADD CONSTRAINT fk_conversation_user FOREIGN KEY (host_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: likes fk_like_comment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT fk_like_comment FOREIGN KEY (comment_id) REFERENCES public.comments(id) ON DELETE CASCADE;


--
-- Name: likes fk_like_post; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT fk_like_post FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE CASCADE;


--
-- Name: likes fk_like_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT fk_like_user FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: messages fk_message_conversation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT fk_message_conversation FOREIGN KEY (conversation_id) REFERENCES public.conversations(id) ON DELETE CASCADE;


--
-- Name: messages fk_message_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT fk_message_user FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: notifications fk_notification_actor; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notification_actor FOREIGN KEY (actor_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: notifications fk_notification_comment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notification_comment FOREIGN KEY (comment_id) REFERENCES public.comments(id) ON DELETE CASCADE;


--
-- Name: notifications fk_notification_post; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notification_post FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE CASCADE;


--
-- Name: notifications fk_notification_target_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notification_target_user FOREIGN KEY (target_user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: posts fk_post_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_post_user FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: social_requests fk_socialrequest_receiver; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.social_requests
    ADD CONSTRAINT fk_socialrequest_receiver FOREIGN KEY (account_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: social_requests fk_socialrequest_sender; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.social_requests
    ADD CONSTRAINT fk_socialrequest_sender FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: users fk_user_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_user_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE;


--
-- Name: user_securities fk_usersecurity_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_securities
    ADD CONSTRAINT fk_usersecurity_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

