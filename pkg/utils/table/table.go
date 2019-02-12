package table

import (
	"fmt"
	"os"

	"github.com/gosuri/uitable"
	"github.com/olekukonko/tablewriter"
)

// Type : table type
type Type int

// table type
const (
	Horizontal Type = 0
	Vertical   Type = 1
)

type table struct {
	Type    Type
	Headers []string
	Rows    [][]string
}

// Table interface
type Table interface {
	SetHeaders(h []string)
	AddRow(r []string) error
	IsEmpty() bool
	Println()
}

// NewTable is init a new default table
func NewTable(t Type) Table {
	return &table{
		Type: t,
	}
}

// SetHeaders is set headers to table
func (t *table) SetHeaders(h []string) {
	t.Headers = h
}

// AddRow is add row to table
func (t *table) AddRow(r []string) error {
	if len(r) == len(t.Headers) {
		t.Rows = append(t.Rows, r)
		return nil
	}
	return fmt.Errorf("NUMBER OF GRID CELLS AND THE  HEAD INCONSISTENCIES")
}

// Println is print table to stdout
func (t *table) Println() {
	switch t.Type {
	case Horizontal:
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(t.Headers)
		for _, r := range t.Rows {
			table.Append(r)
		}
		table.Render()
	case Vertical:
		table := uitable.New()
		//table.MaxColWidth = 200
		table.Wrap = true // wrap columns
		for i, r := range t.Rows {
			table.AddRow(fmt.Sprintf("*** %d ***", i)) // blank
			for k, v := range r {
				table.AddRow(fmt.Sprintf("%s:", t.Headers[k]), v)
			}
			table.AddRow("") // blank
		}
		fmt.Println(table)
	}
}
// IsEmpty is check is check if table rows is more than zero.
func (t *table) IsEmpty() bool {
	if len(t.Rows) > 0 {
		return false
	}
	return true
}
