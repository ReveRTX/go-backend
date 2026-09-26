# Backend Health Check API

 A simple Go backend that provides a health-check endpoint to verify that the server is running successfully.

 ## Features

 - Built with Go's standard library.
- Runs an HTTP server on port `8000`.
- Provides a `/health` endpoint.
- Returns a JSON health status.
- Generates timestamps in **IST (Asia/Kolkata)**.
- Logs health-check requests and server status.

 ## Endpoint

 ### `GET /health`

 Returns:

```
{
  "status": "healthy",
  "message": "Backend is running",
  "timestamp": "2026-09-26T10:38:00+05:30"
}
```

 ## Run

```
go run main.go
```

 The server will be available at:

```
http://localhost:8000
```

 Health check:

```
http://localhost:8000/health
```

 ## Requirements

 - Go 1.18+
- No external dependencies

 ## Timezone

 The backend uses `Asia/Kolkata`, which represents **IST (UTC+05:30)** and applies to locations such as Mumbai and Kolkata.