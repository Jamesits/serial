package list

import (
	"fmt"
	"github.com/alessio/shellescape"
	log "github.com/sirupsen/logrus"
	"go.bug.st/serial/enumerator"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type serialPortDetail struct {
	Name               string `json:"name" csv:"Name"`
	DisplayName        string `json:"display_name" csv:"Display Name"`
	Path               string `json:"path" csv:"Path"`
	PersistentName     string `json:"persistent_name" csv:"Persistent Name"`
	IsUSB              bool   `json:"is_usb" csv:"Is USB?"`
	PID                string `json:"pid,omitempty" csv:"PID"`
	VID                string `json:"vid,omitempty" csv:"VID"`
	DeviceSerialNumber string `json:"serial_number,omitempty" csv:"Serial Number"`
}

type serialPortEnvelope struct {
	Ports *[]serialPortDetail `json:"ports"`
}

func getPortDetail() *[]serialPortDetail {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.WithError(err).Errorln("failed getting serial port details")
		return nil
	}

	portDetail := []serialPortDetail{}

	for _, port := range ports {
		log.Tracef("reading port %s", port.Name)

		path := port.Name
		persistentName := port.Name

		switch runtime.GOOS {
		case "windows":
			path = fmt.Sprintf(`\\.\%s`, port.Name)
			persistentName = port.Name
		case "linux":
			// try using `udevadm` to get the device location
			out, err := exec.Command(fmt.Sprintf("udevadm info -q path -n '%s'", shellescape.Quote(port.Name))).Output()
			if err == nil {
				persistentName = strings.TrimSpace(string(out))
				log.Tracef("Got port location: %s", persistentName)
			} else {
				log.WithError(err).Tracef("udevadm command failure")
			}
		}

		// Product might be missing
		// On Windows, it will be in the format of "Prolific PL2303GT USB Serial COM Port (COM13)"
		displayName := port.Name
		if port.Product != "" {
			displayName = port.Product
		}

		portDetail = append(portDetail, serialPortDetail{
			Name:           port.Name,
			DisplayName:    displayName,
			Path:           path,
			PersistentName: persistentName,

			// Note: if port is not from USB, you won't get PID & VID
			IsUSB:              port.IsUSB,
			PID:                port.PID,
			VID:                port.VID,
			DeviceSerialNumber: port.SerialNumber,
		})
	}

	return &portDetail
}

type marshaller func(interface{}) ([]byte, error)

func portDetailDump(portDetail interface{}, marshaller marshaller) error {
	b, err := marshaller(portDetail)
	if err != nil {
		log.WithError(err).Errorln("formatter error")
	}
	n, err := os.Stdout.Write(b)
	log.Tracef("%d bytes flushed to stdout", n)
	if err != nil {
		log.WithError(err).Errorln("stdout write failed")
	}

	return err
}
