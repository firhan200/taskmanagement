# Task Management Web Application

## Task Management

This Task Management Web Application is a robust tool for managing your tasks efficiently. Built using Golang for the backend, PostgreSQL for the database, and React.js for the frontend, this application is designed to streamline your task management process.
Features
### 1. User Authentication

- **Sign Up**: New users can create an account by providing their full name, email and creating a password.
- **Log In**: Existing users can securely log in using their credentials.
- **Logout**: Users can log out to ensure the security of their accounts.

### 2. Task Management

- **Create Tasks**: Users can create new tasks by specifying a title, description, and due date.
- **Edit Tasks**: Existing tasks can be edited to update details such as title, description and due date.
- **Delete Tasks**: Tasks can be deleted when they are no longer needed.
- **View Tasks**: Users can view a list of their tasks, sorted by created date and due date.

## Technologies Used

- Backend: [Golang](https://go.dev/)
- Frontend: [React.js](https://react.dev/)
- Database: [PostgreSQL](https://www.postgresql.org/)

# Getting Started

To get started with this Task Management Web Application, follow these steps:

- Clone the repository to your local machine.
- Make sure you have [Docker](https://www.docker.com/) installed on your machine
- type `make build` to build server, client and database using docker compose
- then run `make up` to run container or `make up-quite` to run container on the background process
