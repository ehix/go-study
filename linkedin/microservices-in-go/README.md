# Build a microservice with go

[Course](https://www.linkedin.com/learning/build-a-microservice-with-go)

## What you need to know/tooling

- Go
  - focusing on abstractions rather than the language, i.e. it's not optimised or will use proper code style.
  - Go 1.20
  - GOPATH setup

- Database
  - Postgres, from [here](https://www.postgresql.org/download/linux/ubuntu/) or [here](https://ubuntu.com/server/docs/databases-postgresql).
  - Running in Docker locally (script available)

- Tools
  - HTTPie, not a requirement, but apparently cleaner that curl.

## Go for Microservices
- Simplicity
  - Concise, easy to read, easy to maintain
  - Rapid development cycles
- Speed
  - Quick compilation
  - Compiles to a native executable
  - Garbage collection (GC) quick
- Concurrency
  - Built for multicore processors
  - Parallel processing by default

## Prep the environment

1. Check no images are running using `docker ps`
2. Follow the instructions in the `README` supplied.

If encountering the following error: 
```
Error starting userland proxy: listen tcp4 0.0.0.0:5432: bind: address already in use.
```
- Get the PID of the process bound to the port socket, and kill it.
```
$ lsof -i :5432
COMMAND    PID     USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
postgres 42365 postgres    5u  IPv4 238761      0t0  TCP localhost:postgresql (LISTEN)
$ kill 42365
```
3. `exe` into the container: `docker exec -it local-pg /bin/bash`
4. connect to the database: `psql -U postgres`
5. copy the schema and data from the exercise files.
6. confirm successful: `select * from wisdom.customers limit 10;`
```
postgres=# select * from wisdom.customers limit 10;
             customer_id              | first_name | last_name |                  email                  |     phone      |                   address                   
--------------------------------------+------------+-----------+-----------------------------------------+----------------+---------------------------------------------
 8aa4b76a-f66c-4289-b7f6-59f24affe13d | Cally      | Reynolds  | penatibus.et@lectusa.com                | (901) 166-8355 | 556 Lakewood Park, Bismarck, ND 58505
 0d1a8a8c-8e23-49a3-b1f3-baf9cbbb5003 | Sydney     | Bartlett  | nibh@ultricesposuere.edu                | (982) 231-7357 | 4829 Badeau Parkway, Chattanooga, TN 37405
 941d17d0-c3d5-41d1-87b8-9e17c2eb83ea | Hunter     | Newton    | quam.quis.diam@facilisisfacilisis.org   | (831) 996-1240 | 2 Rockefeller Avenue, Waco, TX 76796
 90c49d78-a1e3-4025-b45f-4044c6495109 | Brooke     | Perkins   | sit@vitaealiquetnec.net                 | (340) 732-9367 | 87 Brentwood Park, Dallas, TX 75358
 0b89675a-5cf3-42bf-9934-f810774ad21f | Nolan      | Slater    | sociis.natoque.penatibus@pedeCras.co.uk | (540) 487-5928 | 99 Sage Street, Reno, NV 89505
 4c382c65-150c-4b9e-ba90-58b14adec981 | Germaine   | Green     | ultrices.Vivamus@orciin.edu             | (466) 455-4160 | 6 Jana Park, San Antonio, TX 78240
 aa9a304f-046d-4eed-8140-5287fe59e6ff | Medge      | Ratliff   | nulla.ante@posuerevulputate.org         | (358) 751-8227 | 75 Erie Terrace, Dayton, OH 45454
 837eb73e-07dd-4853-91ea-e44b36e4d11f | Nash       | Vasquez   | ut.nisi@elitAliquam.ca                  | (989) 937-6199 | 39464 Debra Lane, Young America, MN 55557
 f3a66244-2a86-450b-8c52-018145823ece | Michael    | Rutledge  | eget.lacus@sitametorci.org              | (366) 822-4574 | 8231 Crowley Crossing, Cincinnati, OH 45999
 a71bce4a-8ca2-4e54-a1df-68ec524e40e0 | Guy        | Ochoa     | montes.nascetur@semperrutrum.net        | (720) 242-4596 | 92483 Doe Crossing Drive, Lansing, MI 48956
(10 rows)
```

## Set Up Your Project
1. Create a go mod `go mod init github.com/ehix/go-microservices`.
  - Don't run `tidy`
2. Get all dependencies outlined for the course:
  ```shell
  go get github.com/google/uuid
  go get github.com/labstack/echo/v4
  go get github.com/lib/pq
  go get gorm.io/gorm
  go get gorm.io/driver/postgres
  ```
- `Gorm`: Object Relational Mapping (ORM) is a technique used in creating a "bridge" between object-oriented programs and, in most cases, relational databases.

## Set Up the DB Client
- [Use internal packages to reduce your public API surface](https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface)
- Starting with DB layer

## Set Up the Echo Client
- Create webserver itself, wrapped in an interface.
- [What is an Echo Server?](https://medium.com/@himalee.tailor/what-is-an-echoserver-b2bfd3b8deeb)

## Wiring the Service
- Have server wrapped and configure, and client ready.
- Final piece of scaffolding, to allow us to implement the methods we need to for a functioning webservice.
- Test that system is working with.. (downloaded `httpie` to do this `sudo apt install httpie`)
```
$ http :8080/readiness
HTTP/1.1 200 OK
Content-Length: 16
Content-Type: application/json; charset=UTF-8
Date: Thu, 14 Sep 2023 15:27:56 GMT

{
    "status": "OK"
}
```

## `getAll` Operations
- Now scaffolding complete, create a get all function that responds to a get request at the collection route url.
- Test with:
```shell
$ http :8080/customers

HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Thu, 14 Sep 2023 15:54:33 GMT
Transfer-Encoding: chunked

[
    {
        "address": "556 Lakewood Park, Bismarck, ND 58505",
        "customerID": "8aa4b76a-f66c-4289-b7f6-59f24affe13d",
        "emailAddress": "penatibus.et@lectusa.com",
        "firstName": "Cally",
        "lastName": "Reynolds",
        "phoneNumber": "(901) 166-8355"
    },
...
```