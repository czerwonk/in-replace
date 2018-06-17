package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

type replacer func(string) (res string, replaced bool)

func processFile(f *File) error {
	of, err := os.OpenFile(f.Path, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer of.Close()

	content, modified, err := replaceInFileContent(f, of)
	if err != nil {
		return err
	}

	if !modified {
		return nil
	}

	return writeToFile(content, of)
}

func writeToFile(content string, of *os.File) error {
	_, err := of.Seek(0, 0)
	if err != nil {
		return err
	}

	err = of.Truncate(0)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(of)
	_, err = w.WriteString(content)
	if err != nil {
		return err
	}

	err = w.Flush()
	return err
}

func replaceInFileContent(f *File, of *os.File) (string, bool, error) {
	repl, err := replacersForFile(f)
	if err != nil {
		return "", false, err
	}

	return replaceForReader(of, repl)
}

func replaceForReader(reader io.Reader, repl []replacer) (string, bool, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", false, err
	}

	str := string(b)
	modified := false
	for _, r := range repl {
		var m bool
		str, m = r(str)
		if m && !modified {
			modified = true
		}
	}

	return str, modified, nil
}

func replacersForFile(f *File) ([]replacer, error) {
	replacers := make([]replacer, len(f.Replacements))
	for i, r := range f.Replacements {
		re, err := replacerForReplacement(r)
		if err != nil {
			return nil, err
		}

		replacers[i] = re
	}

	return replacers, nil
}

func replacerForReplacement(r *Replacement) (replacer, error) {
	regex, err := regexp.Compile(r.Regex)
	if err != nil {
		return nil, err
	}

	return func(content string) (string, bool) {
		return replace(content, regex, r)
	}, nil
}

func replace(content string, regex *regexp.Regexp, r *Replacement) (res string, replaced bool) {
	str := regex.ReplaceAllStringFunc(content, func(m string) string {
		if len(m) == 0 {
			return m
		}

		if r.Group == 0 {
			replaced = true
			return r.Replacement
		}

		idx := 2 * r.Group
		gp := regex.FindStringSubmatchIndex(m)
		if len(gp) <= int(r.Group)+2 {
			return m
		}

		replaced = true
		return m[0:gp[idx]] + r.Replacement + m[gp[idx+1]:]
	})

	if replaced {
		content = str
	}

	return content, replaced
}
