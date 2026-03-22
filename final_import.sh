#!/bin/bash

BASE_URL="http://localhost:8082/api/v1"
TODAY="2099-12-31"
API_KEY="your_secret_api_key_here" # Change this to match your .env

echo "Creating profile..."
PROFILE_RESPONSE=$(curl -s -X POST "$BASE_URL/profile" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d '{
    "name": "Anugrah Surya Putra",
    "about": "Mobile Engineer with 2+ years of experience building and maintaining cross-platform applications using Flutter and Kotlin Multiplatform. Strong focus on clean architecture, production stability, CI/CD workflows, and close collaboration with backend teams to deliver reliable, high-quality mobile experiences.",
    "address": "Jakarta, Indonesia",
    "email": "anugrahsputra@gmail.com",
    "phone": "+6283812134055"
  }')

PROFILE_ID=$(echo $PROFILE_RESPONSE | jq -r '.data.id')

if [ "$PROFILE_ID" == "null" ] || [ -z "$PROFILE_ID" ]; then
  echo "Failed to create profile: $PROFILE_RESPONSE"
  exit 1
fi

echo "Profile created: $PROFILE_ID"

# 2. LinkedIn URL
curl -s -X POST "$BASE_URL/profile-url" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{\"profile_id\":\"$PROFILE_ID\",\"label\":\"LinkedIn\",\"url\":\"https://linkedin.com/in/anugrahsputra\"}" > /dev/null

# 3. Experience
echo "Adding experiences..."
curl -s -X POST "$BASE_URL/experience" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"company\": \"BRIK Indonesia\",
    \"position\": \"Mobile Engineer\",
    \"description\": [
        \"Maintained two Flutter mobile applications: BRIK Hub and PaintPro Loyalty.\",
        \"Collaborated closely with backend engineers to define requirements.\",
        \"Diagnosed and resolved production issues.\",
        \"Refactored shared components and business logic.\",
        \"Owned app release management across app stores.\",
        \"Led integration of Shorebird code push into CI/CD pipeline.\",
        \"Monitored app performance using Firebase Analytics, Crashlytics, and Performance Monitoring.\"
    ],
    \"start_date\": \"2025-08-01\",
    \"end_date\": \"$TODAY\"
}" > /dev/null

curl -s -X POST "$BASE_URL/experience" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"company\": \"PT. Semesta Arus Teknologi\",
    \"position\": \"Mobile Engineer\",
    \"description\": [
        \"Developed cross-platform mobile applications using Flutter and Compose Multiplatform.\",
        \"Participated in full development lifecycle aligned with SDLC.\",
        \"Applied clean architecture principles.\",
        \"Collaborated with backend engineers to integrate RESTful APIs.\",
        \"Diagnosed and resolved critical bugs and performance issues.\",
        \"Implemented new features and enhancements.\",
        \"Contributed to technical documentation.\"
    ],
    \"start_date\": \"2024-06-01\",
    \"end_date\": \"2025-07-31\"
}" > /dev/null

curl -s -X POST "$BASE_URL/experience" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"company\": \"PT. Inovasi Karya Mahendra (INKARA)\",
    \"position\": \"Mobile Developer Apprenticeship\",
    \"description\": [
        \"Developed app prototype 'Kantin Virtual' for the Faculty of Engineering at Universitas Pasundan.\",
        \"Designed and implemented user interface using Flutter Framework.\",
        \"Completed the project independently within a three-month apprenticeship.\"
    ],
    \"start_date\": \"2023-03-01\",
    \"end_date\": \"2023-05-31\"
}" > /dev/null

# 4. Education
echo "Adding education..."
curl -s -X POST "$BASE_URL/education" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"school\": \"Universitas Pasundan\",
    \"degree\": \"Bachelor of Engineering\",
    \"field_of_study\": \"Computer Science\",
    \"gpa\": 4.0,
    \"start_date\": \"2019-09-01\",
    \"graduation_date\": \"2023-07-01\"
}" > /dev/null

# 5. Skills
echo "Adding skills..."
curl -s -X POST "$BASE_URL/skill" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"tools\": [\"Android Studio\", \"Figma\", \"VS Code\", \"Neovim\"],
    \"technologies\": [\"Android\", \"Kotlin\", \"Compose Multiplatform\", \"Dart\", \"Flutter\", \"Firebase\", \"REST API\", \"Git\"],
    \"hard_skills\": [\"Third-party API Integration\", \"Performance Optimization\"],
    \"soft_skills\": [\"Problem Solving\", \"Adaptability\", \"Communication\"]
}" > /dev/null

