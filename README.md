# Keploy-Session-2

## Description

A custom RESTful API server built with Go and PostgreSQL. It includes endpoints to perform CRUD operations on user data and a React frontend that interacts with these APIs.

## Project Overview

This project is a scalable full-stack user management system designed to perform CRUD operations efficiently and provide a seamless interface for data handling. It features:
  - A Go-based REST API for managing user data and handling client requests.
  - PostgreSQL as the relational database for storing user information with GORM ORM support.
  - A React frontend that communicates with the backend to create, retrieve, update, and delete users.
  - Modular backend architecture with clearly separated layers for database logic, API routing, and business logic.

## Functions

  - `User Management`: Create, retrieve, update, and delete user records with unique IDs.
  - `RESTful API`: Exposes clean API endpoints to interact with user data using standard HTTP methods.
  - `Database Integration`: Persists user information using PostgreSQL with GORM ORM.
  - `Frontend Integration`: React-based interface to consume APIs and manage users interactively.

## Technologies Used

  - Backend: Go (Golang) with Gorilla Mux
  - Database: PostgreSQL with GORM ORM
  - Frontend: React (with Axios for API requests)
  - API Design: RESTful architecture

## Prerequisites

  - Go 1.20+ installed
  - PostgreSQL (local setup or managed service like Neon DB)
  - Node.js and npm (for running the React frontend)

## Database Initialization

  Before starting the application, ensure the PostgreSQL database is properly set up. The following steps must be completed:
    - A database named go-server must be created.
    - Within this database, a table named users will be used to store user records, including:
      - `id` (UUID or auto-incremented)
      - `name` (string)
      - `email` (string)

  *Note: If the table does not exist at runtime, it will be automatically created using GORM's AutoMigrate functionality. However, pre-creating the database is recommended for proper setup and access control.*

## Implementation

1. Clone the repository
     ```
     git clone https://github.com/ananyab1909/custom-api-server.git
     ```

2. Enter into the directory
     ```
     cd PepSales-Task
     ```

3. Start the backend server - Make sure PostgreSQL is running locally or configure your Neon DB connection in db.Connect() function.
    ```
    go run main.go
    ```

4. Start the frontend (React)
    ```
    cd frontend
    npm start
    ```

5. Access the application
  ```
  Backend API: http://localhost:8080
  Frontend UI: http://localhost:3000
  ```

## User Routes

  The API Documentation is provided in API.md

## About Me

Hello, my name is Ananya Biswas. I am an Engineering Student at [Kalinga Institute of Industrial Technology](https://kiit.ac.in/). I enjoy making projects and now that my this project is over, I am open-sourcing the project. Hope you like it! Lastly, I would like to put it out there that I have worked on other projects that you may like. You can check them out at my [Github](https://github.com/ananyab1909/). Give it a whirl and let me know your thoughts.

## Socials
  - Portfolio : https://dub.sh/ananyabiswas
  - LinkedIn : https://linkedin.com/in/ananya-biswas-kiit/
  - Mastodon : https://mastodon.social/@captain_obvious/
  - Twitter : https://x.com/not_average_x/
  - Github : https://github.com/ananyab1909/

 
