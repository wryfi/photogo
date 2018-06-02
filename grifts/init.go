package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/wryfi/photogo/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
