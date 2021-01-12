package resource_test

import (
	"ResourceServer/resource"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

var testResourceFolder = func() string {
	ex, _ := os.Executable()
	return filepath.Dir(ex)
}

func helper_removeTestFile(filepath string) error {
	return os.RemoveAll(filepath)
}

func TestSave_ShouldCreateFileAndSaveString(t *testing.T) {
	folder := testResourceFolder()
	filepath := path.Join(folder, "123.txt")
	readcloser := ioutil.NopCloser(strings.NewReader("2123123123"))
	assert.NoError(t, resource.Save(filepath, readcloser))

	assert.NoError(t, helper_removeTestFile(filepath))
}

func TestSave_ShouldPanicIfReaderIsempty(t *testing.T) {
	folder := testResourceFolder()
	filepath := path.Join(folder, "123.txt")

	assert.Panics(t, func() {
		resource.Save(filepath, nil)
	})

	assert.NoError(t, helper_removeTestFile(filepath))
}

func TestLoad_ShouldSaveResourceFromInternet(t *testing.T) {
	var tt = []struct {
		name string
		url  string
		e    bool
	}{
		{
			name: "Test google play link",
			url:  "https://play-lh.googleusercontent.com/IksIM315GVSASC9YVZrv470XttljyPnS5z9eTxgCQs_OgtRKASyHDIYWha6N2F-qi9Q",
			e:    false,
		},
		{
			name: "Test appstore link",
			url:  "https://is4-ssl.mzstatic.com/image/thumb/Purple114/v4/9d/9f/e1/9d9fe144-421c-bdc5-964c-1689979b53f1/pr_source.png/1242x2208w.png",
			e:    false,
		},
		{
			name: "Test empty link",
			url:  "",
			e:    true,
		},
		{
			name: "Test incorrect link",
			url:  "https://is2-ssl.mzstatic.com/image/thumb/Features4/v4/e7/f0/71/e7f07160-bb0f-496b-fc11-e1c29e36162a/source/3200x",
			e:    true,
		},
	}
	folder := testResourceFolder()
	filespath := path.Join(folder, "resource_test")
	var _ = os.MkdirAll(filespath, 0777)

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			_, err := resource.Load(filespath, test.url)
			if test.e {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
	assert.NoError(t, helper_removeTestFile(filespath))
}
