# api-todo-list

## Getting started

This project requires `Go` and `Docker` to be able to run it. Follow these stesp to get all three containers running.

- Navigate to the root folder in PS of this project
- Run `docker compose up`
- It should start up the `mysql`, `api-todo-list-server` and `api-todo-list-client` containers.
- In the `api-todo-list-client` container you shoul be able to use the `add` command to add new entries or the `list` command to read all the entries in the SQL database.
