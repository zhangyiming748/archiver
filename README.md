# Archiver

Archiver 是一个命令行工具，用于媒体文件管理和格式转换。支持视频转 H265 格式和图片转 AVIF 格式。

## ✨ 功能特性

- 🎬 **视频转换**：自动查找目录下的所有视频文件并转换为 H265 编码
- 📹 **MP4 转换**：将视频文件转换为 H265 MP4 格式
- 🧠 **智能压缩**：智能转换视频为更小的 H265 MP4 格式，节省存储空间
- 🔄 **视频旋转**：支持顺时针旋转视频 90° 或 270°
- 🖼️ **图片转换**：自动查找目录下的所有图片文件并转换为 AVIF 格式
- ⚡ **高效处理**：支持 FHD 模式和多线程处理，提供更高质量的视频转换
- 🔄 **强制覆盖**：支持强制覆盖已存在的文件，避免交互式确认
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

### 视频转换（H265）

将指定目录下的所有视频文件转换为 H265 格式：

```bash
# 基本用法
archiver video --dir /path/to/videos

# 简写形式
archiver video -d /path/to/videos

# 启用 FHD 模式（更高质量）
archiver video -d /path/to/videos --fhd
archiver video -d /path/to/videos -f

# 强制覆盖已存在的文件（避免 FFmpeg 询问）
archiver video -d /path/to/videos --force

# 组合使用多个参数
archiver video -d /path/to/videos --fhd --force
```

**参数说明：**

| 参数 | 简写 | 类型 | 默认值 | 说明 |
|------|------|------|--------|------|
| `--dir` | `-d` | string | 必需 | 要搜索视频文件的目录路径 |
| `--fhd` | `-f` | bool | false | 启用 FHD 模式进行视频转换 |
| `--force` | - | bool | false | 强制覆盖已存在的文件，避免 FFmpeg 交互式确认 |

**使用提示：**

- 如果不使用 `--force` 参数，当输出文件已存在时，FFmpeg 会询问是否覆盖。在非交互模式下可能导致转换失败并产生 0 字节临时文件
- 建议使用 `--force` 参数来自动覆盖已存在的文件，确保转换顺利进行
- FHD 模式会提供更高的视频质量，但转换时间会更长

### MP4 转换（H265 MP4）

将指定目录下的所有视频文件转换为 H265 MP4 格式：

```bash
# 基本用法
archiver mp4 --dir /path/to/videos

# 简写形式
archiver mp4 -d /path/to/videos

# 启用 FHD 模式
archiver mp4 -d /path/to/videos --fhd

# 强制覆盖已存在的文件
archiver mp4 -d /path/to/videos --force

# 组合使用多个参数
archiver mp4 -d /path/to/videos --fhd --force
```

**参数说明：**

| 参数 | 简写 | 类型 | 默认值 | 说明 |
|------|------|------|--------|------|
| `--dir` | `-d` | string | 必需 | 要搜索视频文件的目录路径 |
| `--fhd` | `-f` | bool | false | 启用 FHD 模式进行 MP4 转换 |
| `--force` | - | bool | false | 强制覆盖已存在的文件 |

### 智能压缩转换

智能转换视频为更小的 H265 MP4 格式，优化文件大小：

```bash
# 基本用法
archiver smart --dir /path/to/videos

# 简写形式
archiver smart -d /path/to/videos

# 启用 FHD 模式
archiver smart -d /path/to/videos --fhd

# 强制覆盖已存在的文件
archiver smart -d /path/to/videos --force

# 组合使用多个参数
archiver smart -d /path/to/videos --fhd --force
```

**参数说明：**

| 参数 | 简写 | 类型 | 默认值 | 说明 |
|------|------|------|--------|------|
| `--dir` | `-d` | string | 必需 | 要搜索视频文件的目录路径 |
| `--fhd` | `-f` | bool | false | 启用 FHD 模式进行智能转换 |
| `--force` | - | bool | false | 强制覆盖已存在的文件 |

### 视频旋转

旋转指定目录下的所有视频文件：

```bash
# 基本用法（默认顺时针旋转 90°）
archiver rotate --dir /path/to/videos

# 简写形式
archiver rotate -d /path/to/videos

# 指定旋转方向（90 或 270）
archiver rotate -d /path/to/videos --rotate 90
archiver rotate -d /path/to/videos -r 270

# 完整示例
archiver rotate -d /path/to/videos -r 90
```

**参数说明：**

| 参数 | 简写 | 类型 | 默认值 | 说明 |
|------|------|------|--------|------|
| `--dir` | `-d` | string | `./` | 要旋转视频的目录路径 |
| `--rotate` | `-r` | string | `90` | 旋转方向：90（顺时针90°）或 270（顺时针270°） |

