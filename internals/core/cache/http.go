package cache

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func (cm *CacheManager) PathFor(uri string, contentId *string) string {
	if contentId != nil {
		os.MkdirAll(path.Join(cm.imgDir, *contentId), 0o775)
		return path.Join(cm.imgDir, *contentId, GetHash(uri))
	} else {
		return path.Join(cm.imgDir, "other", GetHash(uri))
	}
}

func (cm *CacheManager) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
	}

	if req.URL.Path != "/local/img" {
		http.NotFound(w, req)
	}

	target := req.URL.Query().Get("u")
	if target == "" {
		http.NotFound(w, req)
	}

	var filePath string

	contentId := req.URL.Query().Get("id")
	if contentId != "" {
		filePath = cm.PathFor(target, &contentId)
	} else {
		filePath = cm.PathFor(target, nil)
	}

	fileHash := GetHash(target)

	// Try cache
	if f, err := os.Open(filePath); err == nil {
		defer f.Close()

		header := make([]byte, 512)
		n, _ := io.ReadFull(f, header)
		contentType := http.DetectContentType(header[:n])
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		if n > 0 {
			//TODO: check cache exarations

			w.Write(header)
		}

		_, _ = io.Copy(w, f)
		return
	}

	if deadline, ok := req.Context().Deadline(); ok {
		cm.Client.Timeout = time.Until(deadline)
	}
	r, _ := http.NewRequestWithContext(req.Context(), "GET", target, nil)
	r.Header.Set("User-Agent", "Raphy-cache/1.0")
	resp, err := cm.Client.Do(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		w.WriteHeader(resp.StatusCode)
		_, _ = io.Copy(w, resp.Body)
		return
	}

	contentType := resp.Header.Get("Content-Type")
	isImage := strings.HasPrefix(strings.ToLower(contentType), "image")
	if isImage {
		tempFilePath := filePath + ".part"
		out, err := os.Create(filePath)
		if err == nil {
			defer func() {
				out.Close()
				_ = os.Rename(tempFilePath, filePath)
			}()

			w.Header().Set("Content-Type", contentType)
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")

			tr := io.TeeReader(resp.Body, out)

			err := cm.addImageCache(fileHash, contentId, resp.ContentLength)
			if err != nil {
				cm.lg.Error("cache error: " + err.Error())
			}

			_, _ = io.Copy(w, tr)
			return
		}

		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		_, _ = io.Copy(w, resp.Body)
	}
}
