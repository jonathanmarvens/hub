package github

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/bmizerany/assert"
)

func TestConfigs_loadFrom(t *testing.T) {
	file, _ := ioutil.TempFile("", "test-gh-config-")
	defer os.RemoveAll(file.Name())

	content := `[[hosts]]
  host = "https://github.com"
  user = "jingweno"
  access_token = "123"
  protocol = "https"`
	ioutil.WriteFile(file.Name(), []byte(content), os.ModePerm)

	cc := &Configs{}
	err := loadFrom(file.Name(), cc)
	assert.Equal(t, nil, err)

	assert.Equal(t, 1, len(cc.Hosts))
	host := cc.Hosts[0]
	assert.Equal(t, "https://github.com", host.Host)
	assert.Equal(t, "jingweno", host.User)
	assert.Equal(t, "123", host.AccessToken)
	assert.Equal(t, "https", host.Protocol)
}

func TestConfigs_saveTo(t *testing.T) {
	file, _ := ioutil.TempFile("", "test-gh-config-")
	defer os.RemoveAll(file.Name())

	host := Host{Host: "https://github.com", User: "jingweno", AccessToken: "123", Protocol: "https"}
	c := Configs{Hosts: []Host{host}}

	err := saveTo(file.Name(), &c)
	assert.Equal(t, nil, err)

	b, _ := ioutil.ReadFile(file.Name())
	content := `[[hosts]]
  host = "https://github.com"
  user = "jingweno"
  access_token = "123"
  protocol = "https"`
	assert.Equal(t, content, string(b))
}
