# 📝 Todo CLI — A Lightweight Task Manager in Go & Cobra

Todo CLI is a fast, minimal, and user-friendly command-line task manager built with **Golang** using the powerful **Cobra CLI framework**. It enables you to manage your daily tasks directly from the terminal using simple and intuitive commands.

With support for adding, listing, updating, and deleting tasks, this CLI application is designed for developers who prefer productivity tools that live in the terminal.

Whether you're organizing daily work or exploring Go-based CLI development, Todo CLI provides a practical and extensible foundation.

---

## 🐍 Why Cobra?

This project uses [Cobra](https://github.com/spf13/cobra), a popular Go library for building modern CLI applications.

Cobra provides structured commands and subcommands, built-in flag handling, and automatic help generation. It enables clean, scalable CLI architecture and is widely used in production tools like Kubernetes and Helm.

---

## ⚙️ Build & Run

### 1️⃣ Build the application

```bash
go build
```

### 2️⃣ Run the CLI

Use the following command structure:

```bash
./todo <subcommand>
```

### Available Subcommands

```bash
./todo add
./todo list
./todo update
./todo delete
```

Use the following command for help:

```bash
./todo --help
```

---

## ➕ add — Add a New Todo Item

Adds a new item to your todo list.  
Priority helps you identify which item to work on first.

### Flags

| Flag         | Shorthand | Description                                                     | Required | Default |
| ------------ | --------- | --------------------------------------------------------------- | -------- | ------- |
| `--name`     | `-n`      | Name of the item you want to add                               | ✅ Yes   | —       |
| `--priority` | `-p`      | Priority level: `Today`, `Tomorrow`, `This week`, `This Month` | ❌ No    | `Today` |
| `--help`     | `-h`      | Help for the add command                                       | ❌ No    | —       |

### Usage

Add a task with default priority (Today):

```bash
./todo add --name "Buy Milk, Eggs and Bread"
```

Add a task with custom priority:

```bash
./todo add --name "Prepare sales presentation" --priority "This week"
```

---

## ✏️ update — Update an Existing Todo Item

Updates the name, priority, or status of an existing task in your todo list.

### Flags

| Flag         | Shorthand | Description                                    | Required |
| ------------ | --------- | ---------------------------------------------- | -------- |
| `--name`     | `-n`      | Name of the item you want to update            | ✅ Yes   |
| `--new_name` | `-e`      | Update the name of your task                   | ❌ No    |
| `--priority` | `-p`      | Update the priority of your task               | ❌ No    |
| `--status`   | `-s`      | Update the status (`Complete` or `Incomplete`) | ❌ No    |
| `--help`     | `-h`      | Help for the update command                    | ❌ No    |

### Usage

Mark a task as complete:

```bash
./todo update -n "Workout" -s "Complete"
```

Change priority:

```bash
./todo update -n "Prepare presentation" -p "This week"
```

Rename a task:

```bash
./todo update -n "Workout" -e "Morning Workout"
```

---

## 📋 list — View Todo Items

Displays tasks in your todo list.

By default, only **Incomplete** tasks are shown.

### Flags

| Flag     | Shorthand | Description                             | Default Behavior |
| -------- | --------- | --------------------------------------- | ---------------- |
| `--all`  | `-a`      | List both Complete and Incomplete tasks | ❌ Not enabled   |
| `--help` | `-h`      | Help for the list command               | —                |

### Usage

List only incomplete tasks (default behavior):

```bash
./todo list
```

List all tasks (Complete and Incomplete):

```bash
./todo list -a
```

---

## 🗑️ delete — Remove a Todo Item

Deletes an item from your todo list.  
The command includes a confirmation prompt to prevent accidental deletions.

### Flags

| Flag     | Shorthand | Description                         | Required |
| -------- | --------- | ----------------------------------- | -------- |
| `--name` | `-n`      | Name of the item you want to delete | ✅ Yes   |
| `--help` | `-h`      | Help for the delete command         | ❌ No    |

### Usage

```bash
./todo delete --name "Buy Milk, Eggs and Bread"
```

You will be prompted for confirmation:

```
Are you sure you want to delete item - Buy Milk, Eggs and Bread?
Please enter Y or N: Y
Item is deleted
```

This confirmation step ensures safer task management and prevents accidental data removal.

---

## 📂 Project Structure

```bash
todo/
│
├── cmd/                     # Cobra command definitions
│   ├── add/
│   │   └── add.go           # add subcommand implementation
│   ├── delete/
│   │   └── delete.go        # delete subcommand implementation
│   ├── list/
│   │   └── list.go          # list subcommand implementation
│   ├── update/
│   │   └── update.go        # update subcommand implementation
│   └── root.go              # root command configuration
│
├── todo_items/
│   └── todo_items.go        # Core todo item logic & data handling
│
├── main.go                  # Application entry point
├── todo_item_list.json      # Persistent storage for todo items
└── README.md                # Project documentation
```

### 🏗️ Architecture Overview

- **cmd/** contains all Cobra commands, organized by subcommand for scalability.
- **todo_items/** contains the core business logic separated from CLI handling.
- **main.go** initializes and executes the root Cobra command.
- **todo_item_list.json** stores task data persistently.
