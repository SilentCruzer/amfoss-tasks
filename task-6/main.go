package main
import (
	"context"
	"fmt"	
	"github.com/vartanbeno/go-reddit/reddit"
	"log"
	)	
		
var ctx = context.Background()

func main() {	
    if err := run(); err != nil {
		log.Fatal(err)		
	}	
}

func run() (err error) {			
    withCredentials := reddit.WithCredentials("id", "secret", "user", "password")
    client, _ := reddit.NewClient(withCredentials)
	posts, _, err := client.Subreddit.TopPosts(ctx, "memes", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		Time: "week",
	})

	if err != nil {
		return
	}

	for _, post := range posts {
		_, err := client.Post.Upvote(context.Background(),"t3_"+post.ID)	
	if err != nil {
    	return err
	}
	
}
	fmt.Println("Upvoted last weeks 100 posts")
	return
}
