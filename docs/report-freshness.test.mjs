import assert from 'node:assert/strict'
import test from 'node:test'
import { staleProbeAge } from './report-freshness.js'

const timing = 'Probe run started: 2026-09-09 07:58 UTC · Report generated: 2026-09-11 07:58 UTC'
const started = Date.parse('2026-09-09T07:58:00Z')
const hour = 60 * 60 * 1000
const day = 24 * hour

test('fresh and future probes do not show a notice, including the threshold', () => {
  for (const age of [-hour, 0, hour, 2 * hour, 9 * day, 10 * day]) {
    assert.equal(staleProbeAge(timing, started + age), null)
  }
})

test('staleness uses the probe time and appears only after ten days', () => {
  assert.equal(staleProbeAge(timing, started + 10 * day + 1), '10 days ago')
  assert.equal(staleProbeAge(timing, started + 12 * day), '12 days ago')
  assert.equal(staleProbeAge(timing, started + 12 * day + 23 * hour), '12 days ago')
})

test('missing and invalid probe dates do not show a notice', () => {
  for (const value of ['', 'Report generated: 2026-09-09 07:58 UTC',
    'Probe run started: 2026-02-30 07:58 UTC', 'Probe run started: 2026-09-09 25:58 UTC']) {
    assert.equal(staleProbeAge(value, started + 20 * day), null)
  }
})
