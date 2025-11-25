# 基础配置（保留原有逻辑 + 新增跨平台配置）
BINARY_NAME=main
DIST_DIR := dist                  # 多平台产物根目录
VERSION := $(shell git describe --tags --always --dirty || echo "v1.0.0")
BUILD_TIME := $(shell date +%Y%m%d_%H%M%S)

# 编译参数：静态编译 + 注入版本信息（不影响原有功能）
LDFLAGS := -ldflags "\
	-s -w \
	-X main.Version=$(VERSION) \
	-X main.BuildTime=$(BUILD_TIME) \
"
CGO_FLAG := CGO_ENABLED=0  # 关闭CGO，生成无依赖静态二进制

# 定义需要编译的目标平台+架构（按需增删）
TARGETS := \
	linux_amd64 \
	linux_arm64 \
	darwin_amd64  \
	darwin_arm64  \
	windows_amd64

# ===================== 原有功能（完全保留） =====================
# 构建前端
build-frontend:
	@echo "Building frontend..."
	cd ui && npm install && npm run build
	@echo "Copying frontend build to assets..."

# 构建后端（本地平台）
build-backend:
	@echo "Building backend (local platform)..."
	go build -o ${BINARY_NAME} ${LDFLAGS} main.go

# 完整构建（前端 + 本地后端）
build: build-frontend build-backend
	@echo "Build completed successfully!"

# 运行项目
run: build
	./${BINARY_NAME}

# 清理构建产物（含多平台产物 + 原有清理逻辑）
clean:
	@echo "Cleaning all build artifacts..."
	go clean
	rm -f ${BINARY_NAME}
	rm -rf ui/dist
	cd ui && rm -rf node_modules
	rm -rf $(DIST_DIR)  # 清理多平台产物目录
	@echo "Clean completed!"

# 开发模式（仅后端，用于开发）
dev:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

# 仅构建前端
frontend: build-frontend

# ===================== 新增：多平台打包功能 =====================
# 默认多平台构建目标（先编译前端，再编译所有平台后端）
.PHONY: build-all-platforms $(TARGETS)
build-all-platforms: build-frontend $(TARGETS)
	@echo "✅ All platforms build completed! Artifacts in $(DIST_DIR)/"

# 模式规则：编译单个平台-架构的后端产物
$(DIST_DIR)/$(BINARY_NAME)-%:
	# 解析平台和架构（如 linux_amd64 → GOOS=linux, GOARCH=amd64）
	$(eval GOOS := $(firstword $(subst _, ,$*)))
	$(eval GOARCH := $(lastword $(subst _, ,$*)))
	# 处理Windows后缀
	$(eval EXT := $(if $(filter windows,$(GOOS)),.exe,))
	# 创建产物目录
	mkdir -p $(DIST_DIR)/$(GOOS)_$(GOARCH)
	# 交叉编译后端（依赖已构建的前端assets）
	@echo "=== Building backend for $(GOOS)/$(GOARCH) ==="
	$(CGO_FLAG) GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
		${LDFLAGS} \
		-o $(DIST_DIR)/$(GOOS)_$(GOARCH)/$(BINARY_NAME)$(EXT) \
		main.go
	# 复制前端静态资源到对应平台目录（可选：如果后端需要携带assets）
	@cp -r assets $(DIST_DIR)/$(GOOS)_$(GOARCH)/
	@echo "✅ $(GOOS)/$(GOARCH) build done: $(DIST_DIR)/$(GOOS)_$(GOARCH)/$(BINARY_NAME)$(EXT)"

# 为每个目标创建快捷规则（如 make linux_amd64）
$(TARGETS): $(DIST_DIR)/$(BINARY_NAME)-%
	@true

# 打包多平台产物（压缩为zip/tar.gz，方便分发）
package: build-all-platforms
	@echo "=== Packaging all artifacts ==="
	@for target in $(TARGETS); do \
		GOOS=$${target%_*}; \
		GOARCH=$${target#*_}; \
		EXT=$$(if [ "$$GOOS" = "windows" ]; then echo ".exe"; else echo ""; fi); \
		cd $(DIST_DIR)/$$target && \
		if [ "$$GOOS" = "windows" ]; then \
			zip -q ../$(BINARY_NAME)-$(VERSION)-$$target.zip $(BINARY_NAME)$$EXT assets/; \
		else \
			tar -zcf ../$(BINARY_NAME)-$(VERSION)-$$target.tar.gz $(BINARY_NAME)$$EXT assets/; \
		fi; \
	done
	@echo "✅ Packaging completed! Compressed files in $(DIST_DIR)/"