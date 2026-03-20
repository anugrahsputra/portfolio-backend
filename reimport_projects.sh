#!/bin/bash

BASE_URL="http://localhost:8082/api/v1"
PROFILE_ID="44c71a08-6055-41fa-ba1a-5b7179cc6886"

echo "Re-importing projects with new date format..."

# 1. Cosmic App
echo "Creating Cosmic App..."
curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"title\": \"Cosmic App KIOSK Touchscreen\",
    \"description\": [
      \"Enhanced and maintained a KIOSK touchscreen application for the Bureau of Publishing (DPR RI).\",
      \"Refactored codebase to improve readability and maintainability.\",
      \"Implemented device power scheduling using AlarmManager.\",
      \"Developed auto-return mechanism for WebView.\",
      \"Integrated WebSocket for real-time action triggering.\",
      \"Implemented heartbeat mechanism for application status.\"
    ],
    \"tech_stacks\": [\"Android\", \"Kotlin\", \"WebView\", \"WebSocket\", \"AlarmManager\"],
    \"company\": \"PT. Semesta Arus Teknologi\",
    \"start_date\": \"2024-10-01\",
    \"end_date\": \"2025-02-01\",
    \"location\": \"South Jakarta, Indonesia\",
    \"is_featured\": true
  }" | jq .

# 2. Quraani
echo "Creating Quraani..."
curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"title\": \"Quraani - Quran Mobile App\",
    \"description\": [
      \"Developed a lightweight, user-friendly Quran app for Android using Flutter.\",
      \"Focused on providing a clean interface for optimal reading experience.\",
      \"Integrated features such as bookmarks, search, and sharing verses.\"
    ],
    \"tech_stacks\": [\"Flutter\", \"Dart\", \"Android\"],
    \"company\": \"Personal Project\",
    \"start_date\": \"2024-04-01\",
    \"end_date\": \"2024-05-01\",
    \"location\": \"Lebak, Indonesia\",
    \"is_featured\": true
  }" | jq .

# 3. E-Market
echo "Creating E-Market..."
curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"title\": \"E-Market Mobile Applications\",
    \"description\": [
      \"Conducted system analysis and user requirements gathering.\",
      \"Developed dual-sided e-marketplace apps (Seller and Buyer).\",
      \"Utilized Firebase for backend integration and data storage.\",
      \"Integrated Google Maps API for user location services.\",
      \"Performed blackbox testing with 100% pass rate.\"
    ],
    \"tech_stacks\": [\"Flutter\", \"Dart\", \"Firebase\", \"Google Maps API\"],
    \"company\": \"Academic Project\",
    \"start_date\": \"2022-12-01\",
    \"end_date\": \"2023-06-01\",
    \"location\": \"Bandung, Indonesia\",
    \"is_featured\": false
  }" | jq .

echo "Projects re-imported successfully!"
