// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type AddMenuReq struct {
	MenuType        string  `json:"menuType"`            // 菜单类型
	Title           string  `json:"title"`               // 菜单标题
	Sort            int32   `json:"sort"`                // 显示排序
	ParentMenuId    int64   `json:"parentMenuId"`        // 上级菜单
	Icon            string  `json:"icon,optional"`       // 图标
	Name            string  `json:"name,optional"`       // 路由名称
	Component       string  `json:"component,optional"`  // 组件路径
	Path            string  `json:"path,optional"`       // 路由地址
	Permission      string  `json:"permission,optional"` // 权限标识
	HideInMenu      bool    `json:"hideInMenu,optional"`
	IgnoreKeepAlive bool    `json:"ignoreKeepAlive,optional"`
	LinkFlag        bool    `json:"linkFlag,optional"`
	Link            string  `json:"link,optional"` // 组件路径
	Disabled        bool    `json:"disabled,optional"`
	SelectApi       []int64 `json:"selectApi,optional"`
}

type AddMenuResp struct {
}

type ApiListData struct {
	Id         int64  `json:"id"`
	Handle     string `json:"handle"`     // 方法名
	Title      string `json:"title"`      // 方法描述
	Path       string `json:"path"`       // 请求路径
	Type       string `json:"type"`       // 类型 1 系统 2 业务
	Action     string `json:"action"`     // 请求方式
	CreateTime string `json:"createTime"` // 创建时间
}

type ApiListReq struct {
	Type     string `form:"type,optional"`       // 系统 业务
	Action   string `form:"action,optional"`     // 请求方法 GET POST DELETE PUT
	PageNum  int64  `form:"pageNum,default=1"`   //  第几页
	PageSize int64  `form:"pageSize,default=10"` // 每页的数量
}

type ApiListResp struct {
	Total int64          `json:"total"`
	Data  []*ApiListData `json:"data"`
}

type DeleteMenuReq struct {
	MenuId int64 `path:"menuId"`
}

type DeleteMenuResp struct {
}

type DeptAddReq struct {
	ParentDeptId int64  `json:"parentDeptId"`
	DeptName     string `json:"deptName"`
	Sort         int32  `json:"sort"`
	Leader       string `json:"leader,optional"`
	Phone        string `json:"phone,optional"`
	Email        string `json:"email,optional"`
	Status       int32  `json:"status"`
}

type DeptAddResp struct {
}

type DeptDeleteReq struct {
	DeptId int64 `path:"deptId"`
}

type DeptDeleteResp struct {
}

