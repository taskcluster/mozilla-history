import { expect, test } from '@playwright/test'
import { execFileSync } from 'node:child_process'
import { fileURLToPath } from 'node:url'

let renderedReport
test.beforeAll(() => {
  // Exercise the current renderer without committing generated report changes.
  renderedReport = execFileSync('go', [
    'run', './audit-worker-versions', 'render', 'WorkerVersions/workers.json',
  ], {
    cwd: fileURLToPath(new URL('..', import.meta.url)),
    encoding: 'utf8',
    maxBuffer: 10 * 1024 * 1024,
  })
})

async function loadReport(page) {
  await page.route('**/WorkerVersions/README.md', route => route.fulfill({
    contentType: 'text/plain; charset=utf-8',
    body: renderedReport,
  }))
  await page.goto('/docs/index.html?local')
  await expect(page.locator('#content h1')).toBeVisible({ timeout: 30_000 })
  await expect(page.locator('#toc-list li').first()).toBeAttached()
}

test('wide navigation is compact, collapses, and persists', async ({ page }, testInfo) => {
  await page.setViewportSize({ width: 1440, height: 900 })
  await loadReport(page)

  const toc = page.locator('#toc')
  const content = page.locator('.content-column')
  const toggle = page.locator('#toc-toggle')
  const expandedToc = await toc.boundingBox()
  const expandedContent = await content.boundingBox()

  expect(expandedToc.x + expandedToc.width).toBeLessThanOrEqual(expandedContent.x)
  expect(expandedContent.x).toBeLessThan(330)
  expect(await page.locator('#toc-list a').evaluateAll(links =>
    links.every(link => getComputedStyle(link).whiteSpace === 'nowrap')
  )).toBe(true)
  const toggleBox = await toggle.boundingBox()
  expect(toggleBox.width).toBeCloseTo(32, 1)
  expect(toggleBox.height).toBeCloseTo(32, 1)
  expect(toggleBox.y).toBeCloseTo(31, 0)
  await expect(toggle).toHaveAttribute('title', 'Collapse table of contents')
  await expect(toggle).toHaveAttribute('aria-label', 'Collapse table of contents')
  await page.screenshot({ path: testInfo.outputPath('wide-expanded.png') })

  await toggle.click()
  await expect(toggle).toHaveAttribute('aria-expanded', 'false')
  await expect(toggle).toHaveAttribute('title', 'Expand table of contents')
  await expect(page.locator('#toc-list')).toBeHidden()
  const collapsedToc = await toc.boundingBox()
  const collapsedContent = await content.boundingBox()
  expect(collapsedToc.width).toBeLessThanOrEqual(32)
  expect(collapsedToc.x).toBe(20)
  expect(await page.locator('#report-layout').evaluate(element =>
    Number.parseFloat(getComputedStyle(element).columnGap)
  )).toBe(16)
  expect(collapsedToc.y).toBeCloseTo(31, 0)
  expect(collapsedContent.x).toBeLessThan(expandedContent.x)
  await page.screenshot({ path: testInfo.outputPath('wide-collapsed.png') })

  await page.reload()
  await expect(toggle).toHaveAttribute('aria-expanded', 'false')
  await expect(page.locator('#toc-list')).toBeHidden()
})

test('narrow navigation defaults collapsed and expands in flow', async ({ page }, testInfo) => {
  await page.setViewportSize({ width: 720, height: 900 })
  await loadReport(page)

  const toc = page.locator('#toc')
  const content = page.locator('.content-column')
  const toggle = page.locator('#toc-toggle')
  await expect(toggle).toHaveAttribute('aria-expanded', 'false')
  await expect(page.locator('#toc-list')).toBeHidden()
  expect((await toc.boundingBox()).width).toBeLessThanOrEqual(32)
  await page.screenshot({ path: testInfo.outputPath('narrow-collapsed.png') })

  await toggle.click()
  await expect(toggle).toHaveAttribute('aria-expanded', 'true')
  const tocBox = await toc.boundingBox()
  const contentBox = await content.boundingBox()

  expect(await toc.evaluate(element => getComputedStyle(element).position)).toBe('static')
  expect(tocBox.y + tocBox.height).toBeLessThanOrEqual(contentBox.y)
  await page.screenshot({ path: testInfo.outputPath('narrow-expanded.png') })

  await page.reload()
  await expect(toggle).toHaveAttribute('aria-expanded', 'true')
  const reloadedTocBox = await toc.boundingBox()
  const reloadedContentBox = await content.boundingBox()
  expect(reloadedTocBox.y + reloadedTocBox.height).toBeLessThanOrEqual(reloadedContentBox.y)
})

