# Auth-Rest
Simple Authentication API server that uses SMS-based OTP for authentication.

## Features
- User Authentication with SMS-Based OTP
- Basic User Management API
- 3 User Roles ("super", "admin" and "user")
- Support for databases supported by GORM

## Configuration
Example configuration is available at [config.yaml](config.yaml.example).

## Run
Recommended database for initial testing and development purposes is SQLite 3, as it is lightweight, has no other dependency and is fully featured.

### First setup
```shell
# an empty file is used to make the mount working 
cp db.sqlite3.empty db.sqlite3

# example config will suffice
cp config.yaml.example config.yaml

# build docker image
docker build -t auth-rest .

# initial setup, migrating the database and creating first superuser
docker run -v ./config.yaml:/app/config.yaml -v ./db.sqlite3:/app/db.sqlite3 auth-rest setup -p "09123456789" -c "/app/config.yaml"
```

### Running the server
#### Using Docker command
```shell
docker run -v ./config.yaml:/app/config.yaml -v ./db.sqlite3:/app/db.sqlite3 -p 4000:4000 auth-rest run -c "/app/config.yaml"
```
#### Using Docker Compose
```shell
docker compose up
```

## API Documentation
After running the server, Swagger will be accessible at "http://localhost:4000/swagger/". Also, you can find the
OpenAPI spec in [here](internal/docs/swagger.yaml).

## SMS Providers
Currently, there is one provider implemented, and it prints SMS content to logs (for testing and development purposes).

## License
MIT License
