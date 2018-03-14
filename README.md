# autodeploy [![Build Status][1]][2] [![Go Report Card][3]][4] [![License: MIT][5]][6]
[1]: https://img.shields.io/travis/seankhliao/autodeploy.svg?style=flat-square
[2]: https://travis-ci.org/seankhliao/autodeploy
[3]: https://goreportcard.com/badge/github.com/seankhliao/autodeploy?style=flat-square
[4]: https://goreportcard.com/report/github.com/seankhliao/autodeploy
[5]: https://img.shields.io/badge/License-MIT-blue.svg?longCache=true&style=flat-square
[6]: LICENSE

minimal tool to update/build/deploy on github webhooks

it does something similar to
```sh
git pull && \
go build && \
screen -XS SESSION quit ; \
screen -dmS SESSION command and options
```

## Install
```sh
go get github.com/seankhliao/autodeploy
```

## Usage
```sh
autodeploy -dir /path/to/dir ./command with options
```
