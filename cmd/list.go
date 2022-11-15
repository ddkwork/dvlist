package main

import (
	"github.com/hujun-open/dvlist"
	"github.com/hujun-open/dvlist/packets"
	"sort"
)

type (
	Interface interface{ dvlist.Data }
	object    struct{ rows [][]string }
)

func newObject() Interface { return &object{rows: mockRows()} }

func (o *object) Header() []string {
	return []string{
		packets.NamePacketField.Method(),
		packets.NamePacketField.Scheme(),
		packets.NamePacketField.Host(),
		packets.NamePacketField.Path(),
		packets.NamePacketField.ContentType(),
		packets.NamePacketField.ContentLength(),
		packets.NamePacketField.Status(),
		packets.NamePacketField.Notes(),
		packets.NamePacketField.Process(),
		packets.NamePacketField.PadTime(),
	}
}

func (o *object) SetColumnWidth(id int, width float32) {
	//TODO implement me
	panic("implement me")
}

func (o *object) Row(id int) []string { return o.rows[id] }

func (o *object) AddRow(row []string) {
	//TODO implement me
	panic("implement me")
}

func (o *object) Len() int { return len(o.rows) }

func (o *object) Sort(id int, ascend bool) {
	sort.Slice(o.rows[id], func(i, j int) bool {
		if ascend {
			return i > j
		}
		return i < j
	})
}

func (o *object) Filter(columnValue string, id int) {
	//TODO implement me
	panic("implement me")
}
