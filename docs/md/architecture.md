## 架构

internal：内部架构，请求调用链路从上往下
routers：路由
service：服务和入参校验
dao：封装调用 model
model：数据库 model
----
middleware：中间件
- access_log：访问日志
- app_info：程序信息
- context_timeout：请求超时
- jwt：鉴权
- limiter：限流
- recovery：故障恢复和邮件报警
- translations：国际化


pkg：内容为通用组件内容，贯穿整个project，如：
- app：
    - response 响应处理；
    - pagination 分页处理；
    - form 错误信息国际化；
    - jwt 鉴权处理;
- errcode：通用错误码
- logger：日志
- setting：project 配置<br/>配置读入 global 中
- upload：文件上传；从安全角度上说，文件资源不应当与应用资源放在一起
- util：工具库；md5序列化
- email：邮件服务
- limiter：流量限流
- tracer：链路追踪


configs：yaml 配置文件<br>
scripts：存放 sql 等脚本<br>
storage：存放日志等文件<br>
