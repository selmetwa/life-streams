# Project life-streams

Beep boop boop

## Todo
- Go [x]
- Sqlite [x]
- templ [x]
- htmx [x]
- Sql [x]
- View transitions [x]
- Auth [x]
- Html drag and drop [x]
- Html popover [x]
- Semantic html [x]
- dark mode [x]

schema
 - streams
  - created_at
  - updated_at
  - id
  - user_id
  - description
  - title
  - priority
  - stream_id
  - position

 - task
  - id
  - user_id
  - stream_id
  - type
  - priority
  - title
  - description
  - due_date
  - created_at
  - updated_at
  - position

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```