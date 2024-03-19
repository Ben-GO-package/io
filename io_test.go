package io

import (
	"testing"
)

func Test_mapArray2tsv(t *testing.T) {
	var data = []map[string]string{
		{"First Name": "John", "Last Name": "Doe"},
		{"First Name": "Jane", "Last Name": "Doe"},
	}
	var columns = []string{"First Name", "Last Name"}

	err := mapArray2tsv(data, columns, "test.tsv")
	if err != nil {
		t.Error("Error:", err)
	}
}
