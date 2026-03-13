<h1 align="center">NOFX</h1>

<p align="center">
  <strong>你的个人 AI 交易助手。</strong><br/>
  <strong>任何市场。任何模型。用 USDC 付费，无需 API Key。</strong>
</p>

<p align="center">
  <a href="https://github.com/woodyyan/nofx/stargazers"><img src="https://img.shields.io/github/stars/woodyyan/nofx?style=for-the-badge" alt="Stars"></a>
  <a href="https://github.com/woodyyan/nofx/releases"><img src="https://img.shields.io/github/v/release/woodyyan/nofx?style=for-the-badge" alt="Release"></a>
  <a href="https://github.com/woodyyan/nofx/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-AGPL--3.0-blue.svg?style=for-the-badge" alt="License"></a>
  <a href="https://t.me/nofx_dev_community"><img src="https://img.shields.io/badge/Telegram-Community-blue?style=for-the-badge&logo=telegram" alt="Telegram"></a>
</p>

<p align="center">
  <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go" alt="Go"></a>
  <a href="https://reactjs.org/"><img src="https://img.shields.io/badge/React-18+-61DAFB?style=flat&logo=react" alt="React"></a>
  <a href="https://x402.org"><img src="https://img.shields.io/badge/x402-USDC%20Payments-2775CA?style=flat" alt="x402"></a>
  <a href="https://claw402.ai"><img src="https://img.shields.io/badge/Claw402-AI%20Gateway-FF6B35?style=flat" alt="Claw402"></a>
  <a href="https://blockrun.ai"><img src="https://img.shields.io/badge/BlockRun-x402%20Provider-8B5CF6?style=flat" alt="BlockRun"></a>
</p>

<p align="center">
  <a href="../../../README.md">English</a> ·
  <a href="README.md">中文</a> ·
  <a href="../ja/README.md">日本語</a> ·
  <a href="../ko/README.md">한국어</a> ·
  <a href="../ru/README.md">Русский</a> ·
  <a href="../uk/README.md">Українська</a> ·
  <a href="../vi/README.md">Tiếng Việt</a>
</p>

> **语言声明：** 本中文版本文档仅为方便海外华人社区阅读而提供，不代表本软件面向中国大陆、香港、澳门或台湾地区用户开放。如您位于上述地区，请勿使用本软件。

---

NOFX 是一个开源的**自主式** AI 交易助手。与需要手动配置模型、管理 API Key、接入数据源的传统 AI 工具不同 —— NOFX 的 AI **自主感知市场、自选模型、自动获取数据**。零人工干预。你只需设定策略，AI 负责一切。

**完全自主**：AI 自行决定使用哪个模型、获取什么市场数据、何时交易。无需手动配置模型，无需管理各种服务的 API Key。只需充值 USDC 钱包，一键启动。

