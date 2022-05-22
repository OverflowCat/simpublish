<p align="center">
  <img src="https://user-images.githubusercontent.com/81074/169682369-7ac181d8-5096-4cac-8899-e13d70a311a6.png" alt="External v2">
</p>

# Simpublish

基于 Local First 理念，简悦导出文件全部存储在本地同步目录的 `output` 中。而开启导出 HTML 功能时，阅读模式下的文章就会离线备份到本地。因此，我们可以使用成熟的静态文件托管服务托管 `output`   中的 HTML 文件，将本地的文档备份变成在线的资料库。令人可喜的是，这样的静态托管服务很多，且免费额度宽裕，适合个人使用。

结合简悦的 [自动化](http://ksria.com/simpread/docs/#/自动化) 和 [外部链接](https://github.com/Kenshin/simpread/discussions/3725#discussioncomment-2691470) 功能，我们也很容易在阅读文章的过程中，随时把标注的内容输出与输出的外部链接绑定，实现从阅读笔记到外部资料库的深度跳转。

**Simpublish** 就是基于此理念开发的一个 **自建服务方案（建站方案）**，目前支持 Vercel。

## 演示

- https://simpublish-demo.vercel.app 密码：`123456`

## 教程

请前往 [简悦官方教程](https://github.com/Kenshin/simpread/discussions/3960#discussioncomment-2795341)

## 部署

### Deploy 方案 

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2FOverflowCat%2Fsimpublish&env=SIMPUBLISH_PASSWD&envDescription=%E8%AE%BF%E9%97%AE%E5%AF%86%E7%A0%81&project-name=my-simpublish-site&repository-name=my-simpublish-site&demo-title=Simpublish%20Demo&demo-url=https%3A%2F%2Fsimpublish.vercel.app%2F)

详细教程请前往 [Deploy 方案](https://github.com/Kenshin/simpread/discussions/3960#discussioncomment-2797922)

### Fork 方案

详细教程请前往 [Fork 方案](https://github.com/Kenshin/simpread/discussions/3960#discussioncomment-2797923)

## 区别

[官方资料库（阅读列表）](https://kb.simpread.pro/#/page/阅读列表) · [静态部署](https://github.com/Kenshin/simpread/discussions/3823) · [建站方案](https://github.com/Kenshin/simpread/discussions/3960) 的区别 👇 

|        | 阅读列表 | 文章页 | 密码 | 隐私                 | 环境准备 | 难易程度       | 备注             |
|--------|------|-----|----|--------------------|------|------------|----------------|
| 官方方案   | ✔    | ✔   | ✔  | 数据在用户的坚果云中，简悦只负责读取 | 不需要  | 需要授权以及开通权限 [教程](https://kb.simpread.pro/#/page/开通开放平台) | 不需要任何额外环境      |
| 静态部署方案 | ✔    | ✖   | ✖  | 数据在用户的 Github      | Git、Github、Vercel 等 | 简单         | 需要掌握 git 等操作环境 |
| 建站方案   | ✔    | ✔   | ✔  | 数据在用户的 Github      | Git、Github、Vercel 等 | 中等         | 需要掌握 git 等操作环境 |

## 附录

1. [静态建站方案](https://github.com/Kenshin/simpread/discussions/3823)

2. [官方资料库方案](https://kb.simpread.pro/#/page/建立资料库)

## 功能一览

- [x] 密码验证 

- [x] 阅读列表 

- [x] 正文（可直接读取 `idx` 或 `title`） 

- [x] 自定义样式 

- [x] 暗色模式 

- [ ] 更多自定义

## 简单说明

### 文件

首先需要清空项目中的 `api/_output` 中的示例文件，然后将简悦自动化的 `output` 文件夹中的 HTML 或 Markdown 文件放入 `api/_output` 即可。该操作可以通过各类自动化软件完成，如 Windows 的 [DropIt](http://www.dropitproject.com/) 和 macOS 的 [Hazel](https://www.noodlesoft.com/)。

### 密码

API 将读取环境变量 `SIMPUBLISH_PASSWD` 的值作为登录密码，可以在 Vercel 中 Project Settings 里的 Environment Variables 中设置（下图所示）。

1. 通过上方按钮部署时，也会提示设置这一环境变量。

2. 如果留空则没有密码，在登录界面直接点击「验证密码」即可。

![image](https://user-images.githubusercontent.com/81074/169682571-696bbbc0-762f-47a8-8b78-ed596d9d60e3.png)

### 自定义

- [`src/custom.js`](src/custom.js) 中设定了显示在文章列表顶部的文章名称和 slogan。

- [`src/custom.css`](src/custom.css) 定义了主题色，可以自行更改。并可在其后附上自定义的 CSS，将会全局生效。

## 开发

### 安装依赖

`pnpm install`

### 调试

在没有后端的情况下，可以将 `List.svelte` 中的 `isDebug` 设置为 true 以跳过登录界面，载入 mock 的文章列表。

```bash
pnpm run dev

# 在浏览器中打开
pnpm run dev -- --open
```

### 构建

```bash
pnpm run build
```
