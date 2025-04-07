# 全栈计算器项目

基于 ConnectRPC 的现代全栈计算器实现，包含 Go 后端服务和 Next.js 前端

## 技术栈
- **后端**: Go + ConnectRPC
- **前端**: Next.js + TypeScript
- **通信协议**: Connect Protocol (基于 HTTP/2)

## 功能特性
- [x] 基础四则运算（加减乘除）
- [x] 实时错误处理（除零错误）
- [x] 类型安全的 RPC 通信
- [x] 响应式 UI 设计

## 快速开始

### 前置要求
- Go 1.23+
- Node.js 20+
- buf CLI (`brew install buf`)

### 1. 克隆仓库
```bash
git clone https://github.com/WUZHUHE38/calculator-fullstack.git
cd fullstack-calculator
```

### 2. 安装依赖

**后端依赖**:

```
cd calculator-backend && go mod tidy
```

**前端依赖**:

```
cd calculator-frontend && npm install
```

### 3. 生成 Protobuf 代码

```
buf generate proto
npx buf generate proto
```

### 4. 启动后端服务

```
cd backend && go run main.go
```

服务默认运行在 `http://localhost:8080`

### 5. 启动前端应用

```
cd frontend && npm run start
```

访问 `http://localhost:3000`

## 部署说明

1. 编译后端：

```
cd calculator-backend && go build -o calculator
```

   2.生产构建前端：

```
cd calculator-frontend && npm run build
```

## 许可证

MIT License