package pkg

import (
	"embed"
	"gobee/internal/database"
	album_service "gobee/internal/services/album"
	albumphoto_service "gobee/internal/services/albumphoto"
	api_interface_service "gobee/internal/services/api_interface"
	auth_service "gobee/internal/services/auth"
	comment_service "gobee/internal/services/comment"
	common_service "gobee/internal/services/common"
	essay_service "gobee/internal/services/essay"
	file_service "gobee/internal/services/file"
	flink_service "gobee/internal/services/flink"
	friend_circle_service "gobee/internal/services/friend_circle"
	friendcirclerecord_service "gobee/internal/services/friendcirclerecord"
	member_service "gobee/internal/services/member"
	memberlevel_service "gobee/internal/services/memberlevel"
	payorder "gobee/internal/services/pay_order"
	permission_service "gobee/internal/services/permission"
	post_service "gobee/internal/services/post"
	product_service "gobee/internal/services/product"
	role_service "gobee/internal/services/role"
	setting_service "gobee/internal/services/setting"
	storagestrategy_service "gobee/internal/services/storagestrategy"
	user_service "gobee/internal/services/user"
	visit_service "gobee/internal/services/visit"
	wallet_service "gobee/internal/services/wallet"
	"gobee/pkg/config"
)

const autoMigrate = true

type ServiceMap struct {
	ApiInterfaceService       api_interface_service.ApiInterfaceService
	CommentService            comment_service.CommentService
	CommonService             common_service.CommonService
	FriendCircleService       friend_circle_service.FriendCircleService
	PermissionService         permission_service.PermissionService
	PostService               post_service.PostService
	ProductService            product_service.ProductService
	RoleService               role_service.RoleService
	SettingService            setting_service.SettingService
	UserService               user_service.UserService
	EssayService              essay_service.EssayService
	FlinkService              flink_service.FlinkService
	VisitService              visit_service.VisitService
	PayOrderService           payorder.PayOrderService
	AlbumService              album_service.AlbumService
	AuthService               auth_service.AuthService
	FileService               file_service.FileService
	AlbumPhotoService         albumphoto_service.AlbumPhotoService
	FriendCircleRecordService friendcirclerecord_service.FriendCircleRecordService
	StorageStrategyService    storagestrategy_service.StorageStrategyService
	WalletService             wallet_service.WalletService
	MemberService             member_service.MemberService
	MemberLevelService        memberlevel_service.MemberLevelService
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
	productService := product_service.NewProductServiceImpl(db)
	roleService := role_service.NewRoleServiceImpl(db)
	settingService := setting_service.NewSettingServiceImpl(db)
	userService := user_service.NewUserServiceImpl(db)
	essayService := essay_service.NewEssayServiceImpl(db)
	flinkService := flink_service.NewFlinkServiceImpl(db)
	visitService := visit_service.NewVisitServiceImpl(db)
	payOderService := payorder.NewPayOrderServiceImpl(db)
	albumService := album_service.NewAlbumServiceImpl(db)
	authService := auth_service.NewAuthServiceImpl(db)
	fileService := file_service.NewFileServiceImpl(db)
	albumPhotoService := albumphoto_service.NewAlbumPhotoServiceImpl(db)
	friendCircleRecordService := friendcirclerecord_service.NewFriendCircleRecordServiceImpl(db)
	storageStrategyService := storagestrategy_service.NewStorageStrategyServiceImpl(db)
	walletService := wallet_service.NewWalletServiceImpl(db)
	memberService := member_service.NewMemberServiceImpl(db)
	memberLevelService := memberlevel_service.NewMemberLevelServiceImpl(db)

	//执行
	permissionService.LoadPermissionsFromDef(moduleDefs)

	serviceMap := ServiceMap{
		ApiInterfaceService:       apiInterfaceService,
		CommentService:            commentService,
		CommonService:             commonService,
		FriendCircleService:       friendCircleService,
		PermissionService:         permissionService,
		PostService:               postService,
		ProductService:            productService,
		RoleService:               roleService,
		SettingService:            settingService,
		UserService:               userService,
		EssayService:              essayService,
		FlinkService:              flinkService,
		VisitService:              visitService,
		PayOrderService:           payOderService,
		AlbumService:              albumService,
		AuthService:               authService,
		FileService:               fileService,
		AlbumPhotoService:         albumPhotoService,
		FriendCircleRecordService: friendCircleRecordService,
		StorageStrategyService:    storageStrategyService,
		WalletService:             walletService,
		MemberService:             memberService,
		MemberLevelService:        memberLevelService,
	}

	return serviceMap
}
