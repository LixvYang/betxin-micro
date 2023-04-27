package mixinsnapshot

import (
	"time"

	"github.com/shopspring/decimal"
)

type MixinSnapshot struct {
	Id             int             `db:"Id"`
	SnapshotId     string          `gorm:"type:varchar(36)" json:"snapshot_id"`
	TraceId        string          `gorm:"type:varchar(36);not null;" json:"trace_id"`
	AssetId        string          `gorm:"type:varchar(36);index" json:"asset_id"`
	OpponentID     string          `gorm:"type:varchar(36)" json:"opponent_id"`
	Amount         decimal.Decimal `gorm:"type:decimal(32, 8)" json:"amount"`
	Memo           string          `gorm:"type:varchar(255)" json:"memo"`
	Type           string          `gorm:"type:varchar(255)" json:"type"`
	OpeningBalance decimal.Decimal `gorm:"type:decimal(32, 8)" json:"opening_balance"`
	ClosingBalance decimal.Decimal `gorm:"type:decimal(32, 8)" json:"closing_balance"`

	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}
/**

// {
	SnapshotID:6b6a80a2-f69e-4b08-8379-8feac5f6d73a 
	CreatedAt:2022-12-17 12:36:28.595421 +0000 UTC 
	TraceID:8bdeaa3a-a9a7-4df2-a7e4-d9fc29d52763 
	UserID: 
	AssetID:31d2ea9c-95eb-3355-b65b-ba096853bc18 
	ChainID: 
	OpponentID:6a87e67f-02fb-47cf-b31f-32a13dd5b3d9 
	Source: 
	Amount:0.1 
	OpeningBalance:1.9 
	ClosingBalance:2 
	Memo:eyJ0aWQiOiAiZTVkYjYyMGUtMmVkMC00YjVjLTg4M2YtZTkxNjFmM2UxOGQwIiwgInllc19yYXRpbyI6IGZhbHNlLCAibm9fcmF0aW8iOiB0cnVlfQ== 
	Type:transfer 
	Sender: 
	Receiver: 
	TransactionHash: 
	SnapshotHash: 
	SnapshotAt:<nil>
	Asset:<nil>}
*/
