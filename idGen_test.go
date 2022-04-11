package idGen

import (
	"testing"
)

func TestStatic(t *testing.T) {
	val := "testing_value"
	salt := "156vs/d1ce5_35ctRF&^£$fDFS5RV4531dfv1r4w5e6f1d1\\5sdD5we(fE\",fe3s5EF"

	tests := []struct {
		cfg  Config
		want string
	}{
		{
			Config{64, salt, "", ""},
			"NwiKlYB4xFOeEc5Fy-0I4NGCg0nvTWWRc-aAftaZ9KUTjVN5n1RAcYr23CEFm-Mr",
		},
		{
			Config{64, salt, LowerOnly, ""},
			"nwiklyb4xfoeec5fy-0i4ngcg0nvtwwrc-aaftaz9kutjvn5n1racyr23cefm-mr",
		},
		{
			Config{64, salt, UpperOnly, ""},
			"NWIKLYB4XFOEEC5FY-0I4NGCG0NVTWWRC-AAFTAZ9KUTJVN5N1RACYR23CEFM-MR",
		},

		{
			Config{64, salt, "", AlphaOnly},
			"NwiKlYBxFOeEcFyINGCgnvTWWRcaAftaZKUTjVNnRAcYrCEFmMrwxOcDBJryYZKT",
		},
		{
			Config{64, salt, LowerOnly, AlphaOnly},
			"nwiklybxfoeecfyingcgnvtwwrcaaftazkutjvnnracyrcefmmrwxocdbjryyzkt",
		},
		{
			Config{64, salt, UpperOnly, AlphaOnly},
			"NWIKLYBXFOEECFYINGCGNVTWWRCAAFTAZKUTJVNNRACYRCEFMMRWXOCDBJRYYZKT",
		},

		{
			Config{64, salt, "", NumericOnly},
			"4504095123896559496345396855855513224654178015435189919615859574",
		},
		{
			Config{64, salt, LowerOnly, NumericOnly},
			"4504095123896559496345396855855513224654178015435189919615859574",
		},
		{
			Config{64, salt, UpperOnly, NumericOnly},
			"4504095123896559496345396855855513224654178015435189919615859574",
		},

		{
			Config{64, salt, "", AlphanumericOnly},
			"NwiKlYB4xFOeEc5Fy0I4NGCg0nvTWWRcaAftaZ9KUTjVN5n1RAcYr23CEFmMrwxO",
		},
		{
			Config{64, salt, LowerOnly, AlphanumericOnly},
			"nwiklyb4xfoeec5fy0i4ngcg0nvtwwrcaaftaz9kutjvn5n1racyr23cefmmrwxo",
		},
		{
			Config{64, salt, UpperOnly, AlphanumericOnly},
			"NWIKLYB4XFOEEC5FY0I4NGCG0NVTWWRCAAFTAZ9KUTJVN5N1RACYR23CEFMMRWXO",
		},
	}

	for _, test := range tests {
		got := Static(val, &test.cfg)

		if got != test.want {
			t.Errorf("expected %s, got %s\n", test.want, got)
		}
	}
}

func TestRandom(t *testing.T) {
	desiredIDLength := 100
	output := make(map[string]string)

	for i := 0; i < 20; i++ {
		got := Random(&Config{desiredIDLength, "", "", ""})
		gotLen := len(got)

		if gotLen != desiredIDLength {
			t.Errorf("length expected %d, received length is %d", desiredIDLength, gotLen)
		}

		_, exist := output[got]
		if exist {
			t.Errorf("Random() generated the same ID twice. ID: %s", got)
		}

		output[got] = ""
	}
}
