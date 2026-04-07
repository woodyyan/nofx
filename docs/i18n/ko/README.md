<h1 align="center">NOFX</h1>

<p align="center">
  <strong>당신만의 AI 트레이딩 어시스턴트.</strong><br/>
  <strong>모든 시장. 모든 모델. API 키 없이 USDC로 결제.</strong>
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
</p>

<p align="center">
  <a href="../../../README.md">English</a> ·
  <a href="../zh-CN/README.md">中文</a> ·
  <a href="../ja/README.md">日本語</a> ·
  <a href="README.md">한국어</a> ·
  <a href="../ru/README.md">Русский</a> ·
  <a href="../uk/README.md">Українська</a> ·
  <a href="../vi/README.md">Tiếng Việt</a>
</p>

---

NOFX는 오픈소스 **자율형** AI 트레이딩 어시스턴트입니다. 수동으로 모델을 설정하고, API 키를 관리하고, 데이터 소스를 연결해야 하는 기존 AI 도구와 달리 — NOFX의 AI는 **시장을 스스로 인식하고, 모델을 스스로 선택하고, 데이터를 스스로 가져옵니다**. 인간 개입 제로. 전략만 설정하면 나머지는 AI가 처리합니다.

**완전 자율**: AI가 어떤 모델을 사용할지, 어떤 시장 데이터를 가져올지, 언제 거래할지를 스스로 결정합니다. 수동 모델 설정 불필요. 여러 서비스의 API 키 관리 불필요. USDC 지갑에 충전하고 실행하기만 하면 됩니다.

### 공식 링크

