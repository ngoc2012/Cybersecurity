package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadImages(t *testing.T) {
	// Create a test server with a simple HTML page containing an image
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><body><img src="test.jpg"></body></html>`))
	}))
	defer server.Close()

	// Create a temporary directory for the test
	tempDir := t.TempDir()

	// Call the downloadImages function
	err := downloadImages(server.URL, false, 1, tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the image was downloaded
	imagePath := filepath.Join(tempDir, "test.jpg")
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		t.Fatalf("Expected image to be downloaded, but it was not found")
	}
}

func TestSaveImage(t *testing.T) {
	// Create a test server with a simple image
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test image content"))
	}))
	defer server.Close()

	// Create a temporary directory for the test
	tempDir := t.TempDir()

	// Call the saveImage function
	err := saveImage(server.URL, tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the image was saved
	imagePath := filepath.Join(tempDir, "test image content")
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		t.Fatalf("Expected image to be saved, but it was not found")
	}

	// Check the content of the saved image
	content, err := ioutil.ReadFile(imagePath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if string(content) != "test image content" {
		t.Fatalf("Expected image content to be 'test image content', got %s", string(content))
	}
}

func TestDownloadImagesRecursive(t *testing.T) {
	// Create a test server with a simple HTML page containing a link to another page with an image
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`<html><body><a href="/page2">Link</a></body></html>`))
		} else if r.URL.Path == "/page2" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`<html><body><img src="test.png"></body></html>`))
		}
	}))
	defer server.Close()

	// Create a temporary directory for the test
	tempDir := t.TempDir()

	// Call the downloadImages function with recursive option
	err := downloadImages(server.URL, true, 2, tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the image was downloaded
	imagePath := filepath.Join(tempDir, "test.png")
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		t.Fatalf("Expected image to be downloaded, but it was not found")
	}
}

func TestDownloadImagesMultipleFormats(t *testing.T) {
	// Create a test server with a simple HTML page containing multiple image formats
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><body>
			<img src="test.jpg">
			<img src="test.png">
			<img src="test.gif">
			<img src="test.bmp">
		</body></html>`))
	}))
	defer server.Close()

	// Create a temporary directory for the test
	tempDir := t.TempDir()

	// Call the downloadImages function
	err := downloadImages(server.URL, false, 1, tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the images were downloaded
	formats := []string{"test.jpg", "test.png", "test.gif", "test.bmp"}
	for _, format := range formats {
		imagePath := filepath.Join(tempDir, format)
		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			t.Fatalf("Expected image %s to be downloaded, but it was not found", format)
		}
	}
}

func TestDownloadImagesInvalidURL(t *testing.T) {
	// Create a temporary directory for the test
	tempDir := t.TempDir()

	// Call the downloadImages function with an invalid URL
	err := downloadImages("http://invalidurl", false, 1, tempDir)
	if err == nil {
		t.Fatalf("Expected error for invalid URL, got none")
	}
}
