# gomono

## ⚠️ Status: Not Released

This project is a work in progress. APIs and behavior may change without notice.
Do not use in production.

## What is this?

**gomono** is a CLI tool that flattens modular Go code into a single-file `main.go` script.

This is useful for environments like competitive programming or tooling pipelines where you want to bundle all logic into a single executable file.

---

## ✨ Features

* Parses Go modules and resolves local imports
* Merges code from dependencies into a single file
* CLI with flexible flags for specifying project root and main entry point

---

## 🔧 Installation

```sh
go install github.com/rkitamu/gomono/cmd/gomono@latest
```

Or clone and build manually:

```sh
git clone https://github.com/rkitamu/gomono.git
cd gomono
go build -o bin/gomono ./cmd/gomono
```

---

## 🚀 Usage

```sh
gomono --root ./your_project --main ./your_project/cmd/main.go
```

| Flag     | Description                        | Required |
| -------- | ---------------------------------- | -------- |
| `--root` | Path to the project root directory | ✅        |
| `--main` | Path to the entry point `main.go`  | ✅        |

---

## 💠 VSCode Support

You can debug and build `gomono` easily with VSCode:

### `.vscode/launch.json`

Runs the tool with example arguments:

```json
{
  "configurations": [
    {
      "name": "Launch gomono",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/cmd/gomono",
      "args": ["--root", "./example_project", "--main", "./example_project/cmd/main.go"]
    }
  ]
}
```

### `.vscode/tasks.json`

Builds the binary to `bin/gomono`:

```json
{
  "label": "Build gomono",
  "type": "shell",
  "command": "go",
  "args": ["build", "-o", "bin/gomono", "./cmd/gomono"]
}
```

---

## 📁 Project Structure

```
gomono/
├── cmd/
│   └── gomono/        # CLI entry point (main.go)
├── internal/
│   ├── cmd/           # cobra root command
│   ├── parser/        # Go source parser (WIP)
│   ├── merger/        # AST merging logic (WIP)
│   └── util/          # Utilities
├── .vscode/           # Editor configs
├── bin/               # Build output (gitignored)
├── go.mod
└── README.md
```

---

## 🧪 Development

```sh
go run ./cmd/gomono --root ./example_project --main ./example_project/cmd/main.go
```

---

## 📄 License

MIT
