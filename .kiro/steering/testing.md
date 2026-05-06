---
inclusion: always
---

# Testing Principles

## Unit Tests
- Write unit tests for business logic, handlers, and utility functions
- Use Go's standard `testing` package
- Keep tests focused and fast

## Integration Tests
- Do NOT write heavyweight integration tests
- Avoid tests that require spinning up databases, external services, or complex infrastructure

## Visual Regression / E2E
- Use Playwright for visual regression checks and basic UI verification
- Focus on catching unintended visual changes rather than full E2E flows
- Keep Playwright tests lightweight and maintainable

## General
- Prefer testing behavior over implementation details
- Tests should be runnable via `make test` (unit) and `make test-visual` (Playwright)
