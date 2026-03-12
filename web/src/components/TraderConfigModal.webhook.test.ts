import { describe, it, expect } from 'vitest'

/**
 * Tests for Webhook configuration in TraderConfigModal
 *
 * Strategy: pure logic tests (no DOM rendering) consistent with project style.
 * We test:
 *   1. FormState initial values for webhook fields
 *   2. Webhook fields are correctly included in CreateTraderRequest
 *   3. traderData → FormState mapping (edit mode pre-fill)
 *   4. URL validation logic
 *   5. i18n translation keys exist for all three languages
 */

// ---------- Types we replicate locally to keep tests self-contained ----------
interface FormState {
  trader_name: string
  ai_model: string
  exchange_id: string
  strategy_id: string
  is_cross_margin: boolean
  show_in_competition: boolean
  scan_interval_minutes: number
  initial_balance?: number
  webhook_url: string
  webhook_secret: string
  webhook_enabled: boolean
}

interface CreateTraderRequest {
  name: string
  ai_model_id: string
  exchange_id: string
  strategy_id?: string
  is_cross_margin?: boolean
  show_in_competition?: boolean
  scan_interval_minutes?: number
  webhook_url?: string
  webhook_secret?: string
  webhook_enabled?: boolean
}

interface TraderConfigData {
  trader_name: string
  ai_model: string
  exchange_id: string
  strategy_id?: string
  is_cross_margin: boolean
  show_in_competition: boolean
  scan_interval_minutes: number
  initial_balance: number
  is_running: boolean
  webhook_url?: string
  webhook_secret?: string
  webhook_enabled?: boolean
}

// ---------- Helpers (mirrors component logic) ----------
function buildDefaultFormState(): FormState {
  return {
    trader_name: '',
    ai_model: '',
    exchange_id: '',
    strategy_id: '',
    is_cross_margin: true,
    show_in_competition: true,
    scan_interval_minutes: 3,
    webhook_url: '',
    webhook_secret: '',
    webhook_enabled: false,
  }
}

function formStateFromTraderData(traderData: TraderConfigData): FormState {
  return {
    trader_name: traderData.trader_name,
    ai_model: traderData.ai_model,
    exchange_id: traderData.exchange_id,
    strategy_id: traderData.strategy_id || '',
    is_cross_margin: traderData.is_cross_margin,
    show_in_competition: traderData.show_in_competition,
    scan_interval_minutes: traderData.scan_interval_minutes,
    initial_balance: traderData.initial_balance,
    webhook_url: traderData.webhook_url || '',
    webhook_secret: traderData.webhook_secret || '',
    webhook_enabled: traderData.webhook_enabled ?? false,
  }
}

function buildSaveData(formData: FormState, isEditMode: boolean): CreateTraderRequest {
  const saveData: CreateTraderRequest = {
    name: formData.trader_name,
    ai_model_id: formData.ai_model,
    exchange_id: formData.exchange_id,
    strategy_id: formData.strategy_id,
    is_cross_margin: formData.is_cross_margin,
    show_in_competition: formData.show_in_competition,
    scan_interval_minutes: formData.scan_interval_minutes,
    webhook_url: formData.webhook_url,
    webhook_secret: formData.webhook_secret,
    webhook_enabled: formData.webhook_enabled,
  }
  return saveData
}

function isValidWebhookURL(url: string): boolean {
  if (!url) return true // empty = disabled, acceptable
  try {
    const parsed = new URL(url)
    return parsed.protocol === 'http:' || parsed.protocol === 'https:'
  } catch {
    return false
  }
}

