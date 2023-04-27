package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MixinsnapshotModel = (*customMixinsnapshotModel)(nil)

type (
	// MixinsnapshotModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMixinsnapshotModel.
	MixinsnapshotModel interface {
		mixinsnapshotModel
	}

	customMixinsnapshotModel struct {
		*defaultMixinsnapshotModel
	}
)

// NewMixinsnapshotModel returns a model for the database table.
func NewMixinsnapshotModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MixinsnapshotModel {
	return &customMixinsnapshotModel{
		defaultMixinsnapshotModel: newMixinsnapshotModel(conn, c, opts...),
	}
}
