package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var sizes = []int{16, 18, 22, 24, 32, 42, 48, 64, 84, 96, 128}

func main() {
	inputDirs := []string{"scalable/apps", "scalable/emblems"}

	var svgFiles []string
	for _, dir := range inputDirs {
		files, err := filepath.Glob(filepath.Join(dir, "*.svg"))
		if err != nil {
			log.Fatalf("Failed to read SVG files from %s: %v", dir, err)
		}
		svgFiles = append(svgFiles, files...)
	}

	for _, s := range sizes {
		for _, dir := range inputDirs {
			baseName := strings.TrimPrefix(dir, "scalable/")
			sizeDir := fmt.Sprintf("%dx%d/%s", s, s, baseName)
			os.MkdirAll(sizeDir, os.ModePerm)
			cleanDirectory(sizeDir)
		}
	}

	var wg sync.WaitGroup

	for _, icon := range svgFiles {
		wg.Add(1)
		go func(icon string) {
			defer wg.Done()
			optimizeAndGenerateIcon(icon)
		}(icon)
	}

	wg.Wait()
	fmt.Println("PNG generation completed.")
}

func cleanDirectory(dir string) {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		log.Printf("Failed to clean directory %s: %v", dir, err)
		return
	}
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			log.Printf("Failed to remove %s: %v", file, err)
		}
	}
}

func optimizeAndGenerateIcon(icon string) {
	optimizedIcon := optimizeSVG(icon)
	iconname := filepath.Base(icon[:len(icon)-4])
	category := strings.TrimPrefix(filepath.Dir(icon), "scalable/")

	for _, s := range sizes {
		sizeDir := fmt.Sprintf("%dx%d/%s", s, s, category)
		outputFile := filepath.Join(sizeDir, iconname+".png")

		fmt.Printf("Generating %s...\n", outputFile)
		cmd := exec.Command("rsvg-convert", "-w", fmt.Sprint(s), "-h", fmt.Sprint(s), "-o", outputFile, optimizedIcon)
		if err := cmd.Run(); err != nil {
			log.Printf("Failed to generate %s: %v", outputFile, err)
		}
	}
}

func optimizeSVG(icon string) string {
	optimizedIcon := icon
	fmt.Printf("Optimizing %s...\n", icon)

	cmd := exec.Command("svgo", "-i", icon, "-o", optimizedIcon)
	if err := cmd.Run(); err != nil {
		log.Printf("Failed to optimize %s: %v", icon, err)
		return icon
	}

	return optimizedIcon
}
