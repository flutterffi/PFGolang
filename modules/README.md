# Modules

This folder is for `go mod` and `go work` practice.

## Goal

Understand how Go organizes code across modules and workspaces.

## How To Use It

```bash
go run ./modules/01_single_module_app
go run ./modules/02_flags_cli --name Learner
go run ./modules/03_workspace_alpha
go run ./modules/04_workspace_beta
```

## Focus Rule

Do not just run the examples. Also inspect:

- `go.mod`
- `go.work`
- module import paths
- workspace membership
