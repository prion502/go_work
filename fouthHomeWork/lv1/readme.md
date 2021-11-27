#context包
Context包实现功能函数
Login函数:实现登录注册界面,form表单传数据

DoLogin函数:
通过post传过来的值，查询cookie。

1.如果不存在该cookie,则显示注册，设置该cookie

2.如果存在该cookie，检查cookie的值，如果和密码相等，则显示登录成功

3.如果和密码不相等，则显示登录失败,密码错误
#templates文件夹
渲染的html模板


