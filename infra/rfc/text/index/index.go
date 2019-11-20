package index

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

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

type Metadata struct {
	Status         string
	ObsoletingIDs  []int
	ObsoletedByIDs []int
	UpdatingIDs    []int
	UpdatedByIDs   []int
}

func (m *Metadata) Scan(state fmt.ScanState, _ rune) error {
	for {
		state.SkipSpace()
		if _, _, err := state.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		state.UnreadRune()

		datum := new(metaDatum)
		if _, err := fmt.Fscan(state, datum); err != nil {
			return err
		}

		store, ok := m.stores()[datum.key]
		if !ok {
			continue
		}
		if err := store(datum.val); err != nil {
			return err
		}
	}

	return nil
}

func (m *Metadata) stores() map[string]func(string) error {
	return map[string]func(string) error{
		"Status": func(val string) error {
			m.Status = val
			return nil
		},
		"Obsoletes": func(val string) error {
			ids, err := m.splitIDs(val)
			if err != nil {
				return fmt.Errorf("failed to store obsoleting ids: %s", err)
			}
			m.ObsoletingIDs = ids

			return nil
		},
		"Obsoleted by": func(val string) error {
			ids, err := m.splitIDs(val)
			if err != nil {
				return fmt.Errorf("failed to store obsoleted by ids: %s", err)
			}
			m.ObsoletedByIDs = ids

			return nil
		},
		"Updates": func(val string) error {
			ids, err := m.splitIDs(val)
			if err != nil {
				return fmt.Errorf("failed to store updating ids: %s", err)
			}
			m.UpdatingIDs = ids

			return nil
		},
		"Updated by": func(val string) error {
			ids, err := m.splitIDs(val)
			if err != nil {
				return fmt.Errorf("failed to store updated by ids: %s", err)
			}
			m.UpdatedByIDs = ids

			return nil
		},
	}
}

func (m *Metadata) splitIDs(s string) ([]int, error) {
	splited := strings.Split(s, ",")
	ids := make([]int, len(splited))
	for i, raw := range splited {
		var err error
		ids[i], err = strconv.Atoi(strings.TrimPrefix(strings.TrimLeft(raw, " "), "RFC"))
		if err != nil {
			return nil, err
		}
	}

	return ids, nil
}

type metaDatum struct {
	key, val string
}

func (m *metaDatum) Scan(state fmt.ScanState, _ rune) error {
	if _, _, err := state.ReadRune(); err != nil {
		return err
	}

	read, err := state.Token(true, func(char rune) bool {
		return char != ' '
	})
	if err != nil {
		return err
	}
	m.key = strings.TrimRight(string(read), ":")
	if reader, ok := m.keyReaders()[m.key]; ok {
		read, err = reader(state)
		if err != nil {
			return err
		}
		m.key += " " + string(read)
	}

	read, err = state.Token(true, func(char rune) bool {
		return char != ')'
	})
	if err != nil {
		return err
	}
	m.val = string(read)

	if _, _, err := state.ReadRune(); err != nil {
		return err
	}

	return nil
}

func (m *metaDatum) keyReaders() map[string]func(fmt.ScanState) ([]byte, error) {
	return map[string]func(fmt.ScanState) ([]byte, error){
		"Obsoleted": func(state fmt.ScanState) ([]byte, error) {
			return state.Token(true, func(char rune) bool {
				return char != ' '
			})
		},
		"Updated": func(state fmt.ScanState) ([]byte, error) {
			return state.Token(true, func(char rune) bool {
				return char != ' '
			})
		},
	}
}
