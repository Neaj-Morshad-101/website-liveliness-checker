package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// readWebsitesFromFile reads website URLs from a file.
func readWebsitesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %v", filename, err)
	}
	defer file.Close()

	var websites []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			websites = append(websites, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filename, err)
	}
	return websites, nil
}

// checkWebsite checks if a website is up using an HTTP GET request and logs errors.
func checkWebsite(client *http.Client, url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Error checking %s: %v", url, err) // Print error as requested
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true
	}
	log.Printf("Website %s returned non-success status code: %d", url, resp.StatusCode) // Log non-success status
	return false
}

func main() {
	websites, err := readWebsitesFromFile("websites.txt")
	if err != nil {
		log.Fatalf("Error reading websites file: %v", err)
	}
	if len(websites) == 0 {
		log.Println("No websites to check.")
		return
	}

	// Configure HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	upCount := 0

	// Check each website concurrently
	for _, url := range websites {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			if checkWebsite(client, u) {
				mu.Lock()
				upCount++
				mu.Unlock()
			}
		}(url)
	}

	// Wait for all checks to complete
	wg.Wait()

	// Print summary
	fmt.Printf("Total websites checked: %d\n", len(websites))
	fmt.Printf("Websites Up: %d\n", upCount)
	fmt.Printf("Websites Down: %d\n", len(websites)-upCount)
}
