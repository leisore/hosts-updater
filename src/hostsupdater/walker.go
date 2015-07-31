package hostsupdater

import (
	"io"
)

type Author struct {
	Name  string
	Email string
}

type Walker interface {
	Name() string
	Version() string
	License() string
	Desc() string
	Authors() []Author
	WalkedHosts() (io.Reader, error)
}

var walker_registry map[string]Walker

func init() {
	walker_registry = make(map[string]Walker)
}

func RegisterWalker(w Walker) {
	walker_registry[w.Name()] = w
}

func GetWalkers() []Walker {
	walkers := make([]Walker, 0)
	for _, v := range walker_registry {
		walkers = append(walkers, v)
	}
	return walkers
}
