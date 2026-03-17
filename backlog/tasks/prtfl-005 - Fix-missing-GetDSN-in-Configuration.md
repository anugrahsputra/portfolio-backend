---
id: PRTFL-005
title: Fix missing GetDSN() in Configuration
status: Done
assignee: []
created_date: '2026-03-17 18:51'
labels: []
milestone: m-0 - foundation-&-infrastructure.md
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Fix the missing GetDSN() method in config/config.go which is currently called in config/database.go, causing potential build issues.
<!-- SECTION:DESCRIPTION:END -->
