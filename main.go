package main

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/suessflorian/gqlfetch"
)

var (
	token      string
	host       string
	apiPath    string
	outputPath string
)

func main() {
	workingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while getting working directory, %v\n", err)
		os.Exit(1)
	}

	flag.StringVar(&token, "token", "", "Bearer token to use it in the header of the http request in case the server is protected")
	flag.StringVar(&host, "host", "", "URL of the host")
	flag.StringVar(&apiPath, "api-path", "api/graphql/", "Path of the graphql api")
	flag.StringVar(&outputPath, "output", filepath.Join(workingDirectory, "schema.graphql"), "Destination to save fetched schema file")
	flag.Parse()

	if host == "" {
		fmt.Println("host is a mandatory flag")
		os.Exit(1)
	}

	if err := introspect(); err != nil {
		fmt.Printf("Error while introspecting schema, %v\n", err)
		os.Exit(1)
	}
}

func introspect() error {
	graphqlURL, err := getGraphqlAPIBaseURL()
	if err != nil {
		return err
	}

	schema, err := gqlfetch.BuildClientSchemaWithHeaders(
		context.Background(),
		graphqlURL,
		getAuthenticationHeaders(),
		true,
	)
	if err != nil {
		return err
	}

	absOutputPath, err := filepath.Abs(outputPath)
	if err != nil {
		return err
	}

	// public file mode
	var publicFileMode fs.FileMode = 0o644
	if err = os.WriteFile(absOutputPath, []byte(schema), publicFileMode); err != nil {
		return err
	}

	return nil
}

func getGraphqlAPIBaseURL() (string, error) {
	return url.JoinPath(host, apiPath)
}

func getAuthenticationHeaders() http.Header {
	headers := http.Header{}
	if token != "" {
		headers.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	}

	return headers
}
