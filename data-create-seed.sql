-- Complete Social Network Database Creation and Seeding Script
-- This script creates all tables with proper constraints and seeds the data

-- Drop all tables if they exist
DROP TABLE IF EXISTS public.messages CASCADE;
DROP TABLE IF EXISTS public.likes CASCADE;
DROP TABLE IF EXISTS public.comments CASCADE;
DROP TABLE IF EXISTS public.notifications CASCADE;
DROP TABLE IF EXISTS public.social_requests CASCADE;
DROP TABLE IF EXISTS public.posts CASCADE;
DROP TABLE IF EXISTS public.conversations CASCADE;
DROP TABLE IF EXISTS public.user_securities CASCADE;
DROP TABLE IF EXISTS public.users CASCADE;
DROP TABLE IF EXISTS public.roles CASCADE;

-- Create roles table
CREATE TABLE public.roles (
                              id character varying(100) NOT NULL,
                              role_name character varying(100) NOT NULL,
                              active_status boolean NOT NULL DEFAULT true,
                              created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                              CONSTRAINT roles_pkey PRIMARY KEY (id)
);

-- Create users table
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
                              updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                              CONSTRAINT users_pkey PRIMARY KEY (id),
                              CONSTRAINT users_email_key UNIQUE (email),
                              CONSTRAINT fk_user_role FOREIGN KEY (role_id)
                                  REFERENCES public.roles (id) MATCH SIMPLE
                                  ON UPDATE NO ACTION
                                  ON DELETE CASCADE
);

-- Create user_securities table
CREATE TABLE public.user_securities (
                                        user_id character varying(100) NOT NULL,
                                        access_token text,
                                        refresh_token text,
                                        action_token text,
                                        fail_access integer DEFAULT 0,
                                        last_fail timestamp without time zone,
                                        CONSTRAINT user_securities_pkey PRIMARY KEY (user_id),
                                        CONSTRAINT fk_usersecurity_user FOREIGN KEY (user_id)
                                            REFERENCES public.users (id) MATCH SIMPLE
                                            ON UPDATE NO ACTION
                                            ON DELETE CASCADE
);

-- Create posts table
CREATE TABLE public.posts (
                              id character varying(100) NOT NULL,
                              author_id character varying(100),
                              content text NOT NULL,
                              attachment text,
                              is_private boolean NOT NULL,
                              is_hidden boolean NOT NULL,
                              status boolean NOT NULL,
                              created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                              CONSTRAINT posts_pkey PRIMARY KEY (id),
                              CONSTRAINT fk_post_user FOREIGN KEY (author_id)
                                  REFERENCES public.users (id) MATCH SIMPLE
                                  ON UPDATE NO ACTION
                                  ON DELETE CASCADE
);

-- Create comments table
CREATE TABLE public.comments (
                                 id character varying(100) NOT NULL,
                                 author_id character varying(100) NOT NULL,
                                 post_id character varying(100) NOT NULL,
                                 content text NOT NULL,
                                 created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                 updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                 CONSTRAINT comments_pkey PRIMARY KEY (id),
                                 CONSTRAINT fk_comment_post FOREIGN KEY (post_id)
                                     REFERENCES public.posts (id) MATCH SIMPLE
                                     ON UPDATE NO ACTION
                                     ON DELETE CASCADE,
                                 CONSTRAINT fk_comment_user FOREIGN KEY (author_id)
                                     REFERENCES public.users (id) MATCH SIMPLE
                                     ON UPDATE NO ACTION
                                     ON DELETE CASCADE
);

-- Create conversations table
CREATE TABLE public.conversations (
                                      id character varying(100) NOT NULL,
                                      name character varying(200),
                                      host_id character varying(100),
                                      members character varying(500),
                                      is_group boolean,
                                      is_delete boolean,
                                      created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                      updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                      CONSTRAINT conversations_pkey PRIMARY KEY (id),
                                      CONSTRAINT fk_conversation_user FOREIGN KEY (host_id)
                                          REFERENCES public.users (id) MATCH SIMPLE
                                          ON UPDATE NO ACTION
                                          ON DELETE CASCADE
);

