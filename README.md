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

### 2. Modules

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

### 3. How To Practice

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
  modules/        # module and workspace practice
  go.mod          # root module
  go.work         # workspace file including the root and nested modules
```

## Commands You Will Use Often

```bash
go run ./foundations/06_functions
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

- pointers
- generics
- file I/O
- HTTP servers
- JSON encoding
- interfaces in larger designs
- table-driven tests
- benchmarks
- context and cancellation

Build slowly, stay curious, and use this repo as your personal Go gym.
