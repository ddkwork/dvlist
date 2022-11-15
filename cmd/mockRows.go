package main

import (
	"fmt"
	"github.com/hujun-open/dvlist/packets"
	"net/http"
)

func fmtRow(row packets.Row) []string {
	return []string{
		row.Method,
		row.Scheme,
		row.Host,
		row.Path,
		row.ContentType,
		fmt.Sprint(row.ContentLength),
		row.Status,
		row.Note,
		row.Process,
		fmt.Sprint(row.PadTime),
	}
}

func mockRows() [][]string {
	rows := make([][]string, 0)
	for i := 0; i < 100; i++ {
		row := fmtRow(packets.Row{
			Method:        http.MethodConnect,
			Scheme:        "wss",
			Host:          "www.baidu.com",
			Path:          "cmsocket",
			ContentType:   "json",
			ContentLength: int64(i),
			Status:        "ok",
			Note:          "aes key",
			Process:       "steam.exe",
			PadTime:       10,
		})
		rows = append(rows, row)
	}
	return rows
}
