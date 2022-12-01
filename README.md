# 🧙 Nemeton Leaderboard 🏆

[![nemeton github bannner](./etc/nemeton-banner.jpg)](https://nemeton.okp4.network)

[![version](https://img.shields.io/github/v/release/okp4/nemeton-leaderboard?style=for-the-badge&logo=github)](https://github.com/okp4/nemeton-leaderboard/releases)
[![lint](https://img.shields.io/github/workflow/status/okp4/nemeton-leaderboard/Lint?label=lint&style=for-the-badge&logo=github)](https://github.com/okp4/nemeton-leaderboard/actions/workflows/lint.yml)
[![build](https://img.shields.io/github/workflow/status/okp4/nemeton-leaderboard/Build?label=build&style=for-the-badge&logo=github)](https://github.com/okp4/nemeton-leaderboard/actions/workflows/build.yml)
[![test](https://img.shields.io/github/workflow/status/okp4/nemeton-leaderboard/Test?label=test&style=for-the-badge&logo=github)](https://github.com/okp4/nemeton-leaderboard/actions/workflows/test.yml)
[![codecov](https://img.shields.io/codecov/c/github/okp4/nemeton-leaderboard?style=for-the-badge&token=6NL9ICGZQS&logo=codecov)](https://codecov.io/gh/okp4/nemeton-leaderboard)
[![conventional commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge&logo=conventionalcommits)](https://conventionalcommits.org)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg?style=for-the-badge)](https://github.com/okp4/.github/blob/main/CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg?style=for-the-badge)](https://opensource.org/licenses/BSD-3-Clause)

> 🧙 Leaderboard service repository for the [Nemeton program][https://nemeton.okp4.network] - the [OKP4](https://okp4.network/) incentivized testnet program.

## Purpose

Here you'll find the source code of the leaderboard service for the [OKP4 Nemeton program][https://nemeton.okp4.network] - the [OKP4] incentivized testnet program that has started on November 2, 2022.

The service is in charge to count the points of each druids and expose them to the [web interface](https://github.com/okp4/nemeton-web).

## Prerequisites

- Be sure you have [Golang](https://go.dev/doc/install) installed.
- [Docker](https://docs.docker.com/engine/install/) as well if you want to use the Makefile.

## Build

```sh
make build
```
