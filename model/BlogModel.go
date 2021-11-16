package model

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

var (
	cacheShorturlShortenPrefix = "cache#article_sn#title#"
)

type (
	BlogModel struct {
		sqlc.CachedConn
		table string
	}

	Blog struct {
		ArticleSn string `db:"article_sn"` // shorten key
		Title     string `db:"title"`      // original url
		Author    string `db:"author"`     // original url
		Content   string `db:"content"`    // original url
		CreateAt  string `db:"create_at"`  // original url
		UpdatedAt string `db:"updated_at"` // original url
		UserId    string `db:"user_id"`    // original url
	}
)

func NewShorturlModel(conn sqlx.SqlConn, c cache.CacheConf, table string) *BlogModel {
	return &BlogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      table,
	}
}

func (m *BlogModel) FindOne(id int) (*Blog, error) {
	shorturlShortenKey := fmt.Sprintf("%s%v", cacheShorturlShortenPrefix, id)
	var resp Blog
	err := m.QueryRow(&resp, shorturlShortenKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select article_sn, title from ` + m.table + ` where id = ? limit 1`
		fmt.Println(query)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlc.ErrNotFound
	default:
		return nil, err
	}
}
