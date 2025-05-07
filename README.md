# ⏱️ Timekeeper

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/dl/)

> A simple and fast command-line tool to track time spent on tasks for different customers — written in Go.

Timekeeper is a small evening project built to get hands-on experience with Go. While the code may not be perfect, it’s functional and improving with every commit. 😄

---

## 📚 Index

- [⏱️ Timekeeper](#️-timekeeper)
  - [📚 Index](#-index)
  - [🚀 Features](#-features)
  - [📦 Installation](#-installation)
    - [🔧 Prerequisite](#-prerequisite)
    - [🛠️ Clone \& Build](#️-clone--build)
  - [📦 Build for specific Platforms](#-build-for-specific-platforms)
  - [⚙️ Command-Line Overview](#️-command-line-overview)
    - [✅ `add` — Add a Time Entry](#-add--add-a-time-entry)
    - [📄 `list` — View Logged Entries](#-list--view-logged-entries)
      - [`--by-date` (summary mode)](#--by-date-summary-mode)
    - [🔢 `calculate` — Calculate Duration Between Two Times](#-calculate--calculate-duration-between-two-times)
      - [🧪 Example](#-example)
    - [🛠️ `config` — Manage Configuration](#️-config--manage-configuration)
      - [✅ Add a Customer](#-add-a-customer)
      - [🚚 Move File Based on Date Range](#-move-file-based-on-date-range)
    - [💡 `completion` — Shell Autocompletion](#-completion--shell-autocompletion)
  - [🧩 Shell Autocompletion Setup](#-shell-autocompletion-setup)
    - [🐚 Bash](#-bash)
    - [🐚 Zsh](#-zsh)
    - [🐟 Fish](#-fish)
    - [💻 PowerShell](#-powershell)
  - [✨ Quick Examples](#-quick-examples)

---

## 🚀 Features

- Track time by **customer**, **task**, and **date**
- Supports durations in `h`, `m`, and decimals (`1.5h`, `30m`, `1,5h`)
- CSV storage (`entries.csv`) for time logs
- List, filter, and group entries with `list` command
- Built-in customer validation using `customers.csv`
- Calculate time between two `HH:MM` timestamps
- Auto-rounding to nearest 15-minute increment
- Bash, Zsh, Fish, and PowerShell shell autocompletion support

---

## 📦 Installation

### 🔧 Prerequisite

Install Cobra CLI (only needed for development):

```bash
go install github.com/spf13/cobra/cobra@latest
```

### 🛠️ Clone & Build

```bash
git clone https://github.com/leonlipinski/timekeeper.git
cd timekeeper
go build -o /usr/local/bin/timekeeper
```

---

## 📦 Build for specific Platforms

```bash
GOOS=darwin  GOARCH=arm64   go build -o builds/timekeeper-darwin-arm64
GOOS=linux   GOARCH=amd64   go build -o builds/timekeeper-linux-amd64
GOOS=windows GOARCH=amd64   go build -o builds/timekeeper-windows-amd64.exe
```

---

## ⚙️ Command-Line Overview

```bash
timekeeper [command] [flags]
```

---

### ✅ `add` — Add a Time Entry

Add a new tracked time entry to your `entries.csv`.

```bash
timekeeper add -c <customer> -t <task> -m <duration> [-d <date>]
```

| Flag         | Shorthand | Description                                             | Example          |
|--------------|-----------|---------------------------------------------------------|------------------|
| `--customer` | `-c`      | Name of the customer (must exist in `customers.csv`)    | `"Globex"`       |
| `--task`     | `-t`      | Task description                                        | `"Write tests"`  |
| `--time`     | `-m`      | Time spent: `30m`, `1.5h`, `1,5h`                        | `"1.5h"`         |
| `--date`     | `-d`      | Entry date: `YYYY-MM-DD` or `"yesterday"`               | `"2025-04-11"`   |

> ⏱️ Time is rounded to the nearest 15-minute increment unless `0`.

---

### 📄 `list` — View Logged Entries

List and group time entries from `entries.csv`.

```bash
timekeeper list [-d <date>] [--by-date]
```

| Flag         | Shorthand | Description                                   | Example              |
|--------------|-----------|-----------------------------------------------|----------------------|
| `--date`     | `-d`      | Filter by a specific date (e.g., `"2025-04-11"`) | `-d 2025-04-11`      |
| `--by-date`  | `-b`      | Show only total minutes per date              | `--by-date`          |

#### `--by-date` (summary mode)

Show total minutes worked per day, regardless of task or customer:

```bash
timekeeper list --by-date
```

**Example Output:**
```
Combined time entries by date:
Date: 2025-04-10, Total minutes worked: 02:15
Date: 2025-04-11, Total minutes worked: 05:30
```

> Combine with `--date` to limit to one day:

```bash
timekeeper list -d 2025-04-11 -b
```

---

### 🔢 `calculate` — Calculate Duration Between Two Times

Calculate minutes between a start and end time (e.g., work session):

```bash
timekeeper calculate -s <start-time> -e <end-time>
```

| Flag       | Shorthand | Description                   | Example   |
|------------|-----------|-------------------------------|-----------|
| `--start`  | `-s`      | Start time in `HH:MM` format  | `08:59`   |
| `--end`    | `-e`      | End time in `HH:MM` format    | `17:01`   |

#### 🧪 Example

```bash
timekeeper calculate -s 08:59 -e 17:01
```

**Output:**

```
Without rounding: 482
Total minutes spent: 480
Total time spent: 08:00
```

---

### 🛠️ `config` — Manage Configuration

Add customers or archive old logs.

```bash
timekeeper config [--customer <name>] [--move]
```

| Flag         | Shorthand | Description                                                                    | Example                       |
|--------------|-----------|--------------------------------------------------------------------------------|-------------------------------|
| `--customer` | `-c`      | Add a new customer to `customers.csv`                                          | `-c "NewClient"`              |
| `--move`     | `-m`      | Rename `entries.csv` to `entries-<min>-to-<max>.csv` based on date range       | `--move`                      |

#### ✅ Add a Customer

```bash
timekeeper config -c "NewClient"
```

This will append `"NewClient"` to your `~/.config/timekeeper/customers.csv`.

#### 🚚 Move File Based on Date Range

```bash
timekeeper config --move
```

If `entries.csv` exists, it will be renamed to something like:

```
entries-2025-03-01-to-2025-04-11.csv
```

---

### 💡 `completion` — Shell Autocompletion

Generate shell autocompletion scripts for enhanced CLI usage:

```bash
timekeeper completion [bash|zsh|fish|powershell]
```

---

## 🧩 Shell Autocompletion Setup

### 🐚 Bash

```bash
# Temporary:
source <(timekeeper completion bash)

# macOS (Homebrew):
timekeeper completion bash > /usr/local/etc/bash_completion.d/timekeeper

# Linux:
timekeeper completion bash > /etc/bash_completion.d/timekeeper
```

---

### 🐚 Zsh

```bash
autoload -U compinit && compinit
timekeeper completion zsh > "${fpath[1]}/_timekeeper"
```

---

### 🐟 Fish

```bash
timekeeper completion fish > ~/.config/fish/completions/timekeeper.fish
```

---

### 💻 PowerShell

```powershell
timekeeper completion powershell | Out-String | Invoke-Expression
```

Add the above to your `$PROFILE` to enable on every session.

---

## ✨ Quick Examples

```bash
# Add entries
timekeeper add -c "AcmeCorp" -t "Debugging" -m "1.5h" -d yesterday
timekeeper add -c "Globex" -t "Project kickoff" -m "45m"

# View entries
timekeeper list
timekeeper list --by-date
timekeeper list -d 2025-04-11 -b

# Time calculation
timekeeper calculate -s 08:59 -e 17:01

# Add customer
timekeeper config -c "Internal"

# Archive time entries
timekeeper config --move
```
