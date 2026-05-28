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

### 4. Advanced Practice

The `advanced/` directory is for patterns you will meet in real Go projects.

Run them like this:

```bash
go run ./advanced/19_generics
go run ./advanced/20_context_and_cancellation
go run ./advanced/21_worker_pool
go run ./advanced/22_real_http_server
go run ./advanced/23_json_file_project
go run ./advanced/24_cli_todo_app --list
```

Topics covered:

- generic types and functions
- context timeouts and cancellation
- worker pool concurrency
- real HTTP routes and local servers
- JSON persistence to files
- CLI argument parsing with a small app

### 5. How To Practice

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
  advanced/       # real-world patterns and mini-projects
  pkg/            # reusable packages for training and testing
  modules/        # module and workspace practice
  go.mod          # root module
  go.work         # workspace file including the root and nested modules
```

## Commands You Will Use Often

```bash
go run ./foundations/06_functions
go run ./intermediate/16_json_basics
go run ./advanced/24_cli_todo_app --add "practice interfaces"
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

- interfaces in larger designs
- benchmarks
- dependency injection
- database access
- custom packages with versioning
- middleware and routing
- graceful shutdown with signals
- streaming and pipelines
- mock-based testing for services

## Advanced Lesson Notes

- `advanced/22_real_http_server` uses `httptest` to prove the handler works first, then tries a real server on `http://127.0.0.1:8080`. In restricted sandboxes, the live socket step may be blocked, which is also a useful lesson about environments.
- `advanced/23_json_file_project` writes data to a temporary JSON file so you can learn persistence without polluting the repository.
- `advanced/24_cli_todo_app` stores tasks in memory for each run. Later you can upgrade it to save tasks to disk.

Build slowly, stay curious, and use this repo as your personal Go gym.
