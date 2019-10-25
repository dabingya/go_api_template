# go_api_template



1. 配置文件读取viper

1. 日志使用logrus，同时覆盖logurs日志为日志增加访问ip与url的字段

1. 日志分为info.log run.log 按天切割

1. 加载中间件跨域、cache、add header

1. 全局路由404判断

1. 新增配置文件连接MySQL

1. 新增参数校验，支持自定义校验规则、自定义错误输出

1. 增加rest返回状态码
