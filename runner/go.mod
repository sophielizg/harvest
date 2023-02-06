module github.com/sophielizg/harvest/runner

go 1.19

require (
	github.com/gocolly/colly v1.2.0
	github.com/sophielizg/harvest/common/config v0.0.0-20230202071100-3e66946710af
	github.com/sophielizg/harvest/common/harvest v0.0.0-20230202071100-3e66946710af
	github.com/sophielizg/harvest/common/mysql v0.0.0-20230202071100-3e66946710af
)

require (
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/antchfx/htmlquery v1.2.6 // indirect
	github.com/antchfx/jsonquery v1.3.2 // indirect
	github.com/antchfx/xmlquery v1.3.14 // indirect
	github.com/antchfx/xpath v1.2.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/saintfish/chardet v0.0.0-20230101081208-5e3ef4b5456d // indirect
	github.com/sophielizg/harvest/common/utils v0.0.0-20230202071100-3e66946710af // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/temoto/robotstxt v1.1.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/sophielizg/harvest/common/config => ../common/config

replace github.com/sophielizg/harvest/common/harvest => ../common/harvest

replace github.com/sophielizg/harvest/common/mysql => ../common/mysql

replace github.com/sophielizg/harvest/common/utils => ../common/utils
