# Rest API with Golang

## Description

## Table of Content

## Built with

This RestAPI was built-in with:

## Getting Started

### Prerequisites

Make sure Go is installed. Click [here](https://go.dev/doc/install) to download.

Clone the repository.

```bash
git clone git@github.com:aogallo/go-server.git
```

This repository connects to Postgresql database, so create a `.env` file and add the following variables if you have Postgresql locally:

```
 DATABASE_URL="host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
```

## How to user the project

First you can clone the repository [go server]("https://github.com/aogallo/go-server")

You need to create an `.env` file, and add the following values:

```
DATABASE_URL="postgresql://<USER>:<PASSWORD>@<HOST>/<DATABASE_NAME>?sslmode=require"
```
