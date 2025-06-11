# URL Shortener

This is a simple URL shortener service written in Go using the Templ templating engine. It allows users to shorten long URLs into shorter, more manageable links.

## Features

- Shorten URLs
- Redirect to original URLs
- Customizable short URLs (future implementation)
- Simple and clean user interface

## Technologies Used

- Go
- Templ
- PostgreSQL (default, configurable)

## Installation

1.  Clone the repository:

    ```bash
    git clone <repository_url>
    cd shortenurl
    ```

2.  Install dependencies:

    ```bash
    go mod download
    ```

3.  Configure the database:

    - The application uses PostgreSQL by default. You can configure the database connection in the `config/db.go` file.
    - To use a different database, you'll need to update the database connection details and potentially modify the `models/url.go` file.

4.  Set environment variables:

    - Create a `.env` file in the root directory.
    - Add the following environment variables:

      ```
      DATABASE_URL=postgresql://username:password@host:port/database
      PORT=8080
      ```

5.  Run the application:

    ```bash
    go run main.go
    ```

## Usage

1.  Open your web browser and go to `http://localhost:8080` (or the port you configured in the `.env` file).
2.  Enter the URL you want to shorten in the input field.
3.  Click the "Shorten" button.
4.  The shortened URL will be displayed on the results page.

## Project Structure

```
.
├── .env                    # Environment variables
├── .gitignore              # Git ignore file
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies file
├── main.go                 # Main application file
├── readme.md               # This file
├── config/                 # Configuration files
│   └── db.go               # Database configuration
├── controllers/            # Controllers
│   └── url_controller.go   # URL controller
├── models/                 # Data models
│   └── url.go              # URL model
├── routes/                 # Routes
│   └── routes.go           # Application routes
├── static/                 # Static files
│   └── style.css           # CSS file
├── views/                  # Templ templates
│   ├── index_templ.go      # Index template (Go file)
│   ├── index.templ         # Index template
│   ├── layout_templ.go     # Layout template (Go file)
│   ├── layout.templ        # Layout template
│   ├── list_url_templ.go   # List URL template (Go file)
│   ├── list_url.templ      # List URL template
│   ├── result_templ.go     # Result template (Go file)
│   └── result.templ        # Result template
└── ...
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

[MIT](LICENSE)
