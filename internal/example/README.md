## Generation Swagger 
```shell 
swag  init -d .\cmd\http -o .\docs\swagger
```


## Build Docker Container

```shell
docker build -t http -f deployments/Dockerfile .
```