package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/takuma123-type/TakumaSQL-Go/src/domain/tabledm"
)

type FileTableRepository struct {
	filePath string
	tables   map[string]*tabledm.Table
}

func NewFileTableRepository(filePath string) *FileTableRepository {
	// 初期化時に、ファイルから既存のテーブルをロードします
	repo := &FileTableRepository{
		filePath: filePath,
		tables:   make(map[string]*tabledm.Table),
	}
	repo.loadTables()
	return repo
}
func (r *FileTableRepository) CreateTable(ctx context.Context, name string) (*tabledm.Table, error) {
	// ログを追加して処理の流れを確認
	log.Println("Attempting to create table:", name)

	if _, exists := r.tables[name]; exists {
		log.Println("Table already exists:", name)
		return nil, fmt.Errorf("Table %s already exists", name)
	}

	table := tabledm.NewTable(name)
	r.tables[name] = table

	// ファイルに書き込む
	data, err := json.MarshalIndent(r.tables, "", "  ")
	if err != nil {
		log.Println("Failed to marshal tables:", err)
		return nil, fmt.Errorf("Failed to marshal tables: %v", err)
	}

	err = ioutil.WriteFile(r.filePath, data, 0644)
	if err != nil {
		log.Println("Failed to write file:", err)
		return nil, fmt.Errorf("Failed to write file: %v", err)
	}

	log.Println("Table created and saved successfully:", name)
	return table, nil
}

// テーブルをファイルからロード
func (r *FileTableRepository) loadTables() {
	// ファイルが存在しない場合は新規作成
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		r.tables = make(map[string]*tabledm.Table)
		return
	}

	// ファイルから読み込み
	data, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}

	// JSONをパース
	err = json.Unmarshal(data, &r.tables)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
	}
}

// 特定のテーブルを取得
func (r *FileTableRepository) FindTableByName(ctx context.Context, name string) (*tabledm.Table, error) {
	log.Println("Looking for table:", name)
	table, exists := r.tables[name]
	if !exists {
		log.Println("Table not found:", name)
		return nil, fmt.Errorf("Table %s not found", name)
	}
	log.Println("Table found:", name)
	return table, nil
}

// すべてのテーブルを取得
func (r *FileTableRepository) GetAllTables(ctx context.Context) []string {
	tableNames := []string{}
	for name := range r.tables {
		tableNames = append(tableNames, name)
	}
	return tableNames
}
