package tablepp

import (
	"fmt"
	"strconv"
	"strings"
)

// Alignment type to fspecify to column alignment
type Alignment uint

// AlignmentLeft in column
const AlignmentLeft = Alignment(1)

// AlignmentRight in column
const AlignmentRight = Alignment(2)

// Header info of table
type Header struct {
	Name      string
	Width     int
	Alignment Alignment
}

// Row info of table
type Row []interface{}

// Table info
type Table struct {
	Name    string
	Headers []Header
	Rows    []Row
}

// return pretty print string of table
func (t *Table) String() string {

	// to form table border line
	borderLine := "\n"

	// pattern to form row line
	rowLinePattern := "\n"

	totalWidth := 0
	headerStrings := make([]interface{}, 0)
	for _, header := range t.Headers {
		borderLine += "+" + strings.Repeat("-", header.Width)
		totalWidth += header.Width
		if header.Alignment == AlignmentLeft {
			rowLinePattern += "|%-" + strconv.Itoa(header.Width) + "s"
		} else {
			rowLinePattern += "|%" + strconv.Itoa(header.Width) + "s"
		}
		headerStrings = append(headerStrings, header.Name)
	}

	rowLinePattern += "|"
	borderLine += "+"

	// forming table name line
	space := strings.Repeat(" ", ((totalWidth - len(t.Name)) / 2))
	nameLine := fmt.Sprintf("\n| %-"+strconv.Itoa(totalWidth)+"s|", space+t.Name+space)

	//forming table header line
	headerLine := fmt.Sprintf(rowLinePattern, headerStrings...)

	// forming table printstring
	printStr := borderLine

	printStr += nameLine

	printStr += borderLine

	printStr += headerLine

	printStr += borderLine

	for _, row := range t.Rows {
		rowLine := fmt.Sprintf(rowLinePattern, row...)
		printStr += rowLine
	}
	printStr += borderLine

	return printStr
}

// AddRow to table
func (t *Table) AddRow(row []interface{}) {
	t.Rows = append(t.Rows, row)
}
