# Fizzbuzz Go server
## Description
This is a simple Fizzbuzz server written in Golang.
We define a Fizzbuzz server a server that:
- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where:
  - all multiples of int1 are replaced by str1
  - all multiples of int2 are replaced by str2
  - all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

Bonus: add a statistics endpoint allowing users to know what the most frequent request has been.
This endpoint should:
- Accept no parameter
- Returns:
  - The parameters corresponding to the most used request
  - The number of hits for this request

## Installation
### With Docker
To run the project, a redis server must be up and running.
I have written a simple docker-compose.yml that runs the Go server with its Redis.
Enter the following command to run it:
```bash
make docker
```

### Local
Assuming you have an already running Redis database, you can run:
```bash
make dev
```

## Available Environment Variables
| Variable          | Type     | Default value | Possible Values |
|-------------------|----------|---------------|-----------------|
| PORT              | Int      | 8000          |                 |
| HEALTH_CHECK_PORT | Int      | 8001          |                 |
| DEBUG_LEVEL       | String   | INFO          | PANIC, FATAL, ERROR, WARN, NOTICE, INFO, DEBUG, TRACE |
| REDIS_HOST        | String   | localhost     |                 |
| REDIS_PORT        | Int      | 6379          |                 |

##  Endpoints
### Fizzbuzz
To get the generated Fizzbuzz, you can call the `/fizzbuzz` POST endpoint.
This endpoint returns the generated string
```bash
curl http://localhost:8080/fizzbuzz -X POST -d '{"int1":3,"int2":5,"limit":100,"str1":"fizz","str2":"buzz"}'
```

Accepted parameters:
```js
{
  int1:  number,
  int2:  number,
  limit: number,
  str1:  string,
  str2:  string,
}
```

Calling this endpoint with any other method than *POST* will return a `501 NOT IMPLEMENTED` response.

### Statistics
The `/statistics` GET endpoint returns the parameters corresponding to the most used request and the number of hits.
```bash
curl http://localhost:8080/statistics
```
The result is formatted in JSON as:
```json
{
  "params": {
    "int1": 3,
    "int2": 5,
    "limit": 1000,
    "str1": "fizz",
    "str2": "buzz",
  },
  "score": 10
}
```
Calling this endpoint with any other method than *GET* will return a `501 NOT_IMPLEMENTED` response.

### Health Check
There is an health check server running on a different port (default to `8081`) on the root (`/`) endpoint.
This endpoint returns `ok`

```bash
curl http://localhost:8081/ # ok
```

Any other path on that server will result with a `404 NOT_FOUND`.

## Running tests
```bash
make test
```
