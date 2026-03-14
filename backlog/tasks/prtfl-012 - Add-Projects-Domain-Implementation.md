---
id: PRTFL-012
title: Add Projects Domain Implementation
status: To Do
assignee: []
created_date: '2026-03-13 15:34'
labels: []
milestone: m-2
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement a new 'Projects' domain to showcase personal and professional projects. This should follow the existing hexagonal architecture: Schema -> SQLC -> Domain -> Repository -> Usecase -> DTO/Handler/Route.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Define 'projects' table in PostgreSQL schema with fields like name, description (TEXT[]), tools/technologies (TEXT[]), and URL
- [ ] #2 Implement SQL queries for CRUD operations in projects.sql and generate code with SQLC
- [ ] #3 Implement Domain models, Repository, Usecase, and Mapper layers for Projects
- [ ] #4 Implement REST API Handlers, DTOs, and Routes for Projects (/api/v1/project)
- [ ] #5 Verify implementation with unit tests for mapper, usecase, and handlers
<!-- AC:END -->
