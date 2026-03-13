---
id: PRTFL-011
title: Refactor Experience description to array of strings
status: Done
assignee: []
created_date: '2026-03-13 15:19'
updated_date: '2026-03-13 15:20'
labels: []
milestone: m-2
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Refactor the 'description' field in the Experience entity from a single TEXT string to an array of strings (TEXT[]). This better represents a list of responsibilities or achievements and allows for cleaner data handling in the frontend.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Update the database schema to store experience responsibilities/description as an array of strings (TEXT[])
- [x] #2 Update the Experience domain model and DTOs to use []string for description
- [x] #3 Update the repository layer and SQL queries to handle the new data type
- [x] #4 Migrate existing string descriptions to single-element arrays (optional but recommended for data integrity)
- [x] #5 Update API handlers to accept and return the array of strings for descriptions
<!-- AC:END -->
