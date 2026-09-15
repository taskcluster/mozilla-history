const SORT_KEY_WIDTH = 16

function paddedInteger(value) {
  return value.replace(/^0+(?=\d)/, '').padStart(SORT_KEY_WIDTH, '0')
}

export function versionSortKey(value) {
  const version = value.trim()
  if (!/^\d+(?:\.\d+)*$/.test(version)) return ''

  // The leading letter makes sortable compare this as text rather than
  // parseFloat-ing only the first dotted component.
  return `v${version.split('.').map(paddedInteger).join('.')}`
}

export function rangeSortKey(value) {
  const match = value.trim().match(/^(\d+)(?:\s*[–-]\s*(\d+))?$/)
  if (!match) return ''

  const minimum = paddedInteger(match[1])
  const maximum = paddedInteger(match[2] || match[1])
  return `r${maximum}.${minimum}`
}

export function numericSortKey(value) {
  const number = value.trim()
  if (!/^\d+$/.test(number)) return ''
  return `n${paddedInteger(number)}`
}

function setColumnSortKeys(table, columnIndex, keyForValue) {
  table.querySelectorAll('tbody tr').forEach(row => {
    const cell = row.cells[columnIndex]
    if (cell) cell.dataset.sort = keyForValue(cell.textContent)
  })
}

export function decorateSortableTable(table) {
  table.classList.add('sortable', 'n-last')

  table.querySelectorAll('thead th').forEach((header, columnIndex) => {
    const title = header.textContent.trim()
    if (title === 'Version') {
      setColumnSortKeys(table, columnIndex, versionSortKey)
    } else if (title === 'Configured Workers' || title === 'Configured Capacity') {
      setColumnSortKeys(table, columnIndex, rangeSortKey)
      header.title = 'Sorted by maximum, then minimum'
    } else if (title === 'Slots per Worker') {
      setColumnSortKeys(table, columnIndex, numericSortKey)
    }
  })
}

export function decorateSortableTables(root = document) {
  root.querySelectorAll('table').forEach(decorateSortableTable)
}
