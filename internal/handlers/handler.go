package handlers

import (
	"gobee/internal/database"
	album_handler "gobee/internal/handlers/album"
	albumphoto_handler "gobee/internal/handlers/albumphoto"
	apiinterface_handler "gobee/internal/handlers/apiinterface"
	auth_handler "gobee/internal/handlers/auth"
	comment_handler "gobee/internal/handlers/comment"
	common_handler "gobee/internal/handlers/common"
	essay_handler "gobee/internal/handlers/essay"
	file_handler "gobee/internal/handlers/file"
	flink_handler "gobee/internal/handlers/flink"
	flinkgroup_handler "gobee/internal/handlers/flinkgroup"
	friendcircle_handler "gobee/internal/handlers/friendcircle"
	initialize_handler "gobee/internal/handlers/initialize"
	member_handler "gobee/internal/handlers/member"
	memberlevel_handler "gobee/internal/handlers/memberlevel"
	payorder_handler "gobee/internal/handlers/payorder"
	post_handler "gobee/internal/handlers/post"
	product_handler "gobee/internal/handlers/product"
	role_handler "gobee/internal/handlers/role"
	route_handler "gobee/internal/handlers/route"
	setting_handler "gobee/internal/handlers/setting"
	storagestrategy "gobee/internal/handlers/storagestrategy"
	user_handler "gobee/internal/handlers/user"
	visit_handler "gobee/internal/handlers/visit"
	wallet_handler "gobee/internal/handlers/wallet"
	"gobee/pkg"
)

type HandlerMap struct {
	AlbumHandler           album_handler.AlbumHandler
	AlbumPhotoHandler      albumphoto_handler.AlbumPhotoHandler
	ApiInterfaceHandler    apiinterface_handler.ApiInterfaceHandler
	AuthHandler            auth_handler.AuthHandler
	CommentHandler         comment_handler.CommentHandler
	CommonHandler          common_handler.CommonHandler
	FileHandler            file_handler.FileHandler
	FlinkHandler           flink_handler.FlinkHandler
	FlinkGroupHandler      flinkgroup_handler.FlinkGroupHandler
	FriendCircleHandler    friendcircle_handler.FriendCircleHandler
	InitializeHandler      initialize_handler.InitializeHandler
	MemberHandler          member_handler.MemberHandler
	MemberLevelHandler     memberlevel_handler.MemberLevelHandler
	PayOrderHandler        payorder_handler.PayOrderHandler
	PostHandler            post_handler.PostHandler
	ProductHandler         product_handler.ProductHandler
	RoleHandler            role_handler.RoleHandler
	RouteHandler           route_handler.RouteHandler
	SettingHandler         setting_handler.SettingHandler
	UserHandler            user_handler.UserHandler
	EssayHandler           essay_handler.EssayHandler
	StorageStrategyHandler storagestrategy.StorageStrategyHandler
	VisitHandler           visit_handler.VisitHandler
	WalletHandler          wallet_handler.WalletHandler
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
	friendCircleHandler := friendcircle_handler.NewFriendCircleHandlerImpl(serviceMap.FriendCircleService)
	initializeHandler := initialize_handler.NewInitializeHandlerImpl(serviceMap.UserService, serviceMap.SettingService)
	payOrderHandler := payorder_handler.NewPayOrderHandlerImpl(database.DB, serviceMap.PayOrderService)
	postHandler := post_handler.NewPostHandlerImpl(serviceMap.PostService)
	productHandler := product_handler.NewProductHandlerImpl(serviceMap.ProductService)
	roleHandler := role_handler.NewRoleHandlerImpl(serviceMap.RoleService)
	routeHandler := route_handler.NewRouteHandlerImpl()
	settingHandler := setting_handler.NewSettingHandlerImpl(serviceMap.SettingService)
	userHandler := user_handler.NewUserHandlerImpl(serviceMap.UserService, serviceMap.RoleService)
	essayHandler := essay_handler.NewEssayHandlerImpl(serviceMap.EssayService)
	storageStrategyHandler := storagestrategy.NewStorageStrategyHandlerImpl(serviceMap.StorageStrategyService)
	visitHandler := visit_handler.NewVisitHandlerImpl(serviceMap.VisitService)
	walletHandler := wallet_handler.NewWalletHandlerImpl(serviceMap.WalletService)
	memberHandler := member_handler.NewMemberHandlerImpl(serviceMap.UserService, serviceMap.MemberService)
	memberLevelHandler := memberlevel_handler.NewMemberLevelHandlerImpl(serviceMap.MemberLevelService)

	handlerMap := HandlerMap{
		AlbumHandler:           albumHandler,
		AlbumPhotoHandler:      albnumPhotoHandler,
		ApiInterfaceHandler:    apiInterfaceHandler,
		AuthHandler:            authHandler,
		CommentHandler:         commentHandler,
		CommonHandler:          commonHandler,
		FileHandler:            fileHandler,
		FlinkHandler:           flinkHandler,
		FlinkGroupHandler:      flinkGroupHandler,
		FriendCircleHandler:    friendCircleHandler,
		InitializeHandler:      initializeHandler,
		MemberHandler:          memberHandler,
		MemberLevelHandler:     memberLevelHandler,
		PayOrderHandler:        payOrderHandler,
		PostHandler:            postHandler,
		ProductHandler:         productHandler,
		RoleHandler:            roleHandler,
		RouteHandler:           routeHandler,
		SettingHandler:         settingHandler,
		UserHandler:            userHandler,
		EssayHandler:           essayHandler,
		StorageStrategyHandler: storageStrategyHandler,
		VisitHandler:           visitHandler,
		WalletHandler:          walletHandler,
	}

	return handlerMap

}
