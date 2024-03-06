# Event Booking API

## Overview

Comprehensive Go programming course (Go - The Complete Guide), covering fundamentals to advanced topics. Includes practical project: building an event booking API.

## Installation

To set up the Event Booking API on your local machine, follow these steps:

1. **Clone the Repository**

   ```sh
   git clone https://github.com/thatguylionel/event-booking-api
   ```

2. **Install Dependencies**
   Navigate to the project directory and install the required dependencies:

   ```sh
   cd event-booking-api
   go mod tidy
   ```

3. **Run the Application**

   ```sh
   go run .
   ```

## Usage

The API endpoints include:

- **POST /events** - Creates a new event. Requires authentication.
- **PUT /events/:id** - Updates an existing event identified by its ID. Requires authentication.
- **DELETE /events/:id** - Deletes an event identified by its ID. Requires authentication.
- **POST /events/:id/register** - Registers the authenticated user for an event identified by its ID.
- **DELETE /events/:id/register** - Cancels the registration of the authenticated user for an event identified by its ID.

## Acknowledgments

- Special thanks to Maximilian Schwarzmüller for conducting such a well orchestrated course! 
[Go - The Complete Guide
](https://www.udemy.com/course/go-the-complete-guide)

## depedencies used for this application

- **Web Framework**: `github.com/gin-gonic/gin`
- **Database**: `github.com/mattn/go-sqlite3`
- **Cryptography**: `golang.org/x/crypto`
- **Authentication**: `github.com/golang-jwt/jwt/v5`

To install a dependency, simply run the `go get` command, for example:
`go get -u github.com/gin-gonic/gin`
