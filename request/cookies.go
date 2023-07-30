package request

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"

	"github.com/sophielizg/harvest/cache"
	"golang.org/x/net/publicsuffix"
)

type cookiesForKey = map[url.URL][]*http.Cookie

type remoteCookieJar struct {
	cacheId  string
	cache    cache.Cache
	cacheTtl int
	jar      *cookiejar.Jar
	cookies  cookiesForKey
	sync.RWMutex
}

func NewRemoteCookieJar(cookiesCache cache.Cache, cacheId string, cacheTtl int) (http.CookieJar, error) {
	cookiesRaw, err := cookiesCache.Get(cacheId)
	if err != nil {
		return nil, err
	}

	cookies, ok := cookiesRaw.(cookiesForKey)
	if !ok {
		return nil, InvalidCookieItemFormatError
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}

	for u, cookiesForUrl := range cookies {
		jar.SetCookies(&u, cookiesForUrl)
	}

	return &remoteCookieJar{
		cacheId:  cacheId,
		cache:    cookiesCache,
		cacheTtl: cacheTtl,
		jar:      jar,
		cookies:  cookies,
	}, nil
}

func (jar *remoteCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.Lock()
	defer jar.Unlock()

	jar.jar.SetCookies(u, cookies)
	jar.cookies[*u] = jar.jar.Cookies(u)
}

func (jar *remoteCookieJar) Cookies(u *url.URL) []*http.Cookie {
	return jar.jar.Cookies(u)
}

func (jar *remoteCookieJar) Save() error {
	jar.RLock()
	defer jar.RUnlock()

	return jar.cache.Put(jar.cacheId, jar.cookies, jar.cacheTtl)
}

func getClientWithCookies(cookiesCache cache.Cache, cacheId string, cacheTtl int) (*http.Client, error) {
	jar, err := NewRemoteCookieJar(cookiesCache, cacheId, cacheTtl)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Jar: jar,
	}, nil
}

func setCookiesFromClient(client *http.Client) error {
	remoteJar, ok := client.Jar.(*remoteCookieJar)
	if !ok {
		return InvalidCookieJarTypeError
	}

	return remoteJar.Save()
}
