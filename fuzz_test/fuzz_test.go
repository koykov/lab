//go:build go1.18
// +build go1.18

package fuzz_test

const (
	envTypeUnknown = 0
	envTypeWeb     = 1
)

type kdmReq struct {
	envType int
}

func FuzzGetData(f *testing.F) {
	f.Add("v=default&page=https%253A%252F%252Fucoz.ru%252F&domain=ditky.info&blockID=322502&width=333&height=330&gdpr=0&gdprConsent=&limit=1&format=json&sspUid=135a27f6-ed24-4faa-9c6f-099ec375c9ea")
	f.Fuzz(func(t *testing.T, queryRaw string) {
		req, err := kdmHandler(queryRaw)
		if err != nil {
			f.Error(err)
		}
		if req == nil {
			f.Skip()
		}
	})
}

func kdmHandler(rawQuery string) (*kdmReq, error) {
	req := &kdmReq{
		envType: envTypeWeb,
	}

	query, err := url.ParseQuery(raw)
	if err != nil {
		return nil, err
	}

	var page string
	if page = query.Get("page"); len(page) == 0 {
		page = query.Get("domain")
	}
	if len(page) == 0 {
		req.envType = envTypeUnknown
	}

	return req, err
}
