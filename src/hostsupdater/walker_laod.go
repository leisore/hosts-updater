package hostsupdater

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

const _load_site = "http://laod.cn/hosts"
const _load_update_regex = "http://laod.cn/hosts/[0-9]*-google-hosts.html"
const _load_upload_regex = "http://laod.cn/wp-content/uploads/[0-9]*/[0-9]*/[0-9]*-hosts.txt"

type laod_walker struct {
}

func init() {
	RegisterWalker(laod_walker{})
}

func (lw laod_walker) Name() string {
	return "laod_walker"
}

func (lw laod_walker) Version() string {
	return "0.0.1"
}

func (lw laod_walker) License() string {
	return "Apache License 2"
}

func (lw laod_walker) Desc() string {
	return fmt.Sprintf("walking hosts from website [%s]", _load_site)
}

func (lw laod_walker) Authors() []Author {
	return []Author{Author{"leisore", "leisore@foxmail.com"}}
}

func (lw laod_walker) WalkedHosts() (io.Reader, error) {

	fmt.Printf("\t->:%s\n", _load_site)
	hostsUrl, err := findRegexUrl(_load_site, _load_update_regex)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\t  ->:%s\n", hostsUrl)
	uploadUrl, err := findRegexUrl(hostsUrl, _load_upload_regex)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\t    ->:%s\n", uploadUrl)
	if resp, err := http.Get(uploadUrl); err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

func findRegexUrl(srcUrl string, regx string) (string, error) {
	resp, err := http.Get(srcUrl)
	if err != nil {
		return "", err
	}

	defer func() { resp.Body.Close() }()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if regx, err := regexp.Compile(regx); err != nil {
		return "", err
	} else {

		return regx.FindString(string(bytes)), nil
	}
}
