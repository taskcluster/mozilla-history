import assert from 'node:assert/strict'
import test from 'node:test'

import { numericSortKey, rangeSortKey, versionSortKey } from './table-sort.js'

test('version sort keys order dotted numeric components naturally', () => {
  const versions = ['108.0.0', '9.10.0', '100.0.1', '9.2.0', '99.2.0']
  assert.deepEqual(
    versions.sort((left, right) => versionSortKey(left).localeCompare(versionSortKey(right))),
    ['9.2.0', '9.10.0', '99.2.0', '100.0.1', '108.0.0'],
  )
})

test('non-version values receive an empty nulls-last sort key', () => {
  assert.equal(versionSortKey('No artifacts found'), '')
  assert.equal(versionSortKey('—'), '')
})

test('range sort keys order by maximum and then minimum', () => {
  const ranges = ['0–100', '1–4', '0–10', '4', '0–4']
  assert.deepEqual(
    ranges.sort((left, right) => rangeSortKey(left).localeCompare(rangeSortKey(right))),
    ['0–4', '1–4', '4', '0–10', '0–100'],
  )
})

test('missing ranges receive an empty nulls-last sort key', () => {
  assert.equal(rangeSortKey('—'), '')
  assert.equal(rangeSortKey('unknown'), '')
})

test('numeric sort keys order slots numerically and identify missing values', () => {
  const slots = ['16', '2', '128', '1']
  assert.deepEqual(
    slots.sort((left, right) => numericSortKey(left).localeCompare(numericSortKey(right))),
    ['1', '2', '16', '128'],
  )
  assert.equal(numericSortKey('—'), '')
})
