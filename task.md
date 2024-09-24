我发现我过分关注layout
这块可以之后rewrite再去看，可以先一把梭

暂时关注的是容器+dns同步化
    - 自动化部署到生产 + dns 等配置
    - Paas平台 和 运维面板1panel区别还是很大的，Pass是以项目为核心，以项目做原子而不是容器/网站
```
my_project/
├── cmd/
│   └── my_app/
│       └── main.go         # 应用入口
├── internal/
│   ├── model/              # 存放模型
│   │   └── user.go
│   ├── api/                # 存放 API 处理逻辑
│   │   └── user_handler.go
│   ├── dto/                # 存放数据传输对象
│   │   └── user_dto.go
│   └── service/            # 业务逻辑
│       └── user_service.go
├── pkg/                    # 可重用的库
└── go.mod                  # Go 模块文件
```




这监控怎么这么像云管，我测
- Moniter
    - system status
      - 概览,http求平均 宿主机器的 cpu / 内存 / 磁盘空间 / 网卡流量
        - Wrapper
      
      - ws宿主机器的 cpu / 内存 / 磁盘空间
      - service status
      
    - APM: 
      - 监控应用程序的性能，包括请求处理时间、错误率等。
      - 数据库相关
    - 告警通知
      - 自定义阈值和条件
      - 多渠道，乐 (email / sms / 平台站内信)
      - tracker 
    - open-Api     
    - dashboard 
    - 资源分析(低优先级)
- 健康检查
  - 相关依赖
  - 中间件相关连接

- Login 
  - [] 网关单点

- 一键部署
  -  ci/cd build
  - deploy

- 日志
  - log manager 
