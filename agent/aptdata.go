package agent

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func RequiresUpdate(expires int64) bool {
	files, err := ioutil.ReadDir("/var/lib/apt/lists")
	if err != nil {
		log.Printf("ERROR failed to read apt list directory")
		return true
	}

	for _, f := range files  {
		if ! strings.HasPrefix(f.Name(), "deb.") {
			continue
		}
		if f.IsDir() {
			continue
		}

		if listRequiresUpdate(f.Name(), expires) {
			return true
		}
	}
	return false
}

func listRequiresUpdate(listPath string, expires int64) bool {
	content, err := os.ReadFile(listPath)
	if err != nil {
		log.Printf("ERROR could not read list file %s: %v", listPath, err)
		return true
	}

}
