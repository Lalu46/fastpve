package vmdownloader

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/kspeeder/blobDownload/blobDownloader"
)

// Integration-style check that the GHCR package for Win10 Chinese Simplified
// is reachable and contains at least one file. Skips by default to avoid
// network usage and large downloads.
func TestGHCRWin10Metadata(t *testing.T) {
	if os.Getenv("RUN_GHCR_META") == "" {
		t.Skip("set RUN_GHCR_META=1 to enable GHCR metadata check")
	}

	ref, err := ghcrWindowsReference(Win10, "Chinese (Simplified)")
	if err != nil {
		t.Fatalf("ghcrWindowsReference returned error for Win10: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	api, spec, err := buildRegistryClient(ref)
	if err != nil {
		t.Fatalf("buildRegistryClient: %v", err)
	}
	dl, err := blobDownloader.New(ctx, api, spec)
	if err != nil {
		t.Fatalf("init blob downloader: %v", err)
	}
	files := dl.Files()
	if len(files) == 0 {
		t.Fatalf("no files found in GHCR package")
	}
	if files[0].Size <= 0 {
		t.Fatalf("unexpected file size: %+v", files[0])
	}
}
