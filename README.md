# Go Boring Avatars

[![Go Report Card](https://goreportcard.com/badge/github.com/hcarriz/go-boring-avatars)](https://goreportcard.com/report/github.com/hcarriz/go-boring-avatars)
[![Go Reference](https://pkg.go.dev/badge/github.com/hcarriz/go-boring-avatars.svg)](https://pkg.go.dev/github.com/hcarriz/go-boring-avatars)

Go Boring Avatars is a tiny, zero-dependency go package that generates custom, SVG-based avatars from any username and color palette.

## Install

> go get github.com/hcarriz/go-boring-avatars

## Usage

```go
package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	goboringavatars "github.com/hcarriz/go-boring-avatars"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Generate a random avatar for every request.
		avatar, _ := goboringavatars.New(time.Now().String())

		// Serve the image.
		w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
		io.WriteString(w, avatar)

	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		slog.Error("closing server", slog.String("error", err.Error()))
		os.Exit(1)
	}

}

```

## Config

See the [docs](https://pkg.go.dev/github.com/hcarriz/go-boring-avatars) for configuration options.
