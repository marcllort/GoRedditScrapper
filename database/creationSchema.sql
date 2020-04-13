create table redditDB.POSTS_TABLE
(
    id            int auto_increment
        primary key,
    subreddit     varchar(255)                         not null,
    title         varchar(255)                         not null,
    published     tinyint(1) default 0                 null,
    publishedDate datetime                             null,
    date          datetime   default CURRENT_TIMESTAMP null
);
