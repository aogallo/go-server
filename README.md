# Rest API with Golang

## Table of Contents

- [About Project](#about-project)
  > - [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  - []()

## About Project

## Tech Stack

- [Golang](https://go.dev/): Programing Language
- [Gorm](https://gorm.io/index.html): The fantastic ORM library for Golang
- [Postgresql](https://www.postgresql.org/): Database

## Getting Started

### Prerequisites

Make sure Go is installed. Click [here](https://go.dev/doc/install) to download.

### Environment Variables

This repository connects to Postgresql database, so create a `.env` file and add the following variables if you have Postgresql locally:

```
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

Start the server

```bash
go run .
```

### Running Tests
