# @eth-optimism/chain-mon

[![codecov](https://codecov.io/gh/ethereum-optimism/pepe/branch/develop/graph/badge.svg?token=0VTG7PG7YR&flag=chain-mon-tests)](https://codecov.io/gh/ethereum-optimism/pepe)

`chain-mon` is a collection of chain monitoring services.

## Installation

Clone, install, and build the Pepe monorepo:

```
git clone https://github.com/ethereum-optimism/pepe.git
yarn install
yarn build
```

## Running a service

Copy `.env.example` into a new file named `.env`, then set the environment variables listed there depending on the service you want to run.
Once your environment variables have been set, run via:

```
yarn start:<service name>
```

For example, to run `drippie-mon`, execute:

```
yarn start:drippie-mon
```
