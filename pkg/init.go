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
	"gobee/pkg/config"
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
	dbType := config.GetString(config.DATABASE_TYPE)
	dbConfig := database.DBConfig{
		DBType: dbType,
		DBUrl:  config.GetString(config.DATABASE_URL),
	}
	//初始化数据库
	db, err := database.InitializeDB(dbConfig, autoMigrate)
	if err != nil {
		panic(err)
	}

	// 初始化 service
	apiInterfaceService := api_interface_service.NewApiInterfaceServiceImpl(db)
	commentService := comment_service.NewCommentServiceImpl(db)
	commonService := common_service.NewCommonServiceImpl(db, user_service.NewUserServiceImpl(db), post_service.NewPostServiceImpl(db), comment_service.NewCommentServiceImpl(db))
	friendCircleService := friend_circle_service.NewFriendCircleServiceImpl(db)
	permissionService := permission_service.NewPermissionServiceImpl(db)
	postService := post_service.NewPostServiceImpl(db)
	roleService := role_service.NewRoleServiceImpl(db)
	settingService := setting_service.NewSettingServiceImpl(db)
	userService := user_service.NewUserServiceImpl(db)
	essayService := essay_service.NewEssayServiceImpl(db)
	flinkService := flink_service.NewFlinkServiceImpl(db)

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
