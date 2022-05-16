# Simpublish

简介

## 一键部署

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2FOverflowCat%2Fsimpublish&env=SIMPUBLISH_PASSWD&envDescription=%E8%AE%BF%E9%97%AE%E5%AF%86%E7%A0%81&project-name=my-simpublish-site&repository-name=my-simpublish-site&demo-title=Simpublish%20Demo&demo-url=https%3A%2F%2Fsimpublish.vercel.app%2F)

_Demo: https://simpublish.vercel.app_

## 使用说明

首先需要清空项目中的 `api/_output` 中的示例文件，然后将简悦自动化的 `output` 文件夹中的 HTML 或 Markdown 文件放入 `api/_output` 即可。该操作可以通过各类自动化软件完成，如 Windows 的 [DropIt](http://www.dropitproject.com/) 和 macOS 的 [Hazel](https://www.noodlesoft.com/)。

### 自定义

[`src/custom.css`](src/custom.css) 定义了主题色，可以自行更改。并可在其后附上自定义的 CSS，将会全局生效。

[`src/custom.js`](src/custom.js) 中设定了显示在文章列表顶部的文章名称和 slogan。

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

## 构建

```bash
pnpm run build
```
