package handlers

import (
	"gobee/internal/database"
	album_handler "gobee/internal/handlers/content/album"
	albumphoto_handler "gobee/internal/handlers/content/albumphoto"
	category_handler "gobee/internal/handlers/content/category"
	comment_handler "gobee/internal/handlers/content/comment"
	essay_handler "gobee/internal/handlers/content/essay"
	flink_handler "gobee/internal/handlers/content/flink"
	flinkgroup_handler "gobee/internal/handlers/content/flinkgroup"
	friendcircle_handler "gobee/internal/handlers/content/friendcircle"
	post_handler "gobee/internal/handlers/content/post"
	tag_handler "gobee/internal/handlers/content/tag"
	file_handler "gobee/internal/handlers/infra/file"
	migration_handler "gobee/internal/handlers/infra/migration"
	storagestrategy "gobee/internal/handlers/infra/storagestrategy"
	visit_handler "gobee/internal/handlers/infra/visit"
	coupon_handler "gobee/internal/handlers/mall/coupon"
	couponusage_handler "gobee/internal/handlers/mall/couponusage"
	member_handler "gobee/internal/handlers/mall/member"
	memberlevel_handler "gobee/internal/handlers/mall/memberlevel"
	payorder_handler "gobee/internal/handlers/mall/payorder"
	product_handler "gobee/internal/handlers/mall/product"
	wallet_handler "gobee/internal/handlers/mall/wallet"
	apiinterface_handler "gobee/internal/handlers/system/apiinterface"
	auth_handler "gobee/internal/handlers/system/auth"
	common_handler "gobee/internal/handlers/system/common"
	initialize_handler "gobee/internal/handlers/system/initialize"
	notification_handler "gobee/internal/handlers/system/notification"
	role_handler "gobee/internal/handlers/system/role"
	route_handler "gobee/internal/handlers/system/route"
	setting_handler "gobee/internal/handlers/system/setting"
	user_handler "gobee/internal/handlers/system/user"
	"gobee/pkg"
)

type HandlerMap struct {
	AlbumHandler           album_handler.AlbumHandler
	AlbumPhotoHandler      albumphoto_handler.AlbumPhotoHandler
	ApiInterfaceHandler    apiinterface_handler.ApiInterfaceHandler
	AuthHandler            auth_handler.AuthHandler
	CategoryHandler        category_handler.CategoryHandler
	CommentHandler         comment_handler.CommentHandler
	CommonHandler          common_handler.CommonHandler
	CouponHandler          coupon_handler.CouponHandler
	CouponUsageHandler     couponusage_handler.CouponUsageHandler
	FileHandler            file_handler.FileHandler
	FlinkHandler           flink_handler.FlinkHandler
	FlinkGroupHandler      flinkgroup_handler.FlinkGroupHandler
	FriendCircleHandler    friendcircle_handler.FriendCircleHandler
	InitializeHandler      initialize_handler.InitializeHandler
	MemberHandler          member_handler.MemberHandler
	MemberLevelHandler     memberlevel_handler.MemberLevelHandler
	MigrationHandler       migration_handler.MigrationHandler
	NotificationHandler    notification_handler.NotificationHandler
	PayOrderHandler        payorder_handler.PayOrderHandler
	PostHandler            post_handler.PostHandler
	ProductHandler         product_handler.ProductHandler
	RoleHandler            role_handler.RoleHandler
	RouteHandler           route_handler.RouteHandler
	SettingHandler         setting_handler.SettingHandler
	TagHandler             tag_handler.TagHandler
	UserHandler            user_handler.UserHandler
	EssayHandler           essay_handler.EssayHandler
	StorageStrategyHandler storagestrategy.StorageStrategyHandler
	VisitHandler           visit_handler.VisitHandler
	WalletHandler          wallet_handler.WalletHandler
}

