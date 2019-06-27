package grifts

import (
	"github.com/boskiv/dob_api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
