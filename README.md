# gin_video_server
一个基于gin的视频网站go服务（用于学习交流）

### 要感谢的人
- 项目主要是我github上看到了煎鱼的blog，基本思路也是根据他的博客学习的，大家可以去他的主页学习
https://github.com/EDDYCJY/blog

- 这个项目参考了慕课网艾文西的《Go语言实战流媒体视频网站》，有机会建议大家可以学习他的课程
http://coding.imooc.com/class/227.html

### 部署
#### 环境
- 确保自己的开发机器上安装了docker，及其相关组件： docker-machine, docker-compose
- 将代码git clone到本地开发机器的$GOPATH/src目录下

#### 构建镜像

##### 构建数据库（mysql）镜像并初始化
1. 从docker hub 拉取最新的mysql镜像
```shell
docker pull mysql
```
2. 进入mysql目录,运行
```shell
docker-compose up -d
```
3. 连接到mysql并执行initdb.sql


##### 构建程序镜像

1. 运行脚本，构建项目镜像
```shell
./docker_server_create.sh
```
2. 在gin_video_server目录下运行
```shell
docker-compose up -d
``` 

##### 登陆一下
http://127.0.0.1/api/v1

### 远程云服务器部署

1. 创建docker-machine到远程云服务器的主机
2. 在云主机上安装docker-compose（也可以在本地使用docker stack，大家自己看看怎么弄，这里就不细说了）
3. 将mysql构建相关文件拷贝到云主机hxaliyunecs
```shell
docker-machine scp -r mysql hxaliyunecs:~/mysql/
``` 
4. 将runtime拷贝的云主机
```shell
docker-machine scp -r runtime/templates hxaliyunecs:~/gin_video_server/runtime/templates/
```
5. 设置docker-machine的env为云服务器
```shell
eval $(docker-machine env hxaliyunecs)
```
6. 在本地为远程云服务器构建项目镜像
```shell
./docker_server_create.sh
```
7. 进入远程云服务器，拉取最新的mysql镜像
```shell
docker-machine ssh hxaliyunecs
docker pull mysql
```
8. 进入远程云服务器，进入mysql目录,运行
```shell
docker-compose up -d
```
9. 进入远程云服务器,连接到mysql并执行initdb.sql
10. 进入远程云服务器, 在gin_video_server目录下运行
```shell
docker-compose up -d
```
11. 登陆一下

http://{远程服务器IP}/api/v1
