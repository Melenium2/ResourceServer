package server

import (
	"ResourceServer/service"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

var pictures = []string{
	"https://play-lh.googleusercontent.com/6hiR7BeQTcyhQoOoNeasulh0WM8tKFLF09naV7t4LAGFjMs20PxKU2RONu1cZIVxY7yy",
	"https://play-lh.googleusercontent.com/2nDcr8fp7m-kLiSN31yUGruOHAZGMHO9u5h-V5EdXqsFwe6FEVU_saDX9ji2d6TNvw",
	"https://play-lh.googleusercontent.com/kAxtdEKXBf5bASmakmC_8qTnUue0vqUMkvP_fofgyWGnfaiQRnEQmLS8sTkvtoVDyww",
	"https://play-lh.googleusercontent.com/0aWZJEYQ5EsvbdaUQO4uSAMu8NODT0XwufU90i1N85BEsZy25CH9rXe6QbnDZkCRQLc",
	"https://play-lh.googleusercontent.com/xz_RxPMYZi9_AugnrHuNnmdrRrZDvALPK5CmauMlgZt9K5bBRfHtFjh2tFSiR3deTg",
	"https://play-lh.googleusercontent.com/5BricghADLXngbOmZArsPqCdzgmtDmq5ES1pPNjQ-ns1mvzG5XTdqPfUEzMHXFKgiA",
	"https://play-lh.googleusercontent.com/dAUvmf-D-R8cCkd5RKOb0_2nyOD9mqHi4VzKyv37tf_Sx1tzYI6UV--jMUL_IFL0Uw",
	"https://play-lh.googleusercontent.com/sjP4ezqr43irX-44ROYGzRuoTli2PTiyB1W7wdS1sp8-PuWlG1Y_K2Faya0h18azVA",
	"https://play-lh.googleusercontent.com/PmxedM7qGD0kHsUiFDXGmS61eKzeCB6dHj57y12SMfQ8B8tB5k8FH4H3jI_OJFcyrQ",
	"https://play-lh.googleusercontent.com/cA8FFkPoLnZCDVdMuqi0D5IzPIq4Pt9TNw-lhCbr4PMnG0f8mbfjNyLMuegCKnrcez3F",
	"https://play-lh.googleusercontent.com/jwWo5D5eTV0UCwOBToLl-J5W9FJmsvAKjaDYhkHhH0D5EbZkm54wkIPa4jL2lL-3Yx8",
	"https://play-lh.googleusercontent.com/5wRkd5tVwZKvbI3zqKlsN1-yOoKSRwpEYl3x8_lgCUH8KNgVvrIYcks2QsIl4yOZFe4",
	"https://play-lh.googleusercontent.com/njuA2vwuWjLqBTeHsG-vtu-Xd8t2qgnjEUHPT6u4FDJIneb30eVSwc5ZXhLuJPqvnw",
	"https://play-lh.googleusercontent.com/5Gu108fDTyz1RGAs09ggwzZ6GgAx8C-AXVVGCcznMQCg98Dr_5pAcK3O3eSRtsdbog4",
	"https://play-lh.googleusercontent.com/kKQ8yRyR4YGnKHoJra1i5B0G4WHAuSIpqvV97hP3YQc-Qmn3pgocYiCH_UPdfTkpwok",
	"https://play-lh.googleusercontent.com/bnbCzTJ4gGugbFqFjH1e7bLxe7I7Ze3e_6LNpPYY2nuwmy64WEpQloJEjcQB9AHcDsE",
	"https://play-lh.googleusercontent.com/ra0x8DKrViL6lhfcVMBvSs7C4WOS7H4eR1KVLC7HuOdLkv5C6fdOUp5M5wvUlocpXOo",
	"https://play-lh.googleusercontent.com/-lidNEWNyB5YDUcoeHXrAFyvoZKMXPlwlMhCka_-oT-2qqzfCMm_gcdkCTCN5Z1Vbw",
}

var picturesAlternative = []string{
	"https://is1-ssl.mzstatic.com/image/thumb/Purple124/v4/47/67/e0/4767e03c-a8cd-c5f3-a4d0-899c83740bad/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/1f/fa/10/1ffa10ba-01db-015e-eb30-678770d1a04f/pr_source.png/2048x2732w.png",
	"https://is4-ssl.mzstatic.com/image/thumb/Purple114/v4/04/e9/d4/04e9d4f5-7cf1-b669-5973-28a4eb770a0f/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/91/3e/7b/913e7be3-4616-6de5-957a-b5f11b2b68ba/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple114/v4/33/f9/39/33f939dc-bba3-7894-5213-236ef967678f/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/59/06/d0/5906d02e-b301-c9d3-edfd-5344d4c26da6/pr_source.png/2048x2732w.png",
}

func TestServer(t *testing.T) {

	wrongItemPictures := make([]string, len(pictures))
	for i, v := range pictures {
		wrongItemPictures[i] = v
	}
	wrongItemPictures[3] = "https://play-lh.googleusercontent.com/5Gu108fDTyz1RGAs09ggwzZ6GgAx8C-AXVVGCcznMQCg98Dr_5pAcK3O3eS123123123"

	var tt = []struct {
		name          string
		route         string
		method        string
		body          []string
		query         string
		expectedCode  int
		expectedBody  string
	}{
		{
			name:          "load route error because empty query",
			route:         "/load",
			method:        "GET",
			body:          nil,
			query:         "",
			expectedCode:  404,
			expectedBody:  "incorrect resource link",
		},
		{
			name:          "load route incorrect link with img",
			route:         "/load",
			method:        "GET",
			body:          nil,
			query:         "https://is4-ssl.mzstatic.com/image/thumb/Purple114/v4/9d/9f/e1/9d9fe144-421c-bdc5-989979b53f1/pr_source.png/1242x2208w.png",
			expectedCode:  404,
			expectedBody:  "url https://is4-ssl.mzstatic.com/image/thumb/Purple114/v4/9d/9f/e1/9d9fe144-421c-bdc5-989979b53f1/pr_source.png/1242x2208w.png",
		},
		{
			name:          "load route success",
			route:         "/load",
			method:        "GET",
			body:          nil,
			query:         "https://is4-ssl.mzstatic.com/image/thumb/Purple114/v4/9d/9f/e1/9d9fe144-421c-bdc5-964c-1689979b53f1/pr_source.png/1242x2208w.png",
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			name:          "load batch error because empty body",
			route:         "/load/batch",
			method:        "POST",
			body:          nil,
			query:         "",
			expectedCode:  404,
			expectedBody:  "",
		},
		{
			name:          "load batch error because wrong links",
			route:         "/load/batch",
			method:        "POST",
			body:          wrongItemPictures,
			query:         "",
			expectedCode:  200,
			expectedBody:  "none",
		},
		{
			name:          "load batch success",
			route:         "/load/batch",
			method:        "POST",
			body:          pictures,
			query:         "",
			expectedCode:  200,
			expectedBody:  "",
		},
	}

	app := fiber.New()
	c := NewConfig()
	c.ServeFolder = "./resources"
	worker := service.New(c.ServeFolder, 5)
	server := New(app, worker, c)
	_ = server.InitRoutes()

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			var body io.Reader
			if test.method == "POST" {
				b, _ := json.Marshal(test.body)
				body = bytes.NewReader(b)
			}
			req, _ := http.NewRequest(test.method, test.route, body)
			if test.method == "GET" {
				q := req.URL.Query()
				q.Add("link", test.query)
				req.URL.RawQuery = q.Encode()
			}

			res, err := server.app.Test(req, -1)

			assert.Equal(t, test.expectedCode, res.StatusCode)
			bodyResp, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err)
			t.Log(string(bodyResp))
			assert.True(t, strings.Contains(string(bodyResp), test.expectedBody))
		})
	}

	_ = os.RemoveAll(c.ServeFolder)
}

