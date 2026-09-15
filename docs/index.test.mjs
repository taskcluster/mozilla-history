import assert from 'node:assert/strict'
import { readFile } from 'node:fs/promises'
import test from 'node:test'

const html = await readFile(new URL('./index.html', import.meta.url), 'utf8')

test('the report page has one syntactically valid inline module', () => {
  const modules = [...html.matchAll(/<script type="module">([\s\S]*?)<\/script>/g)]
  assert.equal(modules.length, 1)

  const sourceWithoutImports = modules[0][1].replace(/^\s*import .*$/gm, '')
  assert.doesNotThrow(() => new Function(sourceWithoutImports))
})

test('Quick and local previews use the bundled report', () => {
  assert.match(html, /query\.has\('local'\)/)
  assert.match(html, /'127\.0\.0\.1', 'localhost', '\[::1\]'/)
  assert.match(html, /loopbackHost \|\| window\.location\.hostname/)
  assert.match(html, /hostname\.endsWith\('\.quick\.mozilla\.cloud'\)/)
  assert.match(html, /new URL\('\.\.\/WorkerVersions\/', baseUrl\)/)
})

test('report metadata is inserted at the end of the report introduction', () => {
  assert.doesNotMatch(html, /<div class="content-column">\s*<p class="source"/)
  assert.match(html, /const firstSection = content\.querySelector\('h2'\)/)
  assert.match(html, /if \(firstSection\) firstSection\.before\(context\)/)
  assert.match(html, /Source code: <a href="\$\{repoUrl\}">GitHub<\/a>/)
  assert.match(html, /Current worker-pool report \(Markdown\)/)
  assert.match(html, /Current worker-pool data \(JSON\)/)
  assert.match(html, /Historical worker-pool data \(JSON\)/)
  assert.match(html, /startsWith\('Probe run started:'\)/)
  assert.match(html, /reportTiming\.insertBefore\(coverage, closingParen\)/)
  assert.match(html, /else reportTiming\.append\(coverage\)/)
  assert.match(html, /coverage\.dataset\.withRevision === 'true' \? `, \$\{summary\}` : ` \(\$\{summary\}\)`/)
  assert.match(html, /\$\{dates\.length\} snapshots from \$\{dates\[0\]\} to \$\{dates\[dates\.length - 1\]\}/)
  assert.doesNotMatch(html, /\.report-context \{[\s\S]*?font-size:/)
  assert.doesNotMatch(html, /Disclaimer:/)
})

test('report artifacts are revalidated at their stable URLs', () => {
  assert.match(html, /fetch\(url, \{ cache: 'no-cache' \}\)/)
  assert.doesNotMatch(html, /cache: 'no-store'/)
})

test('the merged page retains report and upstream table enhancements', () => {
  assert.match(html, /import \{ decorateSortableTables \} from '\.\/table-sort\.js'/)
  assert.match(html, /import \{ enhanceTables \} from '\.\/table-utils\.js'/)
  assert.match(html, /id="toc-list"/)
  assert.match(html, /function buildTableOfContents\(\)/)
  assert.doesNotMatch(html, /\.\/app\.js/)
})

test('the table of contents uses a collapsible, non-overlay layout', () => {
  assert.match(html, /id="report-layout" class="report-layout"/)
  assert.match(html, /grid-template-columns: max-content minmax\(0, 1fr\)/)
  assert.match(html, /grid-template-columns: 32px minmax\(0, 1fr\)/)
  assert.match(html, /\.toc \{[\s\S]*?position: sticky/)
  assert.doesNotMatch(html, /\.toc \{[\s\S]*?position: fixed/)
  assert.match(html, /@media screen and \(max-width: 900px\)[\s\S]*?grid-template-columns: minmax\(0, 1fr\)/)
  assert.match(html, /@media screen and \(min-width: 901px\)[\s\S]*?white-space: nowrap/)
  assert.match(html, /\.toc ul \{[\s\S]*?list-style: none/)
  assert.match(html, /table\.tu-sticky \{[\s\S]*?width: max-content[\s\S]*?overflow: visible/)
})

test('the table of contents control is accessible and persistent', () => {
  assert.match(html, /<button id="toc-toggle"[\s\S]*?aria-expanded="true"[\s\S]*?aria-controls="toc-list"/)
  assert.match(html, /aria-label="Collapse table of contents"/)
  assert.match(html, /title="Collapse table of contents"/)
  assert.match(html, /<svg class="toc-collapse-icon"[\s\S]*?<svg class="toc-expand-icon"/)
  assert.match(html, /const TOC_STORAGE_KEY = 'mozilla-history:toc-collapsed'/)
  assert.match(html, /window\.localStorage\.getItem\(TOC_STORAGE_KEY\)/)
  assert.match(html, /window\.localStorage\.setItem\(TOC_STORAGE_KEY, String\(collapsed\)\)/)
  assert.match(html, /preference \?\? window\.matchMedia\('\(max-width: 900px\)'\)\.matches/)
  assert.match(html, /list\.hidden = collapsed/)
  assert.match(html, /toggle\.setAttribute\('aria-expanded', String\(!collapsed\)\)/)
  assert.match(html, /initTableOfContents\(\)\s*\n\s*init\(\)/)
})

test('table of contents targets receive accessible permalinks', () => {
  assert.match(html, /'#content h1'/)
  assert.match(html, /className = 'heading-permalink'/)
  assert.match(html, /permalink\.href = `#\$\{id\}`/)
  assert.match(html, /permalink\.setAttribute\('aria-label', `Link to \$\{label\}`\)/)
  assert.match(html, /permalink\.title = `Permalink to \$\{label\}`/)
  assert.match(html, /if \(heading\.tagName === 'H1'\) return 'Overview'/)
  assert.match(html, /if \(label === 'Worker pools'\) return 'Pool details'/)
  assert.match(html, /label\.startsWith\('Worker pools by '\)/)
  assert.match(html, /link\.textContent = tableOfContentsLabel\(heading, label\)/)
})
