package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/lucaiatropulus/social/internal/dao"
	"github.com/lucaiatropulus/social/internal/store"
)

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	userRole, err := store.Roles.GetRoleByName(ctx, "user")

	if err != nil {
		log.Println("Error getting roles: ", err)
		return
	}

	users := generateUsers(100, userRole.ID)

	transaction, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, transaction, user); err != nil {
			_ = transaction.Rollback()
			log.Println("Error creating user: ", err)
			return
		}
	}

	transaction.Commit()

	posts := generatePosts(200, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating post: ", err)
			return
		}
	}

	log.Println("Seeding has been completed")
}

func generateUsers(num int, roleID int64) []*dao.User {
	usernames := []string{"marius", "george", "andrei", "alex", "georgi", "ioana", "cristiana", "mihaela", "alina", "cezar", "cristi", "marian"}
	users := make([]*dao.User, num)

	for i := 0; i < num; i++ {
		users[i] = &dao.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			RoleID:   roleID,
		}
	}

	return users
}

func generatePosts(num int, users []*dao.User) []*dao.Post {
	titles := []string{
		"Amazing Guide to Go",
		"Ultimate Docker Tips",
		"Beginner's Guide to Web Development",
		"Expert Cloud Computing Strategies",
		"Complete Machine Learning Roadmap",
		"Advanced AI Trends in 2024",
		"Essential Data Science Concepts",
		"Practical Cybersecurity Tips",
		"Comprehensive Blockchain Basics",
		"Definitive DevOps Best Practices",
		"Go Programming: From Beginner to Expert",
		"How to Build Scalable Web Apps",
		"The Future of AI and Machine Learning",
		"Cybersecurity Threats You Should Know",
		"Mastering SQL for Data Analysis",
		"Top 10 Programming Languages in 2024",
		"Understanding Kubernetes and Containers",
		"Blockchain Explained: A Simple Guide",
		"Effective Debugging Techniques for Developers",
		"Microservices Architecture: Pros and Cons",
		"Building REST APIs with Go",
		"Data Structures and Algorithms Simplified",
		"Web3 and the Future of the Internet",
		"How to Improve Your Coding Skills",
		"Best VS Code Extensions for Developers",
		"The Power of Functional Programming",
		"Serverless Computing: What You Need to Know",
		"Cybersecurity Best Practices for Businesses",
		"Mastering Git and GitHub for Collaboration",
		"The Rise of AI Chatbots",
		"Why Go Is the Future of Backend Development",
		"Cloud Computing Trends in 2024",
		"Introduction to GraphQL for API Development",
		"The Importance of Testing in Software Development",
		"How to Write Clean and Maintainable Code",
		"Getting Started with Machine Learning in Python",
		"Best Open-Source Tools for Developers",
		"Introduction to Rust: A Safer Systems Language",
		"Scaling Applications with Load Balancing",
		"Building a Simple CRUD App in Go",
		"Understanding OAuth for Authentication",
		"Common Mistakes Every Programmer Should Avoid",
		"Automating Tasks with Python Scripts",
		"Artificial Intelligence vs Machine Learning",
		"The Future of Remote Work for Developers",
		"How to Ace Your Technical Interviews",
		"Cloud Security Best Practices",
		"Top 5 Books Every Programmer Should Read",
		"Why You Should Learn Go in 2024",
		"Software Engineering Career Paths Explained",
	}

	tags := []string{
		"go", "golang", "docker", "kubernetes", "microservices",
		"webdev", "backend", "frontend", "fullstack", "main",
		"restapi", "graphql", "database", "sql", "nosql",
		"postgresql", "mongodb", "mysql", "redis", "caching",
		"devops", "ci/cd", "git", "github", "gitlab",
		"cloud", "aws", "gcp", "azure", "serverless",
		"machinelearning", "ai", "deeplearning", "datascience", "bigdata",
		"security", "authentication", "jwt", "oauth", "encryption",
		"performance", "scalability", "architecture", "designpatterns", "bestpractices",
		"testing", "unittesting", "integrationtesting", "logging", "monitoring",
	}

	posts := make([]*dao.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &dao.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: titles[rand.Intn(len(titles))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*dao.User, posts []*dao.Post) []*dao.Comment {
	comms := []string{
		"Great article, very informative!",
		"Thanks for sharing, this really helped.",
		"I didn't understand the part about authentication. Can you explain?",
		"Awesome guide! Keep up the good work.",
		"This solved my issue. Much appreciated!",
		"Can you provide a code example for this?",
		"This method seems outdated. Any newer alternatives?",
		"Well written and easy to follow.",
		"Not sure if this works in the latest Go version.",
		"How does this compare to using Python?",
		"I love using Go for backend development!",
		"Security is crucial—please add a section on best practices.",
		"Why use Redis over PostgreSQL for caching?",
		"I tried this but got an error. Any suggestions?",
		"Clear and concise explanation. Thanks!",
		"Do you have a GitHub repo for this?",
		"Please write more about microservices architecture.",
		"This was exactly what I needed. Thanks!",
		"I think there's a typo in your code snippet.",
		"What about performance benchmarks?",
		"Can this be used with Docker and Kubernetes?",
		"Looking forward to more posts like this!",
		"This approach is good, but I prefer using GraphQL.",
		"How do you handle error logging in this setup?",
		"This article could use some practical examples.",
		"Nice work! What resources do you recommend for learning more?",
		"This worked perfectly for my project. Thank you!",
		"I found an issue when running this on Windows.",
		"Can you compare this with an alternative approach?",
		"I wish you covered more about database optimizations.",
		"Your explanation was super helpful!",
		"How would you scale this for a high-traffic app?",
		"This needs an update for Go 1.20.",
		"Please add unit tests to the example.",
		"This was a lifesaver! Thank you.",
		"Why did you choose this method over another?",
		"Can you add more comments in your code examples?",
		"This helped me debug my API. Thanks!",
		"I'm having trouble implementing this. Any tips?",
		"Your articles are always helpful!",
		"This would be great as a video tutorial too.",
		"Can you cover error handling in more detail?",
		"Would this work with gRPC as well?",
		"Please provide a performance comparison with another method.",
		"I learned something new today. Thanks!",
		"This doesn't work with my setup. Any ideas?",
		"Clear and well-structured explanation.",
		"I appreciate the effort you put into this!",
		"Any recommended tools for debugging Go applications?",
		"Looking forward to more advanced topics!",
	}

	comments := make([]*dao.Comment, num)

	for i := 0; i < num; i++ {
		comments[i] = &dao.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comms[rand.Intn(len(comms))],
		}
	}

	return comments
}
