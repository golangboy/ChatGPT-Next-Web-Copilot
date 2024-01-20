<h1 align="center">copilot-gpt4-service</h1>

<p align="center">
⚔️ 将 Github Copilot 转换为 ChatGPT
</p>

<p align="center">
简体中文 | <a href="README.md">English</a>
</p>

## 如何使用

1. 安装并启动 copilot-gpt4-service 服务，如本地启动后，API 默认地址为：`http://127.0.0.1:8080`;
2. 获取你的 GitHub 账号 Github Copilot Plugin Token（详见下文）；
3. 安装第三方客户端，如：[ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web)，在设置中填入 copilot-gpt4-service 的 API 地址和 Github Copilot Plugin Token，即可使用 GPT-4 模型进行对话。

## 部署方式

### 最佳实践方式

经社区验证和讨论，最佳实践方式为：

1. 本地部署，仅个人使用（推荐）；
2. 自用服务器集成 [ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web) 部署，服务不公开；
3. 服务器部署，公开但个人使用 (例如多客户端使用场景 [Chatbox](https://github.com/Bin-Huang/chatbox), [OpenCat APP](https://opencat.app/), [ChatX APP](https://apps.apple.com/us/app/chatx-ai-chat-client/id6446304087))。

### 不建议方式

1. 以公共服务的方式提供接口

    多个 Token 在同一个 IP 地址进行请求，容易被判定为异常行为

2. 同客户端 Web(例如 ChatGPT-Next-Web) 以默认 API 以及 API Key 的方式提供公共服务

    同一个 Token 请求频率过高，容易被判定为异常行为

3. Serverless 类型的提供商进行部署

    服务生命周期短，更换 IP 地址频繁，容易被判定为异常行为

4. 其他滥用行为或牟利等行为。

### ⚠️ 非常重要

**非常重要：以上不建议的方式，均可能会导致 Github Copilot 被封禁，且封禁后可能无法解封。**

**非常重要：以上不建议的方式，均可能会导致 Github Copilot 被封禁，且封禁后可能无法解封。**

**非常重要：以上不建议的方式，均可能会导致 Github Copilot 被封禁，且封禁后可能无法解封。**

## 客户端

使用 copilot-gpt4-service，需要配合第三方客户端，目前已测试支持以下客户端：

-   [ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web) (推荐)
-   [Chatbox](https://github.com/Bin-Huang/chatbox)：支持 Windows, Mac, Linux 平台
-   [OpenCat APP](https://opencat.app/)：支持 iOS、Mac 平台
-   [ChatX APP](https://apps.apple.com/us/app/chatx-ai-chat-client/id6446304087) ：支持 iOS、Mac 平台

## 服务端

copilot-gpt4-service 服务的部署方式目前包含 Docker 部署、源码部署、Kubernetes 部署实现，下面分别介绍。

### 配置方式

使用环境变量或环境变量配置文件 `config.env` 配置服务（环境变量优先级高于 `config.env`），默认配置项如下：

```env
HOST=localhost # 服务监听地址
PORT=8080 # 服务监听端口
CACHE=true # 是否启用持久化
CACHE_PATH=db/cache.sqlite3 # 持久化缓存的路径（仅当 CACHE=true 时有效）
DEBUG=false # 是否启用调试模式，启用后会输出更多日志
LOGGING=true # 是否启用日志
LOG_LEVEL=info # 日志级别，可选值：panic、fatal、error、warn、info、debug、trace（注意：仅当 LOGGING=true 时有效）
```

### Docker 部署

#### 一键部署方式

```bash
docker run -d \
  --name copilot-gpt4-service \
  --restart always \
  -p 8080:8080 \
  -e HOST=0.0.0.0 \
  aaamoon/copilot-gpt4-service:latest
```

#### 代码构建方式

```bash
git clone https://github.com/aaamoon/copilot-gpt4-service && cd copilot-gpt4-service
# 可在 docker-compose.yml 中修改端口
docker compose up -d
```

如需更新容器，可在源代码文件夹重新拉取代码及构建镜像，命令如下：

```bash
git pull && docker compose up -d --build
```

### Kubernetes 部署

支持通过 Kubernetes 部署，具体部署方式如下：

```bash
helm repo add aaamoon https://charts.kii.la && helm repo update # 源由 github pages 提供
helm install copilot-gpt4-service aaamoon/copilot-gpt4-service
```

## 支持 HTTPS

<details> <summary> 使用 Caddy 支持 HTTPS </summary>

<p>

[Caddy](https://caddyserver.com/docs/) 可以很方便地为端口服务提供 HTTPS 支持，自动管理证书，省心省力。

以下是一个 Debian/Ubuntu 系统上使用 Caddy 的示例，其他系统请参考 [Caddy 官方文档](https://caddyserver.com/docs/)。

### 安装 Caddy

```bash
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https curl
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy
```

### 配置 Caddy

```bash
sudo vi /etc/caddy/Caddyfile
```

假如你准备使用的域名为 `your.domain.com`，请确保以下条件：

-   请先进行 DNS 解析，将你的域名解析到服务器 IP 地址。
-   开放 80 端口和 443 端口，并且端口没有被其他程序占用，如 Nginx、Xray 等。

然后在 Caddyfile 中添加以下内容：

```bash
your.domain.com {
    reverse_proxy localhost:8080
}
```

### 启动 Caddy

执行以下命令启动 Caddy：

```bash
# 启动 Caddy
sudo systemctl start caddy

# 设置 Caddy 开机自启
sudo systemctl enable caddy

# 查看 Caddy 运行状态
sudo systemctl status caddy
```

如果一切顺利，那此时就可以通过 `https://your.domain.com` 访问 copilot-gpt4-service 服务了。

</p>

</details>

## 与 ChatGPT-Next-Web 一起安装

```bash
helm install copilot-gpt4-service aaamoon/copilot-gpt4-service \
  --set chatgpt-next-web.enabled=true \
  --set chatgpt-next-web.config.OPENAI_API_KEY=[ your openai api key ] \ # copilot 获取的 token
  --set chatgpt-next-web.config.CODE=[ backend access code ] \    # next gpt web ui 的访问密码
  --set chatgpt-next-web.service.type=NodePort \
  --set chatgpt-next-web.service.nodePort=30080
```

## 获取 Copilot Token

首先，你的账号需要开通 Github Copilot 服务

获取 Github Copilot Plugin Token 的方式目前有两种方式：

1. 通过 Python 脚本获取，只需要 requests 库（推荐）。
2. 通过安装 [Github Copilot CLI](https://githubnext.com/projects/copilot-cli/) 授权获取（推荐）。
3. 通过第三方接口授权获取，不推荐，因为不安全。

### 通过 Python 脚本获取

首先确保安装了 Python 3.7+，然后安装 requests 库：

```bash
pip install requests
```

然后执行

**Linux/MacOS 平台获取**

```bash
python3 <(curl -fsSL https://raw.githubusercontent.com/aaamoon/copilot-gpt4-service/master/shells/get_copilot_token.py)
```

可通过设置环境变量或修改脚本第3行的字典设置代理。

**Windows 平台获取**

下载脚本，双击运行即可：[get_copilot_token.py](https://raw.githubusercontent.com/aaamoon/copilot-gpt4-service/master/shells/get_copilot_token.py)。

可修改脚本第3行的字典设置代理。


### 通过 Github Copilot CLI 授权获取

**Linux/MacOS 平台获取**

```bash
# 如下脚本会自动安装 Github Copilot CLI 并通过授权获取 Github Copilot Plugin Token
bash -c "$(curl -fsSL https://raw.githubusercontent.com/aaamoon/copilot-gpt4-service/master/shells/get_copilot_token.sh)"
```

**Windows 平台获取**

下载批处理脚本，双击运行即可：[get_copilot_token.bat](https://raw.githubusercontent.com/aaamoon/copilot-gpt4-service/master/shells/get_copilot_token.bat)。

## 常见问题

### 模型支持情况

据测试：模型参数支持 GPT-4 和 GPT-3.5-turbo，实测使用其他模型均会以默认的 3.5 处理（对比 OpenAI API 的返回结果，猜测应该是最早的版本 GPT-4-0314 和 GPT-3.5-turbo-0301）

### 如何判断是不是 GPT-4 模型

鲁迅为什么暴打周树人？

-   GPT-3.5 会一本正经的胡说八道
-   GPT-4 表示鲁迅和周树人是同一个人

我爸妈结婚时为什么没有邀请我？

-   GPT-3.5 他们当时认为你还太小，所以没有邀请你。
-   GPT-4 他们结婚时你还没出生。

### HTTP 响应状态码解析说明

-   401: 使用的 Github Copilot Plugin Token 过期了或者错误，请重新获取
-   403: 使用的账号没有开通 Github Copilot

## 鸣谢

### 贡献者

<a href="https://github.com/aaamoon/copilot-gpt4-service/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aaamoon/copilot-gpt4-service&anon=0" />
</a>

## 开源协议

[MIT](https://opensource.org/license/mit/)
