package main  // Web app launcher

import (
	"context" // New import for server shutdown
	"fmt"
	"log"
	"net/http" // New import for http.Server
	"net/url"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Global variables for user selection and server control
var (
	selectedFilePath string
	serverPort       string = "8080"
	httpServer       *http.Server    // Holds the running HTTP server instance
	serverCtx        context.Context // Context for server shutdown
	serverCancel     context.CancelFunc
)

// openBrowser launches the user's default web browser with the given URL.
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin": // macOS
		cmd = "open"
		args = []string{url}
	default: // Linux and others
		cmd = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmd, args...).Start()
}

// startServerAndLaunchBrowser runs the HTTP server and opens the browser.
func startServerAndLaunchBrowser(window fyne.Window, htmlFilePath string, port string, statusLabel *widget.Label) {
	if httpServer != nil {
		statusLabel.SetText("Server already running.")
		return
	}

	dirPath := path.Dir(htmlFilePath)
	fileName := path.Base(htmlFilePath)
	localURL := fmt.Sprintf("http://localhost:%s/%s", port, url.PathEscape(fileName))
	
	// Set up the HTTP server
	httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: http.FileServer(http.Dir(dirPath)),
	}

	// Create a context for graceful shutdown
	serverCtx, serverCancel = context.WithCancel(context.Background())

	// Start the HTTP server in a goroutine
	go func() {
		log.Printf("Starting server on %s, serving directory: %s", localURL, dirPath)
		
		// Use window.Canvas().Content().Refresh() to update UI from goroutine safely
		// Or better: create a channel-based update mechanism
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Update status label on main thread
	statusLabel.SetText(fmt.Sprintf("✅ Server running at: %s", localURL))

	// Wait a moment for the server to start (crucial step)
	time.Sleep(500 * time.Millisecond)

	// Launch the default web browser
	if err := openBrowser(localURL); err != nil {
		log.Printf("Failed to launch browser: %v. Please open manually: %s", err, localURL)
		dialog.ShowInformation("Error",
			fmt.Sprintf("Failed to launch browser. Please open the following address manually in your browser:\n%s", localURL),
			window)
	}
}

// stopServer gracefully shuts down the running HTTP server.
func stopServer(statusLabel *widget.Label) {
	if httpServer != nil {
		log.Println("Shutting down server...")
		statusLabel.SetText("Shutting down server...")
		
		// Use the context to allow up to 5 seconds for existing connections to close
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
			statusLabel.SetText(fmt.Sprintf("❌ Shutdown Error: %v", err))
		} else {
			log.Println("Server stopped successfully.")
			statusLabel.SetText("Server stopped.")
		}
		httpServer = nil
		serverCancel() // Cancel the server context
	}
}

func main() {
	// --- GUI Setup ---
	a := app.New()
	w := a.NewWindow("Web App Launcher")
	w.Resize(fyne.NewSize(600, 360))
	w.CenterOnScreen()

	// --- 1. Port Input ---
	portEntry := widget.NewEntry()
	portEntry.SetPlaceHolder("Enter server port (e.g., 8080)")
	portEntry.SetText(serverPort)
	portEntry.OnChanged = func(s string) {
		if _, err := strconv.Atoi(s); err == nil {
			serverPort = s
		}
	}
	portFormItem := widget.NewFormItem("Server Port:", portEntry)

	// --- 2. File Selection ---
	pathLabel := widget.NewLabel("No file selected.")
	fileButton := widget.NewButton("Select HTML App File", func() {
		fd := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil || read == nil {
				return
			}
			selectedFilePath = read.URI().Path()
			pathLabel.SetText(fmt.Sprintf("Selected: %s", path.Base(selectedFilePath)))
		}, w)
		fd.Show()
	})
	fileFormItem := widget.NewFormItem("HTML File:", container.NewHBox(fileButton, pathLabel))

	// --- 3. Status Label ---
	statusLabel := widget.NewLabel("Ready to launch.")

	// --- 4. Launch/Stop Buttons ---
	launchButton := widget.NewButton("Start Server & Launch Browser", func() {
		if selectedFilePath == "" {
			dialog.ShowError(fmt.Errorf("Please select an HTML file."), w)
			return
		}
		// Pass the status label to update text
		startServerAndLaunchBrowser(w, selectedFilePath, serverPort, statusLabel)
	})
	
	stopButton := widget.NewButton("Stop Server", func() {
		stopServer(statusLabel)
	})
	
	// --- 5. Clean Shutdown on Window Close ---
	w.SetOnClosed(func() {
		stopServer(statusLabel)
	})

	// --- Layout ---
	content := container.New(
		layout.NewVBoxLayout(),
		widget.NewForm(portFormItem, fileFormItem),
		widget.NewSeparator(),
		launchButton,
		stopButton,
		widget.NewSeparator(),
		container.NewCenter(statusLabel),
	)

	w.SetContent(content)
	w.ShowAndRun()
	
	// The w.ShowAndRun() call blocks here until the user closes the window.
	// When it returns, the main function exits, and the process terminates.
}
