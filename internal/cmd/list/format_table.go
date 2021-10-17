package list

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"sort"
)

func ternaryString(in bool, outTrue string, outFalse string) string {
	if in {
		return outTrue
	} else {
		return outFalse
	}
}

func formatTable() {
	p := *getPortDetail()
	sort.Slice(p, func(i int, j int) bool {
		return p[i].Name < p[j].Name
	})

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Display Name", "Path", "Persistent Name", "USB?", "PID", "VID", "Serial Number"})
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
				port.PID,
				port.VID,
				port.DeviceSerialNumber,
			},
		})
	}
	t.Render()
}
