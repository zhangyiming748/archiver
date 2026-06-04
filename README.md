# Archiver

Archiver 是一个命令行工具，用于媒体文件管理和格式转换。支持视频转 H265 格式和图片转 AVIF 格式。

## ✨ 功能特性

- 🎬 **视频转换**：自动查找目录下的所有视频文件并转换为 H265 编码
- 🖼️ **图片转换**：自动查找目录下的所有图片文件并转换为 AVIF 格式
- ⚡ **高效处理**：支持 FHD 模式，提供更高质量的视频转换
- 🌍 **跨平台**：支持 Linux、macOS 和 Windows

## 📥 快速下载

### 从 GitHub Releases 下载

|平台|架构|下载链接|
|:---:|:---:|:---:|
|Linux|amd64|[archiver_linux_amd64](https://github.com/zhangyiming748/archiver/releases/latest/download/archiver_linux_amd64)|
|Linux|arm64|[archiver_linux_arm64](https://github.com/zhangyiming748/archiver/releases/latest/download/archiver_linux_arm64)|
|macOS|amd64|[archiver_darwin_amd64](https://github.com/zhangyiming748/archiver/releases/latest/download/archiver_darwin_amd64)|
|macOS|arm64(AppleSilicon)|[archiver_darwin_arm64](https://github.com/zhangyiming748/archiver/releases/latest/download/archiver_darwin_arm64)|
|Windows|amd64|[archiver_windows_amd64.exe](https://github.com/zhangyiming748/archiver/releases/latest/download/archiver_windows_amd64.exe)|
|Windows|arm64|[archiver_windows_arm64.exe](https://github.com/zhangyiming748/archiver/releases/latest/download/archiver_windows_arm64.exe)|

**一键下载命令：**

```bash
# Linux/macOS
wget https://github.com/zhangyiming748/archiver/releases/latest/download/archiver_$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m | sed 's/x86_64/amd64/; s/aarch64/arm64/') -O archiver && chmod +x archiver

# Windows PowerShell (amd64)
Invoke-WebRequest -Uri "https://github.com/zhangyiming748/archiver/releases/latest/download/archiver_windows_amd64.exe" -OutFile "archiver.exe"

# Windows PowerShell (arm64)
Invoke-WebRequest -Uri "https://github.com/zhangyiming748/archiver/releases/latest/download/archiver_windows_arm64.exe" -OutFile "archiver.exe"
```

### 使用 Go 安装

```bash
go install github.com/zhangyiming748/archiver@latest
```

### 从源码编译

```bash
git clone https://github.com/zhangyiming748/archiver.git
cd archiver
go build -o archiver
```

## 🚀 使用方法

### 基本命令

```bash
# 查看版本信息
archiver version

# 查看帮助
archiver --help
```

### 视频转换

将指定目录下的所有视频文件转换为 H265 格式：

```bash
# 基本用法
archiver video --dir /path/to/videos

# 简写形式
archiver video -d /path/to/videos

# 启用 FHD 模式（更高质量）
archiver video -d /path/to/videos --fhd
archiver video -d /path/to/videos -f
```

**参数说明：**

- `-d, --dir`：必需，要搜索视频文件的目录路径
- `-f, --fhd`：可选，启用 FHD 模式进行视频转换

### 图片转换

将指定目录下的所有图片文件转换为 AVIF 格式：

```bash
# 基本用法
archiver image --dir /path/to/images

# 简写形式
archiver image -d /path/to/images

# 启用 FHD 模式
archiver image -d /path/to/images --fhd
archiver image -d /path/to/images -f
```

**参数说明：**
- `-d, --dir`：必需，要搜索图片文件的目录路径
- `-f, --fhd`：可选，启用 FHD 模式进行图片转换

## 📋 使用示例

```bash
# 转换 ~/Videos 目录下的所有视频
archiver video -d ~/Videos

# 转换 ~/Pictures 目录下的所有图片，使用 FHD 模式
archiver image -d ~/Pictures -f

# 查看当前版本
archiver version
```

## 🔧 技术栈

- **语言**：Go
- **CLI 框架**：Cobra
- **视频处理**：FFmpeg
- **图片处理**：libavif

## 📝 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📮 反馈

如有问题或建议，请通过 GitHub Issues 联系我们。
