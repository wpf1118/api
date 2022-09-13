package goods

import (
	"github.com/wpf1118/api/pkg/service"
)

type Goods struct {
	ID   uint   `json:"id"`
	Name string `json:"name" gorm:"size:20;comment:分类名称"`
}

type Req struct {
	service.Pagination `search:"-"`
	Name               string `json:"name" search:"type:contains;column:name;table:goods" comment:"名称"`

	//UserId             int    `form:"userId" search:"type:exact;column:user_id;table:sys_user" comment:"用户ID"`
	//Username           string `form:"username" search:"type:contains;column:username;table:sys_user" comment:"用户名"`
	//NickName           string `form:"nickName" search:"type:contains;column:nick_name;table:sys_user" comment:"昵称"`
	//Phone              string `form:"phone" search:"type:contains;column:phone;table:sys_user" comment:"手机号"`
	//RoleId             string `form:"roleId" search:"type:exact;column:role_id;table:sys_user" comment:"角色ID"`
	//Sex                string `form:"sex" search:"type:exact;column:sex;table:sys_user" comment:"性别"`
	//Email              string `form:"email" search:"type:contains;column:email;table:sys_user" comment:"邮箱"`
	//PostId             string `form:"postId" search:"type:exact;column:post_id;table:sys_user" comment:"岗位"`
	//Status             string `form:"status" search:"type:exact;column:status;table:sys_user" comment:"状态"`
}
