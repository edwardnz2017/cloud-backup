# Cloud File Backup CLI

## Setup

To use this application, you need to create your own OAuth 2.0 `credentials.json` file from the Google Cloud Console., and store it in the base of this directory Follow these steps to generate the credentials:

### Step 1: Go to Google Cloud Console

1. Open [Google Cloud Console](https://console.cloud.google.com/).
2. If you don’t have a Google account, create one.
3. Once logged in, click on the **Project Selector** at the top and select **New Project**.
4. Give your project a name (e.g., `My Cloud Backup CLI`) and click **Create**.

### Step 2: Enable Google Drive API

1. In the **Google Cloud Console**, navigate to **APIs & Services** > **Library**.
2. Search for **Google Drive API**.
3. Click on **Google Drive API** and press **Enable**.

### Step 3: Create OAuth 2.0 Credentials

1. Go to **APIs & Services** > **Credentials** on the left sidebar.
2. Click **Create Credentials** and choose **OAuth 2.0 Client ID**.
3. **Configure the OAuth Consent Screen**:
   - If prompted, select **External** as the user type.
   - Fill in your app details (e.g., App Name, Support Email).
   - You can skip the remaining steps by clicking **Save and Continue** until the consent screen is configured.
4. For **Application Type**, choose **Desktop App**.
5. Enter a name for your OAuth credentials (e.g., `My Cloud Backup CLI Credentials`), and click **Create**.
6. Click **Download** to download the `credentials.json` file.

### Step 4: Add Test Users

Since your project is still in testing mode and hasn't been verified by Google, you'll need to add test users to access the app.

1. Go to **APIs & Services** > **OAuth consent screen**.
2. Scroll down to the **Test Users** section.
3. Click **+Add Users**.
4. Enter the email address of the Google account(s) that will use the app in testing mode.
5. Click **Save** to add the users.

### Step 5: Rename to `credentials.json` and add to Your Project Directory

1. After downloading the `credentials.json` file, place it in the root of your project directory where the Go application will be able to access it.
   ```bash
   /cloud-backup
     ├── main.go
     └── credentials.json  # Place the credentials here
   ```
