package downloader

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestFile(t *testing.T, path string, content []byte) {
	err := os.WriteFile(path, content, 0644)
	require.NoError(t, err)
}

func createTestDir(t *testing.T, dirPath string) {
	err := os.MkdirAll(dirPath, 0755)
	require.NoError(t, err)
}

func cleanup(paths ...string) {
	for _, path := range paths {
		os.RemoveAll(path)
	}
}

func TestDownloadLocal(t *testing.T) {
	for _, test := range []struct {
		src, dst string
	}{
		{"a/a.txt", "b/b.txt"},
		{"a/a.txt", "b/"},
		{"a/", "b/"},
	} {
		// testDir := t.TempDir()
		testDir := "./test"
		createTestDir(t, testDir+"/a")
		createTestFile(t, testDir+"/a/a.txt", []byte("hello"))

		err := DownloadLocal(context.Background(), testDir+"/"+test.src, testDir+"/"+test.dst)
		cleanup(testDir)
		require.NoError(t, err)
	}
}

func TestDownloadHttp(t *testing.T) {
	err := DownloadUrl(context.Background(), "https://raw.githubusercontent.com/gosuda/unipath/main/go.mod", "test/go.mod")
	require.NoError(t, err)
}

func TestOpenBrowser(t *testing.T) {
	err := OpenBrowser("https://github.com/gosuda")
	require.NoError(t, err)
}
