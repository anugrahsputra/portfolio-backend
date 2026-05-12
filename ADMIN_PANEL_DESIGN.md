# Admin Panel Design Document

## 1. Overview
The Admin Panel is a secure, web-based dashboard designed to manage the content of the personal portfolio system. It acts as a client to the existing Go-based Portfolio Backend API, allowing the profile owner to create, read, update, and delete (CRUD) entities such as their profile details, experiences, education, skills, languages, and projects.

## 2. Technical Stack & Architecture
Based on the project's design constraints, the Admin Panel will be built using the following technologies:
- **Framework**: React (TypeScript) or Next.js.
- **Styling**: Vanilla CSS (to prioritize flexibility and maintainability, avoiding external styling libraries unless necessary).
- **State Management**: React Context or lightweight tools like Zustand.
- **Data Fetching**: Native `fetch` API or `Axios`, combined with `React Query` or `SWR` for caching and state synchronization.

## 3. Authentication Strategy
The Portfolio Backend API protects its mutation endpoints (POST, PUT, DELETE) using a static API Key.
- **Mechanism**: The Admin Panel will require the user to provide the API Key upon launching the app or logging in.
- **Storage**: The API Key will be stored securely in the browser's `sessionStorage` or `localStorage`.
- **Transmission**: All requests to protected endpoints will include the header `X-API-Key` populated with the stored key.
- **Read Operations**: All GET operations are public and do not require the API key.

## 4. Entity Management & API Integration
The panel will interact with the following API endpoints exposed under `/api/v1/`:

| Entity | GET (Public) | POST (Protected) | PUT (Protected) | DELETE (Protected) |
| :--- | :--- | :--- | :--- | :--- |
| **Profile** | `/profile/:id` | `/profile` | `/profile/:id` | `/profile/:id` |
| **Profile URLs** | `/profile_url/:id` | `/profile_url` | `/profile_url/:id` | `/profile_url/:id` |
| **Experience** | `/experience/:profile_id` | `/experience` | `/experience/:id` | `/experience/:id` |
| **Education** | `/education/:profile_id` | `/education` | `/education/:id` | `/education/:id` |
| **Skills** | `/skill/:profile_id` | `/skill` | `/skill/:id` | `/skill/:id` |
| **Languages** | `/language/:profile_id` | `/language` | `/language/:id` | `/language/:id` |
| **Projects** | `/project/:profile_id` | `/project` | `/project/:id` | `/project/:id` |

*Note: The `contact_form` endpoint (`POST /contact_form/`) is public and intended for the public-facing portfolio application, not the Admin Panel.*

## 5. User Interface (Views/Pages)
The Admin Panel will be structured into the following distinct sections:

1. **Authentication (Login) View**
   - A simple prompt to enter the `X-API-Key`.
   - Validates the key by attempting a benign protected request (or relying on subsequent request failures).

2. **Dashboard / Overview**
   - High-level metrics (e.g., total projects, experiences).
   - Quick links to manage different entities.

3. **Profile Management**
   - A comprehensive form to edit core profile details (Name, Title, Bio, Email).
   - A sub-section to manage associated Profile URLs (e.g., GitHub, LinkedIn links).

4. **Experience & Education Management**
   - List views displaying current entries in chronological order.
   - Modals or separate pages to add/edit entries.
   - Forms will capture organization, title, start/end dates, and a list of descriptions/bullet points.

5. **Projects Management**
   - Grid or list view of projects.
   - Form fields for project name, description, tech stack, and URLs (e.g., live link, source code).

6. **Skills & Languages Management**
   - Simple list views to manage skills and languages, including their respective proficiency levels.

## 6. Implementation Milestones

### Milestone 1: Setup & Authentication
- Initialize the Next.js project with Tailwind CSS and shadcn/ui.
- Implement the basic layout (Sidebar, Main Content Area) utilizing shadcn/ui components.
- Create the Login view and secure API request wrapper (injecting `X-API-Key`).

### Milestone 2: Core Profile & URLs
- Build the data fetching hooks for `Profile` and `ProfileUrl`.
- Implement forms to create and update the main profile data.

### Milestone 3: Historical Data (Experience & Education)
- Implement CRUD interfaces for `Experience` and `Education`.
- Add sorting and validation for dates.

### Milestone 4: Portfolio Content (Projects, Skills, Languages)
- Implement CRUD interfaces for `Projects`, `Skills`, and `Languages`.
- Finalize the dashboard overview and UI polish.
## Milestone 4: Portfolio Content (Projects, Skills, Languages)
- Implement CRUD interfaces for `Projects`, `Skills`, and `Languages`.
- Finalize the dashboard overview and UI polish.
