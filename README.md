# products-manager

Simple project implementing a products discounts system with gRPC.

### calculator-service

Service that implements the gRPC server with the discount logic write in Golang.

> **Dissclaimer**
>
> I'm not a Golang developer, until this project I just knew some syntax and concepts about the language, but the result made me happy.

### products-api

A simple Web API write in Python with [FastAPI](https://fastapi.tiangolo.com/) that list products communicating with the Calculator via gRPC to get the discounts.


## How To Run

With docker and docker-compose installed, just run:

```bash
$ make run
```

Then you have got both applications running. To access the list of products use the endpoint `GET /product`.

> To use discount you need to send a header `X-USER-ID` containing the uuid of the user.
> The user creation is not implemented on the API, so you are going to need create at the database.
>
> Database URL should be `postgresql://postgres:root@localhost:5432/products_manager`


## Running Tests

```bash
$ make test
```
