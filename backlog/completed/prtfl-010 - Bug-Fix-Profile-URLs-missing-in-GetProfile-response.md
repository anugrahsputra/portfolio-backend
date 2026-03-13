---
id: PRTFL-010
title: 'Bug Fix: Profile URLs missing in GetProfile response'
status: Done
assignee: []
created_date: '2026-03-13 14:35'
updated_date: '2026-03-13 14:35'
labels: []
milestone: m-2
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The GetProfile API endpoint was returning an empty array for URLs even when they existed in the database. This was caused by a mismatch between the lowercase JSON keys returned by PostgreSQL's json_agg/json_build_object and the expected PascalCase field names in the Go domain.ProfileUrl struct.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Identify the root cause of empty URL arrays in the Profile response
- [x] #2 Implement a fix (either via JSON tags or SQL alias casing) to ensure correct unmarshaling
- [x] #3 Verify that the GetProfile endpoint returns the associated URLs correctly
<!-- AC:END -->
