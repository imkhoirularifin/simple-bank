## Simple Bank

Simple Bank API

### Installation steps:

1. Clone project from https://github.com/imkhoirularifin/simple-bank.git

```shell
git clone https://github.com/imkhoirularifin/simple-bank.git
```

2. Change directory

```shell
cd simple-bank
```

3. Install Dependencies

```shell
go get ./...
```

4. Create a copy of your .env file

```shell
cp env.example .env
```

5. Create Database in postgresql

```example
example name = simple-bank
```

6. Configure you own .env

```
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_PASSWORD=
DB_NAME=simple-bank
API_PORT=

LOG_FILE=logger.log

SESSION_KEY="wonderfulIndonesia"
SESSION_MAX_AGE=604800 # expire in 7 day, data stored in seconds
```

7. Configure DB_URL in Makefile

```example
DB_URL=postgresql://{user}:{password}@{host}:{port}/simple-bank?sslmode=disable
```

8. Migrate up

```shell
make migrateup
```

9. Run the Project

```shell
go run .
```

10. Import postman collection from this project file, all json request are in there

Note: This project is far from perfect, there is still a lot of inefficient code and lack of security on the backend, I don't use JWT for authentication, I only store user information in the session which is in the client cookie, I hope to develop this project deeper to implement the best practices in all aspects.
