package functions

import (
	"log"

	"github.com/koki-develop/go-fzf"
)

func FzfSelector(namespaces []string) string {
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(namespaces, func(i int) string { return namespaces[i] })
	if err != nil {
		log.Fatal(err)
	}

	return namespaces[idxs[0]]
}
