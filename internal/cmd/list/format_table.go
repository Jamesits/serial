package list

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/term"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const TerminalDefaultWidth int = 80

func ternaryString(in bool, outTrue string, outFalse string) string {
	if in {
		return outTrue
	} else {
		return outFalse
	}
}

func getTerminalWidth() int {
	width, _, err := term.GetSize(0)
	if err != nil {
		return width
	} else {
		return TerminalDefaultWidth
	}
}

var portSortLogicalClassifier = regexp.MustCompilePOSIX(`([[:alnum:]]+[[:alpha:]])([[:digit:]]+)$`)

// sortPort can compare 2 port names in a logical method
// Traditional sorting: "COM1", "COM11", "COM2"
// Logical sorting: "COM1", "COM2", "COM11"
func sortPort(portNameA string, portNameB string) bool {
	matchA := portSortLogicalClassifier.FindStringSubmatch(portNameA)
	matchB := portSortLogicalClassifier.FindStringSubmatch(portNameB)

	for {
		if len(matchA) < 3 || len(matchB) < 3 {
			break
		}

		numA, err := strconv.Atoi(matchA[2])
		if err != nil {
			break
		}

		numB, err := strconv.Atoi(matchB[2])
		if err != nil {
			break
		}

		// use logical order
		if matchA[1] == matchB[1] {
			return numA < numB
		} else {
			return matchA[1] < matchB[1]
		}
	}

	// if either name is not in the format of "COM1" or "ttyUSB1" then fallback to the traditional order
	return portNameA < portNameB
}

func formatTable() {
	p := *getPortDetail()
	sort.Slice(p, func(i int, j int) bool {
		return sortPort(p[i].Name, p[j].Name)
	})

	t := table.NewWriter()
	t.SetAutoIndex(true)
	t.SetAllowedRowLength(getTerminalWidth())
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Path", "Display Name", "USB?", "VID", "PID", "Serial"})
	for _, port := range p {
		//t.AppendSeparator()
		t.AppendRows([]table.Row{
			{
				port.Path,
				port.DisplayName,
				ternaryString(port.IsUSB, "Yes", "No"),
				port.VID,
				port.PID,
				port.DeviceSerialNumber,
			},
		})
	}
	t.Render()
}

func formatTableWide() {
	p := *getPortDetail()
	sort.Slice(p, func(i int, j int) bool {
		return sortPort(p[i].Name, p[j].Name)
	})

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAllowedRowLength(getTerminalWidth())
	t.AppendHeader(table.Row{"#", "Name", "Display Name", "Path", "Persistent Name", "USB?", "VID", "PID", "Serial Number"})
	for i, port := range p {
		//t.AppendSeparator()
		t.AppendRows([]table.Row{
			{
				i,
				port.Name,
				port.DisplayName,
				port.Path,
				port.PersistentName,
				ternaryString(port.IsUSB, "Yes", "No"),
				port.VID,
				port.PID,
				port.DeviceSerialNumber,
			},
		})
	}
	t.Render()
}
