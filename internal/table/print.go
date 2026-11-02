package table

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

// Tab character string
const tab = "\t"

// Default tabwriter parameters
const (
	tabMinWidth = 0
	tabWidth    = 0
	tabPadding  = 2
	tabPadChar  = ' '
	tabFlags    = 0
)

// Tabular represents any resource that can be printed as a table row
type Tabular interface {
	TabularRows() [][]string
}

// NewStandardWriter initializes with standard options
func newStandardWriter(output io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(
		output,
		tabMinWidth,
		tabWidth,
		tabPadding,
		tabPadChar,
		tabFlags,
	)
}

func PrintTable[T Tabular](output io.Writer, headers []string, items []T) error {
	w := newStandardWriter(output)

	fmt.Fprintln(w, strings.Join(headers, tab))
	for _, item := range items {
		for _, row := range item.TabularRows() {
			fmt.Fprintln(w, strings.Join(row, tab))
		}
	}
	return w.Flush()
}
