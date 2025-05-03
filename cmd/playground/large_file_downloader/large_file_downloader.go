package large_file_downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

// ChunkJob represents a chunk of a file to be downloaded
type ChunkJob struct {
	URL         string
	FileName    string
	StartByte   int64
	EndByte     int64
	ChunkNumber int
	TotalChunks int
}

// Result contains information about a completed download chunk
type Result struct {
	ChunkNumber int
	BytesRead   int64
	Error       error
}

func RunLargeFileDownloader() {
	// Configure download parameters
	url := "https://releases.ubuntu.com/22.04.3/ubuntu-22.04.3-desktop-amd64.iso" // 1GB test file // 1GB test file Replace with actual file URL
	outputPath := "downloaded-file.bin"                                           // Local file path
	numWorkers := 5                                                               // Number of concurrent downloaders
	chunkSize := int64(5 * 1024 * 1024)                                           // 5MB chunks

	err := downloadLargeFile(url, outputPath, numWorkers, chunkSize)
	if err != nil {
		log.Fatalf("Download failed: %v", err)
	}

	fmt.Println("Download completed successfully!")
}

/*

The Timeline For The Concurrency Works Like This:
1. Workers start processing chunks and sending results to the buffered channel
2. Main goroutine starts reading results with the range loop
3. Once all workers finish (signaled by WaitGroup), results channel is closed
4. Main goroutine continues reading remaining results from the channel
5. When all results are read, the range loop exits

*/

func downloadLargeFile(url, outputPath string, numWorkers int, chunkSize int64) error {
	// Create a temporary directory for chunks
	tempDir, err := os.MkdirTemp("", "download-chunks")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up at the end

	// Get the file size
	fileSize, err := getFileSize(url)
	if err != nil {
		return fmt.Errorf("failed to get file size: %v", err)
	}

	fmt.Printf("File size: %d bytes (%.2f MB)\n", fileSize, float64(fileSize)/(1024*1024))

	// Calculate the number of chunks
	totalChunks := (fileSize + chunkSize - 1) / chunkSize // Ceiling division

	// Create channels for jobs and results
	jobs := make(chan ChunkJob, totalChunks)
	results := make(chan Result, totalChunks)

	// Start worker pool
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, &wg, jobs, results)
	}

	// Create all jobs
	go func() {
		for i := int64(0); i < totalChunks; i++ {
			startByte := i * chunkSize
			endByte := (i+1)*chunkSize - 1
			if endByte >= fileSize {
				endByte = fileSize - 1
			}

			chunkFileName := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", i))

			jobs <- ChunkJob{
				URL:         url,
				FileName:    chunkFileName,
				StartByte:   startByte,
				EndByte:     endByte,
				ChunkNumber: int(i),
				TotalChunks: int(totalChunks),
			}
		}
		close(jobs)
	}()

	// Start a goroutine to collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	progressTracker := make([]bool, totalChunks)
	completed := 0

	// Start time for download speed calculation
	startTime := time.Now()

	for result := range results {
		if result.Error != nil {
			return fmt.Errorf("chunk %d failed: %v", result.ChunkNumber, result.Error)
		}

		progressTracker[result.ChunkNumber] = true
		completed++

		// Calculate progress and speed
		percentage := float64(completed) / float64(totalChunks) * 100
		elapsed := time.Since(startTime).Seconds()
		bytesDownloaded := calculateBytesDownloaded(tempDir)
		speed := float64(bytesDownloaded) / elapsed / (1024 * 1024) // MB/s

		fmt.Printf("\rProgress: %.1f%% (Chunks: %d/%d) - Speed: %.2f MB/s",
			percentage, completed, totalChunks, speed)
	}
	fmt.Println() // New line after progress is complete

	// Merge chunks into final file
	err = mergeChunks(tempDir, outputPath, int(totalChunks))
	if err != nil {
		return fmt.Errorf("failed to merge chunks: %v", err)
	}

	return nil
}

