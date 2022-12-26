package main

import (
	downloader "github.com/RaftechNL/kubectl-switcher/internal/downloader"
	localfs "github.com/RaftechNL/kubectl-switcher/internal/localfs"
)

func main() {
	localfs.LocalFsTest()

	downloader.DownloaderTest("1.25.0", "kubectl-switcher/output")

}
