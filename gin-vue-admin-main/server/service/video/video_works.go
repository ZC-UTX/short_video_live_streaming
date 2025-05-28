
package video

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
    videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
    "gorm.io/gorm"
)

type VideoWorksService struct {}
// CreateVideoWorks 创建作品管理记录
// Author [yourname](https://github.com/yourname)
func (videoWorksService *VideoWorksService) CreateVideoWorks(ctx context.Context, videoWorks *video.VideoWorks) (err error) {
	err = global.GVA_DB.Create(videoWorks).Error
	return err
}

// DeleteVideoWorks 删除作品管理记录
// Author [yourname](https://github.com/yourname)
func (videoWorksService *VideoWorksService)DeleteVideoWorks(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.VideoWorks{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&video.VideoWorks{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteVideoWorksByIds 批量删除作品管理记录
// Author [yourname](https://github.com/yourname)
func (videoWorksService *VideoWorksService)DeleteVideoWorksByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.VideoWorks{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&video.VideoWorks{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateVideoWorks 更新作品管理记录
// Author [yourname](https://github.com/yourname)
func (videoWorksService *VideoWorksService)UpdateVideoWorks(ctx context.Context, videoWorks video.VideoWorks) (err error) {
	err = global.GVA_DB.Model(&video.VideoWorks{}).Where("id = ?",videoWorks.ID).Updates(&videoWorks).Error
	return err
}

// GetVideoWorks 根据ID获取作品管理记录
// Author [yourname](https://github.com/yourname)
func (videoWorksService *VideoWorksService)GetVideoWorks(ctx context.Context, ID string) (videoWorks video.VideoWorks, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&videoWorks).Error
	return
}
// GetVideoWorksInfoList 分页获取作品管理记录
// Author [yourname](https://github.com/yourname)
func (videoWorksService *VideoWorksService)GetVideoWorksInfoList(ctx context.Context, info videoReq.VideoWorksSearch) (list []video.VideoWorks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&video.VideoWorks{})
    var videoWorkss []video.VideoWorks
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Title != nil && *info.Title != "" {
        db = db.Where("title LIKE ?", "%"+ *info.Title+"%")
    }
    if info.Desc != nil && *info.Desc != "" {
        db = db.Where("desc LIKE ?", "%"+ *info.Desc+"%")
    }
    if info.MusicId != nil {
        db = db.Where("music_id = ?", *info.MusicId)
    }
    if info.WorkType != nil && *info.WorkType != "" {
        db = db.Where("work_type = ?", *info.WorkType)
    }
    if info.CheckStatus != nil && *info.CheckStatus != "" {
        db = db.Where("check_status = ?", *info.CheckStatus)
    }
    if info.CheckUser != nil {
        db = db.Where("check_user = ?", *info.CheckUser)
    }
    if info.IpAddress != nil && *info.IpAddress != "" {
        db = db.Where("ip_address = ?", *info.IpAddress)
    }
    if info.WorkPermission != nil && *info.WorkPermission != "" {
        db = db.Where("work_permission = ?", *info.WorkPermission)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&videoWorkss).Error
	return  videoWorkss, total, err
}
func (videoWorksService *VideoWorksService)GetVideoWorksPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
