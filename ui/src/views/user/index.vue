<template>
  <div class="container-fluid p-4">
    <div class="d-flex justify-content-between align-items-center mb-4">
      <h2>用户管理</h2>
      <n-button type="primary"> <i class="bi bi-plus"></i> 添加用户 </n-button>
    </div>

    <div class="card">
      <div class="card-body">
        <div class="table-responsive">
          <n-table>
            <thead>
              <tr>
                <th>ID</th>
                <th>邮箱</th>
                <th>用户名</th>
                <th>手机号</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              <!-- 用户数据将通过 JavaScript 动态加载 -->
            </tbody>
          </n-table>
        </div>
      </div>
    </div>
  </div>

  <!-- 添加/编辑用户模态框 -->
  <n-modal>
    <template #header>
      <h5 class="modal-title" id="userModalTitle">添加用户</h5>
      <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
    </template>
    <template #default>
      <form id="userForm">
        <input type="hidden" id="userId" />
        <div class="mb-3">
          <label for="email" class="form-label">邮箱</label>
          <input type="email" class="form-control" id="email" required />
        </div>
        <div class="mb-3">
          <label for="name" class="form-label">用户名</label>
          <input type="text" class="form-control" id="name" required />
        </div>
        <div class="mb-3">
          <label for="phoneNumber" class="form-label">手机号</label>
          <input type="tel" class="form-control" id="phoneNumber" />
        </div>
        <div class="mb-3">
          <label for="password" class="form-label">密码</label>
          <input type="password" class="form-control" id="password" required />
        </div>
      </form>
    </template>
    <template #action>
      <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
      <button type="button" class="btn btn-primary" onclick="saveUser()">保存</button>
    </template>
  </n-modal>
</template>
