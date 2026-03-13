<h1 align="center">NOFX</h1>

<p align="center">
  <strong>Trợ lý giao dịch AI cá nhân của bạn.</strong><br/>
  <strong>Mọi thị trường. Mọi mô hình. Thanh toán USDC, không cần API key.</strong>
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
  <a href="../zh-CN/README.md">中文</a> ·
  <a href="../ja/README.md">日本語</a> ·
  <a href="../ko/README.md">한국어</a> ·
  <a href="../ru/README.md">Русский</a> ·
  <a href="../uk/README.md">Українська</a> ·
  <a href="README.md">Tiếng Việt</a>
</p>

---

NOFX là trợ lý giao dịch AI **tự chủ** mã nguồn mở. Không giống các công cụ AI truyền thống yêu cầu bạn cấu hình mô hình thủ công, quản lý API key và kết nối nguồn dữ liệu — AI của NOFX **tự nhận diện thị trường, tự chọn mô hình và tự lấy dữ liệu**. Không cần con người can thiệp. Bạn chỉ cần đặt chiến lược, AI xử lý mọi thứ còn lại.

**Hoàn toàn tự chủ**: AI tự quyết định sử dụng mô hình nào, lấy dữ liệu thị trường gì, khi nào giao dịch. Không cần cấu hình mô hình thủ công. Không cần quản lý API key của nhiều dịch vụ. Chỉ cần nạp ví USDC và chạy.

### Liên Kết Chính Thức

