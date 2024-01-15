package templater

import (
	"fmt"
	"regexp"
)

var (
	re0 *regexp.Regexp
	re1 *regexp.Regexp
)

func init() {
	// match %
	re0 = regexp.MustCompile(`%.`)
	// match ${variable}
	re1 = regexp.MustCompile(`\$\{(\w+)(?:\|%(%.+?))?\}`)
}

type Templater struct {
	template string
	args     []string
}

func NewTemplater(template string) *Templater {
	var args []string

	template = re0.ReplaceAllStringFunc(template, func(match string) string {
		if match == "%%" {
			return "%%%%"
		} else {
			return "%" + match
		}
	})

	template = re1.ReplaceAllStringFunc(template, func(match string) string {
		inMatch := re1.FindAllStringSubmatch(match, -1)
		name := inMatch[0][1]
		arg := inMatch[0][2]

		if arg == "" {
			arg = "%v"
		}
		args = append(args, name)
		return arg
	})

	return &Templater{template: template, args: args}
}

func (t *Templater) Parse(values map[string]interface{}) string {
	length := len(t.args)
	if length == 0 {
		return t.template
	}

	argValues := make([]interface{}, length)
	for i, arg := range t.args {
		argValues[i] = values[arg]
	}
	return fmt.Sprintf(t.template, argValues...)
}
