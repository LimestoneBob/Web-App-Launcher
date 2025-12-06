# Web App Launcher 🚀

A lightweight desktop utility built with [Fyne](https://fyne.io/) that lets you **select and launch local HTML applications** in your default web browser. The app spins up a temporary HTTP server, serves your chosen HTML file, and opens it automatically — perfect for testing or running standalone web apps locally.

## ✨ Features
- **Cross-platform browser launch**  
  Works on Windows, macOS, and Linux using native system commands.
- **Simple GUI**  
  Built with Fyne for a clean, user-friendly interface.
- **Customizable port**  
  Choose the server port (default: `8080`) directly in the app.
- **File picker integration**  
  Select any local HTML file to serve.
- **Status updates**  
  Real-time feedback on server state (running, stopped, errors).
- **Graceful shutdown**  
  Stops the HTTP server safely when requested or on window close.

## 🖥️ How It Works
1. Select an HTML file from your system.
2. Specify the server port (optional).
3. Click **Start Server & Launch Browser**.  
   - A local HTTP server is started.  
   - Your default browser opens the file at `http://localhost:<port>/<file>`.  
4. Stop the server anytime with the **Stop Server** button.

## 🔧 Tech Stack
- **Language:** Go
- **GUI Framework:** Fyne v2
- **Web Serving:** `net/http`
- **Graceful Shutdown:** `context` with timeout handling

## 📦 Use Cases
- Quickly preview local HTML apps.
- Test Progressive Web Apps (PWAs) in a browser.
- Share static prototypes without deploying to a remote server.
- Lightweight alternative to heavier dev servers.

---
