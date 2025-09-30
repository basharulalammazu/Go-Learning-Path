<h1 align="center">Go-Learning-Path</h1>

Welcome to my Go (Golang) learning journey. This repo collects small, focused examples that cover core language features. Each file runs independently with its own main function.

## 📌 What’s inside

Hands-on examples for variables, formatting, input, control flow, functions, and more.

## 📂 Lesson Index

Each file is a standalone program. Use the exact filename (many have spaces).

| File                                      | Topic                                   |
| ----------------------------------------- | --------------------------------------- |
| `01. First.go`                            | Print basics (fmt.Print/Println)        |
| `02. Escape Sequences.go`                 | Newline, tab, quotes, and other escapes |
| `03. Static Variable.go`                  | var declarations and basic types        |
| `04. Dynamic Variable.go`                 | Type inference (compiler infers type)   |
| `05. Shortcut Variable Declaration.go`    | Short variable syntax `:=`              |
| `06. Formatting Output.go`                | fmt.Printf and formatting verbs         |
| `07. Constants and Getting User Input.go` | const and simple user input             |
| `08. Number Conversion Calculator.go`     | Numeric conversion + simple calculator  |
| `09. if_else.go`                          | if/else conditionals                    |
| `10. switch_case.go`                      | switch statements                       |
| `11.function.go`                          | Functions: basics and parameters        |
| `12.functionwith_return_value.go`         | Functions with return values            |
| `13.more_function.go`                     | More function examples                  |

> Note: Go doesn’t have “static variables” like some languages; here it means plain `var` declarations.

## 🛠 Prerequisites

- Install Go from https://go.dev/dl/
- Verify in a new PowerShell window:

```powershell
go version
```

If the command isn’t found, ensure Go’s bin folder is on your PATH (e.g., `C:\Program Files\Go\bin`) and open a new PowerShell.

## 🚀 Run a lesson (Windows PowerShell)

1. Open PowerShell and go to the repo folder:

```powershell
cd d:\Github\Go-Learning-Path
```

2. Run a specific file. Quote paths with spaces:

```powershell
go run ".\01. First.go"
go run ".\02. Escape Sequences.go"
go run ".\09. if_else.go"
```

3. Optional: build an exe and run it:

```powershell
go build -o lesson01.exe ".\01. First.go"; .\lesson01.exe
```

> Tip: Avoid `go run .` in this folder because it contains multiple independent main packages.

## 🔧 Troubleshooting

- "'go' is not recognized": Install Go and/or add Go’s bin folder to PATH. Open a new terminal and run `go version`.
- Filename has spaces: Always wrap the filename in quotes, e.g., `go run ".\02. Escape Sequences.go"`.
- Multiple mains error: Run a single file, not the whole directory.

Enjoy exploring Go! If you find something that could be clearer, feel free to open an issue or PR.
