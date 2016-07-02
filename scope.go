// This pacakge IS THREAD UNSAFE
// Use it only for initialization from one gorutine.
package scope

import (
	"strings"
)

type Scope struct {
	parent   *Scope
	fullName string
}

func (s Scope) Name() string {
	return s.fullName
}

// Call as `defer scope.SubScope("myname")()` to close it automatically
func SubScope(name string) func() {
	newScope := &Scope{
		parent:   scope,
		fullName: strings.Trim(scope.fullName+"."+name, "."),
	}
	scope = newScope
	return func() {
		if scope != newScope {
			panic("invalud scope, expected " + newScope.fullName + " got " + scope.fullName)
		} else {
			scope = newScope.parent
		}
	}
}

var scope = &Scope{
	parent:   nil,
	fullName: "",
}

func Current() *Scope {
	return scope
}