# 6. Languages
echo "Adding languages..."
curl -s -X POST "$BASE_URL/language" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{\"profile_id\":\"$PROFILE_ID\",\"language\":\"English\",\"proficiency\":\"professional\"}" > /dev/null

curl -s -X POST "$BASE_URL/language" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{\"profile_id\":\"$PROFILE_ID\",\"language\":\"Bahasa Indonesia\",\"proficiency\":\"native\"}" > /dev/null

# 7. Projects
echo "Adding projects..."
curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
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
    \"end_date\": \"2025-09-30\",
    \"location\": \"South Jakarta, Indonesia\"
}" > /dev/null

curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
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
    \"location\": \"South Jakarta, Indonesia\"
}" > /dev/null

curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"title\": \"Quraani - Quran Mobile App\",
    \"description\": [\"Simple and easy-to-use Quran app for Android built with Flutter.\"],
    \"tech_stacks\": [\"Flutter\", \"Dio\", \"Bloc\", \"SQLite\"],
    \"live_demo_url\": \"https://github.com/anugrahsputra/qurani-app/releases\",
    \"github_repo_url\": \"https://github.com/anugrahsputra/quraani\",
    \"is_live\": true,
    \"is_featured\": true,
    \"image_url\": \"https://raw.githubusercontent.com/anugrahsputra/qurani-app/main/shots/Frame%201.png\",
    \"company\": \"Personal Project\",
    \"start_date\": \"2024-04-01\",
    \"end_date\": \"2024-05-31\",
    \"location\": \"Lebak, Indonesia\"
}" > /dev/null

curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"title\": \"E-Market Seller Mobile Applications\",
    \"description\": [\"Designed for sellers to sell products online, specifically for UMKM in Kecamatan Malingping.\"],
    \"tech_stacks\": [\"Flutter\", \"Firebase\", \"Google Maps API\"],
    \"github_repo_url\": \"https://github.com/anugrahsputra/emarket-seller.git\",
    \"image_url\": \"https://raw.githubusercontent.com/anugrahsputra/emarket-seller/refs/heads/main/shots-seller.png\",
    \"company\": \"Academic Project\",
    \"start_date\": \"2022-12-01\",
    \"end_date\": \"2023-06-30\",
    \"location\": \"Bandung, Indonesia\"
}" > /dev/null

curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"title\": \"E-Market Buyer Mobile Applications\",
    \"description\": [\"Platform for buyers in Kecamatan Malingping to browse and purchase products online.\"],
    \"tech_stacks\": [\"Flutter\", \"Firebase\", \"Google Maps API\"],
    \"github_repo_url\": \"https://github.com/anugrahsputra/emarket-buyer.git\",
    \"image_url\": \"https://raw.githubusercontent.com/anugrahsputra/emarket-buyer/refs/heads/main/assets/screenshot/shots.png\",
    \"company\": \"Academic Project\",
    \"start_date\": \"2022-12-01\",
    \"end_date\": \"2023-06-30\",
    \"location\": \"Bandung, Indonesia\"
}" > /dev/null

curl -s -X POST "$BASE_URL/project" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -d "{
    \"profile_id\": \"$PROFILE_ID\",
    \"title\": \"Get Wallpaper\",
    \"description\": [\"Allows users to browse, download, and set wallpapers from the internet.\"],
    \"tech_stacks\": [\"Flutter\", \"Pexels API\", \"Bloc\"],
    \"live_demo_url\": \"https://weather-app-demo.web.app\",
    \"github_repo_url\": \"https://github.com/anugrahsputra/get_wallpaper.git\",
    \"is_live\": true,
    \"image_url\": \"https://raw.githubusercontent.com/anugrahsputra/get_wallpaper/refs/heads/main/screenshot/Frame%207.png\",
    \"company\": \"Personal Project\",
    \"start_date\": \"2023-12-01\",
    \"end_date\": \"2024-02-29\",
    \"location\": \"Jakarta, Indonesia\"
}" > /dev/null

echo "All data imported successfully!"
