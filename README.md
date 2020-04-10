# Backend services for chili marketplace app

## Services
- Catalog -> microservice for product discovery
- Auth -> microservice for authentication
- Resources -> microservice for storing and retrieving resources

## Folder Structure
```
.
|
├── svc # contains all services
|   ├── Auth
|   |   ...
│   ├── Catalog
│   │   ├── catalogproto   # contain protobuf files for grpc
│   │   │   ├── [binary]
│   │   │   └── [libs]
│   │   ├── cmd    # entrypoint code to build the binary and run service
│   │   ├── config   # contains json config for development env
│   │   ├── internal   # files to be used internally
│   │   │   ├── app  # combines all components into one single app
│   │   │   ├── config  # contains an object that holds configs
│   │   │   ├── constant   # contains all constant variables
│   │   │   ├── converter   # responsible for converting objects
│   │   │   ├── delivery   # layer that interacts directly with clients
│   │   │   ├── entity   # business entities
│   │   │   ├── repo   # layer that interacts with internal storage (db, redis etc)
│   │   │   ├── requests  # contains objects used as a request parameter
│   │   │   ├── usecase # layer that defines business logics
│   │   │   └── config  # contains a struct that holds configs
│   │   ├── docker-compose.yml
│   │   ├── Makefile
│   │   ├── README.md
│   |
.   .
```

