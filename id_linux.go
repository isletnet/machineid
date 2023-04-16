//go:build linux
// +build linux

package machineid

import (
	"crypto/md5"
	"fmt"
)

const (
	// dbusPath is the default path for dbus machine id.
	// dbusPath = "/var/lib/dbus/machine-id"
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	// dbusPathEtc = "/etc/machine-id"

	isletIDPath = "/etc/islet-id"
	uuidPath    = "/proc/sys/kernel/random/uuid"
)

// machineID returns the uuid specified at `/var/lib/dbus/machine-id` or `/etc/machine-id`.
// If there is an error reading the files an empty string is returned.
// See https://unix.stackexchange.com/questions/144812/generate-consistent-machine-unique-id
func machineID() (string, error) {
	var id string
	idbuf, err := readFile(isletIDPath)
	if err != nil {
		// try fallback path
		// buf := &bytes.Buffer{}
		// err := run(buf, os.Stderr, "cat", uuidPath)
		// if err != nil {
		// 	return "", err
		// }
		uuid, err := readFile(uuidPath)
		if err != nil {
			return "", err
		}

		id = fmt.Sprintf("%x", md5.Sum(uuid))
		writeFile(isletIDPath, []byte(id))

	} else {
		id = string(idbuf)
	}
	if err != nil {
		return "", err
	}
	return trim(id), nil
}
