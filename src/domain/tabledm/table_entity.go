package tabledm

import (
	"github.com/takuma123-type/TakumaSQL-Go/src/domain/recorddm"
)

// Table はテーブルエンティティを表します
type Table struct {
	Name    string
	Records []*recorddm.Record
}

// NewTable は新しいテーブルを作成します
func NewTable(name string) *Table {
	return &Table{
		Name:    name,
		Records: []*recorddm.Record{},
	}
}

// InsertRecord はテーブルにレコードを挿入します
func (t *Table) InsertRecord(record *recorddm.Record) {
	t.Records = append(t.Records, record)
}

// SelectAll はテーブル内のすべてのレコードを返します
func (t *Table) SelectAll() []*recorddm.Record {
	return t.Records
}
