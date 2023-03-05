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
REDIS_HOST=<YOUR_HOST> REDIS_PORT=<YOUR_PORT> make dev
```

### Available Environment Variables
| Variable          | Type     | Default value | Possible Values |
|-------------------|----------|---------------|-----------------|
| PORT              | Int      | 8000          |                 |
|-------------------|----------|---------------|-----------------|
| HEALTH_CHECK_PORT | Int      | 8001          |                 |
|-------------------|----------|---------------|-----------------|
| DEBUG_LEVEL       | String   | INFO          | - PANIC         |
|                   |          |               | - FATAL         |
|                   |          |               | - ERROR         |
|                   |          |               | - WARN          |
|                   |          |               | - NOTICE        |
|                   |          |               | - INFO          |
|                   |          |               | - DEBUG         |
|                   |          |               | - TRACE         |
|-------------------|----------|---------------|-----------------|
| REDIS_HOST        | String   | localhost     |                 |
|-------------------|----------|---------------|-----------------|
| REDIS_PORT        | Int      | 6379          |                 |
|-------------------|----------|---------------|-----------------|
