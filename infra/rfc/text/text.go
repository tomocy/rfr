package text

import (
	"fmt"
	"io"

	"github.com/tomocy/rfv/infra/rfc/text/index"
	"github.com/tomocy/rfv/infra/rfc/text/single"
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
	Title               index.Title
	AuthorsAndIssueDate index.AuthorsAndIssueDate
	Metadata            index.Metadata
}

func (l *LineOfIndex) Scan(state fmt.ScanState, _ rune) error {
	_, err := fmt.Fscan(state, &l.ID, &l.Title)
	if err != nil {
		return err
	}
	if l.Title == "Not Issued" {
		return nil
	}

	_, err = fmt.Fscan(state, &l.AuthorsAndIssueDate, &l.Metadata)
	return err
}

type RFC struct {
	Metadata single.Metadata
	Title    string
}

func (r *RFC) Scan(state fmt.ScanState, _ rune) error {
	if _, err := fmt.Fscan(state, &r.Metadata); err != nil {
		return err
	}

	state.SkipSpace()

	read, err := state.Token(true, func(char rune) bool {
		return char != '\n'
	})
	if err != nil {
		return err
	}

	r.Title = string(read)

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
