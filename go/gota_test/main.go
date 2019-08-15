package main

import (
	"fmt"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	df := dataframe.New(
		series.New([]string{"b", "a"}, series.String, "STRING"),
		series.New([]int{1, 2}, series.Int, "INT"),
		series.New([]float64{1.0, 2.0}, series.Float, "FLOAT"),
	)
	fmt.Println(df)

	df2 := dataframe.LoadRecords(
		[][]string{
			[]string{"STRING", "INT", "FLOAT", "BOOL"},
			[]string{"a", "1", "1.0", "true"},
			[]string{"b", "2", "2.0", "false"},
		},
	)
	fmt.Println(df2)

	type User struct {
		STRING  string
		INT     int
		FLOAT   float64
		ignored bool
	}
	users := []User{
		{"a", 1, 1.0, true},
		{"b", 2, 2.0, true},
	}
	df3 := dataframe.LoadStructs(users)
	fmt.Println(df3)

	df4 := dataframe.LoadRecords(
		[][]string{
			[]string{"STRING", "INT", "FLOAT", "BOOL"},
			[]string{"a", "1", "1.0", "true"},
			[]string{"b", "2", "2.0", "false"},
		},
		dataframe.DetectTypes(false),
		dataframe.DefaultType(series.Float),
		dataframe.WithTypes(map[string]series.Type{
			"STRING": series.String,
			"BOOL":   series.Bool,
		}),
	)
	fmt.Println(df4)

	df5 := dataframe.LoadMaps(
		[]map[string]interface{}{
			map[string]interface{}{
				"STRING": "a",
				"INT":    1,
				"FLOAT":  1.0,
				"BOOL":   true,
			},
			map[string]interface{}{
				"STRING": "b",
				"INT":    2,
				"FLOAT":  2.0,
				"BOOL":   true,
			},
		},
	)

	fmt.Println(df5)

	csvStr := `
string,int,float,bool
"a",1,1.0,true
"b",2,2.0,false
`
	df6 := dataframe.ReadCSV(strings.NewReader(csvStr))

	fmt.Println(df6)

	jsonStr := `[{"INT":1,"FLOAT":1.0},{"STRING":"b","INT":2,"FLOAT":2.0}]`
	df7 := dataframe.ReadJSON(strings.NewReader(jsonStr))

	fmt.Println(df7)
}
