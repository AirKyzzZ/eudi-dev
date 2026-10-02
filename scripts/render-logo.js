// Render logo, wordmark and social images from docs/assets/logo-mark.svg. Run node
// scripts/render-logo.js after changing the SVG. The script uses Playwright from
// e2e/node_modules.
const { createRequire } = require('module');
const fs = require('fs');
const path = require('path');

const { chromium } = createRequire(path.join(__dirname, '../e2e/'))('@playwright/test');

const repo = path.resolve(__dirname, '..');
const out = path.join(repo, 'docs/assets');
const rawMark = fs.readFileSync(path.join(out, 'logo-mark.svg'), 'utf8');
const mark = (w) =>
  rawMark.replace(/width="46"\s+height="51"/, `width="${w}" height="${Math.round((w * 51) / 46)}"`);

function page({
  text, bg, cardBg, cardBorder, radius, markWidth, nameSize, gap, pad, logoOnly = false,
}) {
  return `<!doctype html><html><head><meta charset="utf-8"><style>
    * { margin:0; padding:0; box-sizing:border-box; }
    html, body { background: ${bg}; height:100%; }
    body { display:flex; align-items:center; justify-content:center; }
    /* The light card needs an edge, otherwise it dissolves into a white page. */
    .wrap { display:flex; align-items:center; gap:${gap}px; padding:${pad};
      background:${cardBg || 'transparent'}; border-radius:${radius || 0}px;
      border:${cardBorder && cardBorder !== 'transparent' ? `2px solid ${cardBorder}` : 'none'}; }
    .mark { flex:0 0 auto; display:flex; }
    /* Centring on the mark's box puts the word too high: the wallet body is
       the visual mass, the card above it is an accent. The nudge drops the
       word onto the body, and covers the descender space "EUDI Dev Wallet"
       never uses. */
    .name { font-family:"SF Mono","Menlo","Cascadia Code",monospace; font-size:${nameSize}px;
      font-weight:600; letter-spacing:-1px; color:${text}; line-height:1; white-space:nowrap;
      position:relative; top:${Math.round(nameSize * 0.1)}px; }
  </style></head><body>
    <div class="wrap" id="shot">
      <div class="mark">${mark(markWidth)}</div>
      ${logoOnly ? '' : '<div class="name">EUDI Dev Wallet</div>'}
    </div>
  </body></html>`;
}

const wordmark = { markWidth: 220, nameSize: 104, gap: 40 };

// Generate versions for both UI themes and plain black text for print.
const dark = { text: '#c0caf5', surface: '#1a1b26', border: 'transparent' };
const light = { text: '#24283b', surface: '#ffffff', border: '#d0d0d0' };
const black = { text: '#000000', surface: '#ffffff', border: '#d0d0d0' };

(async () => {
  const browser = await chromium.launch();

  for (const variant of [
    { file: 'logo-banner-dark.png', theme: dark, card: true },
    { file: 'logo-banner-light.png', theme: light, card: true },
    { file: 'logo-wordmark-dark.png', theme: dark, card: false },
    { file: 'logo-wordmark-light.png', theme: light, card: false },
    { file: 'logo-wordmark-black.png', theme: black, card: false },
  ]) {
    const p = await browser.newPage({ viewport: { width: 1800, height: 600 }, deviceScaleFactor: 2 });
    await p.setContent(page({
      ...wordmark,
      text: variant.theme.text,
      bg: 'transparent',
      cardBg: variant.card ? variant.theme.surface : null,
      cardBorder: variant.card ? variant.theme.border : null,
      radius: variant.card ? 28 : 0,
      pad: variant.card ? '64px 90px' : '36px 44px',
    }));
    await p.locator('#shot').screenshot({ path: path.join(out, variant.file), omitBackground: true });
    await p.close();
  }

  for (const variant of [
    { file: 'logo-social-dark.png', theme: dark },
    { file: 'logo-social-light.png', theme: light },
  ]) {
    const p = await browser.newPage({ viewport: { width: 1200, height: 630 }, deviceScaleFactor: 2 });
    await p.setContent(page({
      text: variant.theme.text, bg: variant.theme.surface,
      markWidth: 215, nameSize: 100, gap: 38, pad: '0',
    }));
    await p.screenshot({ path: path.join(out, variant.file) });
    await p.close();
  }

  for (const variant of [
    { file: 'logo-mark.png', width: 512, height: 512, markWidth: 420, bg: 'transparent' },
    { file: 'logo-social-mark.png', width: 1280, height: 640, markWidth: 320, bg: '#ffffff' },
  ]) {
    const p = await browser.newPage({
      viewport: { width: variant.width, height: variant.height }, deviceScaleFactor: 1,
    });
    await p.setContent(page({
      logoOnly: true, bg: variant.bg, markWidth: variant.markWidth,
      nameSize: 0, gap: 0, pad: '0',
    }));
    await p.screenshot({
      path: path.join(out, variant.file), omitBackground: variant.bg === 'transparent',
    });
    await p.close();
  }

  const readmeDark = await browser.newPage({
    viewport: { width: 2000, height: 439 }, deviceScaleFactor: 2,
  });
  await readmeDark.setContent(`<!doctype html><html><head><style>
    body { margin: 0; }
    svg { display: block; }
  </style></head><body>
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 4000 878" width="2000" height="439">
      <g transform="translate(96 0)">${mark(740)}</g>
      <text x="940" y="590" fill="#c0caf5" font-family="Arial, Helvetica, sans-serif"
        font-size="400" font-weight="700" textLength="2920" lengthAdjust="spacingAndGlyphs">EUDI Dev Wallet</text>
    </svg>
  </body></html>`);
  await readmeDark.screenshot({ path: path.join(out, 'logo-readme-dark.png'), omitBackground: true });
  await readmeDark.close();

  await browser.close();
  console.log('wrote banner, wordmark, social and logo-only variants to docs/assets');
})();
