package service

import (
	"fmt"
	"strings"

	"github.com/beriholic/cmdg/internal/system"
)

type Prompt struct {
	Basic  string
	Struct []string
}

func NewPrompt() *Prompt {
	prompt := Prompt{
		Basic:  "You first need to help the user generate the command statements they need to use based on their system environment",
		Struct: []string{},
	}

	return &prompt
}

func (p *Prompt) Build(input string) string {
	p.
		AddSystemInfo().
		AddUserInput(input).
		AddResponseStruct()

	return p.Basic + "\n" + strings.Join(p.Struct, "\n")
}

func (p *Prompt) AddSystemInfo() *Prompt {
	systemInfo := system.GetSystemInfo()
	format := `<UserSystemInfo>
  <Name>%s</Name>
  <Version>%s</Version>
</UserSystemInfo>
`
	prompt := fmt.Sprintf(format, systemInfo.Name, systemInfo.Version)

	p.Struct = append(p.Struct, prompt)
	return p
}

func (p *Prompt) AddUserInput(input string) *Prompt {
	format := `<UserInput>%s</UserInput>`
	prompt := fmt.Sprintf(format, input)

	p.Struct = append(p.Struct, prompt)
	return p
}

func (p *Prompt) AddResponseStruct() *Prompt {
	prompt := `return git commit message using this JSON schema:
	           Return {
				 "msg":string
			   	 "cmd": []string
			   }`
	p.Struct = append(p.Struct, prompt)
	return p
}