**使用提示：**

- 旋转操作会直接修改原文件，建议先备份重要视频
- 支持自动检测硬件编码器（NVIDIA NVENC、Intel QSV、AMD AMF），优先使用硬件加速
- 如果没有硬件编码器，将使用 CPU 软件编码（libx264）

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

# 指定线程数（默认 4 线程）
archiver image -d /path/to/images --threads 8
archiver image -d /path/to/images -t 8

# 组合使用多个参数
archiver image -d /path/to/images --fhd --threads 8
```

**参数说明：**

| 参数 | 简写 | 类型 | 默认值 | 说明 |
|------|------|------|--------|------|
| `--dir` | `-d` | string | 必需 | 要搜索图片文件的目录路径 |
| `--fhd` | `-f` | bool | false | 启用 FHD 模式进行图片转换 |
| `--threads` | `-t` | int | 4 | 用于转换的线程数 |

## 📋 使用示例

```bash
# 转换 ~/Videos 目录下的所有视频为 H265
archiver video -d ~/Videos

# 转换视频为 H265 MP4 格式并启用 FHD 模式和强制覆盖
archiver mp4 -d ~/Videos --fhd --force

# 智能压缩视频以节省空间
archiver smart -d ~/Videos --force

# 顺时针旋转视频 90 度
archiver rotate -d ~/Videos -r 90

# 顺时针旋转视频 270 度
archiver rotate -d ~/Videos -r 270

# 转换 ~/Pictures 目录下的所有图片，使用 FHD 模式和 8 线程
archiver image -d ~/Pictures -f -t 8

# 查看当前版本
archiver version

# Windows 示例：转换 D:\AI\舞蹈 目录下的视频
archiver video -d "D:\AI\舞蹈" --force
```

## ❓ 常见问题

### Q1: 为什么会出现 "File already exists. Overwrite?" 错误？

**原因**：FFmpeg 在转换时，如果输出文件已存在，默认会询问是否覆盖。在非交互模式下无法接收用户输入，导致转换失败。

**解决方案**：使用 `--force` 参数自动覆盖已存在的文件：
```bash
archiver video -d /path/to/videos --force
```

### Q2: 为什么会生成 0 字节的临时文件？

**原因**：当 FFmpeg 询问是否覆盖文件时，由于是非交互模式，默认回答 "N"（不覆盖），导致转换失败但临时文件已被创建。

**解决方案**：同样使用 `--force` 参数来避免这个问题。

### Q3: 支持哪些视频格式？

Archiver 支持以下视频格式：
- 标准视频格式：MP4, AVI, MKV, MOV, WMV, FLV 等（通过 filetype 库识别）
- 特殊格式：RMVB, RM, VOB, FLV, TS, M2TS（通过后缀名识别）

### Q4: 支持哪些图片格式？

支持所有常见图片格式，包括：
- JPEG/JPG, PNG, GIF, BMP, TIFF, WebP 等

### Q5: 转换后的文件命名规则是什么？

- **视频（video/mp4/smart）**：根据具体命令不同，会在原文件名后添加相应后缀
  - `video` 命令：添加 `_h265` 后缀，例如 `video.mp4` → `video_h265.mp4`
  - `mp4` 命令：转换为 H265 MP4 格式
  - `smart` 命令：智能压缩为更小的 H265 MP4 格式
- **图片**：原文件名添加 `_avif` 后缀，例如 `photo.jpg` → `photo_avif.avif`
- **旋转**：直接替换原文件，不创建新文件

### Q6: 视频旋转支持哪些角度？

目前支持两种旋转角度：
- **90°**：顺时针旋转 90 度（`-r 90`）
- **270°**：顺时针旋转 270 度（`-r 270`）

旋转操作会自动检测并使用可用的硬件编码器（NVIDIA、Intel、AMD）进行加速。

### Q7: 如何选择合适的转换命令？

- **`video`**：通用 H265 转换，适合大多数场景
- **`mp4`**：转换为标准 H265 MP4 格式，兼容性更好
- **`smart`**：智能压缩模式，在保持画质的前提下尽可能减小文件大小
- **`rotate`**：仅旋转视频，不进行格式转换

## 🔧 技术栈

- **语言**：Go 1.26+
- **CLI 框架**：Cobra v1.8.0
- **文件类型检测**：filetype v1.1.3
- **视频处理**：FFmpeg (H265/HEVC 编码)
- **图片处理**：libavif (AVIF 编码)
- **外部依赖**：github.com/zhangyiming748/archive

## 📝 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📮 反馈

如有问题或建议，请通过 GitHub Issues 联系我们。
