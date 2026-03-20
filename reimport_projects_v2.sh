#!/bin/bash

BASE_URL="http://localhost:8082/api/v1"
PROFILE_ID="488ccb3b-d4ef-457a-9180-8f5ac5fc1654"

echo "Re-importing projects with is_present field..."

# 1. Change Project Name (PRESENT)
echo "Creating Change Project Name..."
curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"title\": \"Change Project Name\",
    \"description\": [\"A powerful CLI tool to rename Flutter/Dart projects and automatically update all package references, and imports.\"],
    \"tech_stacks\": [\"Dart\", \"CLI\", \"Flutter\"],
    \"live_demo_url\": \"https://pub.dev/packages/change_project_name\",
    \"github_repo_url\": \"https://github.com/anugrahsputra/change_project_name\",
    \"is_live\": true,
    \"is_featured\": true,
    \"image_url\": \"https://pub.dev/static/hash-e4t06sub/img/pub-dev-icon-cover-image.png\",
    \"company\": \"Personal Project\",
    \"start_date\": \"2025-09-01\",
    \"end_date\": \"\",
    \"is_present\": true,
    \"location\": \"South Jakarta, Indonesia\"
  }" | jq .

# 2. Cosmic App
echo "Creating Cosmic App..."
curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"title\": \"Cosmic App KIOSK Touchscreen\",
    \"description\": [\"Enhanced and maintained a KIOSK touchscreen application for the Bureau of Publishing (DPR RI).\", \"Refactored codebase to improve readability.\"],
    \"tech_stacks\": [\"Android\", \"Kotlin\", \"WebView\", \"AlarmManager\", \"WebSocket\"],
    \"github_repo_url\": \"https://github.com/anugrahsputra/cosmic-kiosk\",
    \"is_nda\": true,
    \"is_featured\": true,
    \"image_url\": \"https://images.unsplash.com/photo-1512941937669-90a1b58e7e9c?w=800&h=450&fit=crop&crop=center\",
    \"company\": \"PT. Semesta Arus Teknologi\",
    \"start_date\": \"2024-10-01\",
    \"end_date\": \"2025-02-01\",
    \"is_present\": false,
    \"location\": \"South Jakarta, Indonesia\"
  }" | jq .

echo "Projects re-imported successfully!"
