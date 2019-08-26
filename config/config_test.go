package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	expect := ConfData{
		Token:              "YOUR TOKEN",
		Channel:            "Cxxxxxx",
		GitWorkDir:         "./",
		PullRequestComment: "example",
		PullRequestLabels:  []string{"example"},
		MirageUrl:          "http://example.com",
		DockerImage:        "example:master",
	}

	os.Setenv("SLACK_TOKEN", expect.Token)

	c := LoadConfig("./config.yml")

	if c.Token != expect.Token {
		t.Errorf("Token expect: %+v, got: %+v", expect.Token, c.Token)
	}
	if c.Channel != expect.Channel {
		t.Errorf("Channel expect: %+v, got: %+v", expect.Channel, c.Channel)
	}
	if c.GitWorkDir != expect.GitWorkDir {
		t.Errorf("GitWorkDir expect: %+v, got: %+v", expect.GitWorkDir, c.GitWorkDir)
	}
	if c.PullRequestComment != expect.PullRequestComment {
		t.Errorf("PullRequestComment expect: %+v, got: %+v", expect.PullRequestComment, c.PullRequestComment)
	}
	if len(c.PullRequestLabels) != len(expect.PullRequestLabels) ||
		c.PullRequestLabels[0] != expect.PullRequestLabels[0] {
		t.Errorf("PullRequestComment expect: %+v, got: %+v", expect.PullRequestComment, c.PullRequestComment)
	}
	if c.MirageUrl != expect.MirageUrl {
		t.Errorf("MirageUrl expect: %+v, got: %+v", expect.MirageUrl, c.MirageUrl)
	}
	if c.DockerImage != expect.DockerImage {
		t.Errorf("DockerImage expect: %+v, got: %+v", expect.DockerImage, c.DockerImage)
	}

	t.Logf("%#v\n", c)
}
