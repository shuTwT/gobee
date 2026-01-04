package pkg

import (
	"embed"
	"gobee/internal/database"
	api_interface_service "gobee/internal/services/api_interface"
	comment_service "gobee/internal/services/comment"
	common_service "gobee/internal/services/common"
	essay_service "gobee/internal/services/essay"
	flink_service "gobee/internal/services/flink"
	friend_circle_service "gobee/internal/services/friend_circle"
	permission_service "gobee/internal/services/permission"
	post_service "gobee/internal/services/post"
	role_service "gobee/internal/services/role"
	setting_service "gobee/internal/services/setting"
	user_service "gobee/internal/services/user"
)

const autoMigrate = true

type ServiceMap struct {
	ApiInterfaceService api_interface_service.ApiInterfaceService
	CommentService      comment_service.CommentService
	CommonService       common_service.CommonService
	FriendCircleService friend_circle_service.FriendCircleService
	PermissionService   permission_service.PermissionService
	PostService         post_service.PostService
	RoleService         role_service.RoleService
	SettingService      setting_service.SettingService
	UserService         user_service.UserService
	EssayService        essay_service.EssayService
	FlinkService        flink_service.FlinkService
}

func InitializeServices(moduleDefs embed.FS) ServiceMap {
	//初始化数据库
	database.InitializeDB(database.DBConfig{
		DBType:     "sqlite",
		SqliteFile: "",
	}, autoMigrate)
	// 初始化 service
	apiInterfaceService := api_interface_service.NewApiInterfaceServiceImpl(database.DB)
	commentService := comment_service.NewCommentServiceImpl(database.DB)
	commonService := common_service.NewCommonServiceImpl(database.DB, user_service.NewUserServiceImpl(database.DB), post_service.NewPostServiceImpl(database.DB), comment_service.NewCommentServiceImpl(database.DB))
	friendCircleService := friend_circle_service.NewFriendCircleServiceImpl(database.DB)
	permissionService := permission_service.NewPermissionServiceImpl(database.DB)
	postService := post_service.NewPostServiceImpl(database.DB)
	roleService := role_service.NewRoleServiceImpl(database.DB)
	settingService := setting_service.NewSettingServiceImpl(database.DB)
	userService := user_service.NewUserServiceImpl(database.DB)
	essayService := essay_service.NewEssayService(database.DB)
	flinkService := flink_service.NewFlinkServiceImpl(database.DB)

	//执行
	permissionService.LoadPermissionsFromDef(moduleDefs)

	serviceMap := ServiceMap{
		ApiInterfaceService: apiInterfaceService,
		CommentService:      commentService,
		CommonService:       commonService,
		FriendCircleService: friendCircleService,
		PermissionService:   permissionService,
		PostService:         postService,
		RoleService:         roleService,
		SettingService:      settingService,
		UserService:         userService,
		EssayService:        essayService,
		FlinkService:        flinkService,
	}

	return serviceMap
}
