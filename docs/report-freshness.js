export const STALE_PROBE_HOURS = 10 * 24

export function staleProbeAge(timingText, now = Date.now()) {
  const match = /^Probe run started: (\d{4}-\d{2}-\d{2}) (\d{2}:\d{2}) UTC/.exec(timingText)
  if (!match) return null
  const iso = `${match[1]}T${match[2]}:00.000Z`
  const started = Date.parse(iso)
  if (!Number.isFinite(started) || new Date(started).toISOString() !== iso) return null
  const age = now - started
  if (!Number.isFinite(age) || age <= STALE_PROBE_HOURS * 60 * 60 * 1000) return null
  const hours = Math.floor(age / (60 * 60 * 1000))
  const count = hours >= 24 ? Math.floor(hours / 24) : hours
  const unit = hours >= 24 ? 'day' : 'hour'
  return `${count} ${unit}${count === 1 ? '' : 's'} ago`
}
