module github.com/sophielizg/harvest/api

go 1.19

require (
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/sophielizg/harvest/common/harvest v0.0.0-20230115173816-48993f698041
	github.com/sophielizg/harvest/common/mysql v0.0.0-20230115173816-48993f698041
	github.com/swaggo/http-swagger v1.0.0
	github.com/swaggo/swag v1.8.9
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/sophielizg/harvest/common/local v0.0.0-20230206033005-c4d705cf1396 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/sophielizg/harvest/common/harvest => ../common/harvest

replace github.com/sophielizg/harvest/common/mysql => ../common/mysql

replace github.com/sophielizg/harvest/common/local => ../common/local
