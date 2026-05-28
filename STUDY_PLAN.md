# Study Plan

Use this file when you want to study fast without deciding what to do each day.

## Quick Start

If you only have one hour:

1. Run one lesson.
2. Edit one idea in the code.
3. Run it again.
4. Write down one thing you learned.
5. Move to the next numbered directory only after you understand the current one.

## 7-Day Sprint

### Day 1

- `foundations/01` to `foundations/06`
- focus on variables, conditionals, loops, and functions

### Day 2

- `foundations/07` to `foundations/12`
- focus on returns, structs, interfaces, slices, maps, errors, and goroutines

### Day 3

- `intermediate/13` to `intermediate/18`
- focus on pointers, packages, tests, JSON, file I/O, and HTTP basics

### Day 4

- `modules/01` to `modules/04`
- review `go mod`, `go work`, and CLI flags

### Day 5

- `advanced/19` to `advanced/30`
- focus on generics, context, worker pools, benchmarks, and service tests

### Day 6

- `advanced/31` to `advanced/42`
- focus on middleware, DI, logging, auth, metrics, retry, rate limits, and API structure

### Day 7

- `advanced/43` to `advanced/54`
- focus on OpenAPI shape, cache, queue flow, feature flags, pagination, tenancy, background jobs, and deployment-oriented service design

## 14-Day Plan

Use the same order, but split each day into a smaller theme:

1. syntax basics
2. function patterns
3. structs and interfaces
4. errors and concurrency
5. pointers and packages
6. tests and JSON
7. file I/O and HTTP
8. modules and workspaces
9. generics and context
10. worker pools and benchmarks
11. service tests and CLI state
12. middleware, logging, auth
13. config, retry, rate limits, queues
14. multi-tenant and deployment-shaped services

## High-Value Practice Rules

- always predict output before running code
- change one thing every time you run a lesson
- if a lesson uses a package, open the package file too
- after every 5 lessons, rerun `make test`
- when a topic feels weak, clone the lesson into a new numbered folder and build your own version

## Best Commands

```bash
make help
make test
make bench
make lesson LESSON=foundations/06_functions
make lesson LESSON=advanced/42_full_todo_api_project
```