func worker(id int, wg *sync.WaitGroup, jobs <-chan ChunkJob, results chan<- Result) {
	defer wg.Done()

	for job := range jobs {
		log.Printf("Worker %d downloading chunk %d/%d (bytes %d-%d)",
			id, job.ChunkNumber+1, job.TotalChunks, job.StartByte, job.EndByte)

		bytesRead, err := downloadChunk(job)

		results <- Result{
			ChunkNumber: job.ChunkNumber,
			BytesRead:   bytesRead,
			Error:       err,
		}
	}
}

func downloadChunk(job ChunkJob) (int64, error) {
	// Create a file to save the chunk
	out, err := os.Create(job.FileName)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	// Create HTTP request with Range header
	req, err := http.NewRequest("GET", job.URL, nil)
	if err != nil {
		return 0, err
	}

	// Set range header
	rangeHeader := fmt.Sprintf("bytes=%d-%d", job.StartByte, job.EndByte)
	req.Header.Set("Range", rangeHeader)

	// Send the request
	client := &http.Client{
		Timeout: 30 * time.Minute, // Long timeout for large files
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Check if we got the expected range response
	if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Copy data to file
	bytesRead, err := io.Copy(out, resp.Body)
	if err != nil {
		return 0, err
	}

	return bytesRead, nil
}

func mergeChunks(tempDir, outputPath string, totalChunks int) error {
	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Merge all chunks in order
	for i := 0; i < totalChunks; i++ {
		chunkPath := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", i))

		// Open chunk file
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			return err
		}

		// Copy chunk to output file
		_, err = io.Copy(outFile, chunkFile)
		chunkFile.Close() // Close as soon as we're done

		if err != nil {
			return err
		}
	}

	return nil
}

func getFileSize(url string) (int64, error) {
	// Create a HEAD request to get headers only
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Check if the server supports range requests
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("server returned non-200 status: %d %s", resp.StatusCode, resp.Status)
	}

	// Get Content-Length header
	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "" {
		return 0, fmt.Errorf("content length header not found")
	}

	// Parse the content length
	size, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid content length: %s", contentLength)
	}

	return size, nil
}

func calculateBytesDownloaded(tempDir string) int64 {
	var totalSize int64 = 0

	// Walk through the temp directory and sum up file sizes
	filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	return totalSize
}

// For sites that don't support range requests, this is an alternative approach
func downloadWithoutRanges(url, outputPath string) error {
	// Create the output file
	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Create HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Create a buffered writer
	writer := io.Writer(out)

	// Show progress if Content-Length is available
	if resp.ContentLength > 0 {
		fmt.Printf("Downloading %d bytes...\n", resp.ContentLength)
		progress := &ProgressWriter{
			Total:     resp.ContentLength,
			Writer:    writer,
			StartTime: time.Now(),
		}
		writer = progress
	}

	// Copy data with buffer
	_, err = io.Copy(writer, resp.Body)
	return err
}

// ProgressWriter tracks download progress
type ProgressWriter struct {
	Total     int64
	Current   int64
	Writer    io.Writer
	StartTime time.Time
	LastPrint time.Time
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n, err := pw.Writer.Write(p)
	pw.Current += int64(n)

	// Update progress at most 10 times per second
	now := time.Now()
	if now.Sub(pw.LastPrint) >= 100*time.Millisecond {
		pw.LastPrint = now

		// Calculate progress percentage
		percent := float64(pw.Current) / float64(pw.Total) * 100

		// Calculate speed
		elapsed := now.Sub(pw.StartTime).Seconds()
		speed := float64(pw.Current) / elapsed / (1024 * 1024) // MB/s

		fmt.Printf("\rDownloading... %.1f%% (%.2f MB/%.2f MB) at %.2f MB/s",
			percent,
			float64(pw.Current)/(1024*1024),
			float64(pw.Total)/(1024*1024),
			speed)
	}

	return n, err
}