核心差异：**内置 [x402](https://x402.org) 微支付协议**。无需 API Key，充值 USDC 钱包即可按需付费。钱包就是你的身份。

```bash
curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/main/install.sh | bash
```

打开 **http://127.0.0.1:3000**，完成。

---

## x402 如何工作

传统流程：注册账号 → 购买额度 → 获取 API Key → 管理配额 → 轮换密钥。

x402 流程：

```
请求 → 402（返回价格）→ 钱包签名 USDC → 重试 → 完成
```

无需注册。无需 API Key。无需预付费。一个钱包，所有模型。

### 内置 x402 提供商

| 提供商 | 链 | 模型 |
|:---------|:------|:-------|
| <img src="../../../web/public/icons/claw402.png" width="20" height="20" style="vertical-align: middle;"/> **[Claw402](https://claw402.ai)** | Base | GPT-5.4、Claude Opus、DeepSeek、Qwen、Grok、Gemini、Kimi — 15+ 模型 |
| **[BlockRun](https://blockrun.ai)** | Base | 可配置 |
| **[BlockRun Sol](https://sol.blockrun.ai)** | Solana | 可配置 |

同时兼容 **[ClawRouter](https://github.com/BlockRunAI/ClawRouter)** —— 智能 LLM 路由，自动选择每次请求最便宜的模型（41+ 模型，节省 74-100%，<1ms 路由）。

---

## 功能概览

| 功能 | 描述 |
|:--------|:------------|
| **多 AI** | DeepSeek、Qwen、GPT、Claude、Gemini、Grok、Kimi — 随时切换 |
| **多交易所** | Binance、Bybit、OKX、Bitget、KuCoin、Gate、Hyperliquid、Aster、Lighter |
| **策略工作室** | 可视化构建器 — 币种来源、指标、风控 |
| **AI 竞赛** | AI 实时竞争，排行榜排名 |
| **Telegram Agent** | 与交易助手对话 — 流式输出、工具调用、记忆 |
| **回测实验室** | 历史模拟，权益曲线和性能指标 |
| **仪表板** | 实时持仓、盈亏、AI 决策日志与思维链 |

### 市场

加密货币 · 美股 · 外汇 · 贵金属

### 交易所 (CEX)

| 交易所 | 状态 | 注册 (手续费折扣) |
|:---------|:------:|:------------------------|
| <img src="../../../web/public/exchange-icons/binance.jpg" width="20" height="20" style="vertical-align: middle;"/> **Binance** | ✅ | [注册](https://www.binance.com/join?ref=NOFXENG) |
| <img src="../../../web/public/exchange-icons/bybit.png" width="20" height="20" style="vertical-align: middle;"/> **Bybit** | ✅ | [注册](https://partner.bybit.com/b/83856) |
| <img src="../../../web/public/exchange-icons/okx.svg" width="20" height="20" style="vertical-align: middle;"/> **OKX** | ✅ | [注册](https://www.okx.com/join/1865360) |
| <img src="../../../web/public/exchange-icons/bitget.svg" width="20" height="20" style="vertical-align: middle;"/> **Bitget** | ✅ | [注册](https://www.bitget.com/referral/register?from=referral&clacCode=c8a43172) |
| <img src="../../../web/public/exchange-icons/kucoin.svg" width="20" height="20" style="vertical-align: middle;"/> **KuCoin** | ✅ | [注册](https://www.kucoin.com/r/broker/CXEV7XKK) |
| <img src="../../../web/public/exchange-icons/gate.svg" width="20" height="20" style="vertical-align: middle;"/> **Gate** | ✅ | [注册](https://www.gatenode.xyz/share/VQBGUAxY) |

### 交易所 (Perp-DEX)

| 交易所 | 状态 | 注册 (手续费折扣) |
|:---------|:------:|:------------------------|
| <img src="../../../web/public/exchange-icons/hyperliquid.png" width="20" height="20" style="vertical-align: middle;"/> **Hyperliquid** | ✅ | [注册](https://app.hyperliquid.xyz/join/AITRADING) |
| <img src="../../../web/public/exchange-icons/aster.svg" width="20" height="20" style="vertical-align: middle;"/> **Aster DEX** | ✅ | [注册](https://www.asterdex.com/en/referral/fdfc0e) |
| <img src="../../../web/public/exchange-icons/lighter.png" width="20" height="20" style="vertical-align: middle;"/> **Lighter** | ✅ | [注册](https://app.lighter.xyz/?referral=68151432) |

### AI 模型 (API Key 模式)

| AI 模型 | 状态 | 获取 API Key |
|:---------|:------:|:------------|
| <img src="../../../web/public/icons/deepseek.svg" width="20" height="20" style="vertical-align: middle;"/> **DeepSeek** | ✅ | [获取 API Key](https://platform.deepseek.com) |
| <img src="../../../web/public/icons/qwen.svg" width="20" height="20" style="vertical-align: middle;"/> **通义千问** | ✅ | [获取 API Key](https://dashscope.console.aliyun.com) |
| <img src="../../../web/public/icons/openai.svg" width="20" height="20" style="vertical-align: middle;"/> **OpenAI (GPT)** | ✅ | [获取 API Key](https://platform.openai.com) |
| <img src="../../../web/public/icons/claude.svg" width="20" height="20" style="vertical-align: middle;"/> **Claude** | ✅ | [获取 API Key](https://console.anthropic.com) |
| <img src="../../../web/public/icons/gemini.svg" width="20" height="20" style="vertical-align: middle;"/> **Gemini** | ✅ | [获取 API Key](https://aistudio.google.com) |
| <img src="../../../web/public/icons/grok.svg" width="20" height="20" style="vertical-align: middle;"/> **Grok** | ✅ | [获取 API Key](https://console.x.ai) |
| <img src="../../../web/public/icons/kimi.svg" width="20" height="20" style="vertical-align: middle;"/> **Kimi** | ✅ | [获取 API Key](https://platform.moonshot.cn) |

### AI 模型 (x402 模式 — 无需 API Key)

15+ 模型通过 [Claw402](https://claw402.ai) 或 [BlockRun](https://blockrun.ai) 接入 — 只需一个 USDC 钱包

---

## 安装

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/main/install.sh | bash
```

### Railway (云部署)

[![Deploy on Railway](https://railway.com/button.svg)](https://railway.com/deploy/nofx?referralCode=nofx)

### Docker

```bash
# 下载并启动
curl -O https://raw.githubusercontent.com/woodyyan/nofx/main/docker-compose.prod.yml
docker compose -f docker-compose.prod.yml up -d
```

### Windows

安装 [Docker Desktop](https://www.docker.com/products/docker-desktop/)，然后：

```powershell
curl -o docker-compose.prod.yml https://raw.githubusercontent.com/woodyyan/nofx/main/docker-compose.prod.yml
docker compose -f docker-compose.prod.yml up -d
```

### 从源码构建

```bash
# 前置条件: Go 1.21+, Node.js 18+, TA-Lib
# macOS: brew install ta-lib
# Ubuntu: sudo apt-get install libta-lib0-dev

git clone https://github.com/woodyyan/nofx.git && cd nofx
go build -o nofx && ./nofx          # 后端
cd web && npm install && npm run dev  # 前端 (新终端)
```

### 更新

```bash
curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/main/install.sh | bash
```

此命令会拉取最新官方镜像并自动重启服务。

### 手动安装 (开发者)

#### 前置条件

- **Go 1.21+**
- **Node.js 18+**
- **TA-Lib** (技术指标库)

```bash
# 安装 TA-Lib
# macOS
brew install ta-lib

# Ubuntu/Debian
sudo apt-get install libta-lib0-dev
```

#### 安装步骤

```bash
# 1. 克隆仓库
git clone https://github.com/woodyyan/nofx.git
cd nofx

# 2. 安装后端依赖
go mod download

# 3. 安装前端依赖
cd web
npm install
cd ..

# 4. 构建并启动后端
go build -o nofx
./nofx

# 5. 启动前端 (新终端)
cd web
npm run dev
```

访问 Web 界面: **http://127.0.0.1:3000**

---

## 配置

1. **AI** — 添加 API Key 或配置 x402 钱包
2. **交易所** — 连接交易所 API 凭证
3. **策略** — 在策略工作室构建
4. **交易员** — 组合 AI + 交易所 + 策略
5. **交易** — 从仪表板启动

1. **安装 Docker Desktop**
   - 从 [docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop/) 下载
   - 运行安装程序并重启电脑
   - 启动 Docker Desktop 并等待就绪

2. **运行 NOFX**
   ```powershell
   # 打开 PowerShell 运行：
   curl -o docker-compose.prod.yml https://raw.githubusercontent.com/woodyyan/nofx/main/docker-compose.prod.yml
   docker compose -f docker-compose.prod.yml up -d
   ```

3. **访问**：在浏览器打开 **http://127.0.0.1:3000**

### 方法二：WSL2（适合开发）

1. **安装 WSL2**
   ```powershell
   # 以管理员身份打开 PowerShell
   wsl --install
   ```
   安装完成后重启电脑。

2. **从 Microsoft Store 安装 Ubuntu**
   - 打开 Microsoft Store
   - 搜索 "Ubuntu 22.04" 并安装
   - 启动 Ubuntu 并设置用户名/密码

3. **在 WSL2 中安装依赖**
   ```bash
   # 更新系统
   sudo apt update && sudo apt upgrade -y

   # 安装 Go
   wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
   source ~/.bashrc

   # 安装 Node.js
   curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
   sudo apt-get install -y nodejs

   # 安装 TA-Lib
   sudo apt-get install -y libta-lib0-dev

   # 安装 Git
   sudo apt-get install -y git
   ```

4. **克隆并运行 NOFX**
   ```bash
   git clone https://github.com/woodyyan/nofx.git
   cd nofx

   # 构建并运行后端
   go build -o nofx && ./nofx

   # 在另一个终端运行前端
   cd web && npm install && npm run dev
   ```

5. **访问**：在 Windows 浏览器打开 **http://127.0.0.1:3000**

### 方法三：WSL2 + Docker（两全其美）

1. **安装 Docker Desktop 并启用 WSL2 后端**
   - Docker Desktop 安装时勾选 "Use WSL 2 based engine"
   - 在 Docker Desktop 设置 → Resources → WSL Integration 中启用你的 Linux 发行版

2. **在 WSL2 终端运行**
   ```bash
   curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/main/install.sh | bash
   ```

---

## 服务器部署

### 快速部署 (HTTP/IP 访问)

默认情况下，传输加密已**禁用**，可直接通过 IP 地址访问 NOFX：

```bash
# 部署到你的服务器
curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/main/install.sh | bash
```

通过 `http://你的服务器IP:3000` 访问 - 立即可用。

### 增强安全 (HTTPS)

如需增强安全性，在 `.env` 中启用传输加密：

```bash
TRANSPORT_ENCRYPTION=true
```

启用后，浏览器会使用 Web Crypto API 在传输前加密 API 密钥。此功能需要：
- `https://` - 任何有 SSL 证书的域名
- `http://localhost` - 本地开发

### Cloudflare 快速配置 HTTPS

1. **添加域名到 Cloudflare** (免费计划即可)
   - 访问 [dash.cloudflare.com](https://dash.cloudflare.com)
   - 添加域名并更新 DNS 服务器

2. **创建 DNS 记录**
   - 类型: `A`
   - 名称: `nofx` (或你的子域名)
   - 内容: 你的服务器 IP
   - 代理状态: **已代理** (橙色云朵)

3. **配置 SSL/TLS**
   - 进入 SSL/TLS 设置
   - 加密模式选择 **灵活**

   ```
   用户 ──[HTTPS]──→ Cloudflare ──[HTTP]──→ 你的服务器:3000
   ```

4. **启用传输加密**
   ```bash
   # 编辑 .env 并设置
   TRANSPORT_ENCRYPTION=true
   ```

5. **完成！** 通过 `https://nofx.你的域名.com` 访问

---

## 初始配置 (Web 界面)

启动系统后，通过 Web 界面进行配置:

1. **配置 AI 模型** - 添加你的 AI API 密钥 (DeepSeek, OpenAI 等)
2. **配置交易所** - 设置交易所 API 凭证
3. **创建策略** - 在策略工作室配置交易策略
4. **创建交易员** - 组合 AI 模型 + 交易所 + 策略
5. **开始交易** - 启动你配置的交易员

所有配置都通过 Web 界面完成 - 无需编辑 JSON 文件。

---

## Web 界面功能

### 竞赛页面
- 实时 ROI 排行榜
- 多 AI 性能对比图表
- 实时盈亏追踪和排名

### 仪表板
- TradingView 风格 K 线图
- 实时持仓管理
- AI 决策日志与思维链推理
- 权益曲线追踪

### 策略工作室
- 币种来源配置 (静态列表、AI500 池、OI Top)
- 技术指标 (EMA, MACD, RSI, ATR, 成交量, OI, 资金费率)
- 风控设置 (杠杆、仓位限制、保证金使用率)
- AI 测试与实时提示词预览

---

## 常见问题

### TA-Lib 未找到
```bash
# macOS
brew install ta-lib

# Ubuntu
sudo apt-get install libta-lib0-dev
```

### AI API 超时
- 检查 API 密钥是否正确
- 检查网络连接
- 系统超时时间为 120 秒

### 前端无法连接后端
- 确保后端运行在 http://localhost:8080
- 检查端口是否被占用

---

## 文档

| | |
|:--|:--|
| [架构概览](../../architecture/README.md) | 系统设计和模块索引 |
| [策略模块](../../architecture/STRATEGY_MODULE.md) | 币种选择、AI 提示词、执行 |
| [回测模块](../../architecture/BACKTEST_MODULE.md) | 历史模拟、指标计算 |
| [常见问题](../../faq/README.md) | FAQ |
| [快速开始](../../getting-started/README.md) | 部署指南 |

---

## 贡献

查看 [贡献指南](../../../CONTRIBUTING.md) · [行为准则](../../../CODE_OF_CONDUCT.md) · [安全政策](../../../SECURITY.md)

### 贡献者空投计划

所有贡献在 GitHub 上追踪。当 NOFX 产生收入时，贡献者将获得空投。

**解决 [置顶 Issue](https://github.com/woodyyan/nofx/issues) 的 PR 获得最高奖励！**

| 贡献类型 | 权重 |
|:-------------|:------:|
| 置顶 Issue PR | ★★★★★★ |
| 代码提交 (合并的 PR) | ★★★★★ |
| Bug 修复 | ★★★★ |
| 功能建议 | ★★★ |
| Bug 报告 | ★★ |
| 文档 | ★★ |

---

## 链接

- **GitHub Issues**: [提交 Issue](https://github.com/woodyyan/nofx/issues)
- **开发者社区**: [Telegram 群组](https://t.me/nofx_dev_community)

---

## License

[AGPL-3.0](../../../LICENSE)

[![Star History Chart](https://api.star-history.com/svg?repos=woodyyan/nofx&type=Date)](https://star-history.com/#woodyyan/nofx&Date)
