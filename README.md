<div align="center">
<img src="https://s2.loli.net/2025/07/16/KhZIYaePVlFEMm8.png" style="width:100px;" width="100"/>
<h2>魔链</h2>
</div>

### 一、产品概述

- 一款用 Go 编写的多功能系统命令行补全工具，单个可执行文件，无需运行时依赖。
- 像 BusyBox 那样借助符号链接「补全」系统里缺少的指令，替你返回静态文本、执行脚本或调起真正的程序。
- 支持带逻辑和参数的 Shell 脚本、可执行程序和静态文本三类补全内容，工业现场与测试环境里临时补齐依赖特别好用。
- 可用于软件行为验证、程序依赖的临时补全、教学演示等场景，本地优先、离线可用。

核心价值：当某个环境里缺了一条命令，不必装一整套环境，一条 `magiclink` 就能把这条命令「变」出来。

### 二、软件界面

命令行工具，界面即终端输出，无参数运行时打印 Logo 与完整用法：

![](https://s2.loli.net/2025/07/16/ea6wYMVC9glNSzk.png)

### 三、功能说明

#### 补全机制

- 把二进制放到 `bin` 目录或任何 PATH 位置后，把想补的指令放进程序同级的 `mlink` 文件夹。
- `mlink` 目录结构，`exe` 放可执行程序、`sh` 放可收参数的脚本、`static` 放无后缀文本：

```text
mlink
├── exe
│   └── socat
├── sh
│   ├── add.sh
│   └── hello.sh
└── static
    └── help
```

- `magiclink hello` 直接打印 `static/hello` 的文本内容；`magiclink socat` 调起对应程序；`magiclink add` 执行 Shell 脚本。
- 除静态文本外，其余补全都支持携带任意数量参数。
- `ln -s magiclink xxx` 可把某条指令挂载到全局，之后直接敲 `xxx` 即可调用，且支持 `sudo` 传播。
- 需要补全的局部文件夹里也可以放 `mlink`，局部优先级高于程序同级目录，实现动态伪装与补全。

#### BusyBox 快捷功能

- `magiclink busybox_check`：检查当前系统的 BusyBox 是否可用。
- `magiclink busybox_install`：一键安装 BusyBox。
- `magiclink busybox_mount`：用装好的 BusyBox 批量挂载补全指令。
- 也可以手动把对应平台的 BusyBox 二进制直接放进 `/bin`。

### 四、安装与下载

最新版本：`1.0.20260920`

从 [Releases](https://github.com/Mutantcat-Working-Group/MagicLink/releases) 下载对应平台安装包：

| 平台 | 架构 | 资产 |
| --- | --- | --- |
| Linux | amd64 / arm64 | `.tar.gz` 或 `.AppImage` |
| macOS | Intel / Apple Silicon | `.tar.gz` 或 `.dmg`（ad-hoc 签名） |
| Windows | amd64 | `.tar.gz` 或 `-setup.exe`（NSIS 安装包） |

命令行用法直接用 `.tar.gz` 解压出的二进制即可；想要点开就用的图形化安装体验，可选手里的 AppImage、dmg 或安装包。另附 `checksums.txt` 供校验。

版本号使用纯日期递增（如 `1.0.20260920`），推送同族标签（`v` 前缀可选）后，GitHub Actions 会自动构建三平台安装包并发布 Release。

### 五、快速上手

1. 下载并解压对应平台的包，把 `magiclink` 二进制放进 `/bin` 或任何 PATH 目录。
2. 在二进制同级目录建 `mlink` 文件夹，按 `exe` / `sh` / `static` 分类放入要补的指令。
3. 运行 `magiclink hello` 验证静态文本补全，运行 `magiclink add 1 2` 验证带参脚本。
4. 用 `ln -s magiclink xxx` 把指令挂到全局，之后直接使用 `xxx` 调用。
5. 需要批量补全时依次尝试 `busybox_check`、`busybox_install`、`busybox_mount`。

挂载效果：

![](https://s2.loli.net/2025/07/16/C37q2IP4dljBnks.png)

### 六、从源码构建

```bash
git clone https://github.com/Mutantcat-Working-Group/MagicLink.git
cd MagicLink
go build -ldflags "-X main.version=1.0.20260920" -o magiclink .
```

仓库同时提供 Dockerfile，可在容器内完成交叉编译与验证。

### 七、开源协议

本项目基于 Apache-2.0 协议开源，详见仓库中的 `LICENSE`。
