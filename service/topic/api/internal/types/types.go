// Code generated by goctl. DO NOT EDIT.
package types

type Category struct {
	Id           uint64 `json:"id"`
	CategoryName string `json:"category_name"`
}

type Topic struct {
	Id           int64     `json:"id"`
	Tid          string    `json:"tid"`
	Cid          uint64    `json:"cid"`
	Category     *Category `json:"category"`
	Title        string    `json:"title"`
	Intro        string    `json:"intro"`
	Content      string    `json:"content"`
	YesRatio     float64   `json:"yes_ratio"`
	NoRatio      float64   `json:"no_ratio"`
	YesPrice     float64   `json:"yes_ratio_price"`
	NoPrice      float64   `json:"no_ratio_price"`
	TotalPrice   float64   `json:"total_price"`
	CollectCount uint64    `json:"collect_count"`
	ReadCount    uint64    `json:"read_count"`
	ImgUrl       string    `json:"img_url"`
	IsStop       int64     `json:"is_stop"`
	EndTime      int64     `json:"end_time"`
	CreatedAt    int64     `json:"created_at"`
	UpdatedAt    int64     `json:"updated_at"`
	DeletedAt    int64     `json:"deleted_at"`
}

type CreateTopicReq struct {
	Cid     uint64 `json:"cid"`
	Tid     string `json:"tid"`
	Title   string `json:"title"`
	Intro   string `json:"intro"`
	Content string `json:"content"`
	ImgUrl  string `json:"img_url"`
	EndTime int64  `json:"end_time"`
}

type CreateTopicResp struct {
	Errmsg
	Data Topic `json:"data"`
}

type DeleteTopicReq struct {
	Tid string `json:"tid"`
}

type DeleteTopicResp struct {
	Errmsg
	Data *string `json:"data"`
}

type ListTopicResp struct {
	Errmsg
	Data []Topic `json:"data"`
}

type Errmsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
