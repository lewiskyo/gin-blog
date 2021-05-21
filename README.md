# gin-blog

内容来源于博客 https://golang2.eddycjy.com/posts/ch2/01-simple-server/ <br>

跟煎鱼学gin<br>

主要为了学习go和了解gin用法.

v1:<br>
- 进行了项目目录结构设计
- blog-service的数据库以及对应三张表创建
- 公共model，标签model，文章model以及文章标签关联model
- 文章和标签接口的路由设计，以及接口处理的空接口
- 错误码公共组件
- 配置统一管理组件
- 数据库连接组件
- 日志写入组件
- 响应处理组件
- swagger生成接口文档
- 业务接口处理以及国际化组件
- 简单的接口校验示例

启动步骤：
1. 到mysql依次执行 scripts/db.sql创建数据库，然后进入blog-service数据库执行scripts/table.sql创建表
2. 项目根目录下创建 storage/logs目录
3. 修改configs/config.yaml中数据库相关配置
4. 根目录下 go run main.go 启动