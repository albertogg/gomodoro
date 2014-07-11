# Gomodoro

A Pomodoro timer done in Golang that uses OS X Notification Center to alert the user.

## Installation

```bash
$ go get -u github.com/albertogg/gomodoro
```

## Usage

Gomodoro responds to simple arguments in the command line, such as: pomodoros
work intervals `[-p]`, short brake intervals `[-s]`, large brake intervals `[-l]`
and a maximum number of runs before the program shuts down, called runs `[-r]`.

To view this options just: `gomodoro --help`

Each option flag has a default value:

- `-p` Pomodoro work interval is set to 25 minutes.
- `-s` Short break interval is set to 5 minutes.
- `-l` Large break interval is set to 30 minutes.
- `-r` Pomodoro runs interval is set to 4 runs.

To modify every default value:

```bash
$ gomodoro -p 30 -s 10 -l 25
```

It's pretty simple.

## License

[MIT](https://github.com/albertogg/gomodoro/blob/master/LICENSE) License
