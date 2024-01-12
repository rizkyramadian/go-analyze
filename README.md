# Go-Analyze

## Description

Personal Tool for code analysis

## Getting Started

```
go install github.com/rizkyramadian/go-analyze
```

### Prerequisites

- Golang version 1.21

## Usage

Right now it only has 1 mode which I call `self-dep` to analyze the file for self dependencies.
This only works on receiver functions.

### Self-dep
```
go-analyze self-dep <file>.go [... Flags]
```
This will print all the functions in the file and it will check if the receiver functions calls other functions in the receiver.

This can be useful if you are breaking down a big struct into smaller and see the dependency mapping.

#### Flags
```
-l --line prints the line number for function call
-d <num> --depth <num> Depth on Dependency Check
-r --dep Prints Dependency Only
-f --fn Prints Function Only
```