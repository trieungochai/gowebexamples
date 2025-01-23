This example demonstrates how to use **WebSockets** in a client-server application. The WebSocket protocol allows for two-way communication between the client (browser) and the server over a single, long-lived connection, enabling real-time data transfer.

### Frontend (HTML + JavaScript) - `websockets.html`

1. **WebSocket Setup**:

   - A WebSocket connection is established to the server running on `ws://localhost:8080/echo`.
   - The `WebSocket` object (`socket`) is created, which connects the client to the server.

2. **Event Handlers**:

   - `socket.onopen`: This function runs once the WebSocket connection is established. It appends the status "Connected" to the output element.
   - `socket.onmessage`: This function is triggered every time a message is received from the server. It appends the message received from the server to the output element.

3. **Sending Data**:

   - The `send()` function is called when the user clicks the "Send" button. It sends the text entered in the input field (`input.value`) to the WebSocket server using `socket.send(input.value)`. After sending the message, the input field is cleared.

4. **Output Display**:
   - Messages received from the server are appended to the `<pre id="output">` element for display.

---

### Backend (Go) - `websockets.go`

1. **Upgrading the HTTP Connection to a WebSocket**:

   - The `websocket.Upgrader` is used to upgrade the incoming HTTP request to a WebSocket connection.
   - The `upgrader.Upgrade(w, r, nil)` call performs this upgrade, allowing the server to handle WebSocket-specific messages (as opposed to regular HTTP requests).

2. **Echo Functionality**:

   - The server listens for messages from the client by calling `conn.ReadMessage()`. This reads the message sent by the browser.
   - The server then logs the received message to the console using `fmt.Printf`.
   - After reading the message, the server sends it back to the client with `conn.WriteMessage(msgType, msg)`. This is an **echo server**: it simply sends back whatever it receives.

3. **HTTP Server**:

   - `http.HandleFunc("/echo", ...)` defines the WebSocket handler on the `/echo` endpoint. This is the route where the client connects using WebSocket.
   - `http.HandleFunc("/", ...)` serves the `websockets.html` file when the root path (`/`) is requested by the client (this is the initial page the client loads).

4. **Start Server**:
   - The server starts on port `8080` with `http.ListenAndServe(":8080", nil)`.

### Key Points:

- **WebSocket**: The WebSocket protocol enables bidirectional communication. Instead of making multiple HTTP requests (which are one-way), WebSockets allow the client and server to communicate over a persistent connection.
- **Upgrade Process**: The `websocket.Upgrader` is a key component that converts an HTTP request into a WebSocket connection. This upgrade is necessary because the WebSocket protocol is different from the HTTP protocol, but WebSockets use the HTTP handshake to establish the connection.

- **Echo Server**: The server in this example simply echoes back any message sent by the client. This is a common pattern used for testing and debugging WebSocket communication.

### How the Application Works:

1. The client (browser) opens a WebSocket connection to the server at `ws://localhost:8080/echo`.
2. The user enters text in the input field and clicks the "Send" button. This triggers the `send()` function, which sends the input value over the WebSocket connection to the server.
3. The server receives the message, prints it to the console, and sends the same message back to the client (echo).
4. The client receives the message from the server and displays it in the `<pre id="output">` element.
