package service_test

import (
	"ResourceServer/service"
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"path/filepath"
	"testing"
)

var removePath = func(path string) error {
	return os.RemoveAll(path)
}

var resourcePath = func() string {
	ex, _ := os.Executable()
	return filepath.Dir(ex)
}

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
	"https://is1-ssl.mzstatic.com/image/thumb/Purple124/v4/91/95/19/919519d3-52d3-df40-5fb5-3fb3aa3d6f18/AppIcon-0-0-1x_U007emarketing-0-0-0-7-0-0-sRGB-0-0-0-GLES2_U002c0-512MB-85-220-0-0.png/1024x1024w.png",
	"https://is1-ssl.mzstatic.com/image/thumb/Purple124/v4/47/67/e0/4767e03c-a8cd-c5f3-a4d0-899c83740bad/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/1f/fa/10/1ffa10ba-01db-015e-eb30-678770d1a04f/pr_source.png/2048x2732w.png",
	"https://is4-ssl.mzstatic.com/image/thumb/Purple114/v4/04/e9/d4/04e9d4f5-7cf1-b669-5973-28a4eb770a0f/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/91/3e/7b/913e7be3-4616-6de5-957a-b5f11b2b68ba/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple114/v4/33/f9/39/33f939dc-bba3-7894-5213-236ef967678f/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/59/06/d0/5906d02e-b301-c9d3-edfd-5344d4c26da6/pr_source.png/2048x2732w.png",
}

func TestResourceService_Load_ShouldLoadSingleResourceToLocalFolder(t *testing.T) {
	filepath := path.Join(resourcePath(), "resource_temp")
	var _ = os.MkdirAll(filepath, 0777)
	s := service.New(filepath)
	ctx := context.Background()

	var tt = []struct {
		name string
		urls []string
		e    bool
	} {
		{
			name: "load pictures consistently",
			urls: pictures,
			e:    false,
		},
		{
			name: "load alternative pictures consistently ",
			urls: picturesAlternative,
			e:    false,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			for _, res := range test.urls {
				r, err := s.Load(ctx, res)
				if !test.e {
					assert.NoError(t, err)
					assert.NotEmpty(t, r)
				} else {
					assert.Error(t, err)
					assert.Empty(t, r)
				}
			}
		})
	}

	assert.NoError(t, removePath(filepath))
}

func TestResourceService_LoadBatch_ShouldLoadSingleReportWithConcurrency(t *testing.T) {
	filepath := path.Join(resourcePath(), "resource_temp")
	var _ = os.MkdirAll(filepath, 0777)
	s := service.New(filepath, 5)
	ctx := context.Background()

	var tt = []struct {
		name string
		urls []string
		e    bool
	} {
		{
			name: "load batch pictures consistently",
			urls: pictures,
			e:    false,
		},
		{
			name: "load batch alternative pictures consistently",
			urls: picturesAlternative,
			e:    false,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			r, err := s.LoadBatch(ctx, test.urls)
			if !test.e {
				assert.Nil(t, err)
				assert.NotEmpty(t, r)
				assert.Equal(t, len(r), len(test.urls))
			} else {
				assert.NotNil(t, err)
				assert.Empty(t, r)
			}
		})
	}

	assert.NoError(t, removePath(filepath))
}
