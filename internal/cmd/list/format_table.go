package list

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/term"
	"os"
	"sort"
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

func formatTable() {
	p := *getPortDetail()
	sort.Slice(p, func(i int, j int) bool {
		return p[i].Name < p[j].Name
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
		return p[i].Name < p[j].Name
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
