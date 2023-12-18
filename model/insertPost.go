package model
import(
	"time"
)
	// CreatePost inserts a new post into the database and returns the post ID.
	func InsertPost(username,title, category, body string, userID int) error {
		
		// Prepare the SQL statement to insert a new post.
		stmt, err := db.Prepare("INSERT INTO Posts (username, title, body, category, created_at) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
	
		// Get the current timestamp.
		createdAt := time.Now().Format("2006-01-02 15:04:05")
	
		// Execute the SQL statement to insert the new post.
		_, err = stmt.Exec(username, title, category, body, createdAt)
		if err != nil {
			return err
		}
	
		return nil
	}