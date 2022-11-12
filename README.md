#bomble-fight-backend

Based on template: https://github.com/leeprovoost/go-rest-api-template

## How to run

Navigate to `cmd/api-service` and run the following:

```
export ENV=LOCAL
export VERSION=VERSION
export PORT=3001
export FIXTURES=./fixtures.json
go build && ./api-service
```

To run postman collection with Newman, run the following after setting the environment variables:

``` 
./api-service.exe & newman run bomble.postman.json -e environment.postman.json
```

To kill the api running in the background, run the following (Windows):
````
TASKKILL //F //IM api-service.exe
```