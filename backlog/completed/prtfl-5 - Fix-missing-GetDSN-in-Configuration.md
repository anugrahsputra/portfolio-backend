---
id: PRTFL-5
title: Fix missing GetDSN() in Configuration
status: Done
assignee: []
created_date: '2026-03-05 18:37'
updated_date: '2026-03-07 20:11'
labels: []
milestone: m-0
dependencies: []
priority: high
ordinal: 2500
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Fix the missing GetDSN() method in config/config.go which is currently called in config/database.go, causing potential build issues.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 config.Config struct has GetDSN() method returning the connection string
- [ ] #2 config/database.go compiles and can connect using GetDSN()
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 Code follows Go idiomatic practices
- [ ] #2 Changes are verified with tests (when applicable)
- [ ] #3 All TODOs addressed or tracked
- [ ] #4 Clean build and no lint errors
<!-- DOD:END -->
