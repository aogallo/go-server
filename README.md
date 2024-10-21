# Rest API with Golang

## Table of Contents

- [About Project](#about-project)
  > - [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Environment variables](#environment-variables)
  - [Run Locally](#run-locally)
  - [Running Tests](#running-tests)

## About Project

### Tech Stack

- [Golang](https://go.dev/): Programing Language
- [Gin](https://gin-gonic.com/docs/): Gin is a HTTP web framework written in Go (Golang)
- [Gorm](https://gorm.io/index.html): The fantastic ORM library for Golang
- [Postgresql](https://www.postgresql.org/): Database

## Getting Started

### Prerequisites

Make sure Go and Postgresql are installed.

- Golang: [Download Golang](https://go.dev/doc/install)
- Postgresql: [Download Postgresql](https://www.postgresql.org/download/)
- Air: [Air Docs](https://github.com/air-verse/air?tab=readme-ov-file#cloud-air---live-reload-for-go-apps)

### Environment Variables

This repository connects to Postgresql database, so create a `.env` file and add the following variables if you have Postgresql locally:

```zsh
 DATABASE_URL="host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
```

### Run Locally

Clone the project

```bash
git clone git@github.com:aogallo/go-server.git
```

Go to the project directory

```bash
cd go-server
```

Start the server in the development environment

```bash
air
```

### Running Tests

To run tests, run the following command

```bash
go test
```
