# User Estimation service

## Feature

| Title | Value |
|:-----:|:-----:|
|language| golang (v1.17) |
|architecture| full clean arch implemented |
|transport| gRPC |
|db| mongodb |
|container| docker |
|orchestrator| docker-compose |
|documentation| openAPI |
|CICD| gitlab friendly |

## Development
### Runing dependencies easly

```bash
docker-compose -p go-challeng-infra up -d
```

### Runing application
```bash
go run main.go \
  --port 3000
```

## Use in prodution

```bash
docker build -t estimate-service:1.0.0 .
docker run \
  --name estimate-service
  -p 3000:3000 \
  -v ${PWD}/config.yaml:/app/config/config.yaml \
  estimate-service:1.0.0
```

## Architecture

![](https://koenig-media.raywenderlich.com/uploads/2019/06/Clean-Architecture-graph.png)
## Documentation

![](./docs/app-swagger.png)