package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Define the skills you are looking for
	skills := []string{"Golang", "backend", "Typescript", "frontend developer", "React"}

	// Indeed URL for remote jobs
	url := "https://www.indeed.com/jobs?q=full+stack+developer+entry+level&l=Remote&vjk=7e5da91fb9690341"

	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	// Find and loop through each job listing
	doc.Find(".jobsearch-SerpJobCard").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title").Text()
		location := s.Find(".location").Text()
		summary := s.Find(".summary").Text()

		// Check if the job listing contains any of the required skills
		if containsSkills(title, summary, skills) {
			fmt.Printf("Title: %s\nLocation: %s\nSummary: %s\n\n", title, location, summary)
		}
	})
}

// Helper function to check if a job listing contains any of the required skills
func containsSkills(title, summary string, skills []string) bool {
	title = strings.ToLower(title)
	summary = strings.ToLower(summary)

	for _, skill := range skills {
		if strings.Contains(title, strings.ToLower(skill)) || strings.Contains(summary, strings.ToLower(skill)) {
			return true
		}
	}
	return false
}
