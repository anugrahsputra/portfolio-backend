---
id: PRTFL-009
title: Core Domain Usecase Implementation
status: Done
assignee: []
created_date: '2026-03-07 20:15'
updated_date: '2026-03-12 18:47'
labels: []
milestone: m-1
dependencies: []
priority: high
ordinal: 375
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the usecase/service layer to encapsulate business logic for all domain entities (Profile, Experience, Education, Skills, Languages). This layer will serve as the intermediary between the API handlers and the repositories, following the Clean/Hexagonal Architecture. This ensures that the business logic is decoupled from the delivery mechanism (REST API) and data access.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Define usecase interfaces for all domain entities (Profile, Experience, Education, Skills, Languages)
- [x] #2 Implement usecase logic for CRUD operations and any business rules
- [x] #3 Ensure usecases are injected with their corresponding repositories via dependency injection
- [x] #4 Implement unit tests for each usecase to verify business logic isolation
- [x] #5 Update main.go to bootstrap the usecases during startup
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 Code follows Go idiomatic practices
- [ ] #2 Changes are verified with tests (when applicable)
- [ ] #3 All TODOs addressed or tracked
- [ ] #4 Clean build and no lint errors
<!-- DOD:END -->
