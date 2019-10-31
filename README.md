# WTFd

[![License](https://img.shields.io/github/license/wtfd-tech/wtfd?style=flat-square)](https://github.com/wtfd-tech/wtfd/blob/master/LICENSE)
[![Latest stable version](https://img.shields.io/github/v/tag/wtfd-tech/wtfd?label=Latest%20Version&style=flat-square)](https://github.com/wtfd-tech/wtfd/releases)  
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fwtfd-tech%2Fwtfd%2Fbadge%3Fref%3Dmaster&style=flat-square)](https://actions-badge.atrox.dev/wtfd-tech/wtfd/goto?ref=master)
[![Codecov](https://img.shields.io/codecov/c/github/wtfd-tech/wtfd?style=flat-square&logo=codecov&label=Coverage)](https://codecov.io/gh/wtfd-tech/wtfd)
[![Dependencies](https://img.shields.io/librariesio/github/wtfd-tech/wtfd?style=flat-square&label=Dependencies)](https://libraries.io/github/wtfd-tech/wtfd)
![Repository Size](https://img.shields.io/github/repo-size/wtfd-tech/wtfd?style=flat-square&label=Repo%20Size)  
[![Last Commit](https://img.shields.io/github/last-commit/wtfd-tech/wtfd?style=flat-square&label=Last%20Commit)](https://github.com/wtfd-tech/wtfd/commits/master)
[![Contributors](https://img.shields.io/github/contributors/wtfd-tech/wtfd?style=flat-square&label=Contributors)](https://github.com/wtfd-tech/wtfd/graphs/contributors)
[![Open Issues](https://img.shields.io/github/issues/wtfd-tech/wtfd?style=flat-square&label=Issues)](https://github.com/wtfd-tech/wtfd/issues)
[![Open PRs](https://img.shields.io/github/issues-pr/wtfd-tech/wtfd?style=flat-square&label=Pull%20Requests)](https://github.com/wtfd-tech/wtfd/pulls)
<!--Micro badger docker image size-->
<!-- Docker hub stars-->

![icon](https://raw.githubusercontent.com/wtfd-tech/wtfd/master/icon.svg)
a [CTFd](https://ctfd.io/)-like Server in go

![demo](https://raw.githubusercontent.com/wtfd-tech/wtfd/master/demo.png)

## Configuration

You need a `config.json` in the same path as your wtfd binary (or the path you're in if you do `go run ./cmd/wtfd.go`)

It shall look like that:

```
{
        "Port": <The Port WTFd should run (ports <1024 need root (don't run WTFd as root))>,
        "challinfodir": "<The (relative or absolute) directory with your challenge infos>",
        "social": "<Html for the down left corner (e.g. \u003ca class=\"link sociallink\" href=\"https://github.com/wtfd-tech/wtfd\"\u003e\u003cspan class=\"mdi mdi-github-circle\"\u003e\u003c/span\u003e GitHub\u003c/a\u003e)>",
        "icon": "<top left icon (e.g. icon.svg)>",
        "firstline": "<first line in the top left>",
	"secondline": "<second line in the top left>",
       	"sshhost": "<The domain of your ssh challenges>"
	"servicedeskaddress": <Mail address for a GitLab service desk instance (or '-' if it is disabled),
	"smtprelaymailwithport": "<Sender address (e.g. wtfdsender@wtfd.tech:25)>",
	"smtprelaymailpassword": "<Password for sender>",
	"ServiceDeskRateLimitInterval": <Interval (in seconds) where access is tracked>,
	"ServiceDeskRateLimitReports": <Maximun ammount of access per user in interval> 
}
```

WTFd will also generate the field `Key` in which the cookie session key will be stored


The Challenge info Dir shall look like that:

```
├── chall-1
│   ├── meta.json
│   ├── README.md
│   └── SOLUTION.md
├── chall-2
│   ├── meta.json
│   ├── README.md
│   └── SOLUTION.md
```

For each Challenge you need a `meta.json`, a `README.md` and a `SOLUTION.md`

The `meta.json` shall look like that:

```
{
	"points": <How many points the challenge should have>,
        "uri": "<Protocol and user of your ssh Challenges (e.g. `ssh://chall-1@%s`>",
	"deps": [<Dependencies the Challenge has>],
	"flag": "<The flag>",
	"author": "<The author of the challenge>"
}
```

The `README.md` and `SOLUTION.md` are markdown files ([syntax](https://github.com/gomarkdown/markdown#extensions)).
The `SOLUTION.md` contents can only be seen by users who already solved the challenge

## Running WTFd

Now you can finally start wtfd by downloading it from the [releases](https://github.com/wtfd-tech/wtfd/releases), giving it permissions `chmod +x wtfd` and running it `./wtfd`

WTFd is HTTP only, if you need HTTPS use a reverse proxy like [Traefik](https://traefik.io/) or [nginx](https://nginx.com/)
