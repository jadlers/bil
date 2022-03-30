package main

import "errors"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	return errors.New("Add some kind of CLI here")
}
