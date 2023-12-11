CREATE TABLE IF NOT EXISTS Users
(
    UserID INTEGER PRIMARY KEY AUTOINCREMENT,
    Username TEXT NOT NULL UNIQUE,
    Email TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Posts
(
    PostID INTEGER PRIMARY KEY AUTOINCREMENT,
    UserID INTEGER,
    Username TEXT,
    Title TEXT,
    Category TEXT,
    Body TEXT,
    Likes INTEGER,
    Dislikes INTEGER,
    Comments TEXT,
    CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Comments
(
    CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
    Username TEXT NOT NULL,
    PostID INTEGER NOT NULL,
    Body TEXT NOT NULL,
    CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Categories
(
    CategoryID INTEGER PRIMARY KEY AUTOINCREMENT,
    Name TEXT NOT NULL,
    Description TEXT
);

CREATE TABLE IF NOT EXISTS Likes
(
    LikeID INTEGER PRIMARY KEY AUTOINCREMENT,
    UserID INTEGER,
    PostID INTEGER,
    CommentID INTEGER
);

CREATE TABLE IF NOT EXISTS Dislikes
(
    DislikeID INTEGER PRIMARY KEY AUTOINCREMENT,
    UserID INTEGER,
    PostID INTEGER,
    CommentID INTEGER
);
