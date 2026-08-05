# GameAP 中文汉化版 v1.0.0-zh

GameAP 游戏服务器管理面板的中文汉化版本。

## 说明

这个版本基于 [GameAP 官方仓库](https://github.com/gameap/gameap) 进行了中文汉化。

**注意：这个版本大概率不会跟进上游更新，基本就是一锤子买卖。** 如果你需要最新功能，建议用原版。

## 汉化内容

- 新增简体中文语言文件 (`zh-CN`)
- 前端界面默认使用中文
- 保留英文和俄文，用户可以自己切换
- 翻译不全的地方会回退英文，不会显示空白

## 使用方式

跟原版一样，Docker 部署最省事：

```bash
docker-compose up -d
```

访问 `http://localhost:8025` 就行。

## 本地部署配置

这个版本还带了一个 `docker-compose.local.yml`，是给本地部署用的，配置了：
- 简体中文镜像
- gRPC 端口只绑定 127.0.0.1
- PostgreSQL 和 Redis 不暴露到宿主机

详细文档看原版：https://github.com/gameap/gameap

## 许可证

MIT，跟原版一样。

## 致谢

感谢 GameAP 团队做了这么好用的面板。
