### 环境依赖：

- **Ubuntu 20.04**
- **git**
- **docker 17.03.0-ce+**
- **docker-compose 1.8**
- **Golang 1.11.x+**
- **make**

### 环境搭建

1. 安装 vim、git、docker(需要Docker版本17.03.0-ce或更高版本)、docker-compose(docker-compose 1.8或更高版本是必需的)

   ```shell
   $ sudo apt install vim
   $ sudo apt install git
   $ sudo apt install docker.io
   $ sudo apt install docker-compose
   ```

2. 将当前用户添加到 docker 组,添加成功后必须注销/退出并重新登录(退出终端重新连接即可)

   ```shell
   $ sudo usermod -aG docker kevin
   ```
   
3. 安装Golang,需要版本1.10.x或更高。

   ```shell
   $ wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
   ```
   
   使用 tar 命令将下载后的压缩包文件解压到指定的 /usr/local/ 路径下
   ```shell
   $ sudo tar -zxvf go1.10.3.linux-amd64.tar.gz -C /usr/local/
   ```

   设置GOPATH & GOROOT环境变量, 通过 go env 查看GOPATH路径
   ```shell
   $ sudo vim /etc/profile
   ```
   
   在profile文件最后添加如下内容:
   ```shell
   $ export GOPATH=$HOME/go
   $ export GOROOT=/usr/local/go
   $ export PATH=$GOROOT/bin:$PATH
   ```
   
   使用 source 命令，使刚刚添加的配置信息生效：
   
   ```shell
   $ source /etc/profile
   ```
   
   通过 go version命令验证是否成功：
   
   ```shell
   $ go version
   ```
   
### 安装步骤
1. 新建目录

   ```shell
   $ mkdir -p $GOPATH/src/github.com/kongyixueyuan.com/
   $ cd $GOPATH/src/github.com/kongyixueyuan.com/
   ```

2. 下载项目,并修改fixtures的从属关系

   ```shell
   $ git clone https://github.com/zhanghaotong/BlockchainWallet.git
   $ mv BlockchainWallet education
   $ cd education
   $ sudo chown -R username:username ./fixtures
   ```

3. 进入fixtures目录并拉取镜像

   ```shell
   $ cd $GOPATH/src/github.com/kongyixueyuan.com/education/fixtures
   $ chmod 777 ./pull_images.sh
   $ ./pull_images.sh
   ```


4. 启动服务

   ```shell
   $ cd $GOPATH/src/github.com/kongyixueyuan.com/education/fixtures
   $ make all
   ```

5. 浏览器访问

   ```url
   http://localhost:9000
   ```

6. 登录用户名及密码

   ```
   用户名：admin
   密码：admin
   ```

7. 停止服务

   Ctrl + C 停止Web服务，然后使用如下命令清空：

   ```shell
   $ make clean
   ```

   ​
