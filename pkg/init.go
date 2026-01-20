package pkg

import (
	"embed"
	"gobee/ent"
	album_service "gobee/internal/services/content/album"
	albumphoto_service "gobee/internal/services/content/albumphoto"
	category_service "gobee/internal/services/content/category"
	comment_service "gobee/internal/services/content/comment"
	doclibrary_service "gobee/internal/services/content/doclibrary"
	doclibrarydetail_service "gobee/internal/services/content/doclibrarydetail"
	essay_service "gobee/internal/services/content/essay"
	flink_service "gobee/internal/services/content/flink"
	friend_circle_service "gobee/internal/services/content/friendcircle"
	knowledgebase_service "gobee/internal/services/content/knowledgebase"
	post_service "gobee/internal/services/content/post"
	tag_service "gobee/internal/services/content/tag"
	file_service "gobee/internal/services/infra/file"
	migration_service "gobee/internal/services/infra/migration"
	permission_service "gobee/internal/services/infra/permission"
	plugin_service "gobee/internal/services/infra/plugin"
	schedulejob_service "gobee/internal/services/infra/schedulejob"
	storagestrategy_service "gobee/internal/services/infra/storagestrategy"
	visit_service "gobee/internal/services/infra/visit"
	coupon_service "gobee/internal/services/mall/coupon"
	couponusage_service "gobee/internal/services/mall/couponusage"
	member_service "gobee/internal/services/mall/member"
	memberlevel_service "gobee/internal/services/mall/memberlevel"
	payorder "gobee/internal/services/mall/payorder"
	product_service "gobee/internal/services/mall/product"
	wallet_service "gobee/internal/services/mall/wallet"
	api_interface_service "gobee/internal/services/system/apiinterface"
	auth_service "gobee/internal/services/system/auth"
	common_service "gobee/internal/services/system/common"
	notification_service "gobee/internal/services/system/notification"
	role_service "gobee/internal/services/system/role"
	setting_service "gobee/internal/services/system/setting"
	user_service "gobee/internal/services/system/user"
)

type ServiceMap struct {
	AlbumService            album_service.AlbumService
	AlbumPhotoService       albumphoto_service.AlbumPhotoService
	ApiInterfaceService     api_interface_service.ApiInterfaceService
	AuthService             auth_service.AuthService
	CategoryService         category_service.CategoryService
	CommentService          comment_service.CommentService
	CommonService           common_service.CommonService
	CouponService           coupon_service.CouponService
	CouponUsageService      couponusage_service.CouponUsageService
	DocLibraryService       doclibrary_service.DocLibraryService
	DocLibraryDetailService doclibrarydetail_service.DocLibraryDetailService
	EssayService            essay_service.EssayService
	FileService             file_service.FileService
	FlinkService            flink_service.FlinkService
	FriendCircleService     friend_circle_service.FriendCircleService
	KnowledgeBaseService    knowledgebase_service.KnowledgeBaseService
	MemberLevelService      memberlevel_service.MemberLevelService
	MemberService           member_service.MemberService
	MigrationService        migration_service.MigrationService
	NotificationService     notification_service.NotificationService
	PayOrderService         payorder.PayOrderService
	PermissionService       permission_service.PermissionService
	PluginService           plugin_service.PluginService
	PostService             post_service.PostService
	ProductService          product_service.ProductService
	RoleService             role_service.RoleService
	ScheduleJobService      schedulejob_service.ScheduleJobService
	SettingService          setting_service.SettingService
	StorageStrategyService  storagestrategy_service.StorageStrategyService
	TagService              tag_service.TagService
	UserService             user_service.UserService
	VisitService            visit_service.VisitService
	WalletService           wallet_service.WalletService
}

func InitializeServices(moduleDefs embed.FS, db *ent.Client) ServiceMap {

	albumService := album_service.NewAlbumServiceImpl(db)
	albumPhotoService := albumphoto_service.NewAlbumPhotoServiceImpl(db)
	apiInterfaceService := api_interface_service.NewApiInterfaceServiceImpl(db)
	authService := auth_service.NewAuthServiceImpl(db)
	categoryService := category_service.NewCategoryServiceImpl(db)
	commentService := comment_service.NewCommentServiceImpl(db)
	commonService := common_service.NewCommonServiceImpl(db, user_service.NewUserServiceImpl(db), post_service.NewPostServiceImpl(db), comment_service.NewCommentServiceImpl(db))
	couponService := coupon_service.NewCouponServiceImpl(db)
	couponUsageService := couponusage_service.NewCouponUsageServiceImpl(db)
	doclibraryService := doclibrary_service.NewDocLibraryServiceImpl(db)
	doclibrarydetailService := doclibrarydetail_service.NewDocLibraryDetailServiceImpl(db)
	essayService := essay_service.NewEssayServiceImpl(db)
	fileService := file_service.NewFileServiceImpl(db)
	flinkService := flink_service.NewFlinkServiceImpl(db)
	friendCircleService := friend_circle_service.NewFriendCircleServiceImpl(db)
	knowledgeBaseService := knowledgebase_service.NewKnowledgeBaseServiceImpl(db)
	memberLevelService := memberlevel_service.NewMemberLevelServiceImpl(db)
	memberService := member_service.NewMemberServiceImpl(db)
	payOderService := payorder.NewPayOrderServiceImpl(db)
	permissionService := permission_service.NewPermissionServiceImpl(db)
	pluginService := plugin_service.NewPluginServiceImpl(db)
	postService := post_service.NewPostServiceImpl(db)
	productService := product_service.NewProductServiceImpl(db)
	roleService := role_service.NewRoleServiceImpl(db)
	settingService := setting_service.NewSettingServiceImpl(db)
	storageStrategyService := storagestrategy_service.NewStorageStrategyServiceImpl(db)
	tagService := tag_service.NewTagServiceImpl(db)
	userService := user_service.NewUserServiceImpl(db)
	visitService := visit_service.NewVisitServiceImpl(db)
	walletService := wallet_service.NewWalletServiceImpl(db)
	migrationService := migration_service.NewMigrationServiceImpl(db)
	notificationService := notification_service.NewNotificationServiceImpl(db)
	scheduleJobService := schedulejob_service.NewScheduleJobServiceImpl(db)

	permissionService.LoadPermissionsFromDef(moduleDefs)

	serviceMap := ServiceMap{
		AlbumService:            albumService,
		AlbumPhotoService:       albumPhotoService,
		ApiInterfaceService:     apiInterfaceService,
		AuthService:             authService,
		CategoryService:         categoryService,
		CommentService:          commentService,
		CommonService:           commonService,
		CouponService:           couponService,
		CouponUsageService:      couponUsageService,
		DocLibraryService:       doclibraryService,
		DocLibraryDetailService: doclibrarydetailService,
		EssayService:            essayService,
		FileService:             fileService,
		FlinkService:            flinkService,
		FriendCircleService:     friendCircleService,
		KnowledgeBaseService:    knowledgeBaseService,
		MemberLevelService:      memberLevelService,
		MemberService:           memberService,
		MigrationService:        migrationService,
		NotificationService:     notificationService,
		PayOrderService:         payOderService,
		PermissionService:       permissionService,
		PluginService:           pluginService,
		PostService:             postService,
		ProductService:          productService,
		RoleService:             roleService,
		ScheduleJobService:      scheduleJobService,
		SettingService:          settingService,
		StorageStrategyService:  storageStrategyService,
		TagService:              tagService,
		UserService:             userService,
		VisitService:            visitService,
		WalletService:           walletService,
	}

	return serviceMap
}