- **Website Chính Thức**: [https://nofxai.com](https://nofxai.com)
- **Bảng Điều Khiển Dữ Liệu**: [https://nofxos.ai/dashboard](https://nofxos.ai/dashboard)
- **Tài Liệu API**: [https://nofxos.ai/api-docs](https://nofxos.ai/api-docs)

> **Cảnh Báo Rủi Ro**: Hệ thống này mang tính thử nghiệm. Giao dịch tự động AI có rủi ro đáng kể. Chỉ nên sử dụng cho mục đích học tập/nghiên cứu hoặc kiểm tra với số tiền nhỏ!

## Cộng Đồng Nhà Phát Triển

Tham gia cộng đồng Telegram: **[NOFX Developer Community](https://t.me/nofx_dev_community)**

---

## Trước Khi Bắt Đầu

Để sử dụng NOFX, bạn cần:

1. **Tài khoản sàn giao dịch** - Đăng ký trên sàn được hỗ trợ và tạo API key với quyền giao dịch
2. **API Key mô hình AI** - Lấy từ nhà cung cấp được hỗ trợ (khuyến nghị DeepSeek để tiết kiệm chi phí)

---

## Sàn Giao Dịch Được Hỗ Trợ

### CEX (Sàn Tập Trung)

| Sàn | Trạng thái | Đăng ký (Giảm phí) |
|----------|--------|-------------------------|
| **Binance** | ✅ Hỗ trợ | [Đăng ký](https://www.binance.com/join?ref=NOFXENG) |
| **Bybit** | ✅ Hỗ trợ | [Đăng ký](https://partner.bybit.com/b/83856) |
| **OKX** | ✅ Hỗ trợ | [Đăng ký](https://www.okx.com/join/1865360) |
| **Bitget** | ✅ Hỗ trợ | [Đăng ký](https://www.bitget.com/referral/register?from=referral&clacCode=c8a43172) |
| **KuCoin** | ✅ Hỗ trợ | [Đăng ký](https://www.kucoin.com/r/broker/CXEV7XKK) |
| **Gate** | ✅ Hỗ trợ | [Đăng ký](https://www.gatenode.xyz/share/VQBGUAxY) |

### Perp-DEX (Sàn Phi Tập Trung)

| Sàn | Trạng thái | Đăng ký (Giảm phí) |
|----------|--------|-------------------------|
| **Hyperliquid** | ✅ Hỗ trợ | [Đăng ký](https://app.hyperliquid.xyz/join/AITRADING) |
| **Aster DEX** | ✅ Hỗ trợ | [Đăng ký](https://www.asterdex.com/en/referral/fdfc0e) |
| **Lighter** | ✅ Hỗ trợ | [Đăng ký](https://app.lighter.xyz/?referral=68151432) |

---

## Mô Hình AI Được Hỗ Trợ

| Mô hình AI | Trạng thái | Lấy API Key |
|----------|--------|-------------|
| **DeepSeek** | ✅ Hỗ trợ | [Lấy API Key](https://platform.deepseek.com) |
| **Qwen** | ✅ Hỗ trợ | [Lấy API Key](https://dashscope.console.aliyun.com) |
| **OpenAI (GPT)** | ✅ Hỗ trợ | [Lấy API Key](https://platform.openai.com) |
| **Claude** | ✅ Hỗ trợ | [Lấy API Key](https://console.anthropic.com) |
| **Gemini** | ✅ Hỗ trợ | [Lấy API Key](https://aistudio.google.com) |
| **Grok** | ✅ Hỗ trợ | [Lấy API Key](https://console.x.ai) |
| **Kimi** | ✅ Hỗ trợ | [Lấy API Key](https://platform.moonshot.cn) |

---

## Bắt Đầu Nhanh

### Tùy chọn 1: Triển khai Docker (Khuyến nghị)

```bash
git clone https://github.com/woodyyan/nofx.git
cd nofx
chmod +x ./start.sh
./start.sh start --build
```

Truy cập giao diện Web: **http://localhost:3000**

### Cập Nhật Phiên Bản Mới

> **💡 Cập nhật thường xuyên.** Chạy lệnh này hàng ngày để nhận các tính năng và bản sửa lỗi mới nhất:

```bash
curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/dev/install.sh | bash
```

Mở **http://127.0.0.1:3000**. Xong.

---

## x402 hoạt động như thế nào

Quy trình truyền thống: đăng ký tài khoản → mua credits → lấy API key → quản lý quota → xoay key.

Quy trình x402:

```
Yêu cầu → 402 (đây là giá) → ví ký USDC → thử lại → xong
```

Không tài khoản. Không API key. Không trả trước. Một ví, tất cả mô hình.

### Nhà cung cấp x402 tích hợp

| Nhà cung cấp | Chain | Mô hình |
|:---------|:------|:-------|
| <img src="../../../web/public/icons/claw402.png" width="20" height="20" style="vertical-align: middle;"/> **[Claw402](https://claw402.ai)** | Base | GPT-5.4, Claude Opus, DeepSeek, Qwen, Grok, Gemini, Kimi — 15+ mô hình |
| **[BlockRun](https://blockrun.ai)** | Base | Có thể cấu hình |
| **[BlockRun Sol](https://sol.blockrun.ai)** | Solana | Có thể cấu hình |

Tương thích với **[ClawRouter](https://github.com/BlockRunAI/ClawRouter)** — bộ định tuyến LLM thông minh tự động chọn mô hình rẻ nhất (41+ mô hình, tiết kiệm 74-100%, <1ms định tuyến).

---

## Tính năng

| Tính năng | Mô tả |
|:--------|:------------|
| **Đa AI** | DeepSeek, Qwen, GPT, Claude, Gemini, Grok, Kimi — chuyển đổi bất cứ lúc nào |
| **Đa Sàn** | Binance, Bybit, OKX, Bitget, KuCoin, Gate, Hyperliquid, Aster, Lighter |
| **Strategy Studio** | Trình xây dựng trực quan — nguồn coin, chỉ báo, kiểm soát rủi ro |
| **AI Competition** | AI cạnh tranh thời gian thực, bảng xếp hạng hiệu suất |
| **Telegram Agent** | Chat với trợ lý giao dịch — streaming, gọi công cụ, bộ nhớ |
| **Backtest Lab** | Mô phỏng lịch sử, đường vốn và chỉ số hiệu suất |
| **Dashboard** | Vị thế trực tiếp, P/L, nhật ký quyết định AI với Chain of Thought |

### Thị trường

Crypto · Cổ phiếu Mỹ · Forex · Kim loại

### Sàn giao dịch (CEX)

| Sàn | Trạng thái | Đăng ký (Giảm phí) |
|:---------|:------:|:------------------------|
| <img src="../../../web/public/exchange-icons/binance.jpg" width="20" height="20" style="vertical-align: middle;"/> **Binance** | ✅ | [Đăng ký](https://www.binance.com/join?ref=NOFXENG) |
| <img src="../../../web/public/exchange-icons/bybit.png" width="20" height="20" style="vertical-align: middle;"/> **Bybit** | ✅ | [Đăng ký](https://partner.bybit.com/b/83856) |
| <img src="../../../web/public/exchange-icons/okx.svg" width="20" height="20" style="vertical-align: middle;"/> **OKX** | ✅ | [Đăng ký](https://www.okx.com/join/1865360) |
| <img src="../../../web/public/exchange-icons/bitget.svg" width="20" height="20" style="vertical-align: middle;"/> **Bitget** | ✅ | [Đăng ký](https://www.bitget.com/referral/register?from=referral&clacCode=c8a43172) |
| <img src="../../../web/public/exchange-icons/kucoin.svg" width="20" height="20" style="vertical-align: middle;"/> **KuCoin** | ✅ | [Đăng ký](https://www.kucoin.com/r/broker/CXEV7XKK) |
| <img src="../../../web/public/exchange-icons/gate.svg" width="20" height="20" style="vertical-align: middle;"/> **Gate** | ✅ | [Đăng ký](https://www.gatenode.xyz/share/VQBGUAxY) |

### Sàn giao dịch (Perp-DEX)

| Sàn | Trạng thái | Đăng ký (Giảm phí) |
|:---------|:------:|:------------------------|
| <img src="../../../web/public/exchange-icons/hyperliquid.png" width="20" height="20" style="vertical-align: middle;"/> **Hyperliquid** | ✅ | [Đăng ký](https://app.hyperliquid.xyz/join/AITRADING) |
| <img src="../../../web/public/exchange-icons/aster.svg" width="20" height="20" style="vertical-align: middle;"/> **Aster DEX** | ✅ | [Đăng ký](https://www.asterdex.com/en/referral/fdfc0e) |
| <img src="../../../web/public/exchange-icons/lighter.png" width="20" height="20" style="vertical-align: middle;"/> **Lighter** | ✅ | [Đăng ký](https://app.lighter.xyz/?referral=68151432) |

### Mô hình AI (Chế độ API Key)

| Mô hình AI | Trạng thái | Lấy API Key |
|:---------|:------:|:------------|
| <img src="../../../web/public/icons/deepseek.svg" width="20" height="20" style="vertical-align: middle;"/> **DeepSeek** | ✅ | [Lấy API Key](https://platform.deepseek.com) |
| <img src="../../../web/public/icons/qwen.svg" width="20" height="20" style="vertical-align: middle;"/> **Qwen** | ✅ | [Lấy API Key](https://dashscope.console.aliyun.com) |
| <img src="../../../web/public/icons/openai.svg" width="20" height="20" style="vertical-align: middle;"/> **OpenAI (GPT)** | ✅ | [Lấy API Key](https://platform.openai.com) |
| <img src="../../../web/public/icons/claude.svg" width="20" height="20" style="vertical-align: middle;"/> **Claude** | ✅ | [Lấy API Key](https://console.anthropic.com) |
| <img src="../../../web/public/icons/gemini.svg" width="20" height="20" style="vertical-align: middle;"/> **Gemini** | ✅ | [Lấy API Key](https://aistudio.google.com) |
| <img src="../../../web/public/icons/grok.svg" width="20" height="20" style="vertical-align: middle;"/> **Grok** | ✅ | [Lấy API Key](https://console.x.ai) |
| <img src="../../../web/public/icons/kimi.svg" width="20" height="20" style="vertical-align: middle;"/> **Kimi** | ✅ | [Lấy API Key](https://platform.moonshot.cn) |

### Mô hình AI (Chế độ x402 — Không cần API Key)

15+ mô hình qua [Claw402](https://claw402.ai) hoặc [BlockRun](https://blockrun.ai) — chỉ cần ví USDC

---

## Cài đặt

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/dev/install.sh | bash
```

### Railway (Cloud)

[![Deploy on Railway](https://railway.com/button.svg)](https://railway.com/deploy/nofx?referralCode=nofx)

### Docker

```bash
curl -O https://raw.githubusercontent.com/woodyyan/nofx/main/docker-compose.prod.yml
docker compose -f docker-compose.prod.yml up -d
```

### Từ mã nguồn

```bash
# Yêu cầu: Go 1.21+, Node.js 18+, TA-Lib
# macOS: brew install ta-lib
# Ubuntu: sudo apt-get install libta-lib0-dev

# Cài đặt TA-Lib (macOS)
brew install ta-lib

# Clone và thiết lập
git clone https://github.com/woodyyan/nofx.git
cd nofx
go mod download
cd web && npm install && cd ..

# Khởi động backend
go build -o nofx && ./nofx

# Khởi động frontend (terminal mới)
cd web && npm run dev
```

---

## Liên kết

| | |
|:--|:--|
| Website | [nofxai.com](https://nofxai.com) |
| Dashboard | [nofxos.ai/dashboard](https://nofxos.ai/dashboard) |
| API Docs | [nofxos.ai/api-docs](https://nofxos.ai/api-docs) |
| Telegram | [nofx_dev_community](https://t.me/nofx_dev_community) |
| Twitter | [@nofx_official](https://x.com/nofx_official) |

> **Cảnh báo rủi ro**: Giao dịch tự động AI có rủi ro đáng kể. Chỉ nên sử dụng cho mục đích học tập/nghiên cứu hoặc số tiền nhỏ.

---

## License

[AGPL-3.0](../../../LICENSE)

---

## Giấy Phép

**GNU Affero General Public License v3.0 (AGPL-3.0)**

---

## Liên Hệ

- **GitHub Issues**: [Gửi Issue](https://github.com/woodyyan/nofx/issues)
- **Cộng đồng Nhà phát triển**: [Nhóm Telegram](https://t.me/nofx_dev_community)
