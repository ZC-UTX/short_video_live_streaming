
package video

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
    videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
    "gorm.io/gorm"
)

type VideoTopicService struct {}
// CreateVideoTopic 创建videoTopic表记录
// Author [yourname](https://github.com/yourname)
func (videoTopicService *VideoTopicService) CreateVideoTopic(ctx context.Context, videoTopic *video.VideoTopic) (err error) {
	err = global.GVA_DB.Create(videoTopic).Error
	return err
}

// DeleteVideoTopic 删除videoTopic表记录
// Author [yourname](https://github.com/yourname)
func (videoTopicService *VideoTopicService)DeleteVideoTopic(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.VideoTopic{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&video.VideoTopic{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteVideoTopicByIds 批量删除videoTopic表记录
// Author [yourname](https://github.com/yourname)
func (videoTopicService *VideoTopicService)DeleteVideoTopicByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.VideoTopic{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&video.VideoTopic{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateVideoTopic 更新videoTopic表记录
// Author [yourname](https://github.com/yourname)
func (videoTopicService *VideoTopicService)UpdateVideoTopic(ctx context.Context, videoTopic video.VideoTopic) (err error) {
	err = global.GVA_DB.Model(&video.VideoTopic{}).Where("id = ?",videoTopic.ID).Updates(&videoTopic).Error
	return err
}

// GetVideoTopic 根据ID获取videoTopic表记录
// Author [yourname](https://github.com/yourname)
func (videoTopicService *VideoTopicService)GetVideoTopic(ctx context.Context, ID string) (videoTopic video.VideoTopic, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&videoTopic).Error
	return
}
// GetVideoTopicInfoList 分页获取videoTopic表记录
// Author [yourname](https://github.com/yourname)
func (videoTopicService *VideoTopicService)GetVideoTopicInfoList(ctx context.Context, info videoReq.VideoTopicSearch) (list []video.VideoTopic, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&video.VideoTopic{})
    var videoTopics []video.VideoTopic
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Title != nil && *info.Title != "" {
        db = db.Where("title LIKE ?", "%"+ *info.Title+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&videoTopics).Error
	return  videoTopics, total, err
}
func (videoTopicService *VideoTopicService)GetVideoTopicPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
