package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// Default values
const (
	defaultDepth   = 5
	defaultPath    = "./data/"
	defaultExts    = ".jpg|.jpeg|.png|.gif|.bmp"
)

func main() {
	recursive := flag.Bool("r", false, "Recursively download images")
	depth := flag.Int("l", defaultDepth, "Maximum depth level of the recursive download")
	path := flag.String("p", defaultPath, "Path where the downloaded files will be saved")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: ./spider [-rlp] URL")
		return
	}

	url := flag.Args()[0]

	if err := downloadImages(url, *recursive, *depth, *path); err != nil {
		fmt.Println("Error:", err)
	}
}

func downloadImages(url string, recursive bool, depth int, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// Default values
const (
	defaultDepth   = 5
	defaultPath    = "./data/"
	defaultExts    = ".jpg|.jpeg|.png|.gif|.bmp"
)

func main() {
	recursive := flag.Bool("r", false, "Recursively download images")
	depth := flag.Int("l", defaultDepth, "Maximum depth level of the recursive download")
	path := flag.String("p", defaultPath, "Path where the downloaded files will be saved")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: ./spider [-rlp] URL")
		return
	}

	url := flag.Args()[0]

	if err := downloadImages(url, *recursive, *depth, *path); err != nil {
		fmt.Println("Error:", err)
	}
}

func downloadImages(url string, recursive bool, depth int, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch %s", url)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					imgURL := a.Val
					if !strings.HasPrefix(imgURL, "http") {
						imgURL = url + imgURL
					}
					if match, _ := regexp.MatchString(defaultExts, imgURL); match {
						if err := saveImage(imgURL, path); err != nil {
							fmt.Println("Error saving image:", err)
						}
					}
				}
			}
		}
		if recursive && depth > 0 && n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link := a.Val
					if !strings.HasPrefix(link, "http") {
						link = url + link
					}
					if err := downloadImages(link, recursive, depth-1, path); err != nil {
						fmt.Println("Error downloading images from:", link, err)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nil
}

func saveImage(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	filePath := filepath.Join(path, filename)

	// Create the directory if it does not exist
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

