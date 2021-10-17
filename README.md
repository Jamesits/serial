# Serial

Serial port connector.

**Work in progress, assume everything broken.**

## CI Status & Badges

| CI | Notes | Status |
| :----- | :----- | :----- |
| GitHub Actions | Very basic checks only | ![example workflow](https://github.com/Jamesits/serial/actions/workflows/build.yml/badge.svg) ![example workflow](https://github.com/Jamesits/serial/actions/workflows/unit-test.yml/badge.svg) |
| Azure DevOps | Actual artifacts generation happens here | [![Build Status](https://dev.azure.com/nekomimiswitch/General/_apis/build/status/serial?branchName=master)](https://dev.azure.com/nekomimiswitch/General/_build/latest?definitionId=92&branchName=master) |
| GitLab CI | For continuously testing GitLab's broken CI design | [![pipeline status](https://gitlab.com/Jamesits/serial/badges/master/pipeline.svg)](https://gitlab.com/Jamesits/serial/-/commits/master) [![coverage report](https://gitlab.com/Jamesits/serial/badges/master/coverage.svg)](https://gitlab.com/Jamesits/serial/-/commits/master) |
| Coveralls | Test coverage is an illusion | [![Coverage Status](https://coveralls.io/repos/github/Jamesits/serial/badge.svg?branch=master)](https://coveralls.io/github/Jamesits/serial?branch=master) |

## Features

- Works entirely under a command line, no GUI required (even on Windows)

## Usage

Please use the in-program help system for a detailed list of the options available.

### Listing Ports

```
serial list --format=table
+---+-----------+----------------------------+------+------+------+-----------+
|   | PATH      | DISPLAY NAME               | USB? | VID  | PID  | SERIAL    |
+---+-----------+----------------------------+------+------+------+-----------+
| 1 | \\.\COM1  | Communications Port (COM1) | No   |      |      |           |
| 2 | \\.\COM2  | Communications Port (COM2) | No   |      |      |           |
| 3 | \\.\COM5  | USB Serial Device (COM5)   | Yes  | 05A6 | 0009 |           |
| 4 | \\.\COM7  | USB-SERIAL CH340 (COM7)    | Yes  | 1A86 | 7523 |           |
| 5 | \\.\COM11 | USB Serial Port (COM11)    | Yes  | 0403 | 6010 | FT4UE4H3B |
+---+-----------+----------------------------+------+------+------+-----------+
```

### Connecting to a Port

For *nix, specify the filename:

```shell
serial open --baudrate=115200 /dev/ttyUSB0
```

For Windows, use either:

```cmd
serial open --baudrate=115200 COM5
serial open --baudrate=115200 \\.\COM5
```

## Installation

### Binary Installation

You can use it as a "green software": just run the executable.

### Installation Packages

(TBD.)

## Building

This project adheres to [Golang standard project layout](https://github.com/golang-standards/project-layout). If you
want to build the program yourself, use the state-of-the-art building process.

For building production binaries, use the script located in `contrib/build`. Refer to the CI configurations for how to
use it.

## Licensing

The program ("serial") is licensed under GPLv3. Dual licensing options and commercial support are available if required. 
