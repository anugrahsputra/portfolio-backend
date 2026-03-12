---
id: PRTFL-2
title: Database Schema & SQLC Configuration
status: Done
assignee: []
created_date: '2026-03-05 18:37'
updated_date: '2026-03-07 20:05'
labels: []
milestone: m-0
dependencies: []
priority: high
ordinal: 3000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Define the PostgreSQL database schema and use SQLC to generate type-safe database access code in Go.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 PostgreSQL schema defined in sql/schema/schema.sql
- [x] #2 SQL queries for all entities defined in sql/queries/*.sql
- [x] #3 sqlc generate produces valid Go code in internal/db/
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 Code follows Go idiomatic practices
- [x] #2 Changes are verified with tests (when applicable)
- [x] #3 All TODOs addressed or tracked
- [x] #4 Clean build and no lint errors
<!-- DOD:END -->
