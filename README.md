# PFGolang

PFGolang is an English-first Go practice repository for daily training.

The goal is simple:

- keep every lesson small and runnable
- practice syntax by reading and editing real code
- learn module and workspace management without extra setup friction
- leave room for advanced exercises later

## Learning Path

### 1. Foundations

Each folder in `foundations/` contains one standalone `main.go` file.

Run a lesson like this:

```bash
go run ./foundations/01_hello_world
```

Suggested order:

1. `01_hello_world`
2. `02_variables_and_constants`
3. `03_numbers_and_strings`
4. `04_conditionals`
5. `05_loops`
6. `06_functions`
7. `07_multiple_returns`
8. `08_structs`
9. `09_methods_and_interfaces`
10. `10_slices_and_maps`
11. `11_error_handling`
12. `12_goroutines_and_channels`

### 2. Intermediate Practice

The `intermediate/` directory starts moving from isolated syntax drills to code organization.

Run them like this:

```bash
go run ./intermediate/13_pointers
go run ./intermediate/14_packages_and_files
go run ./intermediate/15_table_driven_tests
go test ./pkg/... ./intermediate/...
go run ./intermediate/16_json_basics
go run ./intermediate/17_file_io
go run ./intermediate/18_http_server
```

Topics covered:

- pointers and mutation
- splitting logic into reusable packages
- table-driven testing
- JSON encoding and decoding
- file creation, writing, and reading
- small HTTP handlers

### 3. Modules

The root repository is a Go module, and `modules/` also contains extra module examples.

Run them like this:

```bash
go run ./modules/01_single_module_app
go run ./modules/02_flags_cli --name Plato
go run ./modules/03_workspace_alpha
go run ./modules/04_workspace_beta
```

These folders are meant to help you practice:

- `go mod init`
- `go mod tidy`
- `go list`
- `go work init`
- `go work use`
- `go env GOPATH`

### 4. How To Practice

Use the repository in loops:

1. Run a lesson.
2. Predict the output before reading it.
3. Change one rule or value.
4. Run it again.
5. Write a small variation beside it later.

Good modifications to try:

- replace loops with different ranges
- add more fields to structs
- return custom errors
- add extra goroutines
- turn prints into helper functions
- create your own lesson directory with the next number

## Repository Layout

```text
PFGolang/
  foundations/    # syntax drills, one runnable file per lesson
  intermediate/   # slightly more realistic runnable exercises
  pkg/            # reusable packages for training and testing
  modules/        # module and workspace practice
  go.mod          # root module
  go.work         # workspace file including the root and nested modules
```

## Commands You Will Use Often

```bash
go run ./foundations/06_functions
go run ./intermediate/16_json_basics
go run ./modules/02_flags_cli --name Learner
go test ./...
go mod tidy
go work sync
```

## Training Rules For Yourself

- keep examples small
- prefer clear names over clever code
- write one idea per lesson
- when confused, print values and types
- revisit old lessons and refactor them

## Next Expansion Ideas

- generics
- interfaces in larger designs
- benchmarks
- context and cancellation
- dependency injection
- database access
- custom packages with versioning
- middleware and routing
- concurrency patterns with worker pools

Build slowly, stay curious, and use this repo as your personal Go gym.
