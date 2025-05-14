# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Segments is a Go-based web application for tracking and managing segments (rectangles) by size, type, or color. It provides functionality for adding, moving, removing, and reactivating segments. Removed segments are marked as inactive rather than deleted from the database.

## Architecture

- **Web Framework**: Gin (github.com/gin-gonic/gin)
- **Database**: SQLite with GORM ORM
- **Template Engine**: Go's standard text/template
- **Authentication**: Session-based using gin-contrib/sessions

### Core Components

1. **Models**: Data structures and database operations
   - `models/connect.go`: Database connection setup
   - `models/segments.go`: Segment-related models and operations
   - `models/colors.go`: Color type and color models
   - `models/locations.go`: Company, Section, and Rack models and queries
   - `models/users.go`: User authentication and management
   - `models/forms.go`: Form structures for validation
   - `models/validators.go`: Custom validation functions

2. **Controllers**: Request handlers
   - `controllers.go`: Main application controllers
   - `admin/*.go`: Admin panel controllers

3. **Views**: HTML templates using Gin's template engine
   - `templates/app/`: Frontend templates
   - `templates/admin/`: Admin interface templates

4. **Static Assets**: CSS, JavaScript, and other static files

### User Roles

- **Regular Users**: Can view and manage segments
- **Superusers**: Have access to the admin panel

## Development Commands

### Building and Running Locally

```shell
# Build the application
go build -o segments

# Run the application
./segments
```

### Docker Build and Run

```shell
# Build Docker image
docker build -t segments .

# Run Docker container
docker run -p 8080:8080 segments
```

### Database Migrations

Database migrations are handled through GORM's AutoMigrate feature. The migration code in `models/connect.go` is disabled by default (the condition `if 1 == 0`). To run migrations, temporarily change this condition and run the application.

```go
// Enable migrations by changing this line in models/connect.go
if 1 == 1 {  // Change from 1 == 0 to 1 == 1
    db.AutoMigrate(
        &ColorType{},
        &Color{},
        &Company{},
        &Section{},
        &Rack{},
        &OrderNumber{},
        &Segment{},
        &User{},
    )
}
```

## Main Data Structure

The application revolves around these key entities:

1. **Companies**: Top-level organization
2. **Sections**: Subdivisions within a company
3. **Racks**: Physical storage locations within a section
4. **Segments**: The actual rectangles being tracked, containing:
   - Dimensions (width, height)
   - Color information
   - Location (rack)
   - Status (active, defective)
   - Order number (for removed segments)

## Contributing Guidelines

When modifying code:

1. Follow the existing code structure and naming conventions
2. Test changes locally before committing
3. For admin-related functionality, use the admin package
4. For new model fields, update the corresponding template files