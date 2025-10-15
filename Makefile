BINARY_NAME=main

# 构建前端
build-frontend:
	@echo "Building frontend..."
	cd ui && npm install && npm run build
	@echo "Copying frontend build to assets..."
	@mkdir -p assets
	@rm -rf assets/*
	@cp -r ui/dist/* assets/

# 构建后端
build-backend:
	@echo "Building backend..."
	go build -o ${BINARY_NAME} main.go

# 完整构建（前端 + 后端）
build: build-frontend build-backend
	@echo "Build completed successfully!"

# 运行项目
run: build
	./${BINARY_NAME}

# 清理构建产物
clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -rf assets/*
	rm -rf ui/dist
	cd ui && rm -rf node_modules

# 开发模式（仅后端，用于开发）
dev:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

# 仅构建前端
frontend:
	cd ui && npm install && npm run build
	@mkdir -p assets
	@rm -rf assets/*
	@cp -r ui/dist/* assets/