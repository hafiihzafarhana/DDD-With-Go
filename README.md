# DDD Go Language Project

## structure folder:

- cmd
  - main.go
- internal
  - app
    - application.go
  - configs
    - pg_config.go
  - constants
  - domain
    - sub_domain1
      - sub_domain1_entity.go
      - sub_domain1_repository.go
      - sub_domain1_service.go
    - sub_domain2
      - sub_domain2_entity.go
      - sub_domain2_repository.go
      - sub_domain2_service.go
    - sub_domain3
      - sub_domain3_entity.go
      - sub_domain3_repository.go
      - sub_domain3_service.go
  - interfaces
    - api
      - handlers
      - middlewares
      - routes
        - routes
        - route.go
    - dto
- pkg
  - errors
  - utils
- go.mod
- go.sum
- .env

## Prettier For Go Language

`gofmt -w ./`
