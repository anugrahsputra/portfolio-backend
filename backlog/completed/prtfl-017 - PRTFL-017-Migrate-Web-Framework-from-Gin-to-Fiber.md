---
id: PRTFL-017
title: PRTFL-017 - Migrate Web Framework from Gin to Fiber
status: Done
assignee: []
created_date: '2026-05-08 12:28'
updated_date: '2026-05-08 21:35'
labels: []
milestone: m-2 - api-&-quality-assurance.md
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Migrate the web framework from Gin to Fiber to improve performance and adopt a more modern, Express-like API. This is a delivery-layer-only refactor, keeping domain and repository layers intact.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 #1 Fiber core dependency installed and Gin removed from go.mod.
- [ ] #2 #2 Global middleware (CORS, Recovery, Security) migrated to Fiber equivalents.
- [ ] #3 #3 All delivery handlers refactored to use *fiber.Ctx and return error.
- [ ] #4 #4 SetupRouter and Main entry point updated to bootstrap Fiber app.
- [ ] #5 #5 All existing handler unit tests updated to verify Fiber implementation.
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
### Phase 1: Infrastructure & Dependencies
1. Add `github.com/gofiber/fiber/v2` and other necessary fiber middleware packages.
2. Initialize Fiber app in `cmd/api/main.go`.

### Phase 2: Middleware Migration
1. Rewrite `pkg/middleware/recovery.go` to Fiber.
2. Rewrite `pkg/middleware/security.go` or replace with `fiber/middleware/helmet`.
3. Update CORS configuration in `SetupRouter` using `fiber/middleware/cors`.
4. Migrate `AuthMiddleware`.

### Phase 3: Handler Refactoring (Iterative)
Refactor handlers and their respective route registrations for each domain:
1. Profile & ProfileURL
2. Experience & Education
3. Skill & Language
4. Project & Resume
5. Contact Form

### Phase 4: Testing & Cleanup
1. Update `test/handler/*_test.go` to use Fiber's `app.Test()` helper.
2. Run full test suite.
3. Remove Gin dependency from `go.mod`.
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Successfully migrated the entire delivery layer and server infrastructure from Gin to Fiber v3.

Key changes:
- All 9 domain handlers refactored to use Fiber v3 signatures and c.Bind().JSON().
- All 9 domain routes updated to accept fiber.Router and follow Fiber v3 syntax.
- SetupRouter refactored to return *fiber.App, with a centralized ErrorHandler and built-in Fiber middlewares (Logger, Recover, Helmet, CORS).
- cmd/api/main.go updated to use Fiber's app.Listen and app.Shutdown for graceful lifecycle management.
- AuthMiddleware migrated to fiber.Handler.
- Entire test suite in test/handler refactored to use Fiber's high-performance app.Test() helper.
- Removed all Gin-specific dependencies and unused middleware files.
- All tests passing and project builds successfully.
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All API endpoints return expected responses matching existing DTOs.
- [ ] #2 No regressions in Auth or Security middleware logic.
- [ ] #3 Project builds and runs successfully in Docker.
<!-- DOD:END -->
