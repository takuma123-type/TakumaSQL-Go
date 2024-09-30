package tabledm

import (
	"context"
)

// TableRepository インターフェース
type TableRepository interface {
	CreateTable(ctx context.Context, name string) (*Table, error)     // テーブルを作成するメソッド
	FindTableByName(ctx context.Context, name string) (*Table, error) // 名前でテーブルを検索するメソッド
	GetAllTables(ctx context.Context) []string                        // すべてのテーブルを取得するメソッド
}
