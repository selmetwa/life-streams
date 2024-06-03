# Project life-streams

Beep boop boop

## Todo
- Go [x]
- Sqlite [x]
- templ [x]
- htmx [x]
- Sql
- View transitions
- Web components (maybe lit)
- Modern css
- Import maps
- Auth 
- Html drag and drop
- Html popover
- No build js
- Jsdoc 
- Semantic html

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