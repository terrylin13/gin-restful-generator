package model

// 定义 GORM模型 GraphQL模型
type Score struct {
	ScoreID       string  `gorm:"column:score_id;type:varchar(255);primaryKey;comment:考核主表主键id;" json:"score_id"`
	Title         string  `gorm:"column:title;type:varchar(255);comment:自动生成：区块名称-YYMMDD-[安全/质量]考核;not null" json:"title"`
	PosType       uint8   `gorm:"column:pos_type;type:tinyint;comment:考核区块类型;not null" json:"pos_type"`
	PosID         string  `gorm:"column:pos_id;type:varchar(255);comment:区块ID;not null" json:"pos_id"`
	PosName       string  `gorm:"column:pos_name;type:varchar(255);comment:区块名称;not null" json:"pos_name"`
	DeptID        string  `gorm:"column:dept_id;type:varchar(255);comment:部门/班组;not null" json:"dept_id"`
	DeptName      string  `gorm:"column:dept_name;type:varchar(255);comment:部门/班组 名称;not null" json:"dept_name"`
	GroupDeptID   string  `gorm:"column:group_dept_id;type:varchar(255);comment:区块id;not null" json:"group_dept_id"`
	GroupDeptName string  `gorm:"column:group_dept_name;type:varchar(255);comment:区块名称;not null" json:"group_dept_name"`
	CaseType      uint8   `gorm:"column:case_type;type:tinyint;comment:考核案件类型;not null" json:"case_type"`
	StartAT       int64   `gorm:"column:start_at;type:bigint;comment:开始时间;not null" json:"start_at"`
	EndAT         int64   `gorm:"column:end_at;type:bigint;comment:要求完成时间;not null" json:"end_at"`
	Score         float64 `gorm:"column:score;type:DECIMAL(10,2);comment:初始100,扣除子项分(子项状态为91的不扣分),加上整改流程flow_score加减分;not null" json:"score"`
	FlowScore     float64 `gorm:"column:flow_score;type:double;comment:整改流程评价加减分;"  json:"flow_score"`
	FlowVStar     int     `gorm:"column:flow_v_star;type:bigint;comment:整改流程效率星级评价;" json:"flow_v_star"`
	FlowQStar     int     `gorm:"column:flow_q_star;type:bigint;comment:整改流程质量星级评价;" json:"flow_q_star"`
	StarScore     int     `gorm:"column:star_score;type:bigint;comment:星级总分;" json:"star_score"`
	CurActorIDs   string  `gorm:"column:cur_actor_ids;type:varchar(255);comment:考核人员ID;" json:"cur_actor_ids"`
	CurActors     string  `gorm:"column:cur_actors;type:varchar(255);comment:考核人员;" json:"cur_actors"`
	Status        uint8   `gorm:"column:status;type:tinyint;comment:考核状态" json:"status"`
	CreatedBy     string  `gorm:"column:created_by;type:varchar(255);comment:填写人ID;not null" json:"created_by"`
	CreatedByName string  `gorm:"column:created_by_name;type:varchar(255);comment:填写人名字;not null" json:"created_by_name"`
	CreatedAt     int64   `gorm:"autoCreateTime:milli;comment:创建时间;not null;" json:"created_at,omitempty"`
	UpdatedAt     int64   `gorm:"autoUpdateTime:milli;comment:修改时间;not null;" json:"updated_at,omitempty"`
	// ScoreItems    []ScoreItem `gorm:"foreignKey:ScoreID"`
}

func (s *Score) TableName() string {
	return "sq_score"
}
