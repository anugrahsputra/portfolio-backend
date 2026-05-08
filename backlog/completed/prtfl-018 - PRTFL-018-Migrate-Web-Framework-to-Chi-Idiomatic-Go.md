---
id: PRTFL-018
title: PRTFL-018 - Migrate Web Framework to Chi (Idiomatic Go)
status: Done
assignee: []
created_date: '2026-05-08 22:21'
updated_date: '2026-05-08 22:32'
labels: []
milestone: m-2 - api-&-quality-assurance.md
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Migrate the web framework from Fiber to Chi to align with idiomatic Go patterns and the net/http standard library. This improves long-term maintainability and interoperability.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 #1 Chi v5 and go-chi/render installed. Fiber removed.
- [ ] #2 #2 AuthMiddleware refactored to use standard http.Handler signature.
- [ ] #3 #3 All domain handlers refactored to func(w http.ResponseWriter, r *http.Request).
- [ ] #4 #4 SetupRouter updated to use chi.NewRouter() with nested routes.
- [ ] #5 #5 Full test suite updated to use httptest.NewRecorder() (Go standard).
<!-- AC:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Successfully migrated the entire project from Fiber v3 to Chi, aligning with idiomatic Go patterns and the net/http standard library.

Key Achievements:
- All domain handlers refactored to standard func(w http.ResponseWriter, r *http.Request) signatures.
- Implemented standardized ResponseJSON and ResponseError helpers in internal/delivery/handler/response.go.
- Refactored AuthMiddleware to use standard middleware patterns.
- SetupRouter updated to use chi.NewRouter() with nested routes and standard middlewares (Logger, Recoverer, rs/cors).
- cmd/api/main.go restored to use standard http.Server with graceful shutdown context.
- Entire handler test suite refactored back to standard httptest.NewRecorder() and r.ServeHTTP() patterns.
- Removed Fiber v3 dependencies and established a clean, standard library-first delivery layer.

The project now follows idiomatic Go practices, ensuring better interoperability with the Go ecosystem and long-term maintainability.
<!-- SECTION:FINAL_SUMMARY:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 Project uses 100% standard net/http compatible interfaces.
- [ ] #2 Centralized Response helpers implemented for JSON/Error consistency.
<!-- DOD:END -->
