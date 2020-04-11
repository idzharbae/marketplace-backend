# Marketplace Auth Service

## binaries
- GRPC -> endpoints for manipulating file and file ownership data
- REST -> file server to access uploaded files

## Tech Stacks 
- GRPC
- PostgreSQL

## How to run
### initialize (first time only)
`make init`
`make migration-init`
### Install dependencies
`make dep`
### Run docker services
`make docker-run`
### Migrate tables
`make migration-up`
### Run service
`make run-grpc` (grpc server)

`make run-rest` (file server)
