package topic

type Category struct {
	Id           uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryName string `gorm:"type:varchar(20);not null;" json:"category_name"`
}

type Topic struct {
	Id           int64    `json:"id"`
	Tid          string   `gorm:"type:varchar(36);index;" json:"tid"`
	Cid          uint64   `gorm:"type:int;not null" json:"cid"`
	Category     Category `gorm:"foreignKey:Cid" json:"category"`
	Title        string   `gorm:"type:varchar(50);not null;index:title_intro_topic_index" json:"title"`
	Intro        string   `gorm:"type:varchar(255);not null;index:title_intro_topic_index" json:"intro"`
	Content      string   `gorm:"type:varchar(255);default null;index:title_content_topic_index" json:"content"`
	YesRatio     float64  `gorm:"type:decimal(5,2);" json:"yes_ratio"`
	NoRatio      float64  `gorm:"type:decimal(5,2);" json:"no_ratio"`
	YesPrice     float64  `gorm:"type:decimal(16,8);default 0" json:"yes_ratio_price"`
	NoPrice      float64  `gorm:"type:decimal(16,8);default 0" json:"no_ratio_price"`
	TotalPrice   float64  `gorm:"type:decimal(32,8);default 0;" json:"total_price"`
	CollectCount uint64   `json:"collect_count"`
	ReadCount    uint64   `json:"read_count"`
	ImgUrl       string   `json:"img_url"`
	IsStop       bool     `json:"is_stop"`
	EndTime      int64    `json:"end_time"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}
