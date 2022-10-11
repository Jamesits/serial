package list

import (
	"fmt"
	"github.com/alessio/shellescape"
	log "github.com/sirupsen/logrus"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type serialPortDetail struct {
	Available          bool   `json:"available" csv:"Available"`
	Name               string `json:"name" csv:"Name"`
	DisplayName        string `json:"display_name" csv:"Display Name"`
	Path               string `json:"path" csv:"Path"`
	PersistentName     string `json:"persistent_name" csv:"Persistent Name"`
	Interface          string `json:"interface" csv:"Interface"`
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

	var portDetail []serialPortDetail

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

		d := serialPortDetail{
			Available: true,

			Name:           port.Name,
			DisplayName:    displayName,
			Path:           path,
			PersistentName: persistentName,

			// Note: if port is not from USB, you won't get PID & VID
			PID:                port.PID,
			VID:                port.VID,
			DeviceSerialNumber: port.SerialNumber,
		}
		if port.IsUSB {
			d.Interface = "USB"
		} else {
			d.Interface = ""
		}

		portDetail = append(portDetail, d)
	}

	// on Windows, there might be a reference to a COMx port in `HKEY_LOCAL_MACHINE\HARDWARE\DEVICEMAP\SERIALCOMM`, but
	// no corresponding device available from SetupApi. We try to list theses too.
	allPorts, err := serial.GetPortsList()
	if err != nil {
		log.WithError(err).Errorln("error getting serial port list")
		return nil
	}

	for _, portName := range allPorts {
		found := false

		for _, port := range portDetail {
			if port.Name == portName {
				found = true
				break
			}
		}

		if !found {
			portDetail = append(portDetail, serialPortDetail{
				Available:   false,
				Name:        portName,
				DisplayName: portName,
			})
		}
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
