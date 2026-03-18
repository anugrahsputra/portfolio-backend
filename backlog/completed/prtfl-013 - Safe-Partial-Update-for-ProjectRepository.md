---
id: PRTFL-013
title: Safe-Partial-Update-for-ProjectRepository
status: Done
assignee: []
created_date: '2026-03-17 18:51'
updated_date: '2026-03-18 12:20'
labels: []
milestone: m-2 - api-&-quality-assurance.md
dependencies: []
priority: high
ordinal: 1000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Refactor the ProjectRepository.UpdateProject method to safely handle partial updates. The current implementation dereferences pointers without checking for nil, which causes panics when some fields are omitted in the request. The repository should fetch the existing record and merge it with the new values.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Add GetProjectByID query to projects.sql and regenerate db code
- [x] #2 Refactor projectRepository.UpdateProject to fetch-and-merge existing data with update input
- [x] #3 Implement nil checks for all fields in ProjectUpdateInput to prevent runtime panics during partial updates
- [ ] #4 Verify the fix with a unit test that performs a partial update on a project
<!-- AC:END -->
