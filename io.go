package io

import (
	"fmt"
	"os"
	"strings"
)

// mapArray2tsv 将 []map[string]string 转换为 TSV 格式并写入到文件或标准输出流(未提供output文件名时)
func mapArray2tsv(data []map[string]string, output_columns []string, output string) error {
	// 如果output是文件名，则打开文件用于写入
	var writer *os.File
	if output != "" {
		file, err := os.Create(output)
		if err != nil {
			return fmt.Errorf("无法创建输出文件: %w", err)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	// 定义列顺序，假设我们想要按照某个固定的顺序输出键对应的值
	var columns []string
	if len(output_columns) > 0 {
		columns = output_columns
	} else {
		if len(data) > 0 {
			for col := range data[0] {
				columns = append(columns, col)
			}
		}
	}
	// 首行输出列标题
	for i, col := range columns {
		if i > 0 {
			_, _ = writer.Write([]byte("\t"))
		}
		_, _ = writer.WriteString(col)
	}
	_, _ = writer.Write([]byte("\n"))

	// 遍历数据，并按列顺序输出
	for _, item := range data {
		for i, col := range columns {
			value, exists := item[col]
			if exists {
				if i > 0 {
					_, _ = writer.Write([]byte("\t"))
				}
				fit_value := strings.Replace(value, "\n", "^", -1)
				_, _ = writer.WriteString(fit_value)
			} else {
				// 若列不存在于item中，可以填充默认值或者错误信息
				if i > 0 {
					_, _ = writer.Write([]byte("\t"))
				}
				_, _ = writer.WriteString("(Column not found)")
			}
		}
		_, _ = writer.Write([]byte("\n"))
	}

	return nil
}
