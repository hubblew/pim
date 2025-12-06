package installer

import (
	"fmt"

	"github.com/hubblew/pim/internal/ui"
)

// UserPrompter handles user interaction for confirmation prompts.
type UserPrompter interface {
	ConfirmOverwrite(path string) (bool, error)
}

// interactivePrompter prompts the user via stdin for confirmation.
type interactivePrompter struct{}

var _ UserPrompter = (*interactivePrompter)(nil)

func NewInteractivePrompter() UserPrompter {
	return &interactivePrompter{}
}

func (p *interactivePrompter) ConfirmOverwrite(path string) (bool, error) {
	fmt.Printf("File %s already exists. Overwrite?\n", path)

	choice, err := ui.NewChoiceDialog("Please confirm:", ui.ChoicesYesNo()).Run()

	if err != nil {
		return false, fmt.Errorf("failed to get user input: %w", err)
	}
	if choice == nil {
		return false, nil
	}

	return choice.Value.(bool), nil
}

type acceptAllPrompter struct{}

var _ UserPrompter = (*acceptAllPrompter)(nil)

func NewAcceptAllPrompter() UserPrompter {
	return &acceptAllPrompter{}
}

func (p *acceptAllPrompter) ConfirmOverwrite(_ string) (bool, error) {
	return true, nil
}
