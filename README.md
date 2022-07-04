

# DouYin_demo

Project of back-end youth training camp based on ByteDance

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



## 目录

- [上手指南](#上手指南)
    - [开发前的配置要求](#开发前的配置要求)
  
- [文件目录说明](#文件目录说明)
- [使用到的框架](#使用到的框架)
- [贡献者](#贡献者)
    - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
- [作者](#作者)
- [鸣谢](#鸣谢)

### 上手指南

请将所有第三方依赖库导入



###### 开发前的配置要求

1. Golang 1.18.2
2. MySQL  8.0
3. Redis  3.2.1
4. ffmpeg 4.4.1





### 文件目录说明
eg:

```
DouYin_demo
│── /app/
│   ├── /common/                公共模块（请求、响应结构体等）
│   └── /controllers/           业务调度器
│   └── /middleware/            中间件
│   └── /models/                数据库结构体
│   └── /services/              业务层
├── /bootstrap/                 项目启动初始化
├── /config/                    配置结构体
├── /global/                    全局变量
├── /routes/                    路由定义
├── /storage/                   系统日志、文件等静态资源
├── /utils/                     工具函数
├── config.yaml                 配置文件
├── go.mod                      第三方库
├── main.go                     项目启动文件
└── README.md                   README

```











### 使用到的框架

- [Gin](https://gin-gonic.com/)
- [Gorm](https://gorm.io/)
- [Zap](https://github.com/uber-go/zap)
- [Lumberjack](https://github.com/natefinch/lumberjack)
- [MySQL](https://www.mysql.com/)
- [Validator](https://github.com/validator)
- [Viper](https://github.com/spf13/viper)
- [jwt](https://github.com/dgrijalva/jwt-go)
- [Redis](https://redis.io/)

### 贡献者

**alexander.huang**

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

### 作者

alexander.huang

掘金:https://juejin.cn/user/550196363672088  
Email:[alexander.huangai77@gmail.com](alexander.huangai77@gmail.com)

*您也可以在贡献者名单中参看所有参与该项目的开发者。*

### 版权说明


该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/Anzz-bot/DouYin_demo/blob/main/LICENSE.txt)



<!-- links -->

[contributors-shield]: https://img.shields.io/github/contributors/Anzz-bot/DouYin_demo.svg?style=flat-square
[contributors-url]: https://github.com/Anzz-bot/DouYin_demo/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Anzz-bot/DouYin_demo.svg?style=flat-square
[forks-url]: https://github.com/Anzz-bot/DouYin_demo/network/members
[stars-shield]: https://img.shields.io/github/stars/Anzz-bot/DouYin_demo.svg?style=flat-square
[stars-url]: https://github.com/Anzz-bot/DouYin_demo/stargazers
[issues-shield]: https://img.shields.io/github/issues/Anzz-bot/DouYin_demo.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/Anzz-bot/DouYin_demo.svg
[license-shield]: https://img.shields.io/github/license/Anzz-bot/DouYin_demo.svg?style=flat-square
[license-url]: https://github.com/Anzz-bot/DouYin_demo/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian




