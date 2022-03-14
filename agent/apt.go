package agent

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

var updateTTL = 6 * 3600

func getUpdates() {
	log.Info().
		Msg("Updating apt cache")
	err := exec.Command("apt-get", "update").Run()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed updating apt")
	}
}

func upToDate() bool {
	fInfo, err := os.Stat("/var/cache/apt/pkgcache.bin")
	if err != nil {
		log.Error().
			Err(err).
			Msg("Updating apt cache")
		return false
	}
	modTime := fInfo.ModTime().Unix()
	now := time.Now().Unix()

	return (now-modTime < int64(updateTTL))
}

func GetUpdatesAvailable() UpdatesAvailable {
	log.Info().
		Msg("Finding updates")
	if !upToDate() {
		getUpdates()
	}

	hostname, _ := os.Hostname()
	ua := UpdatesAvailable{
		Name:     hostname,
		Packages: []string{},
	}

	// apt list --upgradeable| cut -d / -f1 > foo
	// first line is garbage
	out, err := exec.Command("apt", "list", "--upgradeable").Output()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Getting upgradable packages")
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

	log.Info().
		Strs("packages", ua.Packages).
		Msg("Found upgradable packages")

	return ua

}

func DoUpgrade() error {
	log.Info().Msg("Performing upgrade")
	err := exec.Command("apt-get", "full-upgrade", "-y").Run()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed upgrading packages")
	}
	log.Print("Finished upgrade")

	log.Print("Performing autoremove")
	err = exec.Command("apt-get", "autoremove", "-y").Run()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed running autoremove")
	}
	log.Print("Finished autoremove")
	return err
}

func IsRebootRequired() RebootRequired {
	log.Print("Checking if reboot is required")
	hostname, _ := os.Hostname()
	_, err := os.Stat("/var/run/reboot-required")
	if err != nil {
		return RebootRequired{
			Name:           hostname,
			RebootRequired: false,
		}
	}
	return RebootRequired{
		Name:           hostname,
		RebootRequired: true,
	}
}

func DoReboot() error {
	log.Print("Performing reboot NOW!!!")
	err := exec.Command("reboot").Run()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed rebooting")
	}
	return err
}
