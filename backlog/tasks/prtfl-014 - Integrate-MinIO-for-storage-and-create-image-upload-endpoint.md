---
id: PRTFL-014
title: Integrate MinIO for storage and create image upload endpoint
status: To Do
assignee: []
created_date: '2026-03-18 17:40'
labels: []
milestone: m-0 - foundation-&-infrastructure.md
dependencies: []
priority: medium
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Integrate MinIO as the storage solution for the portfolio backend and provide an endpoint for image uploads (e.g., for profile pictures or project screenshots).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 MinIO client is successfully initialized in the backend.
- [ ] #2 New endpoint POST /api/v1/upload is implemented and functional.
- [ ] #3 Images are correctly stored in the MinIO bucket and accessible via a URL.
- [ ] #4 Error handling for invalid file types and sizes is implemented.
<!-- AC:END -->
