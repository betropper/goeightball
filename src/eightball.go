package eightball

import (
	'os'
	'fmt'
	'encoding/json'
	"github.com/skilstak/go/input"
	c "github.com/skilstak/go/colors"
)

func Ask(prompt string) {
	answer := input.Ask("> " + c.B3)
	return answer
}

func LoadAnswers(path string) {
	data := path + "/answers.json"


}
