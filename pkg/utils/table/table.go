package table

import (
	"fmt"
	"os"

	"github.com/gosuri/uitable"
	"github.com/olekukonko/tablewriter"
)

type Type int

const (
	Horizontal Type = 0
	Vertical   Type = 1
)

type table struct {
	Type    Type
	Headers []string
	Rows    [][]string
}

type Table interface {
	SetHeaders(h []string)
	AddRow(r []string) error
	IsEmpty() bool
	Println()
}

func NewTable(t Type) Table {
	return &table{
		Type: t,
	}
}

func (t *table) SetHeaders(h []string) {
	t.Headers = h
}

func (t *table) AddRow(r []string) error {
	if len(r) == len(t.Headers) {
		t.Rows = append(t.Rows, r)
		return nil
	}
	return fmt.Errorf("Number of grid cells and the head inconsistencies.")
}

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
		for _, r := range t.Rows {
			for k, v := range r {
				table.AddRow(fmt.Sprintf("%s:", t.Headers[k]), v)
			}
			table.AddRow("") // blank
		}
		fmt.Println(table)
	}
}

func (t *table) IsEmpty() bool {
	if len(t.Rows) > 0 {
		return false
	}
	return true
}
