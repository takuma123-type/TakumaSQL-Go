package tableinput

// CreateTableInputは、テーブル作成のリクエストデータを表します
type CreateTableInput struct {
	Name string `json:"name"`
}