test('report headings expose stable, accessible permalinks', async ({ page }, testInfo) => {
  await page.setViewportSize({ width: 1440, height: 900 })
  await loadReport(page)

  const overviewHeading = page.locator('#content h1')
  const overviewLink = page.locator('#toc-list > li').first().locator(':scope > a')
  await expect(overviewHeading).toHaveAttribute('id', 'worker-pool-versions')
  await expect(overviewHeading.locator(':scope > .heading-permalink')).toHaveCount(0)
  await expect(overviewLink).toHaveText('Overview')
  await expect(overviewLink).toHaveAttribute('href', '#worker-pool-versions')

  const heading = page.locator('#content h2').first()
  const label = (await heading.evaluate(element => element.childNodes[0].textContent)).trim()
  const id = await heading.getAttribute('id')
  const permalink = heading.locator(':scope > .heading-permalink')

  expect(label).toBe('Generic Worker')
  await expect(permalink).toHaveAttribute('href', `#${id}`)
  await expect(permalink).toHaveAttribute('aria-label', `Link to ${label}`)
  await expect(permalink).toHaveAttribute('title', `Permalink to ${label}`)
  const sectionItem = page.locator('#toc-list > li').nth(1)
  await expect(sectionItem.locator(':scope > a')).toHaveText(label)
  await expect(sectionItem.locator(':scope > ul > li > a')).toHaveText([
    'By version',
    'By image',
    'Pool details',
  ])
  await expect(page.locator('#content h3').first()).toContainText('Worker pools by version')

  await expect(permalink).toHaveCSS('opacity', '0.35')
  await heading.hover()
  await expect(permalink).toHaveCSS('opacity', '1')
  await page.screenshot({ path: testInfo.outputPath('heading-permalink.png') })
  await permalink.click()
  await expect(page).toHaveURL(new RegExp(`#${id}$`))
})

test('version summary tables default to descending natural order', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 900 })
  await loadReport(page)

  const headings = page.locator('#content h3[id$="-worker-pools-by-version"]')
  await expect(headings).toHaveCount(2)

  for (let index = 0; index < await headings.count(); index += 1) {
    const versions = await headings.nth(index)
      .locator('xpath=following-sibling::table[1]')
      .locator('tbody td:first-child')
      .allTextContents()
    const expected = [...versions].sort((left, right) =>
      right.localeCompare(left, undefined, { numeric: true }),
    )
    expect(versions).toEqual(expected)
  }
})

test('wide tables retain page scrolling and sticky cells', async ({ page }, testInfo) => {
  await page.setViewportSize({ width: 1280, height: 800 })
  await loadReport(page)
  await page.locator('#toc-toggle').click()

  const table = page.locator('#history-content table').last()
  await table.scrollIntoViewIfNeeded()
  expect(await table.evaluate(element => getComputedStyle(element).overflow)).toBe('visible')
  expect(await page.evaluate(() => document.documentElement.scrollWidth > innerWidth)).toBe(true)

  const tableBox = await table.boundingBox()
  const tableTop = await page.evaluate(y => scrollY + y, tableBox.y)
  await page.evaluate(({ x, y }) => scrollTo(x, y), {
    x: 500,
    y: tableTop + 100,
  })

  const firstHeader = table.locator('thead th').first()
  const firstCell = table.locator('tbody td').first()
  expect((await firstHeader.boundingBox()).y).toBeCloseTo(0, 0)
  expect((await firstCell.boundingBox()).x).toBeCloseTo(0, 0)
  await page.screenshot({ path: testInfo.outputPath('sticky-table.png') })
})