// Test works. Image returned. But i can't remove directory  with images
// because has error 'resource busy'. I don't know what to do with this problem yet.
//
//func TestServer_ServingFiles_ShouldGetImageByName(t *testing.T) {
//	app := fiber.New()
//	c := NewConfig()
//	c.ServeFolder = "./resources"
//	worker := service.New(c.ServeFolder, 5)
//	server := New(app, worker, c)
//	_ = server.InitRoutes()
//
//	var servingFilename string
//	{
//		req, _ := http.NewRequest("GET", "/load", nil)
//		{
//			q := req.URL.Query()
//			q.Add("link", pictures[0])
//			req.URL.RawQuery = q.Encode()
//		}
//		res, err := server.app.Test(req, -1)
//		assert.NoError(t, err)
//
//		b, _ := ioutil.ReadAll(res.Body)
//		_ = res.Body.Close()
//
//		servingFilename = string(b)
//		assert.NotEmpty(t, servingFilename)
//	}
//	{
//		servingFilename = strings.ReplaceAll(servingFilename, "\"", "")
//		req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", servingFilename), nil)
//		res, err := server.app.Test(req, -1)
//		assert.NoError(t, err)
//
//		b, _ := ioutil.ReadAll(res.Body)
//		_ = res.Body.Close()
//		assert.NotNil(t, b)
//		assert.True(t, strings.Contains(string(b), "PNG"))
//	}
//
//	assert.NoError(t, os.RemoveAll(c.ServeFolder))
//}
