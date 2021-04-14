# Magnet test task

## Run app

```
make run
```

## Stop app

```
make stop
`

### Available env
- MAGNET_TEST_TASK_HTTP_PORT=8001
- MAGNET_TEST_TASK_SQLLITE_NAME="database.sqlite"``
- MAGNET_TEST_TASK_SQLLITE_PATH="/etc/sqlite"
- MAGNET_TEST_TASK_SQLLITE_MIGRATIONPATH="/etc/sqlite/migrations"

### Run swagger for api docs for develop

Install docker https://docs.docker.com/install/

Run docker image:
```
make swagger
```

Docs here: `http://127.0.0.1/swagger/`

