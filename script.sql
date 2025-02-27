-- Bảng Role
CREATE TABLE role (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    role_name VARCHAR(100) NOT NULL,
    active_status BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Bảng User
CREATE TABLE "user" (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    role_id VARCHAR(100),
    full_name VARCHAR(255) NOT NULL,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    date_of_birth DATE,
    profile_avatar VARCHAR(255),
    bio TEXT,
    friends TEXT,
    followers TEXT,
    followings TEXT,
    block_users TEXT,
    conversations TEXT,
    is_private BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    is_activated BOOLEAN DEFAULT FALSE,
    is_have_to_reset_password BOOLEAN DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_User_Role FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE
);

-- Bảng User Security (1-1 với User)
CREATE TABLE user_security (
    user_id VARCHAR(100) PRIMARY KEY NOT NULL,
    access_token TEXT,
    refresh_token TEXT,
    action_token TEXT,
    fail_access INT DEFAULT 0,
    last_fail TIMESTAMP,
    CONSTRAINT FK_UserSecurity_User FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

-- Bảng Social Request
CREATE TABLE social_request (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    author_id VARCHAR(100) NOT NULL,
    account_id VARCHAR(100) NOT NULL,
    status VARCHAR(10) DEFAULT 'pending' CHECK (status IN ('pending', 'accepted', 'rejected')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_SocialRequest_Sender FOREIGN KEY (author_id) REFERENCES "user"(id) ON DELETE CASCADE,
    CONSTRAINT FK_SocialRequest_Receiver FOREIGN KEY (account_id) REFERENCES "user"(id) ON DELETE CASCADE
);

-- Bảng Post
CREATE TABLE post (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    author_id VARCHAR(100),
    content TEXT NOT NULL,
    is_private BOOLEAN NOT NULL,
    is_hidden BOOLEAN NOT NULL,
    status BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_Post_User FOREIGN KEY (author_id) REFERENCES "user"(id) ON DELETE CASCADE
);

-- Bảng Like
CREATE TABLE "like" (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    author_id VARCHAR(100),
    object_id VARCHAR(100),
    object_type VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_Like_User FOREIGN KEY (author_id) REFERENCES "user"(id) ON DELETE CASCADE
);

-- Bảng Notification
CREATE TABLE notification (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    actor_id VARCHAR(100),
    object_id VARCHAR(100),
    object_type VARCHAR(10),
    action VARCHAR(50),
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_Notification_User FOREIGN KEY (actor_id) REFERENCES "user"(id) ON DELETE CASCADE
);

-- Bảng Conversation
CREATE TABLE conversation (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(200),
    host_id VARCHAR(100),
    members VARCHAR(500),
    is_group BOOLEAN,
    is_delete BOOLEAN DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_Conversation_User FOREIGN KEY (host_id) REFERENCES "user"(id) ON DELETE CASCADE
);

-- Bảng Message
CREATE TABLE message (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    author_id VARCHAR(100) NOT NULL,
    conversation_id VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT Fk_Message_User FOREIGN KEY (author_id) REFERENCES "user"(id) ON DELETE CASCADE,
    CONSTRAINT Fk_Message_Conversation FOREIGN KEY (conversation_id) REFERENCES conversation(id) ON DELETE CASCADE
);
