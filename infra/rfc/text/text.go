package text

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

type Index struct {
	Lines []*LineOfIndex
}

func (i *Index) Scan(state fmt.ScanState, _ rune) error {
	for {
		if _, _, err := state.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		state.UnreadRune()

		line := new(LineOfIndex)
		if _, err := fmt.Fscanln(state, line); err != nil {
			return err
		}
		i.Lines = append(i.Lines, line)
	}

	return nil
}

type LineOfIndex struct {
	ID                  int
	Title               Title
	AuthorsAndIssueDate AuthorsAndIssueDate
	other               other
}

func (l *LineOfIndex) Scan(state fmt.ScanState, _ rune) error {
	_, err := fmt.Fscan(state, &l.ID, &l.Title)
	if err != nil {
		return err
	}
	if l.Title == "Not Issued" {
		return nil
	}

	_, err = fmt.Fscan(state, &l.AuthorsAndIssueDate, &l.other)
	return err
}

type Title string

func (t *Title) Scan(state fmt.ScanState, _ rune) error {
	read, err := state.Token(true, func(char rune) bool {
		return char != '.'
	})
	if err != nil {
		return err
	}

	state.ReadRune()

	*t = Title(read)

	return nil
}

type AuthorsAndIssueDate struct {
	Authors  []string
	IssuedAt time.Time
}

func (a *AuthorsAndIssueDate) Scan(state fmt.ScanState, _ rune) error {
	read, err := state.Token(true, func(char rune) bool {
		return char != '('
	})
	if err != nil {
		return err
	}

	read = bytes.TrimRight(read, ". ")
	idx := bytes.LastIndex(read, []byte{'.'})
	rawAuthors, rawDate := string(read[:idx]), strings.TrimLeft(string(read[idx+1:]), " ")

	a.Authors = func(ss []string, fn func(string) string) []string {
		for i, s := range ss {
			ss[i] = fn(s)
		}

		return ss
	}(strings.Split(rawAuthors, ","), func(s string) string {
		return strings.Trim(s, " ")
	})

	a.IssuedAt, err = time.Parse("January 2006", rawDate)
	if err != nil {
		return fmt.Errorf("failed to parse issue date: %s", err)
	}

	return nil
}

type other string

func (o *other) Scan(state fmt.ScanState, _ rune) error {
	read, err := state.Token(false, func(rune) bool {
		return true
	})
	if err != nil && err != io.EOF {
		return err
	}

	*o = other(read)

	return nil
}
