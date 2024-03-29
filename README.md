# Login with GO ,Angular and PostgreSQL

## Project Overview

The project is a simple login application with GOLANG on the backend, Angular on the frontend, and PostgreSQL as the chosen database.

The goal of this proyect is to provide me a hands-on learning experience on full-stack applications creating a backend from scratch using Go without relying on frameworks, and a frontend with the latest version of Angular.JWT tokens are used for session management.

## Requeriments 

You should have installed in your desktop this : 

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Install 

1. Clone this repository :

    ```bash
    git clone https://github.com/95yoel/angular-golang-login
    ```

2. Run the following commands :

    ```bash

    # Remove existing containers and associated volumes
    docker-compose down -v
    # Start and build the application for the first time
    docker-compose up --build
    ```
3. See the app working :

    To access the application user interface , open your web browser and navigate to :
    ```bash
    http://localhost:4200
    ```
    This will take you to the frontend of the app , where you can register a new user and login.

    To access the backend , navigate to :
    ```bash
    http://localhost:8080
    
    ```
    This takes you to the Golang server , where you can see the list of registered users.

## Stop app

1. To stop the app you can run this command : 

    ```bash
    # This command stops Docker
    docker-compose down

    ```

## Current Status

The proyect has been completed and works well , with only a few improvements remaining.A few things need to be enhanced, such as  changing the way the frontend manages the register.

### Additional Note

The way I have stored the `secret_key` is not secure. I have left it as is because this is not a real application. In a production environment, I would include it in the `.gitignore` to ensure it is not public.

### Areas for Improvement

In future projects or posible updates to this project,I will try to enhance my skills in user authentication. I know that, while the current implementation is secure , it is possible to  improve it more.





