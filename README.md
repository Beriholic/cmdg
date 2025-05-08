# cmdg

A CLI tool to get terminal commands from natural language descriptions.

## Description

cmdg (Command Generator) is a command-line interface tool that converts natural language descriptions into terminal commands, making it easier to find and use terminal commands without memorizing them. Simply describe what you want to do in plain language, and cmdg will suggest the appropriate command.

## Installation

```bash
go install github.com/beriholic/cmdg@latest
```
> Make sure `fastfetch` is installed before installation

## Usage

```bash
cmdg -c "your command description"
```

## Examples

Here are examples of different complexity levels:

### Simple Tasks
Basic file and directory operations:
```bash
cmdg -c "show all files in current directory"

cmdg -c "create a new folder named projects"

cmdg -c "show disk space usage"
```

### Intermediate Tasks
Operations requiring specific parameters or options:
```bash
cmdg -c "find files larger than 100MB in ~/Download folder"

cmdg -c "show top 10 processes using the most memory"

cmdg -c "show memory usage"
```

### Advanced Tasks
Complex operations combining multiple commands:
```bash
cmdg -c "find all error messages in /logs from last 24 hours and save to errors.txt"

cmdg -c "create a compressed backup of ./project folder excluding node_modules"

cmdg -c "find and remove all temporary files older than 30 days"
```