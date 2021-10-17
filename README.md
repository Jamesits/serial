# Serial

Serial port connector.

**Work in progress, assume everything broken.**

## CI Status && Badges

| CI | Notes | Status |
| :----- | :----- | :----- |
| GitHub Actions | Very basic checks only | ![example workflow](https://github.com/Jamesits/serial/actions/workflows/build.yml/badge.svg) ![example workflow](https://github.com/Jamesits/serial/actions/workflows/unit-test.yml/badge.svg) |
| Azure DevOps | Actual artifacts generation happens here | [![Build Status](https://dev.azure.com/nekomimiswitch/General/_apis/build/status/serial?branchName=master)](https://dev.azure.com/nekomimiswitch/General/_build/latest?definitionId=92&branchName=master) |
| GitLab CI | For continuously testing GitLab's broken CI design | [![pipeline status](https://gitlab.com/Jamesits/serial/badges/master/pipeline.svg)](https://gitlab.com/Jamesits/serial/-/commits/master) [![coverage report](https://gitlab.com/Jamesits/serial/badges/master/coverage.svg)](https://gitlab.com/Jamesits/serial/-/commits/master) |
| Coveralls | Test coverage is an illusion | [![Coverage Status](https://coveralls.io/repos/github/Jamesits/serial/badge.svg?branch=master)](https://coveralls.io/github/Jamesits/serial?branch=master) |

## Usage

(There should be a detailed documentation about how to use it, but currently please use the in-program help system which
I assume everyone who has the experience of modern command-line applications will know how to use.)

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