type DeptInfoData struct {
	DeptId       int64  `json:"deptId"`
	ParentDeptId int64  `json:"parentDeptId"`
	DeptName     string `json:"deptName"`
	Sort         int32  `json:"sort"`
	Leader       string `json:"leader"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Status       int32  `json:"status"`
	CreateTime   string `json:"createTime"`
}

type DeptInfoReq struct {
	DeptId int64 `path:"deptId"`
}

type DeptListReq struct {
}

type DeptListResp struct {
	Data []*DeptInfoData `json:"data"`
}

type DeptTreeData struct {
	DeptId       int64           `json:"deptId"`
	Value        int64           `json:"value"` // 还是deptId
	ParentDeptId int64           `json:"parentDeptId"`
	DeptName     string          `json:"deptName"`
	Sort         int32           `json:"sort"`
	Leader       string          `json:"leader"`
	Phone        string          `json:"phone"`
	Email        string          `json:"email"`
	Status       int32           `json:"status"`
	CreateTime   string          `json:"createTime"`
	Children     []*DeptTreeData `json:"children"`
}

type DeptTreeReq struct {
}

type DeptTreeResp struct {
	Data []*DeptTreeData `json:"data"`
}

type DeptUpdateReq struct {
	DeptId       int64  `json:"deptId"`
	ParentDeptId int64  `json:"parentDeptId"`
	DeptName     string `json:"deptName"`
	Sort         int32  `json:"sort"`
	Leader       string `json:"leader,optional"`
	Phone        string `json:"phone,optional"`
	Email        string `json:"email,optional"`
	Status       int32  `json:"status"`
}

type DeptUpdateResp struct {
}

type InitApiReq struct {
}

type InitApiResp struct {
}

type ListMenuData struct {
	MenuId          int64  `json:"menuId"`
	MenuType        string `json:"menuType"`     // 菜单类型
	Title           string `json:"title"`        // 菜单标题
	Sort            int32  `json:"sort"`         // 显示排序
	ParentMenuId    int64  `json:"parentMenuId"` // 上级菜单
	Icon            string `json:"icon"`         // 图标
	Name            string `json:"name"`         // 路由名称
	Component       string `json:"component"`    // 组件路径
	Path            string `json:"path"`         // 路由地址
	Permission      string `json:"permission"`   // 权限标识
	HideInMenu      bool   `json:"hideInMenu"`
	IgnoreKeepAlive bool   `json:"ignoreKeepAlive"`
	LinkFlag        bool   `json:"linkFlag"`
	Link            string `json:"link"` // 组件路径
	Disabled        bool   `json:"disabled"`
}

type ListMenuReq struct {
}

type ListMenuResp struct {
	Data []*ListMenuData `json:"data"`
}

type LoginReq struct {
	Username string `json:"username"` //用户名
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken string `json:"accessToken"`
}

type Menu struct {
	MenuId       int64     `json:"menuId"`
	ParentMenuId int64     `json:"parentMenuId"`
	Name         string    `json:"name"`      // 路由名称
	Path         string    `json:"path"`      // 路由地址
	Component    string    `json:"component"` // 组件地址
	MenuMate     *MenuMate `json:"meta"`
	Children     []*Menu   `json:"children,omitempty"`
}

type MenuInfoData struct {
	MenuId          int64   `json:"menuId"`
	MenuType        string  `json:"menuType"`     // 菜单类型
	Title           string  `json:"title"`        // 菜单标题
	Sort            int32   `json:"sort"`         // 显示排序
	ParentMenuId    int64   `json:"parentMenuId"` // 上级菜单
	Icon            string  `json:"icon"`         // 图标
	Name            string  `json:"name"`         // 路由名称
	Component       string  `json:"component"`    // 组件路径
	Path            string  `json:"path"`         // 路由地址
	Permission      string  `json:"permission"`   // 权限标识
	HideInMenu      bool    `json:"hideInMenu"`
	IgnoreKeepAlive bool    `json:"ignoreKeepAlive"`
	LinkFlag        bool    `json:"linkFlag"`
	Link            string  `json:"link"` // 组件路径
	Disabled        bool    `json:"disabled"`
	SelectApi       []int64 `json:"selectApi"`
}

type MenuInfoReq struct {
	MenuId int64 `path:"menuId"`
}

type MenuMate struct {
	Title                    string `json:"title"`                    // 菜单名称
	Order                    int64  `json:"order"`                    // 排序
	Icon                     string `json:"icon"`                     // 图标
	HideInMenu               bool   `json:"hideInMenu"`               // 是否隐藏
	KeepAlive                bool   `json:"keepAlive"`                // 是否缓存
	MenuVisibleWithForbidden bool   `json:"menuVisibleWithForbidden"` // 是否启用
	Link                     string `json:"link"`                     // 用于配置外链跳转路径，会在新窗口打开。
}

type RoleAddReq struct {
	RoleName      string  `json:"roleName"`      // 角色名称
	RoleKey       string  `json:"roleKey"`       //  权限字符
	Status        int32   `json:"status"`        // 状态 1 停用 2 启用
	Sort          int32   `json:"sort"`          // 排序
	SelectMenus   []int64 `json:"selectMenus"`   // 菜单权限
	DefaultRouter string  `json:"defaultRouter"` // 默认路由
}

type RoleAddResp struct {
}

type RoleDeleteReq struct {
	RoleId int64 `path:"roleId"`
}

type RoleDeleteResp struct {
	RoleKey string `form:"roleKey,optional"` //  权限字符
}

type RoleInfoData struct {
	RoleId        int64   `json:"roleId"`
	RoleName      string  `json:"roleName"`             // 角色名称
	RoleKey       string  `json:"roleKey"`              //  权限字符
	Status        int32   `json:"status"`               // 状态 1 停用 2 启用
	Sort          int32   `json:"sort"`                 // 排序
	DefaultRouter string  `json:"defaultRouter"`        // 默认路由
	SelectMenus   []int64 `json:"selectMenus,optional"` // 菜单权限
	Admin         bool    `json:"admin"`
}

type RoleInfoReq struct {
	RoleId int64 `path:"roleId"`
}

type RoleListData struct {
	RoleId        int64  `json:"roleId"`
	RoleName      string `json:"roleName"`      // 角色名称
	RoleKey       string `json:"roleKey"`       //  权限字符
	Status        int32  `json:"status"`        // 状态 1 停用 2 启用
	Sort          int32  `json:"sort"`          // 排序
	DefaultRouter string `json:"defaultRouter"` // 默认路由
	CreateTime    string `json:"createTime"`    // 创建时间
	Admin         bool   `json:"admin"`
}

type RoleListReq struct {
	RoleName string `form:"roleName,optional"`   // 角色名称
	RoleKey  string `form:"roleKey,optional"`    //  权限字符
	Status   int32  `form:"status,optional"`     // 状态 1 停用 2 启用
	PageNum  int64  `form:"pageNum,default=1"`   //  第几页
	PageSize int64  `form:"pageSize,default=10"` // 每页的数量
}

type RoleListResp struct {
	Total int64           `json:"total"`
	Data  []*RoleListData `json:"data"`
}

type RoleUpdateReq struct {
	RoleId        int64   `json:"roleId"`
	RoleName      string  `json:"roleName"`             // 角色名称
	Status        int32   `json:"status"`               // 状态 1 停用 2 启用
	Sort          int32   `json:"sort"`                 // 排序
	SelectMenus   []int64 `json:"selectMenus,optional"` // 菜单权限
	DefaultRouter string  `json:"defaultRouter"`        // 默认路由
}

type RoleUpdateResp struct {
}

type TreeMenuData struct {
	MenuId       int64           `json:"menuId"`       // 对应menuId
	MenuType     string          `json:"menuType"`     // 菜单类型
	ParentMenuId int64           `json:"parentMenuId"` // 上级菜单
	Name         string          `json:"name"`         // 路由名称
	Component    string          `json:"component"`    // 组件路径
	Path         string          `json:"path"`         // 路由地址
	Permission   string          `json:"permission"`   // 权限标识
	LinkFlag     bool            `json:"linkFlag"`
	Meta         *TreeMenuMeta   `json:"meta"`
	Children     []*TreeMenuData `json:"children"`
	Title        string          `json:"title"` // 菜单标题
}

type TreeMenuMeta struct {
	Title                    string `json:"title"`     // 菜单标题
	Icon                     string `json:"icon"`      // 图标
	KeepAlive                bool   `json:"keepAlive"` // 对应ignoreKeepAlive
	HideInMenu               bool   `json:"hideInMenu"`
	Link                     string `json:"link"`
	MenuVisibleWithForbidden bool   `json:"menuVisibleWithForbidden"` // 对应 disabled
	Order                    int32  `json:"order"`                    // sort
}

type TreeMenuReq struct {
	NeedButton bool `form:"needButton"` // 是否需要按钮
}

type TreeMenuResp struct {
	Data []*TreeMenuData `json:"data"`
}

type UpdateMenuReq struct {
	MenuId          int64   `path:"menuId"`
	MenuType        string  `json:"menuType"`            // 菜单类型
	Title           string  `json:"title"`               // 菜单标题
	Sort            int32   `json:"sort"`                // 显示排序
	ParentMenuId    int64   `json:"parentMenuId"`        // 上级菜单
	Icon            string  `json:"icon,optional"`       // 图标
	Name            string  `json:"name,optional"`       // 路由名称
	Component       string  `json:"component,optional"`  // 组件路径
	Path            string  `json:"path,optional"`       // 路由地址
	Permission      string  `json:"permission,optional"` // 权限标识
	HideInMenu      bool    `json:"hideInMenu,optional"`
	IgnoreKeepAlive bool    `json:"ignoreKeepAlive,optional"`
	LinkFlag        bool    `json:"linkFlag,optional"`
	Link            string  `json:"link,optional"` // 组件路径
	Disabled        bool    `json:"disabled,optional"`
	SelectApi       []int64 `json:"selectApi,optional"`
}

type UpdateMenuResp struct {
}

type UserAddReq struct {
	Status      int32  `json:"status"`               // 状态 1 停用 2 启用
	Username    string `json:"username"`             // 登录名
	Password    string `json:"password"`             // 密码
	Nickname    string `json:"nickname,optional"`    // 昵称
	Description string `json:"description,optional"` // 描述
	Mobile      string `json:"mobile,optional"`      // 手机号
	Email       string `json:"email,optional"`       // 邮箱
	Avatar      string `json:"avatar,optional"`      // 头像
	DeptId      int64  `json:"deptId"`               // 部门id
	RoleId      int64  `json:"roleId"`               // 角色id
}

type UserAddResp struct {
}

type UserDeleteReq struct {
	UserId int64 `path:"userId"`
}

type UserDeleteResp struct {
}

type UserDetailResp struct {
	Avater   string   `json:"avater"`
	Roles    []string `json:"roles"`
	UserId   string   `json:"userId"`
	Username string   `json:"username"`
	Desc     string   `json:"desc"`
	HomePath string   `json:"homePath"`
}

type UserInfoData struct {
	UserId      int64  `json:"userId"`
	Status      int32  `json:"status"`      // 状态 1 停用 2 启用
	Username    string `json:"username"`    // 登录名
	Nickname    string `json:"nickname"`    // 昵称
	Description string `json:"description"` // 描述
	Mobile      string `json:"mobile"`      // 手机号
	Email       string `json:"email"`       // 邮箱
	Avatar      string `json:"avatar"`      // 头像
	DeptId      int64  `json:"deptId"`      // 部门id
	RoleId      int64  `json:"roleId"`      // 角色id
	CreateTime  string `json:"createTime"`
}

type UserInfoReq struct {
	UserId int64 `path:"userId"`
}

type UserListData struct {
	UserId      int64  `json:"userId"`
	Status      int32  `json:"status"`      // 状态 1 停用 2 启用
	Username    string `json:"username"`    // 登录名
	Nickname    string `json:"nickname"`    // 昵称
	Description string `json:"description"` // 描述
	Mobile      string `json:"mobile"`      // 手机号
	Email       string `json:"email"`       // 邮箱
	Avatar      string `json:"avatar"`      // 头像
	DeptId      int64  `json:"deptId"`      // 部门id
	RoleId      int64  `json:"roleId"`      // 角色id
	CreateTime  string `json:"createTime"`
}

type UserListReq struct {
	Username string `form:"username,optional"`   // 登录名
	Nickname string `form:"nickname,optional"`   // 昵称
	Mobile   string `form:"mobile,optional"`     // 手机号
	Email    string `form:"email,optional"`      // 邮箱
	DeptId   int64  `form:"deptId,optional"`     // 部门id
	RoleId   int64  `form:"roleId,optional"`     // 角色id
	PageNum  int64  `form:"pageNum,default=1"`   //  第几页
	PageSize int64  `form:"pageSize,default=10"` // 每页的数量
}

type UserListResp struct {
	Total int64           `json:"total"`
	Data  []*UserListData `json:"data"`
}

type UserUpdateReq struct {
	UserId      int64  `json:"userId"`
	Status      int32  `json:"status"`               // 状态 1 停用 2 启用
	Username    string `json:"username"`             // 登录名
	Password    string `json:"password,optional"`    // 密码
	Nickname    string `json:"nickname,optional"`    // 昵称
	Description string `json:"description,optional"` // 描述
	Mobile      string `json:"mobile,optional"`      // 手机号
	Email       string `json:"email,optional"`       // 邮箱
	Avatar      string `json:"avatar,optional"`      // 头像
	DeptId      int64  `json:"deptId"`               // 部门id
	RoleId      int64  `json:"roleId"`               // 角色id
}

type UserUpdateResp struct {
}
