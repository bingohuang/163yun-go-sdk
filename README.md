163yun-Go-SDK
=============
[网易云](https://www.163yun.com/) [OpenAPI](https://www.163yun.com/help/assets?tab=api) [Go SDK](https://github.com/bingohuang/163yun-go-sdk)

## 简介

网易云 API 提供更灵活的资源控制方式，满足你的定制化需求，如自动化部署，持续集成等功能。

网易云 API 需要使用 API Token 来发起 API 请求。 请到 https://c.163.com 页面登录到你的账户，查看你的 Access Key 和 Access Secret，然后调用 生成 API token 接口 来获取 token。

API 访问 https://open.c.163.com 。API 采用 Restful API 风格，支持内省，支持 API 访问频率限制。

## 功能特点

- 支持内省
- 为每一个请求响应包含一个 Request-Id 头，并使用 UUID 作为该值。通过在客户端、服务器或任何支持服务上记录该值，它能为我们提供一种机制来跟踪、诊断和调试请求。
- 安全
- 所有 API 均采用 Https。
- 频率限制
- 针对授权接口，支持基于用户的频率限制，频率值根据用户的等级级别设置；
- 针对非授权接口，支持基于 IP 的频率限制。

## Access Key

注册后，网易云会颁发 Access Key 和 Access Secret 给客户 。 没有颁发认证授权的将无法调用任何接口 说明：

- Access Key： 用于标识客户身份，在网络请求中会以某种形式传输
- Access Secret ： 作为私钥形式存储于客户方本地， 不在网络传递，它的作用是对客户方发起的请求进行数字签名，保证该请求是来自指定客户的请求，并且是合法的有效的。

使用 Access Key 进行身份识别，加上 Access Secret 进行数字签名，即可完成应用接入与认证授权。 你可以在 Access Key 中创建、下载、删除、禁用你的 Access Key。 网易云支持多 Access Key、Access Secret，保障你的应用安全。

## OpenAPI 接口设计风格

OpenAPI 的设计采用了 Restful API 风格。下面将简要介绍 OpenAPI 接口设计风格。

### 协议

客户端访问 OpenAPI 接口服务，必须采用 HTTPS 协议，即所有的访问接口行为，都需要用 TLS ，通过安全连接来访问。

### 版本号

OpenAPI 将 API 版本号放入了 URL 中。举个例子，当前使用 OpenAPI 创建镜像仓库的URL为 https://open.c.163.com/API/v1/repositories ，其中 v1 指的是当前 API 版本为v1。

### 请求

如果请求的 body 需要携带参数，需要将参数 JSON 格式化，否则会得到错误响应。同时，请求中需要规范地使用HTTP 请求方法动词，即：
1. GET方法：从服务器获取资源（一项或多项）；
2. POST方法：在服务器上新建资源；
3. PUT：在服务器上更新资源（客户端提供改变后的完整资源）；
4. DELETE：在服务器上删除资源。

### 响应

OpenAPI 的接口服务会为每次响应，按照 HTTP 规范返回合适的响应状态码。并且，在错误响应中，除了上述的 Request-ID，还会生成结构化的错误信息，即机器可读错误码（code）和肉眼可读的错误信息（message）。其中错误码的构造规则为：code = HttpStatus.code (长度3) + OpenAPI 服务编码(长度2) + 细分错误码(长度2)。对于正确的响应，针对不同的操作，OpenAPI 接口向用户返回的结果符合以下规范：

1. GET请求：返回资源对象的列表或者单个资源对象；
2. POST请求：返回新建的资源的ID；
3. PUT请求：返回处理结果状态码，比如状态码 200；
4. DELETE请求：返回处理结果状态码，比如状态码 200。

## OpenAPI 使用步骤

OpenAPI 的使用主要分为三步：
1. 得到 Key 和 Secret 。进入网易云的控制台，点开个人基本信息，接着进入 API 子页面，就可以看到个人Access Key 和 App Secret；
2. 获取 Token。使用1中得到的 Access Key 和 App Secret，调用 OpenAPI 中的 Token 生成接口，获取授权Token；
3. 带上2中得到的 Token，便可以访问 OpenAPI 中的全部接口服务。如果请求没有 Token，OpenAPI 的安全机制将会视其为对服务（除了获取 Token 服务和 WebHook 服务）的未授权访问操作，进而会拦截。

> Attention
> OpenAPI 中内置了流控机制，对于恶意攻击 OpenAPI 服务的行为，会采取限制措施。



