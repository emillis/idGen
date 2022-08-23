package _test

import (
	"github.com/emillis/idGen"
	"testing"
)

type customEncoder struct{}

func (ce customEncoder) Encode(s string) string {
	return s
}

var val = "testing_value"
var salt = "156vs/d1ce5_35ctRF&^£$fDFS5RV4531dfv1r4w5e6f1d1\\5sdD5we(fE\",fe3s5EF"

var tests = []struct {
	r    idGen.Requirements
	want string
}{
	{
		idGen.Requirements{64, salt, false, nil, nil, nil},
		"NwiKlYB4xFOeEc5Fy-0I4NGCg0nvTWWRc-aAftaZ9KUTjVN5n1RAcYr23CEFm-Mr",
	},
	{
		idGen.Requirements{64, salt, false, idGen.LowerOnly, nil, nil},
		"nwiklyb4xfoeec5fy-0i4ngcg0nvtwwrc-aaftaz9kutjvn5n1racyr23cefm-mr",
	},
	{
		idGen.Requirements{64, salt, false, idGen.UpperOnly, nil, nil},
		"NWIKLYB4XFOEEC5FY-0I4NGCG0NVTWWRC-AAFTAZ9KUTJVN5N1RACYR23CEFM-MR",
	},

	{
		idGen.Requirements{64, salt, false, nil, idGen.AlphaOnly, nil},
		"NwiKlYBxFOeEcFyINGCgnvTWWRcaAftaZKUTjVNnRAcYrCEFmMrwxOcDBJryYZKT",
	},
	{
		idGen.Requirements{64, salt, false, idGen.LowerOnly, idGen.AlphaOnly, nil},
		"nwiklybxfoeecfyingcgnvtwwrcaaftazkutjvnnracyrcefmmrwxocdbjryyzkt",
	},
	{
		idGen.Requirements{64, salt, false, idGen.UpperOnly, idGen.AlphaOnly, nil},
		"NWIKLYBXFOEECFYINGCGNVTWWRCAAFTAZKUTJVNNRACYRCEFMMRWXOCDBJRYYZKT",
	},

	{
		idGen.Requirements{64, salt, false, nil, idGen.NumericOnly, nil},
		"4504095123896559496345396855855513224654178015435189919615859574",
	},
	{
		idGen.Requirements{64, salt, false, idGen.LowerOnly, idGen.NumericOnly, nil},
		"4504095123896559496345396855855513224654178015435189919615859574",
	},
	{
		idGen.Requirements{64, salt, false, idGen.UpperOnly, idGen.NumericOnly, nil},
		"4504095123896559496345396855855513224654178015435189919615859574",
	},

	{
		idGen.Requirements{64, salt, false, nil, idGen.AlphanumericOnly, nil},
		"NwiKlYB4xFOeEc5Fy0I4NGCg0nvTWWRcaAftaZ9KUTjVN5n1RAcYr23CEFmMrwxO",
	},
	{
		idGen.Requirements{64, salt, false, idGen.LowerOnly, idGen.AlphanumericOnly, nil},
		"nwiklyb4xfoeec5fy0i4ngcg0nvtwwrcaaftaz9kutjvn5n1racyr23cefmmrwxo",
	},
	{
		idGen.Requirements{64, salt, false, idGen.UpperOnly, idGen.AlphanumericOnly, nil},
		"NWIKLYB4XFOEEC5FY0I4NGCG0NVTWWRCAAFTAZ9KUTJVN5N1RACYR23CEFMMRWXO",
	},
}

func TestStatic(t *testing.T) {
	for _, test := range tests {
		got := idGen.Static(val, &test.r)

		if got != test.want {
			t.Errorf("expected %s, got %s\n", test.want, got)
		}
	}
}

func TestRandom(t *testing.T) {
	desiredIDLength := 100
	output := make(map[string]string)

	for i := 0; i < 20; i++ {
		got := idGen.Random(&idGen.Requirements{desiredIDLength, "", false, nil, nil, nil})
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

func TestGenerator_Static(t *testing.T) {
	for _, test := range tests {
		jenny := idGen.NewGenerator(&test.r)

		got := jenny.Static(val)

		if got != test.want {
			t.Errorf("expected %s, got %s\n", test.want, got)
		}
	}
}

func TestGenerator_Random(t *testing.T) {
	desiredIDLength := 100
	output := make(map[string]string)

	jenny := idGen.NewGenerator(&idGen.Requirements{desiredIDLength, "", false, nil, nil, nil})

	for i := 0; i < 20; i++ {
		got := jenny.Random()
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

func TestRequirements_Encoder(t *testing.T) {
	r := idGen.Requirements{
		Length:   len(val),
		OmitSalt: true,
		Encoder:  customEncoder{},
	}

	v := idGen.Static(val, &r)

	if v != val {
		t.Errorf("Expected Statically encoded string to be the same as the one supplied, got encoded \"%s\", supplied \"%s\"", v, val)
	}
}
