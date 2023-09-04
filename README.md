# Task Management Web Application
![home page screenshots](/screenshots/home.png)

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
- open browser on [localhost:3000](http://localhost:3000)
> - server will run on port 8000
> - client/reactjs will run on port 3000
> - postgresdb will run on port 3308 or you can edit it on server [.env](/server/.env) file
# Deep Dive
Let's deep dive into the techinal section of this application
## Application Architecture
it have 3 separate docker container that connect on the same network
![application architecture](/screenshots/architecture.png)
## Database Design
we have **users** and **tasks** entity that have relationship one to many, 1 user can have many tasks
![application db design](/screenshots/entities.png)
## Server Layer
on the server we separate layers into:
- **routes**: we separate routes into private and public routes, user that have logged can acces private routes
- **middlewares**:
  - **CORS**
    used for enable client access into the server
  - **JWT**
    used to authenticate if jwt on the **Authorization: Bearer {token}** is valid
- **controllers**: containing fiber function handler for http request
- **services**: containing business logic
- **repository**: call gorm functions in here
![application layer](/screenshots/server_layer.png)
## Test
### Unit Test
we do unit test on repositories and services layer, with result below:
![unit test](/screenshots/test.png)
### benchmark
for handling such as huge request we do bench test on Create Task function on task controller
![bench test](/screenshots/bench_controller.png)

# Thank You
