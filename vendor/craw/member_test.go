package craw

import (
	"github.com/golib/assert"
	"testing"
)

func TestCrawPkgMembers(t *testing.T) {
	assertion := assert.New(t)
	assertion.Nil(nil)

	CrawPkgMembers("https://golang.org/pkg/testing/")

}
