package tablepp

import (
	"testing"
)

func Test_TablePP(t *testing.T) {
	t.Run("PrettyPrintTable1", func(t *testing.T) {

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
		expected := `
+---------------+---------------+
|           TestTable           |
+---------------+---------------+
|TableCol1      |TableCol2      |
+---------------+---------------+
|Row1Cell1      |Row1Cell2      |
|Row2Cell1      |Row2Cell2      |
+---------------+---------------+`
		actual := table.String()
		t.Log("Actual:")
		t.Logf("%s", table.String())
		if actual != expected {
			t.Log("Expected:")
			t.Errorf("%s", expected)
		}
	})
}
