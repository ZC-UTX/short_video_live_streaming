
// 自动生成模板VideoWorks
package video
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 作品管理 结构体  VideoWorks
type VideoWorks struct {
    global.GVA_MODEL
  Title  *string `json:"title" form:"title" gorm:"comment:标题;column:title;size:100;"`  //标题
  Desc  *string `json:"desc" form:"desc" gorm:"comment:描述;column:desc;size:255;"`  //描述
  MusicId  *int `json:"musicId" form:"musicId" gorm:"comment:选择音乐;column:music_id;"`  //选择音乐
  WorkType  *string `json:"workType" form:"workType" gorm:"default:1;comment:作品类型;column:work_type;size:20;" binding:"required"`  //作品类型
  CheckStatus  *string `json:"checkStatus" form:"checkStatus" gorm:"default:1;comment:审核状态;column:check_status;size:20;" binding:"required"`  //审核状态
  CheckUser  *int `json:"checkUser" form:"checkUser" gorm:"comment:审核人;column:check_user;"`  //审核人
  IpAddress  *string `json:"ipAddress" form:"ipAddress" gorm:"comment:IP地址;column:ip_address;size:20;"`  //IP地址
  WorkPermission  *string `json:"workPermission" form:"workPermission" gorm:"default:1;comment:作品权限;column:work_permission;size:20;" binding:"required"`  //作品权限
  LikeCount  *int `json:"likeCount" form:"likeCount" gorm:"comment:喜欢数量;column:like_count;"`  //喜欢数量
  CommentCount  *int `json:"commentCount" form:"commentCount" gorm:"comment:评论数;column:comment_count;"`  //评论数
  ShareCount  *int `json:"shareCount" form:"shareCount" gorm:"comment:分享数;column:share_count;"`  //分享数
  CollectCount  *int `json:"collectCount" form:"collectCount" gorm:"comment:收藏数;column:collect_count;"`  //收藏数
  BrowseCount  *int `json:"browseCount" form:"browseCount" gorm:"comment:浏览量;column:browse_count;"`  //浏览量
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 作品管理 VideoWorks自定义表名 video_works
func (VideoWorks) TableName() string {
    return "video_works"
}





