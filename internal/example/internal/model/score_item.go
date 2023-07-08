package model

type ScoreItem struct {
	ScoreItemID  string  `gorm:"column:score_item_id;type:varchar(255);primaryKey;comment:考核附表主键id;" json:"score_item_id"`
	ScoreID      string  `gorm:"foreignKey:score_id;column:score_id;type:varchar(255);comment:考核主表id;not null" json:"score_id"`
	Lat          float64 `gorm:"column:lat;type:double;comment:纬度;not null" json:"lat"`
	Lng          float64 `gorm:"column:lng;type:double;comment:经度;not null" json:"lng"`
	Type         string  `gorm:"column:type;type:varchar(255);comment:问题类型（从配置表获取，存中文）;not null" json:"type"`
	Score        float64 `gorm:"column:score;type:double;comment:扣分为正数，最小 0.1;not null" json:"score"`
	Memo         string  `gorm:"column:memo;type:text;comment:案件描述;not null" json:"memo"`
	Images       string  `gorm:"column:images;type:text;comment:图片" json:"images"`
	FyStatus     uint8   `gorm:"column:fy_status;type:tinyint;comment:复议状态" json:"fy_status"`
	Severity     uint8   `gorm:"column:severity;type:tinyint;comment:问题严重性" json:"severity"`
	Result       string  `gorm:"column:result;type:text;comment:整改结果;not null" json:"result"`
	ResultImages string  `gorm:"column:result_images;type:longtext;comment:整改结果图片;not null" json:"result_images"`
	CreatedAt    int64   `gorm:"autoCreateTime:milli;comment:创建时间;not null;" json:"created_at,omitempty"`
	UpdatedAt    int64   `gorm:"autoUpdateTime:milli;comment:修改时间;not null;" json:"updated_at,omitempty"`
}

func (s *ScoreItem) TableName() string {
	return "sq_score_item"
}
