# ValorantICS

This project fetches the schedule for Valorant Champions Tour (VCT) professional match schedules using the `valorantesports.com` GraphQL API and automatically generates region-specific ICS (iCalendar) files. This allows users to easily subscribe to or import upcoming VCT matches for their chosen region directly into their preferred calendar application (e.g., Google Calendar, Apple Calendar, Outlook).

## Access the Calendar

You can subscribe to or download the generated VCT schedule for specific regions directly from GitHub Pages. Replace `[REGION]` with one of the following: `emea`, `americas`, `pacific`, `china`, `international`.

**Subscribe to a regional calendar (recommended for automatic updates):**

Use the following URL structure, replacing `[REGION]` with the desired region:

`https://bafifi.github.io/ValorantICS/valorant_[REGION].ics`

**Available Regional Calendars:**

*   **EMEA:** `https://bafifi.github.io/ValorantICS/valorant_emea.ics`
*   **AMERICAS:** `https://bafifi.github.io/ValorantICS/valorant_americas.ics`
*   **PACIFIC:** `https://bafifi.github.io/ValorantICS/valorant_pacific.ics`
*   **CHINA:** `https://bafifi.github.io/ValorantICS/valorant_china.ics`
*   **INTERNATIONAL:** `https://bafifi.github.io/ValorantICS/valorant_international.ics`

**How to subscribe (using a regional URL as an example):**

*   **Google Calendar:**
    1.  Open Google Calendar.
    2.  On the left, next to "Other calendars," click the plus sign `+`.
    3.  Select "From URL."
    4.  Paste the URL for your desired region (e.g., `https://bafifi.github.io/ValorantICS/valorant_emea.ics`) and click "Add calendar."

*   **Apple Calendar (macOS):**
    1.  Open Calendar.
    2.  Go to `File` > `New Calendar Subscription...`.
    3.  Paste the URL for your desired region and click "Subscribe."
    4.  You can customize the name, color, and refresh frequency.

*   **Outlook (Web):**
    1.  Go to your Outlook Calendar.
    2.  Click "Add calendar" on the left.
    3.  Choose "Subscribe from web."
    4.  Paste the URL for your desired region, give it a name, and click "Import."

*   **Outlook (Desktop):**
    1.  Open Outlook.
    2.  Go to `File` > `Account Settings` > `Account Settings...`.
    3.  Select the "Internet Calendars" tab.
    4.  Click "New...", paste the URL for your desired region, and click "Add."
    5.  Confirm and close.

**Download the latest ICS file (for a one-time import):**

Click the links below to download the ICS file for your desired region. Please note that these files will not automatically update.You can also visit the [ValorantICS project page](https://bafifi.github.io/ValorantICS/) to view all available links and easily copy them.


*   [Download EMEA schedule](https://bafifi.github.io/ValorantICS/valorant_emea.ics)
*   [Download AMERICAS schedule](https://bafifi.github.io/ValorantICS/valorant_americas.ics)
*   [Download PACIFIC schedule](https://bafifi.github.io/ValorantICS/valorant_pacific.ics)
*   [Download CHINA schedule](https://bafifi.github.io/ValorantICS/valorant_china.ics)
*   [Download INTERNATIONAL schedule](https://bafifi.github.io/ValorantICS/valorant_international.ics)

## Features

-   Fetches upcoming Valorant Champions Tour (VCT) match schedules from `valorantesports.com`.
-   Parses detailed match data, including participating teams, match dates, times, and event information.
-   Generates region-specific ICS files compatible with most calendar applications.
-   Automatically deployed and updated via GitHub Actions to GitHub Pages.

## Technologies Used

-   Go
-   `valorantesports.com` GraphQL API
-   GitHub Actions (for automation)
-   GitHub Pages (for hosting the ICS files)

## For Developers

If you want to run this project locally or contribute:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/ValorantICS.git
    cd ValorantICS
    ```

2.  **Run the application:**
    ```bash
    go run main.go
    ```
    This will generate the regional `valorant_[REGION].ics` files in the `output/` directory.
