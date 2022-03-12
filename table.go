package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

var columnHeader = []string{
	//"Name",
	"Symbol",
	"Price (USD)",
	"Volume (USD)",
	"Market (USD)",
	"Change (1h)",
	"Change (24h)",
	"Change (7d)"}

var priceColumnName = "Change (1h)"
var NormalColorHeader = tablewriter.Colors{tablewriter.FgHiWhiteColor, tablewriter.Bold, tablewriter.BgGreenColor}
var HighlightColorHeader = tablewriter.Colors{tablewriter.FgHiWhiteColor, tablewriter.Bold, tablewriter.BgRedColor}
var NormalColorColumn = tablewriter.Colors{tablewriter.FgHiWhiteColor, tablewriter.Bold}
var HighlightColorColumn = tablewriter.Colors{tablewriter.FgCyanColor, tablewriter.Bold}
var RedColorColumn = tablewriter.Colors{tablewriter.FgRedColor, tablewriter.Bold}

func BuildTable() *tablewriter.Table {

	if len(columnHeader) < 1 {
		panic("column header size must > 1 , please check !")
	}
	//columnIndex := sort.StringSlice(columnHeader).Search(priceColumnName)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(columnHeader)

	// todo auto config 👆🏻
	buildHeaderColor(table, len(columnHeader), 4)
	buildColumnColor(table, len(columnHeader), 4)

	// init table
	var data [][]string
	PullByLine(func(line []string) {
		data = append(data, line)
	})

	for _, row := range data {
		if strings.ContainsRune(row[4], '-') {
			table.Rich(row, []tablewriter.Colors{NormalColorColumn, NormalColorColumn, NormalColorColumn, NormalColorColumn, RedColorColumn, NormalColorColumn, NormalColorColumn})
		} else {
			table.Rich(row, []tablewriter.Colors{NormalColorColumn, NormalColorColumn, NormalColorColumn, NormalColorColumn, HighlightColorColumn, NormalColorColumn, NormalColorColumn})
		}
	}

	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	//table.AppendBulk(data) // Add Bulk Data
	clearFromTop()
	table.Render()
	return table
}

func buildHeaderColor(table *tablewriter.Table, len int, highlightIndex int) {
	var headerColors []tablewriter.Colors

	for i := 0; i < len; i++ {
		if i == highlightIndex {
			headerColors = append(headerColors, HighlightColorHeader)
		} else {
			headerColors = append(headerColors, NormalColorHeader)
		}
	}

	table.SetHeaderColor(headerColors...)
}

func buildColumnColor(table *tablewriter.Table, len int, highlightIndex int) {
	var columColors []tablewriter.Colors

	for i := 0; i < len; i++ {
		if i == highlightIndex {
			columColors = append(columColors, HighlightColorColumn)
		} else {
			columColors = append(columColors, NormalColorColumn)
		}
	}

	table.SetColumnColor(columColors...)
}
