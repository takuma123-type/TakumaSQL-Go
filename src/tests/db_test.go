package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/TakumaSQL-Go/src/infra/router"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// テーブルとレコードのルーターを設定
	router.NewTableRouter(r)
	router.NewRecordRouter(r)

	return r
}

func TestCreateTableAndInsertRecord(t *testing.T) {
	// テスト用のルーターをセットアップ
	r := setupRouter()

	// テーブル作成リクエスト
	createTablePayload := []byte(`{"name":"users"}`)
	req, _ := http.NewRequest("POST", "/api/tables", bytes.NewBuffer(createTablePayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// テーブル作成結果を確認
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, w.Code)
	}
	if !bytes.Contains(w.Body.Bytes(), []byte("Table created successfully")) {
		t.Errorf("Expected success message in response body, but got: %s", w.Body.String())
	}

	// レコード作成リクエスト
	createRecordPayload := []byte(`{"fields":{"name":"John Doe","email":"john@example.com"}}`)
	req, _ = http.NewRequest("POST", "/api/records", bytes.NewBuffer(createRecordPayload))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// レコード作成結果を確認
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, w.Code)
	}
	if !bytes.Contains(w.Body.Bytes(), []byte("Record created successfully")) {
		t.Errorf("Expected success message in response body, but got: %s", w.Body.String())
	}
}