func InitHandler(serviceMap pkg.ServiceMap) HandlerMap {
	albumHandler := album_handler.NewAlbumHandlerImpl(serviceMap.AlbumService)
	albnumPhotoHandler := albumphoto_handler.NewAlbumPhotoHandlerImpl(serviceMap.AlbumPhotoService)
	apiInterfaceHandler := apiinterface_handler.NewApiInterfaceHandlerImpl(serviceMap.ApiInterfaceService)
	authHandler := auth_handler.NewAuthHandlerImpl(serviceMap.AuthService)
	categoryHandler := category_handler.NewCategoryHandlerImpl(serviceMap.CategoryService)
	commentHandler := comment_handler.NewCommentHandlerImpl(serviceMap.CommentService)
	commonHandler := common_handler.NewCommonHandlerImpl(serviceMap.CommonService)
	couponHandler := coupon_handler.NewCouponHandlerImpl(serviceMap.CouponService)
	couponUsageHandler := couponusage_handler.NewCouponUsageHandlerImpl(serviceMap.CouponUsageService)
	fileHandler := file_handler.NewFileHandlerImpl(serviceMap.FileService)
	flinkHandler := flink_handler.NewFlinkHandlerImpl(database.DB, serviceMap.FlinkService)
	flinkGroupHandler := flinkgroup_handler.NewFlinkGroupHandlerImpl(database.DB, serviceMap.FlinkService)
	friendCircleHandler := friendcircle_handler.NewFriendCircleHandlerImpl(serviceMap.FriendCircleService)
	initializeHandler := initialize_handler.NewInitializeHandlerImpl(serviceMap.UserService, serviceMap.SettingService)
	payOrderHandler := payorder_handler.NewPayOrderHandlerImpl(database.DB, serviceMap.PayOrderService)
	postHandler := post_handler.NewPostHandlerImpl(serviceMap.PostService)
	productHandler := product_handler.NewProductHandlerImpl(serviceMap.ProductService)
	roleHandler := role_handler.NewRoleHandlerImpl(serviceMap.RoleService)
	routeHandler := route_handler.NewRouteHandlerImpl()
	settingHandler := setting_handler.NewSettingHandlerImpl(serviceMap.SettingService)
	tagHandler := tag_handler.NewTagHandlerImpl(serviceMap.TagService)
	userHandler := user_handler.NewUserHandlerImpl(serviceMap.UserService, serviceMap.RoleService)
	essayHandler := essay_handler.NewEssayHandlerImpl(serviceMap.EssayService)
	storageStrategyHandler := storagestrategy.NewStorageStrategyHandlerImpl(serviceMap.StorageStrategyService)
	visitHandler := visit_handler.NewVisitHandlerImpl(serviceMap.VisitService)
	walletHandler := wallet_handler.NewWalletHandlerImpl(serviceMap.WalletService)
	memberHandler := member_handler.NewMemberHandlerImpl(serviceMap.UserService, serviceMap.MemberService)
	memberLevelHandler := memberlevel_handler.NewMemberLevelHandlerImpl(serviceMap.MemberLevelService)
	migrationHandler := migration_handler.NewMigrationHandlerImpl(serviceMap.MigrationService)
	notificationHandler := notification_handler.NewNotificationHandlerImpl(serviceMap.NotificationService)

	handlerMap := HandlerMap{
		AlbumHandler:           albumHandler,
		AlbumPhotoHandler:      albnumPhotoHandler,
		ApiInterfaceHandler:    apiInterfaceHandler,
		AuthHandler:            authHandler,
		CategoryHandler:        categoryHandler,
		CommentHandler:         commentHandler,
		CommonHandler:          commonHandler,
		CouponHandler:          couponHandler,
		CouponUsageHandler:     couponUsageHandler,
		FileHandler:            fileHandler,
		FlinkHandler:           flinkHandler,
		FlinkGroupHandler:      flinkGroupHandler,
		FriendCircleHandler:    friendCircleHandler,
		InitializeHandler:      initializeHandler,
		MemberHandler:          memberHandler,
		MemberLevelHandler:     memberLevelHandler,
		MigrationHandler:       migrationHandler,
		NotificationHandler:    notificationHandler,
		PayOrderHandler:        payOrderHandler,
		PostHandler:            postHandler,
		ProductHandler:         productHandler,
		RoleHandler:            roleHandler,
		RouteHandler:           routeHandler,
		SettingHandler:         settingHandler,
		TagHandler:             tagHandler,
		UserHandler:            userHandler,
		EssayHandler:           essayHandler,
		StorageStrategyHandler: storageStrategyHandler,
		VisitHandler:           visitHandler,
		WalletHandler:          walletHandler,
	}

	return handlerMap

}
