package handlers

import (
	"gobee/internal/database"
	album_handler "gobee/internal/handlers/album"
	albumphoto_handler "gobee/internal/handlers/album_photo"
	apiinterface_handler "gobee/internal/handlers/api_interface"
	auth_handler "gobee/internal/handlers/auth"
	comment_handler "gobee/internal/handlers/comment"
	common_handler "gobee/internal/handlers/common"
	essay_handler "gobee/internal/handlers/essay"
	file_handler "gobee/internal/handlers/file"
	flink_handler "gobee/internal/handlers/flink"
	flinkgroup_handler "gobee/internal/handlers/flink_group"
	friendcirclerecord_handler "gobee/internal/handlers/friend_circle_record"
	friendcirclerule_handler "gobee/internal/handlers/friend_circle_rule"
	initialize_handler "gobee/internal/handlers/initialize"
	payorder_handler "gobee/internal/handlers/pay_order"
	post_handler "gobee/internal/handlers/post"
	role_handler "gobee/internal/handlers/role"
	route_handler "gobee/internal/handlers/route"
	setting_handler "gobee/internal/handlers/setting"
	storagestrategy "gobee/internal/handlers/storage_strategy"
	user_handler "gobee/internal/handlers/user"
	visit_handler "gobee/internal/handlers/visit"
	wallet_handler "gobee/internal/handlers/wallet"
	"gobee/pkg"
)

type HandlerMap struct {
	AlbumHandler              album_handler.AlbumHandler
	AlbumPhotoHandler         albumphoto_handler.AlbumPhotoHandler
	ApiInterfaceHandler       apiinterface_handler.ApiInterfaceHandler
	AuthHandler               auth_handler.AuthHandler
	CommentHandler            comment_handler.CommentHandler
	CommonHandler             common_handler.CommonHandler
	FileHandler               file_handler.FileHandler
	FlinkHandler              flink_handler.FlinkHandler
	FlinkGroupHandler         flinkgroup_handler.FlinkGroupHandler
	FriendCircleRecordHandler friendcirclerecord_handler.FriendCircleRecordHandler
	FriendCircleRuleHandler   friendcirclerule_handler.FriendCircleRuleHandler
	InitializeHandler         initialize_handler.InitializeHandler
	PayOrderHandler           payorder_handler.PayOrderHandler
	PostHandler               post_handler.PostHandler
	RoleHandler               role_handler.RoleHandler
	RouteHandler              route_handler.RouteHandler
	SettingHandler            setting_handler.SettingHandler
	UserHandler               user_handler.UserHandler
	EssayHandler              essay_handler.EssayHandler
	StorageStrategyHandler    storagestrategy.StorageStrategyHandler
	VisitHandler              visit_handler.VisitHandler
	WalletHandler             wallet_handler.WalletHandler
}

func InitHandler(serviceMap pkg.ServiceMap) HandlerMap {
	albumHandler := album_handler.NewAlbumHandlerImpl(serviceMap.AlbumService)
	albnumPhotoHandler := albumphoto_handler.NewAlbumPhotoHandlerImpl(serviceMap.AlbumPhotoService)
	apiInterfaceHandler := apiinterface_handler.NewApiInterfaceHandlerImpl(database.DB, serviceMap.ApiInterfaceService)
	authHandler := auth_handler.NewAuthHandlerImpl(serviceMap.AuthService)
	commentHandler := comment_handler.NewCommentHandlerImpl(database.DB, serviceMap.CommentService)
	commonHandler := common_handler.NewCommonHandlerImpl(serviceMap.CommonService)
	fileHandler := file_handler.NewFileHandlerImpl(serviceMap.FileService)
	flinkHandler := flink_handler.NewFlinkHandlerImpl(database.DB, serviceMap.FlinkService)
	flinkGroupHandler := flinkgroup_handler.NewFlinkGroupHandlerImpl(database.DB, serviceMap.FlinkService)
	friendCircleRecordHandler := friendcirclerecord_handler.NewFriendCircleRecordHandlerImpl(serviceMap.FriendCircleRecordService)
	friendCircleRuleHandler := friendcirclerule_handler.NewFriendCircleRuleHandlerImpl(serviceMap.FriendCircleRuleService)
	initializeHandler := initialize_handler.NewInitializeHandlerImpl(serviceMap.UserService, serviceMap.SettingService)
	payOrderHandler := payorder_handler.NewPayOrderHandlerImpl(database.DB, serviceMap.PayOrderService)
	postHandler := post_handler.NewPostHandlerImpl(serviceMap.PostService)
	roleHandler := role_handler.NewRoleHandlerImpl(serviceMap.RoleService)
	routeHandler := route_handler.NewRouteHandlerImpl()
	settingHandler := setting_handler.NewSettingHandlerImpl(serviceMap.SettingService)
	userHandler := user_handler.NewUserHandlerImpl(serviceMap.UserService, serviceMap.RoleService)
	essayHandler := essay_handler.NewEssayHandlerImpl(serviceMap.EssayService)
	storageStrategyHandler := storagestrategy.NewStorageStrategyHandlerImpl(serviceMap.StorageStrategyService)
	visitHandler := visit_handler.NewVisitHandlerImpl(serviceMap.VisitService)
	walletHandler := wallet_handler.NewWalletHandlerImpl(serviceMap.WalletService)

	handlerMap := HandlerMap{
		AlbumHandler:              albumHandler,
		AlbumPhotoHandler:         albnumPhotoHandler,
		ApiInterfaceHandler:       apiInterfaceHandler,
		AuthHandler:               authHandler,
		CommentHandler:            commentHandler,
		CommonHandler:             commonHandler,
		FileHandler:               fileHandler,
		FlinkHandler:              flinkHandler,
		FlinkGroupHandler:         flinkGroupHandler,
		FriendCircleRecordHandler: friendCircleRecordHandler,
		FriendCircleRuleHandler:   friendCircleRuleHandler,
		InitializeHandler:         initializeHandler,
		PayOrderHandler:           payOrderHandler,
		PostHandler:               postHandler,
		RoleHandler:               roleHandler,
		RouteHandler:              routeHandler,
		SettingHandler:            settingHandler,
		UserHandler:               userHandler,
		EssayHandler:              essayHandler,
		StorageStrategyHandler:    storageStrategyHandler,
		VisitHandler:              visitHandler,
		WalletHandler:             walletHandler,
	}

	return handlerMap

}
