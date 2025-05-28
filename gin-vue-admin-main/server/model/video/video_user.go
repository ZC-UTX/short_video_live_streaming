
// 自动生成模板VideoUser
package video
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户管理 结构体  VideoUser
type VideoUser struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:名称;column:name;size:20;" binding:"required"`  //名称
  NickName  *string `json:"nickName" form:"nickName" gorm:"comment:昵称;column:nick_name;size:20;" binding:"required"`  //昵称
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 用户管理 VideoUser自定义表名 video_user
func (VideoUser) TableName() string {
    return "video_user"
}





