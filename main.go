package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	// Get the environment variables
	repo := envOrDie("GITHUB_REPO")
	token := envOrDie("GITHUB_TOKEN")
	nums := envOrDie("GITHUB_PR_NUMBER")
	num, err := strconv.Atoi(nums)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: GITHUB_PR_NUMBER must contain a number. It currently has the value %q", nums)
		os.Exit(1)
	}

	// Need the org and project separately for the client
	parts := strings.Split(repo, "/")
	if len(parts) != 2 {
		fmt.Fprintf(os.Stderr, "Error: The repo does not follow the format [org|user]/[project] for the variable %q", repo)
		os.Exit(1)
	}

	t := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	c := context.Background()
	tc := oauth2.NewClient(c, t)
	client := github.NewClient(tc)

	opts := &github.ListOptions{
		Page: 0,

		// 100 is the API max
		PerPage: 100,
	}
	commits, _, err := client.PullRequests.ListCommits(c, parts[0], parts[1], num, opts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Unable to fetch commits")
		os.Exit(1)
	}

	out, err := json.Marshal(commits)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Unable to generate JSON")
		os.Exit(1)
	}

	// Print the value to stdout
	fmt.Println(string(out))
}

// If an manadatory environment variable is not present we are going to die
// early with an error message.
func envOrDie(name string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}

	fmt.Fprintf(os.Stderr, "Error: Environment variable %q has no value", name)
	os.Exit(1)

	// A return that should never be called unless something is very very wrong
	return ""
}
