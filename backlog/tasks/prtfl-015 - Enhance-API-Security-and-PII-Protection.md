---
id: PRTFL-015
title: Enhance API Security and PII Protection
status: Done
assignee: []
created_date: '2026-03-22 09:21'
updated_date: '2026-03-22 09:25'
labels: []
dependencies: []
references:
  - pkg/middleware/auth.go
  - .env.example
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Enhance API security and privacy by implementing authentication for write operations, masking sensitive PII in public responses, adding standard security headers, and preventing internal error leakage.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 API-Key auth middleware for write operations (POST, PUT, DELETE) implemented in pkg/middleware/auth.go
- [ ] #2 Security headers middleware implemented in pkg/middleware/security.go
- [ ] #3 Global error handler middleware implemented in pkg/middleware/error.go
- [ ] #4 Middlewares registered in setup_router.go
- [ ] #5 Sensitive PII (address, email, phone) masked in public Profile GET responses via ProfilePublicResp DTO
- [ ] #6 Handlers updated to use generic error messages for the client
<!-- AC:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Added .env.example with API_KEY field.
<!-- SECTION:NOTES:END -->
