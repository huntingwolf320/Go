### 要求
设计完成cloudgo应用程序，即写一个web小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）

### 完成效果
 1. 静态文件服务
在浏览器中输入localhost:8080/static/
**注意：最后必须以/结尾，表示静态文件夹的根目录**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181115172102251.png)
我们可以点击显示文件进行测试，确认支持静态文件的服务

 2. js请求支持
首先使用curl测试输出
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181115172832322.png)
在浏览器输出localhost:8080/进入主界面
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181115172914841.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzM3Nzc5NjA4,size_16,color_FFFFFF,t_70)
打开浏览器的调试器，显示js代码
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181115172948468.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzM3Nzc5NjA4,size_16,color_FFFFFF,t_70)
点击call进行js调用，可以看到显示了返回的API test字样
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181115173041884.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzM3Nzc5NjA4,size_16,color_FFFFFF,t_70)
 3. 模板输出与表单处理
在my post 表单的两个输入框中输入内容
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181115173331444.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzM3Nzc5NjA4,size_16,color_FFFFFF,t_70)
点击submit提交之后，在show my template中显示了模板输出
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181115173435616.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzM3Nzc5NjA4,size_16,color_FFFFFF,t_70)
 ### 扩展部分
 1. gzip代码阅读分析
[gzip代码链接](https://github.com/phyber/negroni-gzip/blob/master/gzip/gzip.go)
 2. 中间件支持
