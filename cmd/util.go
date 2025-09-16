package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

type ReleaseInfo struct {
	Tag     string `json:"tag_name"`
	Created string `json:"created_at"`
}

type Release struct {
	Tag   string
	Major int
	Minor int
	Patch int
}

func FetchReleases() ([]ReleaseInfo, error) {
	url := "https://api.github.com/repos/boostorg/boost/releases?page=1"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error making GET request to github server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Returned status code by api.github.com/releases is not OK")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error when reading body of a response of api.github.com/releases: %w", err)
	}

	var releases []ReleaseInfo
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return nil, fmt.Errorf("Error when unmarshaling api.github.com/releases json, %w", err)
	}

	return releases, nil
}

func parseTag(tag string) (int, int, int, error) {
	tagRe := regexp.MustCompile(`boost-(\d)\.(\d)\.(\d).*`)
	match := tagRe.FindStringSubmatch(tag)
	if match == nil {
		return 0, 0, 0, fmt.Errorf("Could not find a matching version in %v", tag)
	}
	if len(match) < 2 {
		return 0, 0, 0, fmt.Errorf("Not enough information about version found in %v", tag)
	}

	major, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Error parsing %v in %v: not a number", match[1], tag)
	}
	minor := 0
	patch := 0

	if len(match) >= 3 {
		minor, err = strconv.Atoi(match[2])
		if err != nil {
			return 0, 0, 0, fmt.Errorf("Error parsing %v in %v: not a number", match[2], tag)
		}
	}

	if len(match) >= 4 {
		patch, err = strconv.Atoi(match[3])
		if err != nil {
			return 0, 0, 0, fmt.Errorf("Error parsing %v in %v: not a number", match[3], tag)
		}
	}

	return major, minor, patch, nil
}

func FetchLatestRelease() (string, error) {
	releases, err := FetchReleases()
	if err != nil {
		return "", fmt.Errorf("Error when fetching Latest Release: %w", err)
	}

	var parsedReleases []Release
	for _, release := range releases {
		major, minor, patch, err := parseTag(release.Tag)
		if err != nil {
			return "", fmt.Errorf("Error when parsing tag %v: %w", release.Tag, err)
		}
		newRelease := Release{Tag: release.Tag, Major: major, Minor: minor, Patch: patch}
		parsedReleases = append(parsedReleases, newRelease)
	}

	return "", nil
}
