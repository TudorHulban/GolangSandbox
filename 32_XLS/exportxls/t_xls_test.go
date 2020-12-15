package exportxls

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSteps(t *testing.T) {
	require.Equal(t, "", steps(-1))
	require.Equal(t, "", steps(7))
}

func TestRender(t *testing.T) {
	rowData := make([][]interface{}, 2)

	rowData[0] = append(rowData[0], "1")
	rowData[0] = append(rowData[0], "3")
	rowData[0] = append(rowData[0], "5")

	rowData[1] = append(rowData[1], "1")
	rowData[1] = append(rowData[1], 3)
	rowData[1] = append(rowData[1], "15")

	tt := []struct {
		testName    string
		saveTo      string
		columnNames []string
	}{
		{testName: "no column names", saveTo: "t1.xls", columnNames: []string{}},
		{testName: "with column names", saveTo: "t2.xls", columnNames: []string{"a", "b", "c"}},
	}

	for _, tc := range tt {
		cfg := Cfg{
			SheetName:   "x",
			SaveTo:      tc.saveTo,
			ColumnNames: tc.columnNames,
		}

		x := NewXLS(cfg, rowData)
		x.Process()
		require.Equal(t, 0, len(x.ErrorList))

		log.Println("--------- PASSED:", tc.testName)
	}
}
