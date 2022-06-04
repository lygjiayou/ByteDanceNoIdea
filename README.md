#### douyin

视频上传模块：
1.视频上传至public文件夹，修改了router.go中静态资源位置；

2.封面使用ffmpeg提取视频第一帧，需要安装ffmpeg;测试时可以把api/publish.go:75以下几行注释掉，使用固定封面；

3.修改了model中部分关于video的定义，新版在model下video.go和publish.go中；

4.视频和封面url默认使用http，8080端口，后续可改为根据配置读取

