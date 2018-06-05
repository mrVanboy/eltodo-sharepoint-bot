# Eltodo's Sharepoint Bot for Slack

Simple bot, that connect's to internal sharepoint, parse news feed, check
them for the new articles and calling Slack's webhook.



## Requirements
- Docker
- Connection to the internal corporate Sharepoint
- Headless instance of Chrome ([dockerized version](https://github.com/knq/chrome-headless)  can be used)
- Google Storage Bucket

## Running
Project is completely prepared for run in scheduled Gitlab CI pipeline. So
it is possible to inspired by job *run-from-scheduler:* in `.gitlab-ci.yml`.



## Configuration
Configuration is available vie **environment variables**:
- 'WEBHOOK_MAIN_URL' - webhook to the channel in which new posts will
  appear
- 'WEBHOOK_DEBUG_URL' - webhook to the channel in which will be posted
  errors (if occurs)
- 'SHAREPOINT_URL' - url to the sharepoint (can be with credentials)
- 'TITLE_LINK' - title link that appears in message
- 'GOOGLE_STORAGE_BUCKET' - name of the bucket on google storage
- 'GOOGLE_BUCKET_OBJECT' - name of the object in bucket
- 'CHROME_URL' - url to the json entrypoint of the headless chrome instance.
For example: http://localhost:9222/json
- 'GOOGLE_APPLICATION_CREDENTIALS_JSON' - JSON with Google Cloud
credentials, smth like this:
    ```json
    {
        "type": "service_account",
        "project_id": "uber_project_id",
        "private_key_id": "some_id",
        "private_key": "-----BEGIN PRIVATE KEY-----\nbla-bla-bla==\n-----END PRIVATE KEY-----\n",
        "client_email": "username@appname.iam.gserviceaccount.com",
        "client_id": "number",
        "auth_uri": "https://accounts.google.com/o/oauth2/auth",
        "token_uri": "https://accounts.google.com/o/oauth2/token",
        "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
        "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/email" }
    ```

