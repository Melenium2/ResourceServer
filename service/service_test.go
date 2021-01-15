package service_test

import (
	"ResourceServer/service"
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"
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

var largePictures = []string {
	"https://is2-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/bb/03/09/bb030956-0a17-0ae7-7681-b391d29f85bf/d19e43ae-70ac-45e3-af44-ca4a1865df4d_01_iPhoneXsMax_iPhoneXr_6_5_inch_1242x2688.jpg/1242x2688w.png",
	"https://is2-ssl.mzstatic.com/image/thumb/PurpleSource114/v4/8b/6f/65/8b6f650e-d4ef-7a0a-b61a-a51a07249bd4/f2b3fa7f-7ad7-47f8-905a-2a3f79733c8e_02_iPhoneXsMax_iPhoneXr_6_5_inch_1242x2688.jpg/1242x2688w.png",
	"https://is1-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/dd/5f/63/dd5f63db-754b-be2f-17f8-a5eb1041caf2/94d72773-2c59-4006-9eed-d954b92098b2_03_iPhoneXsMax_iPhoneXr_6_5_inch_1242x2688.jpg/1242x2688w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/90/bc/8e/90bc8ed9-5d0d-27b1-2add-93091d81e037/91e12a90-1aec-4015-9005-461ed0730a08_04_iPhoneXsMax_iPhoneXr_6_5_inch_1242x2688.jpg/1242x2688w.png",
	"https://is4-ssl.mzstatic.com/image/thumb/PurpleSource114/v4/67/f2/d1/67f2d128-e56a-747f-036a-198acbdda4dd/926f20ff-8082-4426-a0b2-bbc472d4a06f_05_iPhoneXsMax_iPhoneXr_6_5_inch_1242x2688.jpg/1242x2688w.png",
	"https://is4-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/05/9c/84/059c84a9-69bc-1dd1-33ba-fbd138b54a30/4d1a7be7-d0f3-4e19-bc2c-03d99582aaeb_01_iPhone6sPlus_iPhone7Plus_iPhone8Plus_5_5_inch_1242x2208.jpg/1242x2208w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/22/b7/6d/22b76da2-977d-c28f-5c55-c1d8f9da5cf6/1e167881-4fed-4891-92b7-481901d1d680_02_iPhone6sPlus_iPhone7Plus_iPhone8Plus_5_5_inch_1242x2208.jpg/1242x2208w.png",
	"https://is4-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/1b/31/3d/1b313d58-5d79-3f46-026e-2d271e16a8a4/ece9ab13-bb63-44d7-98a8-643e4f0fbd76_03_iPhone6sPlus_iPhone7Plus_iPhone8Plus_5_5_inch_1242x2208.jpg/1242x2208w.png",
	"https://is1-ssl.mzstatic.com/image/thumb/PurpleSource114/v4/fb/35/85/fb3585cc-6148-7e46-b615-d1b7d1091cbf/574d1049-0f3a-44c2-9207-4c94edbc162e_04_iPhone6sPlus_iPhone7Plus_iPhone8Plus_5_5_inch_1242x2208.jpg/1242x2208w.png",
	"https://is4-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/f1/77/1e/f1771eb7-5572-a281-3ffb-3742ee6fc949/a23c57b2-cc32-4a3e-9393-6087a1577fed_05_iPhone6sPlus_iPhone7Plus_iPhone8Plus_5_5_inch_1242x2208.jpg/1242x2208w.png",
	"https://is1-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/ac/d6/c0/acd6c011-a246-b508-6d30-42328d6bf8de/43d6bf54-a0ff-4b01-9c2b-7536c71a226e_01_iPadPro_12_9_inch_2048x2732.jpg/2048x2732w.png",
	"https://is2-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/a8/e9/f9/a8e9f987-4a02-f246-6a69-7eccccdcf9e0/88a2da1d-6581-4f12-827d-9dbe7f890030_02_iPadPro_12_9_inch_2048x2732.jpg/2048x2732w.png",
	"https://is2-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/2a/53/db/2a53dbde-e950-f9e3-362e-55edadefd4a0/1ef74314-49c2-4c82-b67b-d5b9f20b2fd4_03_iPadPro_12_9_inch_2048x2732.jpg/2048x2732w.png",
	"https://is5-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/7b/39/ab/7b39ab6f-9f9a-b69b-4c94-f95db73474b9/9f6c3feb-dfaf-417c-b2c9-1b80c8f35e6c_04_iPadPro_12_9_inch_2048x2732.jpg/2048x2732w.png",
	"https://is2-ssl.mzstatic.com/image/thumb/PurpleSource114/v4/1f/5b/54/1f5b5454-4cf0-d65d-663e-807779f8da40/d50dcdcf-71b3-4357-9bc0-8959af5db830_05_iPadPro_12_9_inch_2048x2732.jpg/2048x2732w.png",
	"https://is1-ssl.mzstatic.com/image/thumb/PurpleSource114/v4/7a/ac/3c/7aac3c83-06a3-c3d1-83a9-5a5c05a4aafd/ed2a1477-f1c0-4625-8de6-150ab153554e_01_ipadPro129_inch_2048x2732.jpg/2048x2732w.png",
	"https://is5-ssl.mzstatic.com/image/thumb/PurpleSource114/v4/9d/5c/b2/9d5cb2d3-ed3a-6cbe-1a68-b7a028e13726/3dd54dd3-cddf-4b87-a772-90d49ca0a8bc_02_ipadPro129_inch_2048x2732.jpg/2048x2732w.png",
	"https://is1-ssl.mzstatic.com/image/thumb/PurpleSource114/v4/ce/65/d6/ce65d630-b24c-8b01-da51-cb2b61212fef/4c98e5e6-c1aa-41a7-b74e-37f123e51d1f_03_ipadPro129_inch_2048x2732.jpg/2048x2732w.png",
	"https://is2-ssl.mzstatic.com/image/thumb/PurpleSource124/v4/cf/5d/06/cf5d0606-bb52-8bef-ed28-a15c7d31281e/60fe8c12-5bff-4eff-8f4b-c6cfa27c953c_04_ipadPro129_inch_2048x2732.jpg/2048x2732w.png",
	"https://is5-ssl.mzstatic.com/image/thumb/PurpleSource114/v4/f0/89/bf/f089bf2e-1a27-33cf-d554-c98e3ab33d77/2854ca37-57f6-4ad1-968a-3d94233969ff_05_ipadPro129_inch_2048x2732.jpg/2048x2732w.png",
}

func TestResourceService_Load_ShouldLoadSingleResourceToLocalFolder(t *testing.T) {
	filepath := path.Join(resourcePath(), "./resource_temp")
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
	filepath := path.Join(resourcePath(), "./resource_temp")
	var _ = os.MkdirAll(filepath, 0777)
	s := service.New(filepath, 5)
	ctx := context.Background()

	var tt = []struct {
		name string
		urls []string
		e    bool
	} {
		{
			name: "load batch, pictures concurrency",
			urls: pictures,
			e:    false,
		},
		{
			name: "load batch, alternative pictures concurrency",
			urls: picturesAlternative,
			e:    false,
		},
		{
			name: "load batch, large files concurrency",
			urls: largePictures,
			e: false,
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

func TestResourceService_LoadBatch_MakeNRequestAsyncAndWaitAMinute(t *testing.T) {
	filepath := path.Join(resourcePath(), "./resource_temp")
	var _ = os.MkdirAll(filepath, 0777)
	s := service.New(filepath, 5)
	ctx := context.Background()
	n := 5

	for i := 0; i < n; i++ {
		go func(index int) {
			r, err := s.LoadBatch(ctx, largePictures)
			assert.NoError(t, err)
			assert.NotNil(t, r)
			assert.Equal(t, len(largePictures), len(r))
			t.Log("Done ", index)
		}(i)
	}

	time.Sleep(time.Minute * 2)
}