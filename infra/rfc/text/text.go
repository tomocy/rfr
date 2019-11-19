package text

import (
	"fmt"
	"io"
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
	ID    int
	Title Title
	other other
}

func (l *LineOfIndex) Scan(state fmt.ScanState, _ rune) error {
	_, err := fmt.Fscan(state, &l.ID, &l.Title, &l.other)
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
