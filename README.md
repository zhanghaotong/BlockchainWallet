### 环境依赖：

- **Ubuntu 20.04**
- **git**
- **docker 17.03.0-ce+**
- **docker-compose 1.8**
- **Golang 1.11.x+**
- **make**

### 安装步骤

1. 新建目录

   ```shell
   $ mkdir -p $GOPATH/src/github.com/kongyixueyuan.com/
   $ cd $GOPATH/src/github.com/kongyixueyuan.com/
   ```

2. 下载项目

   ```go
   $ git clone https://github.com/zhanghaotong/BlockchainWallet.git
   ```

3. 进入fixtures目录启动网络

   ```shell
   $ cd $GOPATH/src/github.com/kongyixueyuan.com/BlockchainWallet/fixtures
   $ docker-compose up
   ```

4. 返回至项目根目录

   ```shell
   $ cd $GOPATH/src/github.com/kongyixueyuan.com/BlockchainWallet
   ```


5. 启动服务

   ```shell
   $ make all
   ```

6. 浏览器访问

   ```url
   http://localhost:9000
   ```

7. 登录用户名及密码

   ```
   用户名：admin
   密码：admin
   ```

8. 停止服务

   Ctrl + C 停止Web服务，然后使用如下命令清空：

   ```shell
   $ make clean
   ```

   ​
