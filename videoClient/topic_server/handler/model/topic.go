package model

import (
	"gorm.io/gorm"
	"time"
	"topic_server/basic/config"
)

type VideoTopic struct {
	Id         int32          `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Title      string         `gorm:"column:title;type:varchar(50);comment:话题标题;default:NULL;" json:"title"`            // 话题标题
	CreateUser int16          `gorm:"column:create_user;type:smallint;comment:话题发起人;default:NULL;" json:"create_user"` // 话题发起人
	HotScore   float64        `gorm:"column:hot_score;type:double;comment:实时热度分;default:NULL;" json:"hot_score"`       // 实时热度分
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	CreatedBy  uint64         `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy  uint64         `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;default:NULL;" json:"updated_by"` // 更新者
	DeletedBy  uint64         `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}

// 定义数据库名称
func (v *VideoTopic) TableName() string {
	return "video_topic"
}

type VideoDataTopic struct {
	Id      int32 `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	VideoId int32 `gorm:"column:video_id;type:int;not null;" json:"video_id"`
	TopicId int32 `gorm:"column:topic_id;type:int;not null;" json:"topic_id"`
}

// 定义数据库中间表名称
func (v *VideoDataTopic) TableName() string {
	return "video_data_topic"
}

// 计算每个作品的使用量
func (v *VideoDataTopic) SumVideoDateTopic(id int32) (video []VideoDataTopic, err error) {
	err = config.DB.Where("topic_id = ?", id).Find(&video).Error
	return
}

// 话题列表展示
func (v *VideoTopic) ListVideo() (video []VideoTopic, err error) {
	err = config.DB.Find(&video).Error
	return
}

// 添加话题
func (v *VideoTopic) CreateTopic() (err error) {
	err = config.DB.Create(v).Error
	return
}
