
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type VideoWorksSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Title  *string `json:"title" form:"title"` 
      Desc  *string `json:"desc" form:"desc"` 
      MusicId  *int `json:"musicId" form:"musicId"` 
      WorkType  *string `json:"workType" form:"workType"` 
      CheckStatus  *string `json:"checkStatus" form:"checkStatus"` 
      CheckUser  *int `json:"checkUser" form:"checkUser"` 
      IpAddress  *string `json:"ipAddress" form:"ipAddress"` 
      WorkPermission  *string `json:"workPermission" form:"workPermission"` 
    request.PageInfo
}
