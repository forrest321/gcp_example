# Social Media Trending Topics Scraper

## Overview
This application is designed to scrape trending topics from various social media platforms. It's built to run on Google Cloud Platform (GCP), leveraging serverless architecture for efficient and scalable operation. The application is divided into two primary components: a backend for scraping and processing data, and a frontend for displaying the results.

## Technology Stack
- **Backend**: Written in Go (Golang), hosted on Google Cloud Functions.
- **Frontend**: A React-based web application.
- **Data Storage**: SQL and NoSQL databases for managing scraped data. Depending on the nature of data, Cloud SQL or Firestore can be used.
- **Cloud Hosting**: Google Cloud Platform (GCP) services, including Cloud Functions, Cloud Scheduler for scheduling tasks, and Pub/Sub for event-driven processing.
- **Web Scraping**: Using `go-rod/rod` for scraping dynamic web content.

## Key Packages and Libraries
- `github.com/go-rod/rod`: A Golang library for web scraping, capable of handling JavaScript-heavy websites using headless Chrome.
- `gopkg.in/yaml.v2`: YAML processing library in Go, used for parsing configuration files.
- `github.com/sirupsen/logrus`: Logging library for Go, used for structured logging.
- Additional Go standard libraries for HTTP handling, file I/O, etc.
