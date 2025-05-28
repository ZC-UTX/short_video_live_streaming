
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type VideoUserSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      NickName  *string `json:"nickName" form:"nickName"` 
    request.PageInfo
}
