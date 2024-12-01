# Web Scraper

## Overview
The **Web Scraper** project is a Go-based application designed to efficiently fetch and process articles from multiple news websites. Using Go's concurrency features, it scrapes multiple pages in parallel and provides the collected data through a REST API.

---

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
- [Author](#author)


## Features
- **Web Scraper**: Extracts article titles, links, and timestamps from multiple news platforms.
- **Concurrency**: Utilizes Go's goroutines to scrape multiple pages simultaneously for better performance.
- **REST API**: Serves the scraped data in JSON format through a user-friendly endpoint.
- **Rate Limiting**: Protects the API by controlling the frequency of incoming requests.
- **Automatic Data Refresh**: Periodically re-scrapes the news data at configurable intervals (e.g., every 5 minutes).

## Requirements
- Go 1.23.3 or later
- PostgreSQL database

## Installation
1. **Clone the repository**:
    ```bash
    git clone https://github.com/Mohammed-Nasif/web-scraper.git
    cd web-scraper
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Set up environment variables**:  
   Create a `.env` file in the root directory with the following content:
    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=yourusername
    DB_PASSWORD=yourpassword
    DB_NAME=newsdb
    DB_SSLMODE=disable
    PORT=8080
    BASE_URL=http://localhost
    ```

4. **Run the application**:
    ```bash
    go run .
    ```


## Usage
After starting the application:
1. The scraper will begin fetching articles from the configured news websites.
2. Access the REST API to retrieve the scraped data in JSON format.


## Endpoints
- **GET /articles**: Returns a JSON array of the scraped articles, including:
  - Title
  - Link
  - Timestamp  

Example Response:
```json
[
  {
    "title": "Example Article",
    "link": "https://example.com/article",
    "timestamp": "2024-12-01T10:00:00Z"
  }
]
```

## Environment Variables
The following environment variables need to be configured in the `.env` file:

| **Variable**   | **Description**                            |
|-----------------|-------------------------------------------|
| `DB_HOST`       | The hostname of the database server       |
| `DB_PORT`       | The port number for the database          |
| `DB_USER`       | The username for database authentication  |
| `DB_PASSWORD`   | The password for database authentication  |
| `DB_NAME`       | The name of the database to use           |
| `DB_SSLMODE`    | SSL mode for the database connection      |
| `PORT`          | The port on which the application runs    |
| `BASE_URL`      | The base URL for the application          |

## Project Structure
The project is organized as follows:
```
  web-scraper/
  ├── .env
  ├── controllers/
  │   └── articles_controller.go
  ├── db/
  │   ├── articles.go
  │   └── db.go
  ├── middlewares/
  │   └── rateLimiter.go
  ├── models/
  │   └── article.go
  ├── routes/
  │   └── routes.go
  ├── scraper/
  │   ├── arstechnica.go
  │   ├── scraper.go
  │   ├── techcrunch.go
  │   └── theverge.go
  ├── services/
  │   ├── articles_service.go
  │   └── scraper_service.go
  ├── go.mod
  ├── go.sum
  ├── main.go
  └── Requirements.md
```
## Author
Developed by **Mohammed Nasif**.
<br>
If you have any questions, suggestions, or feedback, feel free to reach out.

