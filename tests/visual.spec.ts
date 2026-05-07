import { test, expect } from '@playwright/test';

test.describe('Landing Page', () => {
  test('hero section is visible with image and text', async ({ page }) => {
    await page.goto('/');
    const hero = page.locator('section').first();
    await expect(hero).toBeVisible();
    await expect(hero.locator('h1')).toContainText(/Entdecke Heidelberg|Discover Heidelberg/);
    await expect(hero.locator('img[alt="Heidelberg Panorama"]')).toBeVisible();
  });

  test('landmark grid uses 4-column layout on desktop', async ({ page }) => {
    await page.setViewportSize({ width: 1280, height: 800 });
    await page.goto('/');
    const grid = page.locator('.grid');
    await expect(grid).toBeVisible();
    await expect(grid).toHaveClass(/lg:grid-cols-4/);
  });

  test('category pills are displayed on cards', async ({ page }) => {
    await page.goto('/');
    const pills = page.locator('.rounded-full').filter({ hasText: /ARCHITEKTUR|NATUR|GESCHICHTE|KULTUR/ });
    await expect(pills.first()).toBeVisible();
    const count = await pills.count();
    expect(count).toBeGreaterThanOrEqual(4);
  });

  test('highlight badges are displayed on featured landmarks', async ({ page }) => {
    await page.goto('/');
    const badges = page.locator('span').filter({ hasText: 'Highlight' });
    await expect(badges.first()).toBeVisible();
    const count = await badges.count();
    expect(count).toBe(3);
  });

  test('highlighted landmarks appear first in grid', async ({ page }) => {
    await page.goto('/');
    const cards = page.locator('section').nth(1).locator('a[href^="/landmarks/"]');
    const firstCard = cards.first();
    await expect(firstCard.locator('span').filter({ hasText: 'Highlight' })).toBeVisible();
  });

  test('learn more link is visible on cards', async ({ page }) => {
    await page.goto('/');
    const learnMore = page.locator('span').filter({ hasText: /Mehr erfahren|Learn more/ });
    await expect(learnMore.first()).toBeVisible();
  });
});

test.describe('Detail Page', () => {
  test('breadcrumb shows Home > Landmark Name', async ({ page }) => {
    await page.goto('/landmarks/1');
    const breadcrumb = page.locator('nav').filter({ hasText: /Startseite|Home/ });
    await expect(breadcrumb).toBeVisible();
    await expect(breadcrumb.locator('a[href="/"]')).toBeVisible();
    await expect(breadcrumb.locator('span')).toContainText(/Heidelberger Schloss|Heidelberg Castle/);
  });

  test('category pill is displayed on detail page', async ({ page }) => {
    await page.goto('/landmarks/1');
    const pill = page.locator('.rounded-full').filter({ hasText: /ARCHITEKTUR|ARCHITECTURE/ });
    await expect(pill).toBeVisible();
  });

  test('highlight badge is displayed for highlighted landmark', async ({ page }) => {
    await page.goto('/landmarks/1');
    const badge = page.locator('span').filter({ hasText: 'Highlight' });
    await expect(badge).toBeVisible();
  });

  test('highlight badge is NOT displayed for non-highlighted landmark', async ({ page }) => {
    await page.goto('/landmarks/3');
    const badge = page.locator('span').filter({ hasText: 'Highlight' });
    await expect(badge).toHaveCount(0);
  });
});
