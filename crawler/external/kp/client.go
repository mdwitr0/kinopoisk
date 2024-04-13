package kp

import (
	"github.com/idoubi/goz"
)

const host = "https://graphql.kinopoisk.ru/graphql/"

type Client struct {
	rest *goz.Request
}

func NewClient() *Client {
	g := goz.NewClient(
		goz.Options{
			BaseURI: host,
			Headers: map[string]interface{}{
				"accept":               "*/*",
				"accept-language":      "ru,en;q=0.9",
				"cache-control":        "no-cache",
				"content-type":         "application/json",
				"cookie":               "yandex_login=evgenmist; yandexuid=5442806361710020766; L=BXJAQ1NZYXAGcWVVCEphXQFZV0Z3YnMBPSESUwMfOjsi.1710026786.15643.358124.58ac81c51940fabface619f4930e7df0; gdpr=0; _ym_uid=1710063786310356406; PHPSESSID=715044b4335aec93729d57e4b63cfdb4; yandex_gid=10522; uid=2388885; yuidss=5442806361710020766; location=1; coockoos=4; mda_exp_enabled=1; my_perpages=%5B%5D; yashr=2385182201710439488; i=6fOySPehsYjbuiQ8yJHlCrStRlT/QCzMpYUXSqqme3jzMJBrbIPqUJeFCpj1VwwTjC1JdmSvT3y2wosuamn3K6v3Vrg=; mobile=no; crookie=pi25k0j8oMZgdtdbk4fINkFkHTzhMCI5KTgQxc0zgWB9Sfw4VLowsIn0YX6XaYwlgAxtKGkc/8mGWgrSjYRQpjjIHeg=; cmtchd=MTcxMjY2NzIwNTMyNA==; spravka=dD0xNzEyODM0ODg2O2k9NTcuMTI5LjIzLjIzNDtEPTQzRTI5NTBFNjU2QTE0Rjk1QTg5M0Q5RTQ5QTlEOUVEMkU5MDQ3NkQzNTQxMjQxNEFBNDg1QjMwOTgwMzY1QkUzQkMyRUYxNTIxMUU1M0FFRjk0NTkxM0M5ODY0QzFCQjYxNzE2ODI0MDRGNEExMDk4NEJCNTNBMTU2NEY3QzM4NTFBQkU2ODdFOTJFNUE3MkEyODYzRDQ3MzI0NjUzOTc2OUMyNjc3Qjt1PTE3MTI4MzQ4ODYzNjM5MDY3MDE7aD0zNGNiN2JlMjkxOWY0ZGJmMThjZmYxNDhhNjI2YTkxZA==; tc=6121; disable_server_sso_redirect=1; _yasc=eHLo9OglxpnJrDRvKGDRD/2p5avE5tp9rFHw7CSfW4uu/f5lClWwfp/Bl+lT6lS5qbo1xaNd; ya_sess_id=3:1712942555.5.0.1710026786164:4UiOWw:18.1.2:1|161327698.0.2.3:1710026786|30:10224006.56403.dq1Vx3BohExZY6RknESxG09oecE; sessar=1.1188.CiDSTa0iuBE3jl3TqLL-AUj-V-BEdktomsVvawucWYHNKg.MUXn68jgiwJENHZUMsYP355rDhWmiwi9ZZxSk1f-cTg; ys=udn.cDpldmdlbm1pc3Q%3D#c_chck.737377711; mda2_beacon=1712942555578; sso_status=sso.passport.yandex.ru:synchronized; no-re-reg-required=1; _ym_isad=1; yp=1713028960.yu.5442806361710020766; ymex=1715534560.oyu.5442806361710020766; _ym_d=1712942602; _yasc=lyWImn2b6l9wlq4geOhxhK/c6ovOAuUS7PuDsG6Zmbv5jTS3wJLRBx2BE8qDB2XEwl0JDTUn; i=5C6DkdihBpTWny75rlrgUGXaqwX1kYc4QTO+R/kMtNx9Q3SMP49a1KpeWsYrmskLZtRGMzmh0wxfyJ/Cde+0PpyOwuQ=; yandexuid=3378069401707313310",
				"dnt":                  "1",
				"origin":               "https://www.kinopoisk.ru",
				"pragma":               "no-cache",
				"referer":              "https://www.kinopoisk.ru/",
				"sec-ch-ua-mobile":     "?0",
				"sec-fetch-dest":       "empty",
				"sec-fetch-mode":       "cors",
				"sec-fetch-site":       "same-site",
				"service-id":           "25",
				"user-agent":           "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
				"x-preferred-language": "ru",
			},
		},
	)

	return &Client{
		rest: g,
	}
}
