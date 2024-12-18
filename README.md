# GinFramework

**GinFramework** 是一个基于 Gin 的模板项目，采用分层架构设计，结合 TOT（The Onion Architecture）和依赖倒置原则，旨在构建灵活、可测试、稳定且解耦的 Web 应用程序。此模板集成了 Wire 进行依赖注入，并实现了 Redis、Mysql 支持，进一步提升了项目的可维护性和扩展性。

## 功能概述

- **分层架构**：采用分层架构设计，确保各模块之间的低耦合性和高内聚性。
- **依赖倒置原则**：通过接口和依赖注入实现模块间的解耦，增强系统的灵活性。
- **路由管理**：集成 Gin 的路由功能，支持灵活的路由定义和管理。
- **中间件**：预设常用中间件，如日志记录、跨域处理和错误处理。
- **配置管理**：支持通过环境变量或配置文件进行配置，确保灵活的环境适配。
- **依赖注入**：集成 Wire 进行依赖注入，简化服务的初始化和管理。
- **请求/响应封装**：提供统一的请求和响应格式，便于接口开发和调试。
- **数据库集成**：预设与 GORM 的集成，支持 MySQL 等数据库的连接与操作。
- **Redis 集成**：实现 Redis 支持，用于缓存和会话管理，提升应用性能。

## 使用方法

1. **克隆仓库**

   ```
   git clone https://github.com/QiuTian-324/GinFramework.git
   cd GinFramework
   ```

2. **安装依赖**

   ```
   go mod tidy
   ```

3. **生成依赖注入代码**

   ```
   wire
   ```

4. **运行项目**

   ```
   go run main.go
   ```

5. **配置项目**

   修改 `config.yaml` 或设置环境变量来配置项目，包括 Redis 配置。

## 项目结构

```
GinFramework/
├── api/                # API 层，处理路由和控制器逻辑
│   ├── chat.go         # 聊天相关的 API
│   ├── router.go       # 路由配置
│   └── user.go         # 用户相关的 API
├── cmd/                # 应用程序入口
│   ├── modules.go      # 模块初始化
│   ├── start.go        # 启动文件
│   ├── wire_gen.go     # Wire 生成的依赖注入代码
│   └── wire.go         # Wire 配置
├── configs/            # 配置文件目录
├── global/             # 全局变量和常量
│   ├── constant.go     # 常量定义
│   └── global.go       # 全局变量
├── internal/           # 内部实现，包含各层的具体实现
│   ├── config/         # 配置管理
│   ├── data/           # 数据层，数据库操作
│   ├── dto/            # 数据传输对象
│   ├── handlers/       # 处理器
│   ├── libs/           # 库和工具
│   ├── middleware/     # 中间件
│   ├── repo/           # 仓储层，数据持久化
│   ├── server/         # 服务器配置
│   └── services/       # 服务层，业务逻辑
├── log/                # 日志管理
├── pkg/                # 包，第三方库
├── test/               # 测试文件
├── tmp/                # 临时文件
├── utils/              # 工具函数
├── go.mod              # Go 模块文件
├── go.sum              # 依赖版本锁定文件
└── main.go             # 主程序入口
```

## 功能计划

未来的更新计划包括但不限于：

- **用户认证模块**：提供 JWT、OAuth2 等认证方式的集成。
- **权限管理**：实现基于角色的权限控制 (RBAC)。
- **API 文档生成**：自动生成 API 文档，支持 Swagger 等格式。
- **多数据库支持**：扩展支持更多数据库，如 PostgreSQL、MongoDB。
- **消息队列集成**：支持 Kafka、RabbitMQ 等消息队列的使用。
- **国际化支持**：提供多语言支持，便于国际化应用开发。

## 贡献

欢迎提交 Issue 和 Pull Request，为项目的发展贡献力量。

## 许可

此项目遵循 MIT 许可。

## 联系
如果您有任何问题或建议，请通过以下方式联系我们：

- 邮箱: yjj4872@gmail.com
- WeChat: Akita324
- QQ: 1240092443
- GitHub: [QiuTian-324](https://github.com/QiuTian-324)
- GitHub Issues: [GinFramework Issues](https://github.com/QiuTian-324/GinFramework/issues)

我们会尽快回复您的问题，并感谢您对我们项目的支持和关注。