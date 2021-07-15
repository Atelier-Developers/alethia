USE alethia;

CREATE TABLE USER
(
    id                     INT AUTO_INCREMENT NOT NULL,
    first_name             VARCHAR(100)       NOT NULL,
    last_name              VARCHAR(100)       NOT NULL,
    username               VARCHAR(50)        NOT NULL,
    password               BINARY(60)         NOT NULL,
    intro                  TEXT(255)          NOT NULL,
    about                  TEXT(1024)         NOT NULL,
    accomplishments        TEXT(1024)         NOT NULL,
    additional_information TEXT(1024)         NOT NULL,
    join_date              DATETIME DEFAULT CURRENT_TIMESTAMP,
    birthdate              DATE               NOT NULL,
#     languages
    PRIMARY KEY (id)
);

CREATE TABLE USER_HISTORY
(
    id          INT AUTO_INCREMENT                                                                         NOT NULL,
    user_id     INT                                                                                        NOT NULL,
    start_date  DATETIME                                                                                   NOT NULL,
    end_date    DATETIME,
    category    ENUM ('work_experience', 'education', 'licenses_and_certificates', 'volunteer_experience') NOT NULL,
    title       VARCHAR(255)                                                                               NOT NULL,
    description TEXT(1024)                                                                                 NOT NULL,
    location    VARCHAR(50)                                                                                NOT NULL,

    PRIMARY KEY (id),
    CONSTRAINT FK_user_user_history FOREIGN KEY (user_id) REFERENCES USER (id)

);

CREATE TABLE POST
(
    id          INT AUTO_INCREMENT NOT NULL,
    is_featured BOOLEAN            NOT NULL,
    description TEXT(1024)         NOT NULL,
    # MEDIA
    created     DATETIME DEFAULT CURRENT_TIMESTAMP,
    poster_id   INT                NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT FK_user_post FOREIGN KEY (poster_id) REFERENCES USER (id)
);


CREATE TABLE REPOST
(
    post_id   INT NOT NULL,
    repost_id INT NOT NULL,

    PRIMARY KEY (post_id, repost_id),

    CONSTRAINT FK_repost_post FOREIGN KEY (repost_id) REFERENCES POST (id),
    CONSTRAINT FK_post_post FOREIGN KEY (post_id) REFERENCES POST (id)
);


CREATE TABLE COMMENT
(
    id           INT AUTO_INCREMENT NOT NULL,
    text         TEXT(1024)         NOT NULL,
    created      DATETIME DEFAULT CURRENT_TIMESTAMP,
    commenter_id INT                NOT NULL,
    post_id      INT                NOT NULL,
    reply_id     INT                NOT NULL,

    PRIMARY KEY (id),
    CONSTRAINT FK_user_comment FOREIGN KEY (commenter_id) REFERENCES USER (id),
    CONSTRAINT FK_post_comment FOREIGN KEY (post_id) REFERENCES POST (id)
);

ALTER TABLE COMMENT
    ADD CONSTRAINT FK_reply_comment
        FOREIGN KEY (reply_id) REFERENCES COMMENT (id);


CREATE TABLE POST_LIKE
(
    post_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (post_id) REFERENCES POST (id),
    FOREIGN KEY (user_id) REFERENCES USER (id),
    created DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (post_id, user_id)
);


CREATE TABLE COMMENT_LIKE
(
    comment_id INT NOT NULL,
    user_id    INT NOT NULL,
    FOREIGN KEY (comment_id) REFERENCES COMMENT (id),
    FOREIGN KEY (user_id) REFERENCES USER (id),
    created    DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (comment_id, user_id)
);


CREATE TABLE CONVERSATION
(
    id          INT AUTO_INCREMENT NOT NULL,
    user1_id    INT                NOT NULL,
    user2_id    INT                NOT NULL,
    is_archived BOOLEAN            NOT NULL,
    is_deleted  BOOLEAN            NOT NULL,
    is_read     BOOLEAN            NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT FK_user1_conversation FOREIGN KEY (user1_id) REFERENCES USER (id),
    CONSTRAINT FK_user2_conversation FOREIGN KEY (user2_id) REFERENCES USER (id)
);

CREATE TABLE MESSAGE
(
    id              INT AUTO_INCREMENT NOT NULL,
    reply_id        INT                NOT NULL,
    conversation_id INT                NOT NULL,
    body            TEXT(1024)         NOT NULL,
    created         DATETIME DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT FK_reply_message FOREIGN KEY (reply_id) REFERENCES MESSAGE (id),
    CONSTRAINT FK_conversation_message FOREIGN KEY (conversation_id) REFERENCES CONVERSATION (id)
);


CREATE TABLE FRIEND
(
    user1_id INT NOT NULL,
    user2_id INT NOT NULL,
    created  DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_user1_friend FOREIGN KEY (user1_id) REFERENCES USER (id),
    CONSTRAINT FK_user2_friend FOREIGN KEY (user2_id) REFERENCES USER (id),

    PRIMARY KEY (user1_id, user2_id)
);

CREATE TABLE INVITE
(
    user1_id INT        NOT NULL,
    user2_id INT        NOT NULL,
    created  DATETIME DEFAULT CURRENT_TIMESTAMP,
    body     TEXT(1024) NOT NULL,
    CONSTRAINT FK_user1_invite FOREIGN KEY (user1_id) REFERENCES USER (id),
    CONSTRAINT FK_user2_invite FOREIGN KEY (user2_id) REFERENCES USER (id),

    PRIMARY KEY (user1_id, user2_id)
);

