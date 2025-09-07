<div align="center">
<h3>Go-envy: variable loader & validator for Go</h3>
<img src="./golang-banner.jpeg" alt="Golang Banner" />
</div>

[![Go Reference](https://pkg.go.dev/badge/github.com/fdhliakbar/go-envy.svg)](https://pkg.go.dev/github.com/fdhliakbar/go-envy)

Go-envy is a simple yet powerful library for loading and validating environment variables in Go applications. It helps developers manage configuration easily by providing typed getters, default values, and strict validation for required variables. With support for loading from .env files and clear error handling, Go-envy makes your Go projects more robust, secure, and developer-friendly.

## Features

- Load env with default
- Require env (panic if missing)
- Typed getters (`int`, `bool`)
- Load from `.env` file

## Installation

```bash
go get github.com/fdhliakbar/go-envy
```

## Output Program

Main.go

```bash
go run ./examples/main.go

DB: postgres://prod-db:5432/mydb
PORT: 9090
DEBUG: true
SECRET: supersecret
```

Unit Test

```bash
cd go-envy; go test ./pkg/envy -v
=== RUN   TestGet
--- PASS: TestGet (0.00s)
=== RUN   TestRequire
--- PASS: TestRequire (0.00s)
=== RUN   TestGetInt
--- PASS: TestGetInt (0.00s)
=== RUN   TestGetBool
--- PASS: TestGetBool (0.00s)
=== RUN   TestLoad
--- PASS: TestLoad (0.01s)
PASS
ok      github.com/fdhliakbar/go-envy/pkg/envy  0.518s
```

## Next Update

- Add docker file
- Make CI/CD for development
- Update CLI
- Eazy to read and implement to your projects

---
