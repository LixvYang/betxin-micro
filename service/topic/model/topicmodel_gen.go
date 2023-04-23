// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	topicFieldNames          = builder.RawFieldNames(&Topic{})
	topicRows                = strings.Join(topicFieldNames, ",")
	topicRowsExpectAutoSet   = strings.Join(stringx.Remove(topicFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	topicRowsWithPlaceHolder = strings.Join(stringx.Remove(topicFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheBetxinTopicIdPrefix                = "cache:betxin:topic:id:"
	cacheBetxinTopicTidPrefix               = "cache:betxin:topic:tid:"
	cacheBetxinTopicTitleIntroContentPrefix = "cache:betxin:topic:title:intro:content:"
)

type (
	topicModel interface {
		Insert(ctx context.Context, data *Topic) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Topic, error)
		FindOneByTid(ctx context.Context, tid string) (*Topic, error)
		FindOneByTitleIntroContent(ctx context.Context, title string, intro string, content string) (*Topic, error)
		Update(ctx context.Context, data *Topic) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTopicModel struct {
		sqlc.CachedConn
		table string
	}

	Topic struct {
		Id            int64        `db:"id"`
		Tid           string       `db:"tid"`
		Cid           int64        `db:"cid"`
		Title         string       `db:"title"`
		Intro         string       `db:"intro"`
		Content       string       `db:"content"`
		YesRatio      float64      `db:"yes_ratio"`
		NoRatio       float64      `db:"no_ratio"`
		YesPrice      float64      `db:"yes_price"`
		NoPrice       float64      `db:"no_price"`
		TotalPrice    float64      `db:"total_price"`
		CollectCount  int64        `db:"collect_count"`
		ReadCount     int64        `db:"read_count"`
		ImgUrl        string       `db:"img_url"`
		IsStop        int64        `db:"is_stop"` // 0没有停止 1 停止了
		RefundEndTime time.Time    `db:"refund_end_time"`
		CreatedAt     time.Time    `db:"created_at"`
		UpdatedAt     time.Time    `db:"updated_at"`
		DeletedAt     sql.NullTime `db:"deleted_at"`
	}
)

func newTopicModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTopicModel {
	return &defaultTopicModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`topic`",
	}
}

func (m *defaultTopicModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	betxinTopicIdKey := fmt.Sprintf("%s%v", cacheBetxinTopicIdPrefix, id)
	betxinTopicTidKey := fmt.Sprintf("%s%v", cacheBetxinTopicTidPrefix, data.Tid)
	betxinTopicTitleIntroContentKey := fmt.Sprintf("%s%v:%v:%v", cacheBetxinTopicTitleIntroContentPrefix, data.Title, data.Intro, data.Content)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, betxinTopicIdKey, betxinTopicTidKey, betxinTopicTitleIntroContentKey)
	return err
}

func (m *defaultTopicModel) FindOne(ctx context.Context, id int64) (*Topic, error) {
	betxinTopicIdKey := fmt.Sprintf("%s%v", cacheBetxinTopicIdPrefix, id)
	var resp Topic
	err := m.QueryRowCtx(ctx, &resp, betxinTopicIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", topicRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTopicModel) FindOneByTid(ctx context.Context, tid string) (*Topic, error) {
	betxinTopicTidKey := fmt.Sprintf("%s%v", cacheBetxinTopicTidPrefix, tid)
	var resp Topic
	err := m.QueryRowIndexCtx(ctx, &resp, betxinTopicTidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `tid` = ? limit 1", topicRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, tid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTopicModel) FindOneByTitleIntroContent(ctx context.Context, title string, intro string, content string) (*Topic, error) {
	betxinTopicTitleIntroContentKey := fmt.Sprintf("%s%v:%v:%v", cacheBetxinTopicTitleIntroContentPrefix, title, intro, content)
	var resp Topic
	err := m.QueryRowIndexCtx(ctx, &resp, betxinTopicTitleIntroContentKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `title` = ? and `intro` = ? and `content` = ? limit 1", topicRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, title, intro, content); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTopicModel) Insert(ctx context.Context, data *Topic) (sql.Result, error) {
	betxinTopicIdKey := fmt.Sprintf("%s%v", cacheBetxinTopicIdPrefix, data.Id)
	betxinTopicTidKey := fmt.Sprintf("%s%v", cacheBetxinTopicTidPrefix, data.Tid)
	betxinTopicTitleIntroContentKey := fmt.Sprintf("%s%v:%v:%v", cacheBetxinTopicTitleIntroContentPrefix, data.Title, data.Intro, data.Content)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, topicRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Tid, data.Cid, data.Title, data.Intro, data.Content, data.YesRatio, data.NoRatio, data.YesPrice, data.NoPrice, data.TotalPrice, data.CollectCount, data.ReadCount, data.ImgUrl, data.IsStop, data.RefundEndTime, data.DeletedAt)
	}, betxinTopicIdKey, betxinTopicTidKey, betxinTopicTitleIntroContentKey)
	return ret, err
}

func (m *defaultTopicModel) Update(ctx context.Context, newData *Topic) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	betxinTopicIdKey := fmt.Sprintf("%s%v", cacheBetxinTopicIdPrefix, data.Id)
	betxinTopicTidKey := fmt.Sprintf("%s%v", cacheBetxinTopicTidPrefix, data.Tid)
	betxinTopicTitleIntroContentKey := fmt.Sprintf("%s%v:%v:%v", cacheBetxinTopicTitleIntroContentPrefix, data.Title, data.Intro, data.Content)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, topicRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Tid, newData.Cid, newData.Title, newData.Intro, newData.Content, newData.YesRatio, newData.NoRatio, newData.YesPrice, newData.NoPrice, newData.TotalPrice, newData.CollectCount, newData.ReadCount, newData.ImgUrl, newData.IsStop, newData.RefundEndTime, newData.DeletedAt, newData.Id)
	}, betxinTopicIdKey, betxinTopicTidKey, betxinTopicTitleIntroContentKey)
	return err
}

func (m *defaultTopicModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheBetxinTopicIdPrefix, primary)
}

func (m *defaultTopicModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", topicRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTopicModel) tableName() string {
	return m.table
}