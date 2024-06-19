package handler

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/xmriz/go-scraper/internal/api"
	"github.com/xmriz/go-scraper/internal/model"
)

// capitalizeFirstLetter capitalizes the first letter of a string
func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

// Format user output
func FormatUserOutput(workerID int, user model.UserFull) string {
	return fmt.Sprintf("Worker %-d - User:\n"+
		"\tNama   : %s. %s %s\n"+
		"\tEmail  : %s\n"+
		"\tGender : %s\n\n",
		workerID, capitalizeFirstLetter(user.Title), user.FirstName, user.LastName, user.Email, capitalizeFirstLetter(user.Gender))
}

// Format post output
func FormatPostOutput(workerID int, post model.Post) string {
	user := fmt.Sprintf("%s %s", post.Owner.FirstName, post.Owner.LastName)
	likes := strconv.Itoa(post.Like)
	tags := strings.Join(post.Tags, ", ")

	return fmt.Sprintf("Worker %-d - Post:\n"+
		"\tPosted by    : %s\n"+
		"\tText         : %s\n"+
		"\tLikes        : %s\n"+
		"\tTags         : %s\n"+
		"\tDate posted  : %s\n\n",
		workerID, user, post.Text, likes, tags, post.PublishDate)
}

// ScrapeData function to scrape user and post data concurrently
func ScrapeData(appID string, workers int) error {
	start := time.Now()

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(2)
	go func() {
		defer wg.Done()
		ScrapeUsers(appID, workers)
	}()

	go func() {
		defer wg.Done()
		ScrapePosts(appID, workers)
	}()

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("\nScraping completed...")
	fmt.Printf("Total time taken with %d workers: %v\n", workers, duration)

	return nil
}

func ScrapeUsers(appID string, workers int) {
	var wg sync.WaitGroup
	pageChan := make(chan int, 10)

	// Create worker pool (Process)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for page := range pageChan {
				users, err := api.FetchUsers(page, appID)
				if err != nil {
					log.Println(err)
					continue
				}

				for _, userPreview := range users {
					userDetail, err := api.FetchUserDetail(userPreview.ID, appID)
					if err != nil {
						log.Println(err)
						continue
					}
					fmt.Print(FormatUserOutput(workerID, userDetail))
				}
			}
		}(i + 1) // i + 1 untuk memberikan ID worker yang dimulai dari 1
	}

	// Send pages to worker pool (Dispatcher)
	for page := 1; page <= 10; page++ {
		pageChan <- page
	}
	close(pageChan)

	wg.Wait()
}

func ScrapePosts(appID string, workers int) {
	var wg sync.WaitGroup
	pageChan := make(chan int, 10)

	// Create worker pool (Process)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for page := range pageChan {
				posts, err := api.FetchPosts(page, appID)
				if err != nil {
					log.Println(err)
					continue
				}

				for _, post := range posts {
					fmt.Print(FormatPostOutput(workerID, post))
				}
			}
		}(i + 1) // i + 1 untuk memberikan ID worker yang dimulai dari 1
	}

	// Send pages to worker pool (Dispatcher)
	for page := 1; page <= 10; page++ {
		pageChan <- page
	}
	close(pageChan)

	wg.Wait()
}
