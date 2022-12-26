package downloader

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func DownloaderTest(version, outputDir string) {
	// Create a Resty Client
	client := resty.New()

	// Setting output directory path, If directory not exists then resty creates one!
	// This is optional one, if you're planning using absoule path in
	// `Request.SetOutput` and can used together.
	client.SetOutputDirectory(outputDir)

	_, err := client.R().
		SetOutput(fmt.Sprintf("kubectl_v%s", version)).
		Get(fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/v%s/bin/linux/amd64/kubectl", version))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("eeeee")
}
