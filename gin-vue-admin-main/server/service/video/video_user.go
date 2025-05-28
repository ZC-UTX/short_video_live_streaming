
package video

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
    videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
    "gorm.io/gorm"
)

type VideoUserService struct {}
// CreateVideoUser 创建用户管理记录
// Author [yourname](https://github.com/yourname)
func (videoUserService *VideoUserService) CreateVideoUser(ctx context.Context, videoUser *video.VideoUser) (err error) {
	err = global.GVA_DB.Create(videoUser).Error
	return err
}

// DeleteVideoUser 删除用户管理记录
// Author [yourname](https://github.com/yourname)
func (videoUserService *VideoUserService)DeleteVideoUser(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.VideoUser{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&video.VideoUser{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteVideoUserByIds 批量删除用户管理记录
// Author [yourname](https://github.com/yourname)
func (videoUserService *VideoUserService)DeleteVideoUserByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.VideoUser{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&video.VideoUser{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateVideoUser 更新用户管理记录
// Author [yourname](https://github.com/yourname)
func (videoUserService *VideoUserService)UpdateVideoUser(ctx context.Context, videoUser video.VideoUser) (err error) {
	err = global.GVA_DB.Model(&video.VideoUser{}).Where("id = ?",videoUser.ID).Updates(&videoUser).Error
	return err
}

// GetVideoUser 根据ID获取用户管理记录
// Author [yourname](https://github.com/yourname)
func (videoUserService *VideoUserService)GetVideoUser(ctx context.Context, ID string) (videoUser video.VideoUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&videoUser).Error
	return
}
// GetVideoUserInfoList 分页获取用户管理记录
// Author [yourname](https://github.com/yourname)
func (videoUserService *VideoUserService)GetVideoUserInfoList(ctx context.Context, info videoReq.VideoUserSearch) (list []video.VideoUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&video.VideoUser{})
    var videoUsers []video.VideoUser
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.NickName != nil && *info.NickName != "" {
        db = db.Where("nick_name LIKE ?", "%"+ *info.NickName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&videoUsers).Error
	return  videoUsers, total, err
}
func (videoUserService *VideoUserService)GetVideoUserPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
