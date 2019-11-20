package single

import "fmt"

type Metadata struct {
	ID int
}

func (m *Metadata) Scan(state fmt.ScanState, _ rune) error {
	if _, err := state.Token(true, func(char rune) bool {
		return char != ':'
	}); err != nil {
		return err
	}
	state.ReadRune()
	state.SkipSpace()

	_, err := fmt.Fscan(state, &m.ID)
	if err != nil {
		return err
	}

	for {
		if _, err := state.Token(true, func(char rune) bool {
			return char != '\n'
		}); err != nil {
			return err
		}
		if _, _, err := state.ReadRune(); err != nil {
			return err
		}

		char, _, err := state.ReadRune()
		if err != nil {
			return err
		}
		if char == '\n' {
			break
		}
	}

	return nil
}
