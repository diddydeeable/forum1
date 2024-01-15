package model

import (
	"time"
)


type Comment struct{
	Id int
	Uuid string
	Topic string
	UserID int
	CreatedAt time.Time
}

type Post struct{
	Id int
	Uuid string
	Body string
	UserID int
	ThreadID int
	CreatedAt time.Time
}

// format the CreatedAt date to display nicely on the screen
func (comment *Comment) CreatedAtDate() string {
	return comment.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}


//get the number of posts in a thread
func (comment *Comment) NumComments() (count int) {
	rows, err := db.Query("SELECT count(*) FROM posts WHERE thread_id = ?", comment.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

