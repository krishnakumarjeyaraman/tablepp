# tablepp
Pretty print table in go lang

Example:

Initialize table like below

Note. Rows can be added via AddRow method on table 

```
table := Table{
			Name: "TestTable",
			Headers: []Header{
				Header{
					Name:      "TableCol1",
					Width:     15,
					Alignment: AlignmentLeft,
				},
				Header{
					Name:      "TableCol2",
					Width:     15,
					Alignment: AlignmentLeft,
				},
			},
			Rows: []Row{
				Row{
					"Row1Cell1",
					"Row1Cell2",
				},
				Row{
					"Row2Cell1",
					"Row2Cell2",
				},
			},
		}
```
if you print table.String(), 

```
+---------------+---------------+
|           TestTable           |
+---------------+---------------+
|TableCol1      |TableCol2      |
+---------------+---------------+
|Row1Cell1      |Row1Cell2      |
|Row2Cell1      |Row2Cell2      |
+---------------+---------------+
```
