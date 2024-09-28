package tabledm

import (
	"container/list"

	"github.com/takuma123-type/TakumaSQL-Go/src/domain/recorddm"
)

// Table はテーブルエンティティを表します
type Table struct {
	Name    string
	Records *list.List
}

// NewTable は新しいテーブルを作成します
func NewTable(name string) *Table {
	return &Table{
		Name:    name,
		Records: list.New(),
	}
}

// InsertRecord はテーブルにレコードを挿入します
func (t *Table) InsertRecord(record recorddm.Record) {
	t.Records.PushBack(record)
}

// SelectAll はテーブル内のすべてのレコードを返します
func (t *Table) SelectAll() []*recorddm.Record {
	var records []*recorddm.Record
	for e := t.Records.Front(); e != nil; e = e.Next() {
		records = append(records, e.Value.(*recorddm.Record))
	}
	return records
}
