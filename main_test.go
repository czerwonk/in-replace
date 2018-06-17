package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	defer os.Remove("tests/test.html")
	prepare(t)

	process("tests/config.yml")

	actual, err := ioutil.ReadFile("tests/test.html")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile("tests/test.html.expect")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(expected), string(actual))
}

func prepare(t *testing.T) {
	b, err := ioutil.ReadFile("tests/test.html.base")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("tests/test.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		t.Fatal(err)
	}
}