-- Create messages table
CREATE TABLE public.messages (
                                 id character varying(100) NOT NULL,
                                 author_id character varying(100) NOT NULL,
                                 conversation_id character varying(100) NOT NULL,
                                 content text NOT NULL,
                                 created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                 CONSTRAINT messages_pkey PRIMARY KEY (id),
                                 CONSTRAINT fk_message_conversation FOREIGN KEY (conversation_id)
                                     REFERENCES public.conversations (id) MATCH SIMPLE
                                     ON UPDATE NO ACTION
                                     ON DELETE CASCADE,
                                 CONSTRAINT fk_message_user FOREIGN KEY (author_id)
                                     REFERENCES public.users (id) MATCH SIMPLE
                                     ON UPDATE NO ACTION
                                     ON DELETE CASCADE
);

-- Create social_requests table
CREATE TABLE public.social_requests (
                                        id character varying(100) NOT NULL,
                                        author_id character varying(100) NOT NULL,
                                        account_id character varying(100) NOT NULL,
                                        created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                        CONSTRAINT social_requests_pkey PRIMARY KEY (id),
                                        CONSTRAINT fk_socialrequest_receiver FOREIGN KEY (account_id)
                                            REFERENCES public.users (id) MATCH SIMPLE
                                            ON UPDATE NO ACTION
                                            ON DELETE CASCADE,
                                        CONSTRAINT fk_socialrequest_sender FOREIGN KEY (author_id)
                                            REFERENCES public.users (id) MATCH SIMPLE
                                            ON UPDATE NO ACTION
                                            ON DELETE CASCADE
);

-- Create likes table with FIXED constraints
-- Using separate columns for post_id and comment_id instead of a single object_id
CREATE TABLE public.likes (
                              id character varying(100) NOT NULL,
                              author_id character varying(100) NOT NULL,
                              post_id character varying(100),
                              comment_id character varying(100),
                              created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                              CONSTRAINT likes_pkey PRIMARY KEY (id),
                              CONSTRAINT fk_like_user FOREIGN KEY (author_id)
                                  REFERENCES public.users (id) MATCH SIMPLE
                                  ON UPDATE NO ACTION
                                  ON DELETE CASCADE,
                              CONSTRAINT fk_like_post FOREIGN KEY (post_id)
                                  REFERENCES public.posts (id) MATCH SIMPLE
                                  ON UPDATE NO ACTION
                                  ON DELETE CASCADE,
                              CONSTRAINT fk_like_comment FOREIGN KEY (comment_id)
                                  REFERENCES public.comments (id) MATCH SIMPLE
                                  ON UPDATE NO ACTION
                                  ON DELETE CASCADE,
    -- Ensure either post_id or comment_id is provided, but not both
                              CONSTRAINT check_like_target CHECK (
                                  (post_id IS NOT NULL AND comment_id IS NULL) OR
                                  (post_id IS NULL AND comment_id IS NOT NULL)
                                  )
);

