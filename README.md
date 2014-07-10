# Gomodoro

A Pomodoro timer done in Golang that uses OS X Notification Center to alert the user.

## Usage

Gomodoro responds to simple arguments in the command line, such as: pomodoros
work intervals `[-p]`, short brake intervals `[-s]`, large brake intervals `[-l]`
and a maximum number of runs before the program shuts down, called runs `[-r]`.

To view this options just: `gomodoro --help`

Each option flag has a default value:

- Pomodoro work interval is set to 25 minutes.
- Short break interval is set to 5 minutes.
- Large break interval is set to 30 minutes.
- Pomodoro runs interval is set to 4 runs.

To modify every default value:

```bash
$ gomodoro -p 30 -s 10 -l 25
```

It's pretty simple.