- **공식 웹사이트**: [https://nofxai.com](https://nofxai.com)
- **데이터 대시보드**: [https://nofxos.ai/dashboard](https://nofxos.ai/dashboard)
- **API 문서**: [https://nofxos.ai/api-docs](https://nofxos.ai/api-docs)

> **위험 경고**: 이 시스템은 실험적입니다. AI 자동 거래에는 상당한 위험이 있습니다. 학습/연구 목적 또는 소액 테스트만 강력히 권장합니다!

## 개발자 커뮤니티

Telegram 개발자 커뮤니티 참여: **[NOFX 개발자 커뮤니티](https://t.me/nofx_dev_community)**

---

## 시작하기 전에

NOFX를 사용하려면 다음이 필요합니다:

1. **거래소 계정** - 지원되는 거래소에 등록하고 거래 권한이 있는 API 자격 증명 생성
2. **AI 모델 API 키** - 지원되는 제공업체에서 획득 (비용 효율성을 위해 DeepSeek 권장)

---

## 지원 거래소

### CEX (중앙화 거래소)

| 거래소 | 상태 | 등록 (수수료 할인) |
|----------|--------|-------------------------|
| **Binance** | ✅ 지원 | [등록](https://www.binance.com/join?ref=NOFXENG) |
| **Bybit** | ✅ 지원 | [등록](https://partner.bybit.com/b/83856) |
| **OKX** | ✅ 지원 | [등록](https://www.okx.com/join/1865360) |
| **Bitget** | ✅ 지원 | [등록](https://www.bitget.com/referral/register?from=referral&clacCode=c8a43172) |
| **KuCoin** | ✅ 지원 | [등록](https://www.kucoin.com/r/broker/CXEV7XKK) |
| **Gate** | ✅ 지원 | [등록](https://www.gatenode.xyz/share/VQBGUAxY) |

### Perp-DEX (탈중앙화 영구 선물 거래소)

| 거래소 | 상태 | 등록 (수수료 할인) |
|----------|--------|-------------------------|
| **Hyperliquid** | ✅ 지원 | [등록](https://app.hyperliquid.xyz/join/AITRADING) |
| **Aster DEX** | ✅ 지원 | [등록](https://www.asterdex.com/en/referral/fdfc0e) |
| **Lighter** | ✅ 지원 | [등록](https://app.lighter.xyz/?referral=68151432) |

---

## 지원 AI 모델

| AI 모델 | 상태 | API 키 받기 |
|----------|--------|-------------|
| **DeepSeek** | ✅ 지원 | [API 키 받기](https://platform.deepseek.com) |
| **Qwen** | ✅ 지원 | [API 키 받기](https://dashscope.console.aliyun.com) |
| **OpenAI (GPT)** | ✅ 지원 | [API 키 받기](https://platform.openai.com) |
| **Claude** | ✅ 지원 | [API 키 받기](https://console.anthropic.com) |
| **Gemini** | ✅ 지원 | [API 키 받기](https://aistudio.google.com) |
| **Grok** | ✅ 지원 | [API 키 받기](https://console.x.ai) |
| **Kimi** | ✅ 지원 | [API 키 받기](https://platform.moonshot.cn) |

---

## 빠른 시작

### 옵션 1: Docker 배포 (권장)

```bash
git clone https://github.com/woodyyan/nofx.git
cd nofx
chmod +x ./start.sh
./start.sh start --build
```

웹 인터페이스 접속: **http://localhost:3000**

### 최신 버전 유지

> **💡 업데이트가 빈번합니다.** 최신 기능과 수정 사항을 받으려면 매일 이 명령을 실행하세요:

```bash
curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/dev/install.sh | bash
```

**http://127.0.0.1:3000** 을 열면 완료.

---

## x402 작동 방식

기존 플로우: 계정 등록 → 크레딧 구매 → API 키 받기 → 쿼터 관리 → 키 교체.

x402 플로우:

```
요청 → 402 (가격 제시) → 지갑이 USDC 서명 → 재시도 → 완료
```

계정 불필요. API 키 불필요. 선불 크레딧 불필요. 지갑 하나로 모든 모델.

### 내장 x402 프로바이더

| 프로바이더 | 체인 | 모델 |
|:---------|:------|:-------|
| <img src="../../../web/public/icons/claw402.png" width="20" height="20" style="vertical-align: middle;"/> **[Claw402](https://claw402.ai)** | Base | GPT-5.4, Claude Opus, DeepSeek, Qwen, Grok, Gemini, Kimi — 15+ 모델 |

---

## 기능

| 기능 | 설명 |
|:--------|:------------|
| **멀티 AI** | DeepSeek, Qwen, GPT, Claude, Gemini, Grok, Kimi, MiniMax — 언제든 전환 |
| **멀티 거래소** | Binance, Bybit, OKX, Bitget, KuCoin, Gate, Hyperliquid, Aster, Lighter |
| **전략 스튜디오** | 비주얼 빌더 — 코인 소스, 지표, 리스크 관리 |
| **AI 토론 아레나** | 여러 AI가 거래 토론 (강세 vs 약세 vs 분석가), 투표, 실행 |
| **AI 경쟁** | AI가 실시간 경쟁, 리더보드 순위 |
| **Telegram 에이전트** | 트레이딩 어시스턴트와 채팅 — 스트리밍, 도구 호출, 메모리 |
| **백테스트 랩** | 과거 시뮬레이션, 자산 곡선 및 성과 지표 |
| **대시보드** | 실시간 포지션, 손익, Chain of Thought AI 결정 로그 |

### 시장

암호화폐 · 미국 주식 · 외환 · 귀금속

### 거래소 (CEX)

| 거래소 | 상태 | 등록 (수수료 할인) |
|:---------|:------:|:------------------------|
| <img src="../../../web/public/exchange-icons/binance.jpg" width="20" height="20" style="vertical-align: middle;"/> **Binance** | ✅ | [등록](https://www.binance.com/join?ref=NOFXENG) |
| <img src="../../../web/public/exchange-icons/bybit.png" width="20" height="20" style="vertical-align: middle;"/> **Bybit** | ✅ | [등록](https://partner.bybit.com/b/83856) |
| <img src="../../../web/public/exchange-icons/okx.svg" width="20" height="20" style="vertical-align: middle;"/> **OKX** | ✅ | [등록](https://www.okx.com/join/1865360) |
| <img src="../../../web/public/exchange-icons/bitget.svg" width="20" height="20" style="vertical-align: middle;"/> **Bitget** | ✅ | [등록](https://www.bitget.com/referral/register?from=referral&clacCode=c8a43172) |
| <img src="../../../web/public/exchange-icons/kucoin.svg" width="20" height="20" style="vertical-align: middle;"/> **KuCoin** | ✅ | [등록](https://www.kucoin.com/r/broker/CXEV7XKK) |
| <img src="../../../web/public/exchange-icons/gate.svg" width="20" height="20" style="vertical-align: middle;"/> **Gate** | ✅ | [등록](https://www.gatenode.xyz/share/VQBGUAxY) |

### 거래소 (Perp-DEX)

| 거래소 | 상태 | 등록 (수수료 할인) |
|:---------|:------:|:------------------------|
| <img src="../../../web/public/exchange-icons/hyperliquid.png" width="20" height="20" style="vertical-align: middle;"/> **Hyperliquid** | ✅ | [등록](https://app.hyperliquid.xyz/join/AITRADING) |
| <img src="../../../web/public/exchange-icons/aster.svg" width="20" height="20" style="vertical-align: middle;"/> **Aster DEX** | ✅ | [등록](https://www.asterdex.com/en/referral/fdfc0e) |
| <img src="../../../web/public/exchange-icons/lighter.png" width="20" height="20" style="vertical-align: middle;"/> **Lighter** | ✅ | [등록](https://app.lighter.xyz/?referral=68151432) |

### AI 모델 (API 키 모드)

| AI 모델 | 상태 | API 키 받기 |
|:---------|:------:|:------------|
| <img src="../../../web/public/icons/deepseek.svg" width="20" height="20" style="vertical-align: middle;"/> **DeepSeek** | ✅ | [API 키 받기](https://platform.deepseek.com) |
| <img src="../../../web/public/icons/qwen.svg" width="20" height="20" style="vertical-align: middle;"/> **Qwen** | ✅ | [API 키 받기](https://dashscope.console.aliyun.com) |
| <img src="../../../web/public/icons/openai.svg" width="20" height="20" style="vertical-align: middle;"/> **OpenAI (GPT)** | ✅ | [API 키 받기](https://platform.openai.com) |
| <img src="../../../web/public/icons/claude.svg" width="20" height="20" style="vertical-align: middle;"/> **Claude** | ✅ | [API 키 받기](https://console.anthropic.com) |
| <img src="../../../web/public/icons/gemini.svg" width="20" height="20" style="vertical-align: middle;"/> **Gemini** | ✅ | [API 키 받기](https://aistudio.google.com) |
| <img src="../../../web/public/icons/grok.svg" width="20" height="20" style="vertical-align: middle;"/> **Grok** | ✅ | [API 키 받기](https://console.x.ai) |
| <img src="../../../web/public/icons/kimi.svg" width="20" height="20" style="vertical-align: middle;"/> **Kimi** | ✅ | [API 키 받기](https://platform.moonshot.cn) |
| <img src="../../../web/public/icons/minimax.svg" width="20" height="20" style="vertical-align: middle;"/> **MiniMax** | ✅ | [API 키 받기](https://platform.minimaxi.com) |

### AI 모델 (x402 모드 — API 키 불필요)

15+ 모델을 [Claw402](https://claw402.ai)로 이용 — USDC 지갑만 있으면 됩니다

---

## 설치

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/dev/install.sh | bash
```

### Railway (클라우드)

[![Deploy on Railway](https://railway.com/button.svg)](https://railway.com/deploy/nofx?referralCode=nofx)

### Docker

```bash
curl -O https://raw.githubusercontent.com/woodyyan/nofx/main/docker-compose.prod.yml
docker compose -f docker-compose.prod.yml up -d
```

### 소스에서

```bash
# 필수 조건: Go 1.21+, Node.js 18+, TA-Lib
# macOS: brew install ta-lib

# TA-Lib 설치 (macOS)
brew install ta-lib

# 클론 및 설정
git clone https://github.com/woodyyan/nofx.git
cd nofx
go mod download
cd web && npm install && cd ..

# 백엔드 시작
go build -o nofx && ./nofx

# 프론트엔드 시작 (새 터미널)
cd web && npm run dev
```

---

## 링크

| | |
|:--|:--|
| 웹사이트 | [nofxai.com](https://nofxai.com) |
| 대시보드 | [nofxos.ai/dashboard](https://nofxos.ai/dashboard) |
| API 문서 | [nofxos.ai/api-docs](https://nofxos.ai/api-docs) |
| Telegram | [nofx_dev_community](https://t.me/nofx_dev_community) |
| Twitter | [@nofx_official](https://x.com/nofx_official) |

> **위험 경고**: AI 자동 거래에는 상당한 위험이 있습니다. 학습/연구 또는 소액 테스트만 권장합니다.

---

## License

1. 암호화폐 시장은 매우 변동성이 높음 - AI 결정이 수익을 보장하지 않음
2. 선물 거래는 레버리지 사용 - 손실이 원금을 초과할 수 있음
3. 극단적인 시장 상황에서 청산 위험 있음

---

## 서버 배포

### 빠른 배포 (IP를 통한 HTTP)

기본적으로 전송 암호화가 **비활성화**되어 HTTPS 없이 IP 주소를 통해 NOFX에 액세스할 수 있습니다:

```bash
# 서버에 배포
curl -fsSL https://raw.githubusercontent.com/woodyyan/nofx/dev/install.sh | bash
```

`http://YOUR_SERVER_IP:3000`을 통해 액세스 - 즉시 작동합니다.

### 향상된 보안 (HTTPS)

보안을 강화하려면 `.env`에서 전송 암호화를 활성화하세요:

```bash
TRANSPORT_ENCRYPTION=true
```

활성화되면 브라우저는 Web Crypto API를 사용하여 전송 전에 API 키를 암호화합니다. 이를 위해 필요한 것:
- `https://` - SSL이 있는 모든 도메인
- `http://localhost` - 로컬 개발

### Cloudflare를 사용한 빠른 HTTPS 설정

1. **Cloudflare에 도메인 추가** (무료 플랜 가능)
   - [dash.cloudflare.com](https://dash.cloudflare.com) 방문
   - 도메인 추가 및 네임서버 업데이트

2. **DNS 레코드 생성**
   - 유형: `A`
   - 이름: `nofx` (또는 서브도메인)
   - 콘텐츠: 서버 IP
   - 프록시 상태: **Proxied** (주황색 구름)

3. **SSL/TLS 구성**
   - SSL/TLS 설정으로 이동
   - 암호화 모드를 **Flexible**로 설정

   ```
   User ──[HTTPS]──→ Cloudflare ──[HTTP]──→ Your Server:3000
   ```

4. **전송 암호화 활성화**
   ```bash
   # .env 편집 및 설정
   TRANSPORT_ENCRYPTION=true
   ```

5. **완료!** `https://nofx.yourdomain.com`을 통해 액세스

---

## 초기 설정 (웹 인터페이스)

시스템을 시작한 후 웹 인터페이스를 통해 구성합니다:

1. **AI 모델 구성** - AI API 키 추가 (DeepSeek, OpenAI 등)
2. **거래소 구성** - 거래소 API 자격 증명 설정
3. **전략 생성** - 전략 스튜디오에서 거래 전략 구성
4. **트레이더 생성** - AI 모델 + 거래소 + 전략 결합
5. **거래 시작** - 구성된 트레이더 시작

모든 구성은 웹 인터페이스를 통해 완료 - JSON 파일 편집 불필요.

---

## 웹 인터페이스 기능

### 경쟁 페이지
- 실시간 ROI 리더보드
- 다중 AI 성능 비교 차트
- 실시간 손익 추적 및 순위

### 대시보드
- TradingView 스타일 캔들스틱 차트
- 실시간 포지션 관리
- Chain of Thought 추론이 포함된 AI 결정 로그
- 자본 곡선 추적

### 전략 스튜디오
- 코인 소스 구성 (정적 목록, AI500 풀, OI Top)
- 기술 지표 (EMA, MACD, RSI, ATR, 거래량, OI, 펀딩 비율)
- 리스크 제어 설정 (레버리지, 포지션 한도, 마진 사용)
- 실시간 프롬프트 미리보기를 포함한 AI 테스트

---

## 일반적인 문제

### TA-Lib을 찾을 수 없음
```bash
# macOS
brew install ta-lib

# Ubuntu
sudo apt-get install libta-lib0-dev
```

### AI API 타임아웃
- API 키가 올바른지 확인
- 네트워크 연결 확인
- 시스템 타임아웃은 120초

### 프론트엔드가 백엔드에 연결할 수 없음
- 백엔드가 http://localhost:8080에서 실행 중인지 확인
- 포트가 점유되어 있지 않은지 확인

---

## 라이선스

이 프로젝트는 **GNU Affero General Public License v3.0 (AGPL-3.0)** 라이선스에 따라 제공됩니다 - [LICENSE](LICENSE) 파일을 참조하세요.

---

## 기여

기여를 환영합니다! 다음을 참조하세요:
- **[기여 가이드](CONTRIBUTING.md)** - 개발 워크플로 및 PR 프로세스
- **[행동 강령](CODE_OF_CONDUCT.md)** - 커뮤니티 가이드라인
- **[보안 정책](SECURITY.md)** - 취약점 보고

---

## 기여자 에어드롭 프로그램

모든 기여는 GitHub에서 추적됩니다. NOFX가 수익을 창출하면 기여자는 기여도에 따라 에어드롭을 받게 됩니다.

**[고정된 Issue](https://github.com/woodyyan/nofx/issues)를 해결하는 PR은 최고 보상을 받습니다!**

| 기여 유형 | 가중치 |
|------------------|:------:|
| **고정된 Issue PR** | ⭐⭐⭐⭐⭐⭐ |
| **코드 커밋** (병합된 PR) | ⭐⭐⭐⭐⭐ |
| **버그 수정** | ⭐⭐⭐⭐ |
| **기능 제안** | ⭐⭐⭐ |
| **버그 보고** | ⭐⭐ |
| **문서** | ⭐⭐ |

---

## 위험 경고

1. 암호화폐 시장은 매우 변동성이 높음 - AI 결정이 수익을 보장하지 않음
2. 선물 거래는 레버리지 사용 - 손실이 원금을 초과할 수 있음
3. 극단적인 시장 상황에서 청산 위험 있음




## 연락처

- **GitHub Issues**: [Issue 제출](https://github.com/woodyyan/nofx/issues)
- **개발자 커뮤니티**: [Telegram 그룹](https://t.me/nofx_dev_community)

---

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=woodyyan/nofx&type=Date)](https://star-history.com/#woodyyan/nofx&Date)