-- Create notifications table with FIXED constraints
CREATE TABLE public.notifications (
                                      id character varying(100) NOT NULL,
                                      actor_id character varying(100) NOT NULL,
                                      target_user_id character varying(100), -- User being notified
                                      post_id character varying(100), -- Optional post reference
                                      comment_id character varying(100), -- Optional comment reference
                                      action character varying(50) NOT NULL, -- like, comment, follow, etc.
                                      is_read boolean DEFAULT false,
                                      created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                      CONSTRAINT notifications_pkey PRIMARY KEY (id),
                                      CONSTRAINT fk_notification_actor FOREIGN KEY (actor_id)
                                          REFERENCES public.users (id) MATCH SIMPLE
                                          ON UPDATE NO ACTION
                                          ON DELETE CASCADE,
                                      CONSTRAINT fk_notification_target_user FOREIGN KEY (target_user_id)
                                          REFERENCES public.users (id) MATCH SIMPLE
                                          ON UPDATE NO ACTION
                                          ON DELETE CASCADE,
                                      CONSTRAINT fk_notification_post FOREIGN KEY (post_id)
                                          REFERENCES public.posts (id) MATCH SIMPLE
                                          ON UPDATE NO ACTION
                                          ON DELETE CASCADE,
                                      CONSTRAINT fk_notification_comment FOREIGN KEY (comment_id)
                                          REFERENCES public.comments (id) MATCH SIMPLE
                                          ON UPDATE NO ACTION
                                          ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX IF NOT EXISTS user_securities_pkey ON public.user_securities(user_id);
CREATE INDEX IF NOT EXISTS posts_author_idx ON public.posts(author_id);
CREATE INDEX IF NOT EXISTS comments_post_idx ON public.comments(post_id);
CREATE INDEX IF NOT EXISTS comments_author_idx ON public.comments(author_id);
CREATE INDEX IF NOT EXISTS likes_author_idx ON public.likes(author_id);
CREATE INDEX IF NOT EXISTS likes_post_idx ON public.likes(post_id);
CREATE INDEX IF NOT EXISTS likes_comment_idx ON public.likes(comment_id);

-- =====================================================
-- DATA SEEDING
-- =====================================================

-- Seed roles
INSERT INTO public.roles (id, role_name, active_status, created_at, updated_at)
VALUES
    ('1', 'Admin', true, '2025-03-26 22:10:34.860966', '2025-03-26 22:10:34.860966'),
    ('2', 'User', true, '2025-03-26 22:10:34.860966', '2025-03-26 22:10:34.860966');

-- Seed users
INSERT INTO public.users (
    id, role_id, full_name, username, email, password,
    date_of_birth, profile_avatar, bio, is_private, is_active, is_activated
)
VALUES
    ('04cabad4-5f5f-4245-9b80-cf49b26dfc49', '2', 'John Doe', 'johndoe', 'johndoe@example.com',
     '$2a$10$LE0npKrAFzlubRxYKC3dTOwQVQu/Faymp4UfqrFPu7K0DIeEal9UK',
     '1990-05-15', 'johndoe.jpg', 'Software developer and tech enthusiast', false, true, true),

    ('40eebcec-dc43-4668-bc94-ec480f49481d', '2', 'diftlow', 'diftlow', 'diftlow@gmail.com',
     '$2a$10$9/56Ienbxo0j5Mpkc9nfoOAxPIp2XqDBfgiRfSDtxqtUTYfex7Ehq',
     '1990-05-15', 'avatar_url.jpg', 'Tech enthusiast', false, true, true),

    ('4e8ce455-e59e-4bf6-88ba-defbe325f21c', '2', 'Tester One', 'tester01', 'tester01@example.com',
     '$2a$10$1XZdqlCyijtLGGICkAbGWuA4PY9NavLIyAXlmhrZhuraf.vOHLzLa',
     '1990-05-15', 'tester01.jpg', 'QA tester', false, true, true),

    ('7a105aeb-9fe9-4812-827d-40d496c58d3e', '2', 'Alex Johnson', 'alexj', 'alexj@example.com',
     '$2a$10$ynOhFEWfbB2avt3JTnJabeZI4.vmhLo931EkuU/7PWc.nNAQSTAcK',
     '1990-05-15', 'alexj.jpg', 'Traveler and photographer', false, true, true),

    ('a00c6143-67e2-42db-8351-d8f05c819e9a', '2', 'Jane Smith', 'janesmith', 'janesmith@example.com',
     '$2a$10$yV57BBmsOOai5FkOxmoAeOf19b7uvChJO2BZ66SyCcKaxN9bZ6uRK',
     '1990-05-15', 'janesmith.jpg', 'Digital artist and designer', false, true, true),

    ('b8248036-6b03-4e1e-8f0d-4d86b6ee860f', '2', 'Sam Wilson', 'samwilson', 'samwilson@example.com',
     '$2a$10$jYadsG7f6lnUgq64msvp5OslkC/jS8qgtPLNZGR29syhq/yAwUzEu',
     '1990-05-15', 'samwilson.jpg', 'Community manager', false, true, true),

    ('d1fd2cef-ac30-42f0-a62b-49416418db4d', '2', 'Nam Nguyen', 'namnguyen', 'namnguyen@example.com',
     '$2a$10$ZfjPp5sEKAgswQECmoI2rurU5hEf7OIQD3H.pnE4Bm4JjUhXCKI9q',
     '1990-05-15', 'namnguyen.jpg', 'Software engineer', false, true, true);

-- Update user relationships
UPDATE public.users
SET
    friends = '["40eebcec-dc43-4668-bc94-ec480f49481d","4e8ce455-e59e-4bf6-88ba-defbe325f21c"]',
    followers = '["40eebcec-dc43-4668-bc94-ec480f49481d","4e8ce455-e59e-4bf6-88ba-defbe325f21c","7a105aeb-9fe9-4812-827d-40d496c58d3e"]',
    followings = '["40eebcec-dc43-4668-bc94-ec480f49481d","a00c6143-67e2-42db-8351-d8f05c819e9a"]'
WHERE id = '04cabad4-5f5f-4245-9b80-cf49b26dfc49';

UPDATE public.users
SET
    friends = '["04cabad4-5f5f-4245-9b80-cf49b26dfc49","4e8ce455-e59e-4bf6-88ba-defbe325f21c"]',
    followers = '["04cabad4-5f5f-4245-9b80-cf49b26dfc49","4e8ce455-e59e-4bf6-88ba-defbe325f21c","a00c6143-67e2-42db-8351-d8f05c819e9a"]',
    followings = '["04cabad4-5f5f-4245-9b80-cf49b26dfc49","4e8ce455-e59e-4bf6-88ba-defbe325f21c","7a105aeb-9fe9-4812-827d-40d496c58d3e"]'
WHERE id = '40eebcec-dc43-4668-bc94-ec480f49481d';

UPDATE public.users
SET
    friends = '["04cabad4-5f5f-4245-9b80-cf49b26dfc49","40eebcec-dc43-4668-bc94-ec480f49481d"]',
    followers = '["40eebcec-dc43-4668-bc94-ec480f49481d","7a105aeb-9fe9-4812-827d-40d496c58d3e"]',
    followings = '["04cabad4-5f5f-4245-9b80-cf49b26dfc49","40eebcec-dc43-4668-bc94-ec480f49481d"]'
WHERE id = '4e8ce455-e59e-4bf6-88ba-defbe325f21c';

UPDATE public.users
SET
    friends = '["b8248036-6b03-4e1e-8f0d-4d86b6ee860f"]',
    followers = '["40eebcec-dc43-4668-bc94-ec480f49481d","d1fd2cef-ac30-42f0-a62b-49416418db4d"]',
    followings = '["b8248036-6b03-4e1e-8f0d-4d86b6ee860f","d1fd2cef-ac30-42f0-a62b-49416418db4d"]'
WHERE id = '7a105aeb-9fe9-4812-827d-40d496c58d3e';

UPDATE public.users
SET
    friends = '[]',
    followers = '["04cabad4-5f5f-4245-9b80-cf49b26dfc49"]',
    followings = '["40eebcec-dc43-4668-bc94-ec480f49481d"]'
WHERE id = 'a00c6143-67e2-42db-8351-d8f05c819e9a';

UPDATE public.users
SET
    friends = '["7a105aeb-9fe9-4812-827d-40d496c58d3e"]',
    followers = '["7a105aeb-9fe9-4812-827d-40d496c58d3e"]',
    followings = '[]'
WHERE id = 'b8248036-6b03-4e1e-8f0d-4d86b6ee860f';

UPDATE public.users
SET
    friends = '[]',
    followers = '["7a105aeb-9fe9-4812-827d-40d496c58d3e"]',
    followings = '[]'
WHERE id = 'd1fd2cef-ac30-42f0-a62b-49416418db4d';

-- Seed user securities
INSERT INTO public.user_securities (
    user_id, access_token, refresh_token, action_token, fail_access, last_fail
)
VALUES
    ('04cabad4-5f5f-4245-9b80-cf49b26dfc49',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZXhhbXBsZS5jb20iLCJleHBpcmUiOjE3NDMxMzMzOTEsInJvbGUiOiIyIiwidXNlcl9pZCI6IjA0Y2FiYWQ0LTVmNWYtNDI0NS05YjgwLWNmNDliMjZkZmM0OSJ9.ugGO_1ycTtAup-sIf2nA65e2Iur4Bp5PXOj8n0HDOzg',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZXhhbXBsZS5jb20iLCJleHBpcmUiOjE3NDM2NTE3OTEsInJvbGUiOiIyIiwidXNlcl9pZCI6IjA0Y2FiYWQ0LTVmNWYtNDI0NS05YjgwLWNmNDliMjZkZmM0OSJ9.2jjCleVQptiQiN4BajVAyW9rPOj-BdE8piLcGTq-ATM',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZXhhbXBsZS5jb20iLCJleHBpcmUiOjE3NDMwNDU5MDgsInJvbGUiOiIyIiwidXNlcl9pZCI6IjA0Y2FiYWQ0LTVmNWYtNDI0NS05YjgwLWNmNDliMjZkZmM0OSJ9.Feiwcq9ieHM3McbTzc-EtLnsmqChp6hw1MgCsoDnCwE',
     0, '1900-01-01 00:00:00'),

    ('40eebcec-dc43-4668-bc94-ec480f49481d',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpZnRsb3dAZ21haWwuY29tIiwiZXhwaXJlIjoxNzQzMTMzMzI4LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0MGVlYmNlYy1kYzQzLTQ2NjgtYmM5NC1lYzQ4MGY0OTQ4MWQifQ.5IbKh8JS_23KAS8bKHTkPizFipiu2pnCEGXVTFbzWFw',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpZnRsb3dAZ21haWwuY29tIiwiZXhwaXJlIjoxNzQzNjUxNzI4LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0MGVlYmNlYy1kYzQzLTQ2NjgtYmM5NC1lYzQ4MGY0OTQ4MWQifQ.ABTVd8OJAY_JVyJ5uP4K76pU6_40sVqN18noly2k2Fk',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpZnRsb3dAZ21haWwuY29tIiwiZXhwaXJlIjoxNzQzMDE1NDQ1LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0MGVlYmNlYy1kYzQzLTQ2NjgtYmM5NC1lYzQ4MGY0OTQ4MWQifQ.bV5IyGDfWArfRs3PpaI-0KJ4y4pefHX75zKKZ0aSTFo',
     0, '1900-01-01 00:00:00'),

    ('4e8ce455-e59e-4bf6-88ba-defbe325f21c',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RlcjAxQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzMTMzMjU0LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0ZThjZTQ1NS1lNTllLTRiZjYtODhiYS1kZWZiZTMyNWYyMWMifQ.DtcT5TMldSjcX9NIiLYQDzgxJlHtS2WvUjB88jB6m4A',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RlcjAxQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzNjUxNjU0LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0ZThjZTQ1NS1lNTllLTRiZjYtODhiYS1kZWZiZTMyNWYyMWMifQ.5yds4XUs1FmpV8Ms8PwlL_VlC-8j9cdDEQ7YvEmtY4g',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RlcjAxQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzMDQ3MjExLCJyb2xlIjoiMiIsInVzZXJfaWQiOiI0ZThjZTQ1NS1lNTllLTRiZjYtODhiYS1kZWZiZTMyNWYyMWMifQ.qTQHgrNrDZkukyO29p2mvXZESrP24Ji_cCu4E932pGE',
     0, '1900-01-01 00:00:00'),

    ('7a105aeb-9fe9-4812-827d-40d496c58d3e',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsZXhqQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzMTMzMjk3LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI3YTEwNWFlYi05ZmU5LTQ4MTItODI3ZC00MGQ0OTZjNThkM2UifQ.x3_X9ROlQt7m3g8hzQHae2o5dzvWv0Q7EOgxFnYlTz8',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsZXhqQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzNjUxNjk3LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI3YTEwNWFlYi05ZmU5LTQ4MTItODI3ZC00MGQ0OTZjNThkM2UifQ.nXWA7jZ57Irn-_mMdXXreBZL1YVznnBdhD9hmHuV5Tw',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsZXhqQGV4YW1wbGUuY29tIiwiZXhwaXJlIjoxNzQzMDQ1OTg1LCJyb2xlIjoiMiIsInVzZXJfaWQiOiI3YTEwNWFlYi05ZmU5LTQ4MTItODI3ZC00MGQ0OTZjNThkM2UifQ.61oCchWKvZu9XkFLZuFWIll_tv58jHEQykXoBI85JCM',
     0, '1900-01-01 00:00:00'),

    ('a00c6143-67e2-42db-8351-d8f05c819e9a',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphbmVzbWl0aEBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzEzMjk0NCwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYTAwYzYxNDMtNjdlMi00MmRiLTgzNTEtZDhmMDVjODE5ZTlhIn0.EDWdLNE87gzWyQ_9krxeloe5VMoZqyQ1fSce3DvoSv8',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphbmVzbWl0aEBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzY1MTM0NCwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYTAwYzYxNDMtNjdlMi00MmRiLTgzNTEtZDhmMDVjODE5ZTlhIn0._ii_alwQQ-ea4b-wDclxhEDbKCEQEr9TwW2uSA9yrg4',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphbmVzbWl0aEBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzA0NTk0MSwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYTAwYzYxNDMtNjdlMi00MmRiLTgzNTEtZDhmMDVjODE5ZTlhIn0.FCKzvixt_6PowuazGKuz6B6cScRd80RdMoQrw6rSjSo',
     0, '1900-01-01 00:00:00'),

    ('b8248036-6b03-4e1e-8f0d-4d86b6ee860f',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXdpbHNvbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzEzMjk1Mywicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYjgyNDgwMzYtNmIwMy00ZTFlLThmMGQtNGQ4NmI2ZWU4NjBmIn0.0bGWCZSeJZEmP0OIQu9UbYqAYFaXx4YWy0Tu-Z1utLM',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXdpbHNvbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzY1MTM1Mywicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYjgyNDgwMzYtNmIwMy00ZTFlLThmMGQtNGQ4NmI2ZWU4NjBmIn0.Q3QfI3dc6VGxftRS1-UEmNT0IZyU4Uj5u0JL_L3_Z5Q',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXdpbHNvbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzA0NTk2MCwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiYjgyNDgwMzYtNmIwMy00ZTFlLThmMGQtNGQ4NmI2ZWU4NjBmIn0.uTw9uwWSuI_Euc2Z_oZFRdS8R7zhJd62EW6nAixpcYQ',
     0, '1900-01-01 00:00:00'),

    ('d1fd2cef-ac30-42f0-a62b-49416418db4d',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbW5ndXllbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzEzMjk2Niwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiZDFmZDJjZWYtYWMzMC00MmYwLWE2MmItNDk0MTY0MThkYjRkIn0.QbToiSj0_1C61x1VLex1cdxUW3EpayEmN5LfJwMrJoU',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbW5ndXllbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzY1MTM2Niwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiZDFmZDJjZWYtYWMzMC00MmYwLWE2MmItNDk0MTY0MThkYjRkIn0.MnzhIzBuQqR0A2uGmtz2-Bihrx_Ix0aEX8Af4vMDBts',
     'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbW5ndXllbkBleGFtcGxlLmNvbSIsImV4cGlyZSI6MTc0MzAwMzg5MSwicm9sZSI6IjIiLCJ1c2VyX2lkIjoiIn0.PeK2JmRqUPH4k9c-1r2XszACCrCwY2-C9b4ET7uTxBg',
     0, '1900-01-01 00:00:00');

-- Seed posts
INSERT INTO public.posts (id, author_id, content, attachment, is_private, is_hidden, status)
VALUES
    ('post-1', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'Welcome to our new social network platform!', NULL, false, false, true),
    ('post-2', '40eebcec-dc43-4668-bc94-ec480f49481d', 'Just finished working on a new project. Excited to share more details soon!', 'project-image.jpg', false, false, true),
    ('post-3', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'Check out my latest digital artwork', 'artwork1.jpg', false, false, true),
    ('post-4', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'Work in progress - new design concept', 'design-wip.jpg', false, false, true),
    ('post-5', '7a105aeb-9fe9-4812-827d-40d496c58d3e', 'Important community guidelines update. Please read and follow.', NULL, false, false, true),
    ('post-6', 'a00c6143-67e2-42db-8351-d8f05c819e9a', 'Beautiful sunset from my trip to the mountains', 'sunset.jpg', false, false, true),
    ('post-7', '40eebcec-dc43-4668-bc94-ec480f49481d', 'This is a private post only for my friends', NULL, true, false, true),
    ('post-8', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'System maintenance scheduled for tomorrow', NULL, false, false, true);

-- Seed comments
INSERT INTO public.comments (id, author_id, post_id, content)
VALUES
    ('comment-1', '40eebcec-dc43-4668-bc94-ec480f49481d', 'post-1', 'Excited to be part of this platform!'),
    ('comment-2', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'post-1', 'The UI design looks great!'),
    ('comment-3', '7a105aeb-9fe9-4812-827d-40d496c58d3e', 'post-2', 'Looking forward to seeing your project!'),
    ('comment-4', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'post-3', 'Amazing artwork, Jane!'),
    ('comment-5', 'a00c6143-67e2-42db-8351-d8f05c819e9a', 'post-3', 'What tools do you use for your digital art?'),
    ('comment-6', '40eebcec-dc43-4668-bc94-ec480f49481d', 'post-5', 'Thanks for the update, will review the guidelines.'),
    ('comment-7', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'post-6', 'What a beautiful view! Where is this?'),
    ('comment-8', '7a105aeb-9fe9-4812-827d-40d496c58d3e', 'post-8', 'Thanks for the heads-up about the maintenance.');

-- Seed conversations
INSERT INTO public.conversations (id, name, host_id, members, is_group, is_delete)
VALUES
    ('conv-1', NULL, '04cabad4-5f5f-4245-9b80-cf49b26dfc49', '["04cabad4-5f5f-4245-9b80-cf49b26dfc49","40eebcec-dc43-4668-bc94-ec480f49481d"]', false, false),
    ('conv-2', NULL, '40eebcec-dc43-4668-bc94-ec480f49481d', '["40eebcec-dc43-4668-bc94-ec480f49481d","4e8ce455-e59e-4bf6-88ba-defbe325f21c"]', false, false),
    ('conv-3', NULL, '04cabad4-5f5f-4245-9b80-cf49b26dfc49', '["04cabad4-5f5f-4245-9b80-cf49b26dfc49","7a105aeb-9fe9-4812-827d-40d496c58d3e"]', false, false),
    ('conv-4', NULL, '40eebcec-dc43-4668-bc94-ec480f49481d', '["40eebcec-dc43-4668-bc94-ec480f49481d","a00c6143-67e2-42db-8351-d8f05c819e9a"]', false, false),
    ('conv-5', NULL, '4e8ce455-e59e-4bf6-88ba-defbe325f21c', '["4e8ce455-e59e-4bf6-88ba-defbe325f21c","a00c6143-67e2-42db-8351-d8f05c819e9a"]', false, false),
    ('conv-6', NULL, '04cabad4-5f5f-4245-9b80-cf49b26dfc49', '["04cabad4-5f5f-4245-9b80-cf49b26dfc49","4e8ce455-e59e-4bf6-88ba-defbe325f21c"]', false, false),
    ('conv-7', NULL, 'b8248036-6b03-4e1e-8f0d-4d86b6ee860f', '["b8248036-6b03-4e1e-8f0d-4d86b6ee860f","d1fd2cef-ac30-42f0-a62b-49416418db4d"]', false, false);

-- Update user conversations
UPDATE public.users
SET conversations = '["conv-1","conv-3","conv-6"]'
WHERE id = '04cabad4-5f5f-4245-9b80-cf49b26dfc49';

UPDATE public.users
SET conversations = '["conv-1","conv-2","conv-4"]'
WHERE id = '40eebcec-dc43-4668-bc94-ec480f49481d';

UPDATE public.users
SET conversations = '["conv-2","conv-5","conv-6"]'
WHERE id = '4e8ce455-e59e-4bf6-88ba-defbe325f21c';

UPDATE public.users
SET conversations = '["conv-3"]'
WHERE id = '7a105aeb-9fe9-4812-827d-40d496c58d3e';

UPDATE public.users
SET conversations = '["conv-4","conv-5"]'
WHERE id = 'a00c6143-67e2-42db-8351-d8f05c819e9a';

UPDATE public.users
SET conversations = '["conv-7"]'
WHERE id = 'b8248036-6b03-4e1e-8f0d-4d86b6ee860f';

UPDATE public.users
SET conversations = '["conv-7"]'
WHERE id = 'd1fd2cef-ac30-42f0-a62b-49416418db4d';

-- Seed messages
INSERT INTO public.messages (id, author_id, conversation_id, content)
VALUES
    ('msg-1', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'conv-1', 'Hi John, welcome to the platform!'),
    ('msg-2', '40eebcec-dc43-4668-bc94-ec480f49481d', 'conv-1', 'Thanks! Excited to be here.'),
    ('msg-3', '40eebcec-dc43-4668-bc94-ec480f49481d', 'conv-2', 'Hey Jane, how''s your new design coming along?'),
    ('msg-4', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'conv-2', 'It''s going well! Should be finished by tomorrow.'),
    ('msg-5', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'conv-3', 'We need to review the new user signups.'),
    ('msg-6', '7a105aeb-9fe9-4812-827d-40d496c58d3e', 'conv-3', 'I''ll prepare a report by end of day.'),
    ('msg-7', '40eebcec-dc43-4668-bc94-ec480f49481d', 'conv-4', 'Let''s discuss the project timeline.'),
    ('msg-8', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'conv-4', 'I can work on the design part.'),
    ('msg-9', 'a00c6143-67e2-42db-8351-d8f05c819e9a', 'conv-4', 'I''ll handle the content creation.'),
    ('msg-10', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'conv-5', 'Alex, can you share those travel photos?'),
    ('msg-11', 'a00c6143-67e2-42db-8351-d8f05c819e9a', 'conv-5', 'Sure! Will send them tonight.'),
    ('msg-12', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'conv-6', 'Hey Tester, how are things going with the QA?'),
    ('msg-13', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'conv-6', 'Going well! Found a couple of minor bugs that I''ve documented.'),
    ('msg-14', 'b8248036-6b03-4e1e-8f0d-4d86b6ee860f', 'conv-7', 'Hi Nam, have you seen the latest updates to the API?'),
    ('msg-15', 'd1fd2cef-ac30-42f0-a62b-49416418db4d', 'conv-7', 'Yes, I''ve been working on integrating them into our service.');

-- Seed likes with FIXED structure
-- Using separate post_id and comment_id columns
INSERT INTO public.likes (id, author_id, post_id, comment_id)
VALUES
    ('like-1', '40eebcec-dc43-4668-bc94-ec480f49481d', 'post-1', NULL),
    ('like-2', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'post-1', NULL),
    ('like-3', '7a105aeb-9fe9-4812-827d-40d496c58d3e', 'post-1', NULL),
    ('like-4', 'a00c6143-67e2-42db-8351-d8f05c819e9a', 'post-1', NULL),
    ('like-5', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'post-2', NULL),
    ('like-6', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'post-2', NULL),
    ('like-7', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'post-3', NULL),
    ('like-8', '40eebcec-dc43-4668-bc94-ec480f49481d', 'post-3', NULL),
    ('like-9', '7a105aeb-9fe9-4812-827d-40d496c58d3e', 'post-3', NULL),
    ('like-10', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', NULL, 'comment-1'),
    ('like-11', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', NULL, 'comment-1'),
    ('like-12', '40eebcec-dc43-4668-bc94-ec480f49481d', NULL, 'comment-4');

-- Seed notifications with FIXED structure
INSERT INTO public.notifications (id, actor_id, target_user_id, post_id, comment_id, action, is_read)
VALUES
    ('notif-1', '40eebcec-dc43-4668-bc94-ec480f49481d', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'post-1', NULL, 'like', false),
    ('notif-2', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', 'post-1', NULL, 'comment', true),
    ('notif-3', '7a105aeb-9fe9-4812-827d-40d496c58d3e', '40eebcec-dc43-4668-bc94-ec480f49481d', 'post-2', NULL, 'comment', false),
    ('notif-4', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'post-3', NULL, 'like', true),
    ('notif-5', '04cabad4-5f5f-4245-9b80-cf49b26dfc49', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'post-3', NULL, 'comment', true),
    ('notif-6', '40eebcec-dc43-4668-bc94-ec480f49481d', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', NULL, NULL, 'follow', false),
    ('notif-7', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', '40eebcec-dc43-4668-bc94-ec480f49481d', NULL, NULL, 'follow', true),
    ('notif-8', 'a00c6143-67e2-42db-8351-d8f05c819e9a', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'post-3', NULL, 'comment', false);

-- Seed social requests
INSERT INTO public.social_requests (id, author_id, account_id)
VALUES
    ('req-1', '7a105aeb-9fe9-4812-827d-40d496c58d3e', '40eebcec-dc43-4668-bc94-ec480f49481d'),
    ('req-2', 'a00c6143-67e2-42db-8351-d8f05c819e9a', '04cabad4-5f5f-4245-9b80-cf49b26dfc49'),
    ('req-3', '4e8ce455-e59e-4bf6-88ba-defbe325f21c', 'a00c6143-67e2-42db-8351-d8f05c819e9a');