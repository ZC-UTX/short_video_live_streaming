
// 自动生成模板VideoTopic
package video
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// videoTopic表 结构体  VideoTopic
type VideoTopic struct {
    global.GVA_MODEL
  Title  *string `json:"title" form:"title" gorm:"comment:话题标题;column:title;size:50;"`  //话题标题
  CreateUser  *int `json:"createUser" form:"createUser" gorm:"comment:话题发起人;column:create_user;size:10;"`  //话题发起人
  HotScore  *float64 `json:"hotScore" form:"hotScore" gorm:"comment:实时热度分;column:hot_score;"`  //实时热度分
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName videoTopic表 VideoTopic自定义表名 video_topic
func (VideoTopic) TableName() string {
    return "video_topic"
}





