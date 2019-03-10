package watch

import (
	"github.com/x1unix/gilbert/logging"
	"github.com/x1unix/gilbert/manifest"
	"github.com/x1unix/gilbert/plugins"
	"github.com/x1unix/gilbert/scope"
)

func NewWatchPlugin(scope *scope.Scope, rawParams manifest.RawParams, log logging.Logger) (plugins.Plugin, error) {
	params, err := parseParams(rawParams, scope)
	if err != nil {
		return nil, err
	}

	p, err := newPlugin(scope, *params, log)
	if err != nil {
		return nil, err
	}
	return p, nil
}