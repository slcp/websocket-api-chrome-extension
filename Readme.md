## Installing Chrome extension

- Navigate to `chrome://extensions` in your Chrome/Chromium based browser
- There will be a `Developer mode` toggle somewhere - turn it on
  - This should give you a `Load Unpacked` option
- Select `Load Unpacked`
- Choose the chrome-extension folder - do not navigate into the folder

## Usage

- As soon as you install the extension it will try and open a websocket connection to the API
- On success an alert wiill pop up with a connection ID in
  - This connection ID can be used by yourself or others to send a message to the extension in your browser
- Make a `POST` request to `[api url]/send/[connectionId]
  - The request should have a body that looks like `{"message": "your message goes here"}`
  - API url will be the API Gateway endpoint (including stage) in the `serverless` output once deployed
