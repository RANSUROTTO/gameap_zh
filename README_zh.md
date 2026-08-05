# GameAP 中文汉化版

基于 [GameAP](https://github.com/gameap/gameap) 的中文汉化版本。

## 说明

这是 GameAP 游戏服务器管理面板的中文汉化版。原版支持英文和俄文，这个版本加了简体中文支持，默认语言改成中文。

**注意：这个版本大概率不会跟进上游更新，基本就是一锤子买卖。** 如果你需要最新功能，建议用原版。

## 汉化内容

- 新增 `zh-CN` 语言文件
- 前端界面默认使用中文
- 保留英文和俄文，用户可以自己切换
- 翻译不全的地方会回退英文，不会显示空白

## 使用方式

跟原版一样，没啥区别。Docker 部署最省事：

```bash
docker-compose up -d
```

访问 `http://localhost:8025` 就行。

详细文档看原版：https://github.com/gameap/gameap

## 许可证

MIT，跟原版一样。

## 致谢

感谢 GameAP 团队做了这么好用的面板。
