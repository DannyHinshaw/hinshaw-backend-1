![Go](https://github.com/DannyHinshaw/hinshaw-backend-1/workflows/Go/badge.svg)

# hinshaw-backend-1

A simple web app to demonstrate a basic registration/authentication flow and backend implementation.


## Architecture Overview

![Architecture Image](https://raw.githubusercontent.com/DannyHinshaw/hinshaw-backend-1/main/.github/images/architecture-overview.png)


## Tech-Stack:

**REST API** - Written in Golang with [echo](https://echo.labstack.com/); fast, simple, lightweight and easily extended web framework.

**Database** - [Postgres](https://www.postgresql.org/), my goto open-source database.

**Cache** - [Redis](https://redis.io/), my goto open-source cache storage.

**NGINX/HTML** - Web app built with basic html, css and javascript (bare-bones), served by a simple NGINX web server.


## Features

- Authentication and authorization with [JWT](https://jwt.io/) using a HMAC-512 key.
- Auth features; registration, login, and logout.
- Users can register with email and password (no duplicate emails, basic password requirements and passwords hashed/salted).
- [CORS](https://echo.labstack.com/middleware/cors) configured as well as [additional middleware](https://echo.labstack.com/middleware/secure) 
for basic protection against XSS, content type sniffing etc.

## 3rd Party Modules

- [echo](https://github.com/labstack/echo) - Web/REST-API framework.
- [pgx](https://github.com/jackc/pgx) - For PostgreSQL driver and interactions.
- [uuid](https://github.com/satori/go.uuid) - For generating v4 UUID's for users etc.
- [jwt-go](https://github.com/alicebob/miniredis) - For unit testing redis.
- [crypto](https://golang.org/x/crypto) - Hashing/salting passwords.
- [testify](https://github.com/stretchr/testify) - For unit test suites.


## Install & Run

**Requirements**

- Docker
- Docker Compose
- Make (optional)

1. Clone repo: `git clone https://github.com/DannyHinshaw/hinshaw-backend-1.git`

2. To start run:
```shell script
make start
```

or

```shell script
docker-compose up --build -d
```

3. Visit `http://localhost:8000` to use the web app.

4. To stop run: 
```shell script
make stop
```

or

```shell script
docker-compose down -v --remove-orphans
```

## Checklist

- [x] Create a web application in a private repository called <lastname>-backend-1.
- [x] Build the necessary REST endpoints that allow a user to sign up with an email address and password, login with an existing email address and password, and log out.
- [x] Design a database schema for your preferred database engine to store and query the credentials. Make sure to include the SQL table creation scripts in your repository.
- [x] Create a basic index page at the root / of the project that allows a user to login.
Also, create a /members area that informs the user they are logged in and allows them to log out. Redirect the user to this page upon successful login.
- [x] Inform the user when their login credentials are not valid.
- [x] Create a Dockerfile and docker-compose.yml that allows us to build and start your project via docker-compose up -d.

## Room for Improvement

Some things left unimplemented simply for the sake of time and not technically being requirements (non-exhaustive list).

- Implement `refresh_token` http-only cookie for JWT auth with API endpoint and web app handling.
- Implement shorter lived JWT `access_token` (when `refresh_token` implemented).
- Keep `access_token` in memory instead of `sessionStorage` for better security against XSS (when `refresh_token` implemented).
- Implement JWT validation client-side.
- Better password requirements, currently the only requirement is that it is greater than 5 characters.
- Implement email validation for user registration.
- Implement 2FA.
- Implement integration tests and CI/CD pipeline.