CREATE TABLE SKILL
(
    id       INT AUTO_INCREMENT                                                                 NOT NULL,
    title    VARCHAR(50)                                                                        NOT NULL,
    category ENUM ('industry_knowledge', 'tools_technologies', 'interpersonal_skills', 'other') NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE USER_SKILL
(
    skill_id INT NOT NULL,
    user_id  INT NOT NULL,
    CONSTRAINT FK_skill_user_skill FOREIGN KEY (skill_id) REFERENCES SKILL (id),
    CONSTRAINT FK_user_user_skill FOREIGN KEY (user_id) REFERENCES USER (id),

    PRIMARY KEY (skill_id, user_id)
);

CREATE TABLE ENDORSE
(
    user_skill_skill_id INT NOT NULL,
    user_skill_user_id  INT NOT NULL,
    endorser_id         INT NOT NULL,

    CONSTRAINT FK_user_skill_endorse FOREIGN KEY (user_skill_skill_id, user_skill_user_id) REFERENCES USER_SKILL (skill_id, user_id),
    CONSTRAINT FK_user_endorse FOREIGN KEY (endorser_id) REFERENCES USER (id),

    PRIMARY KEY (user_skill_skill_id, user_skill_user_id, endorser_id)
);


CREATE TABLE NOTIFICATION
(
    id          INT AUTO_INCREMENT NOT NULL,
    receiver_id INT                NOT NULL,
    created     DATETIME DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_user_notification FOREIGN KEY (receiver_id) REFERENCES USER (id),

    PRIMARY KEY (id)
);


CREATE TABLE NOTIFICATION_BIRTHDAY
(
    notif_id INT NOT NULL,
    user_id  INT NOT NULL,

    CONSTRAINT FK_user_notification_birthday FOREIGN KEY (user_id) REFERENCES USER (id),
    CONSTRAINT FK_notification_notification_birthday FOREIGN KEY (notif_id) REFERENCES NOTIFICATION (id),

    PRIMARY KEY (notif_id)
);

CREATE TABLE NOTIFICATION_VIEW_PROFILE
(
    notif_id INT NOT NULL,
    user_id  INT NOT NULL,

    CONSTRAINT FK_user_notification_view_profile FOREIGN KEY (user_id) REFERENCES USER (id),
    CONSTRAINT FK_notification_notification_view_profile FOREIGN KEY (notif_id) REFERENCES NOTIFICATION (id),

    PRIMARY KEY (notif_id)
);


CREATE TABLE NOTIFICATION_LIKE_POST
(
    notif_id INT NOT NULL,
    user_id  INT NOT NULL,
    post_id  INT NOT NULL,

    CONSTRAINT FK_user_notification_like_post FOREIGN KEY (user_id) REFERENCES USER (id),
    CONSTRAINT FK_notification_notification_like_post FOREIGN KEY (notif_id) REFERENCES NOTIFICATION (id),
    CONSTRAINT FK_post_notification_like_post FOREIGN KEY (post_id) REFERENCES POST (id),

    PRIMARY KEY (notif_id)
);


CREATE TABLE NOTIFICATION_COMMENT
(
    notif_id   INT NOT NULL,
    comment_id INT NOT NULL,

    CONSTRAINT FK_notification_notification_comment FOREIGN KEY (notif_id) REFERENCES NOTIFICATION (id),
    CONSTRAINT FK_comment_notification_comment FOREIGN KEY (comment_id) REFERENCES COMMENT (id),

    PRIMARY KEY (notif_id)
);


CREATE TABLE NOTIFICATION_LIKE_COMMENT
(
    notif_id   INT NOT NULL,
    comment_id INT NOT NULL,
    user_id    INT NOT NULL,

    CONSTRAINT FK_notification_notification_like_comment FOREIGN KEY (notif_id) REFERENCES NOTIFICATION (id),
    CONSTRAINT FK_comment_notification_like_comment FOREIGN KEY (comment_id) REFERENCES COMMENT (id),
    CONSTRAINT FK_user_notification_like_comment FOREIGN KEY (user_id) REFERENCES USER (id),


    PRIMARY KEY (notif_id)
);


CREATE TABLE NOTIFICATION_REPLY
(
    notif_id   INT NOT NULL,
    comment_id INT NOT NULL,

    CONSTRAINT FK_notification_notification_reply FOREIGN KEY (notif_id) REFERENCES NOTIFICATION (id),
    CONSTRAINT FK_comment_notification_reply FOREIGN KEY (comment_id) REFERENCES COMMENT (id),


    PRIMARY KEY (notif_id)
);

CREATE TABLE NOTIFICATION_ENDORSE
(
    notif_id            INT NOT NULL,
    user_skill_skill_id INT NOT NULL,
    user_skill_user_id  INT NOT NULL,
    user_id             INT NOT NULL,

    CONSTRAINT FK_notification_notification_endorse FOREIGN KEY (notif_id) REFERENCES NOTIFICATION (id),
    CONSTRAINT FK_user_skill_notification_endorse FOREIGN KEY (user_skill_skill_id, user_skill_user_id) REFERENCES USER_SKILL (skill_id, user_id),
    CONSTRAINT FK_user_notification_endorse FOREIGN KEY (user_id) REFERENCES USER (id),


    PRIMARY KEY (notif_id)
);

CREATE TABLE NOTIFICATION_CHANGE_WORK
(
    notif_id INT NOT NULL,
    user_id  INT NOT NULL,

    CONSTRAINT FK_user_notification_change_work FOREIGN KEY (user_id) REFERENCES USER (id),
    CONSTRAINT FK_notification_notification_change_work FOREIGN KEY (notif_id) REFERENCES NOTIFICATION (id),

    PRIMARY KEY (notif_id)
);










