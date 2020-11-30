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
    withCredentials := reddit.WithCredentials("sCqmgGAk36K3sA", "aR4t3zCZOAUCev38UyM-6FXi5M2FHA", "blazerev", "waterpower")
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
