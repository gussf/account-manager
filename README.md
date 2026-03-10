# Account Manager

- [Assumptions/Decisions](#assumptionsdecisions)
- [Architecture](#architecture)
- [Setup & Run](#setup--run)
- [Tests](#tests)


## Assumptions/Decisions
- **Amount**: 
    - While the industry standard is to use the `Amount` in cents (int instead of float) to avoid precision problems, the specification describes the request amount as a float type, so I assumed I was expected to follow it.
    - Assumed that the HTTP request can have negative and positive values, so I decided to not to check if < 0.
- **OperationTypeID**: I am not querying the operation types from the database, in this case I opted to just map the ids in the code for simplicity


## Architecture

I used a simple version of clean architecture, where each layer: `domain, repository, api` is separated from each other through the use of interfaces.

Domain contains the business logic, with core entities and services, and has no knowledge of the other layers that use it.

## Setup & Run

#### Configuration

The available configs are declared in `config.yaml`. They can be overwritten by environment variables. For example, by setting `DATABASE_PORT: 1234` in the docker-compose.yml


#### Running with Docker (compose):

```bash
make docker-up
```

## Tests

**Unit Tests**: The domain core entities and their business logic are being unit tested in the `./domain/core` directory.

**Integration Tests**: The remaining logic, which involves dependencies such as postgresql, are being tested by integration tests in the `./test` directory.

#### Running all tests:
```bash
make test
```