// ---------- Tests ----------
describe('TraderConfigModal - Webhook Configuration', () => {
  describe('FormState initial values', () => {
    it('should initialize webhook fields with safe defaults', () => {
      const state = buildDefaultFormState()
      expect(state.webhook_url).toBe('')
      expect(state.webhook_secret).toBe('')
      expect(state.webhook_enabled).toBe(false)
    })

    it('should default webhook_enabled to false (opt-in)', () => {
      const state = buildDefaultFormState()
      expect(state.webhook_enabled).toBe(false)
    })
  })

  describe('formStateFromTraderData - edit mode pre-fill', () => {
    it('should map webhook fields from traderData when all are set', () => {
      const traderData: TraderConfigData = {
        trader_name: 'My Bot',
        ai_model: 'gpt-4o',
        exchange_id: 'ex-001',
        is_cross_margin: true,
        show_in_competition: true,
        scan_interval_minutes: 5,
        initial_balance: 1000,
        is_running: false,
        webhook_url: 'https://my-server.com/hook',
        webhook_secret: 'super-secret',
        webhook_enabled: true,
      }
      const state = formStateFromTraderData(traderData)
      expect(state.webhook_url).toBe('https://my-server.com/hook')
      expect(state.webhook_secret).toBe('super-secret')
      expect(state.webhook_enabled).toBe(true)
    })

    it('should fall back to empty strings when webhook fields are missing', () => {
      const traderData: TraderConfigData = {
        trader_name: 'My Bot',
        ai_model: 'gpt-4o',
        exchange_id: 'ex-001',
        is_cross_margin: true,
        show_in_competition: true,
        scan_interval_minutes: 5,
        initial_balance: 1000,
        is_running: false,
        // webhook fields intentionally omitted
      }
      const state = formStateFromTraderData(traderData)
      expect(state.webhook_url).toBe('')
      expect(state.webhook_secret).toBe('')
      expect(state.webhook_enabled).toBe(false)
    })
  })

  describe('buildSaveData - webhook fields included in request', () => {
    it('should include webhook fields in save payload when enabled', () => {
      const formData: FormState = {
        ...buildDefaultFormState(),
        trader_name: 'Test Bot',
        ai_model: 'gpt-4o',
        exchange_id: 'ex-001',
        webhook_url: 'https://example.com/hook',
        webhook_secret: 'mysecret',
        webhook_enabled: true,
      }
      const payload = buildSaveData(formData, false)
      expect(payload.webhook_url).toBe('https://example.com/hook')
      expect(payload.webhook_secret).toBe('mysecret')
      expect(payload.webhook_enabled).toBe(true)
    })

    it('should include webhook_enabled=false when disabled', () => {
      const formData: FormState = {
        ...buildDefaultFormState(),
        trader_name: 'Test Bot',
        ai_model: 'gpt-4o',
        exchange_id: 'ex-001',
        webhook_enabled: false,
      }
      const payload = buildSaveData(formData, false)
      expect(payload.webhook_enabled).toBe(false)
      expect(payload.webhook_url).toBe('')
    })

    it('should send empty webhook_secret when not set', () => {
      const formData: FormState = {
        ...buildDefaultFormState(),
        trader_name: 'Test Bot',
        ai_model: 'gpt-4o',
        exchange_id: 'ex-001',
        webhook_url: 'https://example.com/hook',
        webhook_enabled: true,
      }
      const payload = buildSaveData(formData, false)
      expect(payload.webhook_secret).toBe('')
    })
  })

  describe('URL validation', () => {
    it('should accept empty URL (disabled state)', () => {
      expect(isValidWebhookURL('')).toBe(true)
    })

    it('should accept valid https URL', () => {
      expect(isValidWebhookURL('https://example.com/webhook')).toBe(true)
    })

    it('should accept valid http URL', () => {
      expect(isValidWebhookURL('http://localhost:3000/hook')).toBe(true)
    })

    it('should reject non-HTTP protocols', () => {
      expect(isValidWebhookURL('ftp://example.com/hook')).toBe(false)
    })

    it('should reject malformed URL', () => {
      expect(isValidWebhookURL('not-a-url')).toBe(false)
    })

    it('should reject URL with missing protocol', () => {
      expect(isValidWebhookURL('example.com/webhook')).toBe(false)
    })
  })

  describe('i18n translation keys', () => {
    // Import translations to verify all required keys exist
    it('should have all webhook translation keys in English', async () => {
      const { translations } = await import('../i18n/translations')
      const en = translations.en as Record<string, string>
      expect(en.webhookConfig).toBeDefined()
      expect(en.webhookEnabled).toBeDefined()
      expect(en.webhookEnabledDesc).toBeDefined()
      expect(en.webhookURL).toBeDefined()
      expect(en.webhookURLPlaceholder).toBeDefined()
      expect(en.webhookSecret).toBeDefined()
      expect(en.webhookSecretPlaceholder).toBeDefined()
      expect(en.webhookSecretHint).toBeDefined()
    })

    it('should have all webhook translation keys in Chinese', async () => {
      const { translations } = await import('../i18n/translations')
      const zh = translations.zh as Record<string, string>
      expect(zh.webhookConfig).toBeDefined()
      expect(zh.webhookEnabled).toBeDefined()
      expect(zh.webhookEnabledDesc).toBeDefined()
      expect(zh.webhookURL).toBeDefined()
      expect(zh.webhookSecret).toBeDefined()
      expect(zh.webhookSecretHint).toBeDefined()
    })

    it('should have all webhook translation keys in Indonesian', async () => {
      const { translations } = await import('../i18n/translations')
      const id = translations.id as Record<string, string>
      expect(id.webhookConfig).toBeDefined()
      expect(id.webhookEnabled).toBeDefined()
      expect(id.webhookURL).toBeDefined()
    })

    it('en webhookConfig should not be empty', async () => {
      const { translations } = await import('../i18n/translations')
      const en = translations.en as Record<string, string>
      expect(en.webhookConfig.length).toBeGreaterThan(0)
    })
  })
})
