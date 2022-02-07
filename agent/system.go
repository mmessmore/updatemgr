package agent

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var updateTTL = 6 * 3600

func getUpdates() {
	log.Println("INFO updating apt cache")
	err := exec.Command("apt-get", "update").Run()
	if err != nil {
		log.Printf("ERROR updating apt: %v", err)
	}
}

func upToDate() bool {
	fInfo, err := os.Stat("/var/cache/apt/pkgcache.bin")
	if err != nil {
		log.Printf("ERROR opening apt cache: %v", err)
		return false
	}
	modTime := fInfo.ModTime().Unix()
	now := time.Now().Unix()

	return (now-modTime < int64(updateTTL))
}

func GetUpdatesAvailable() UpdatesAvailable {
	log.Println("INFO finding available updates")
	if !upToDate() {
		getUpdates()
	}

	hostname, _ := os.Hostname()
	ua := UpdatesAvailable{
		Name:     hostname,
		Packages: []string{},
	}

	//apt list --upgradeable| cut -d / -f1 > foo
	// first line is garbage
	out, err := exec.Command("apt", "list", "--upgradeable").Output()
	if err != nil {
		log.Printf("ERROR getting upgradable packages: %v", err)
	}

	outLines := strings.Split(string(out), "\n")
	for i := range outLines {
		// skip line 1 that says "Listing..."
		if i == 0 {
			continue
		}
		if outLines[i] == "" {
			continue
		}
		pkg := strings.Split(outLines[i], "/")[0]
		ua.Packages = append(ua.Packages, pkg)
	}

	log.Printf("Updates Available\n%s", ua.Marshall())

	return ua

}

func Upgrade() error {
	log.Println("INFO performing upgrade")
	err := exec.Command("apt-get", "upgrade", "-y").Run()
	if err != nil {
		log.Printf("ERROR upgrading packages: %v", err)
	}
	return err
}
