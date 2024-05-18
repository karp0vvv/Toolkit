package rutubesegment

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func RutubeGetStreamURL(videoURL string) (string, error) {
	// Используем yt-dlp для получения прямого URL потока
	cmd := exec.Command("yt-dlp", "-f", "best", "-g", videoURL)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Возвращаем URL без пробелов и символов новой строки
	return strings.TrimSpace(out.String()), nil
}

func RutubeSplitVideoStreamToFrames(streamURL, outputPattern string) error {
	var wg sync.WaitGroup

	cmd := exec.Command("ffmpeg", "-i", streamURL, "-vf", "fps=1", outputPattern)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd.Wait()
	}()

	wg.Wait()
	return nil
}
