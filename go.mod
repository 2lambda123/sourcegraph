module github.com/sourcegraph/sourcegraph

go 1.16

require (
	cloud.google.com/go v0.82.0
	cloud.google.com/go/pubsub v1.3.1
	cloud.google.com/go/storage v1.10.0
	github.com/Masterminds/semver v1.5.0
	github.com/NYTimes/gziphandler v1.1.1
	github.com/OneOfOne/xxhash v1.2.8 // indirect
	github.com/PuerkitoBio/rehttp v1.1.0
	github.com/RoaringBitmap/roaring v0.5.1
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/avelino/slugify v0.0.0-20180501145920-855f152bd774
	github.com/aws/aws-sdk-go-v2 v1.3.2
	github.com/aws/aws-sdk-go-v2/config v1.1.2
	github.com/aws/aws-sdk-go-v2/credentials v1.1.2
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.0.3
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.1.2
	github.com/aws/aws-sdk-go-v2/service/kms v1.2.2
	github.com/aws/aws-sdk-go-v2/service/s3 v1.2.1
	github.com/aws/smithy-go v1.3.1
	github.com/beevik/etree v1.1.0
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff
	github.com/certifi/gocertifi v0.0.0-20200211180108-c7c1fbc02894 // indirect
	github.com/cockroachdb/errors v1.8.4
	github.com/cockroachdb/redact v1.0.9 // indirect
	github.com/cockroachdb/sentry-go v0.6.1-cockroachdb.2
	github.com/containerd/containerd v1.4.0 // indirect
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/coreos/go-semver v0.3.0
	github.com/crewjam/saml v0.4.14
	github.com/davecgh/go-spew v1.1.1
	github.com/daviddengcn/go-colortext v1.0.0
	github.com/derision-test/glock v0.0.0-20210316032053-f5b74334bb29
	github.com/derision-test/go-mockgen v1.1.2
	github.com/dghubble/gologin v2.2.0+incompatible
	github.com/dgraph-io/ristretto v0.0.3
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dineshappavoo/basex v0.0.0-20170425072625-481a6f6dc663
	github.com/dnaeon/go-vcr v1.0.1
	github.com/efritz/pentimento v0.0.0-20190429011147-ade47d831101
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/fatih/color v1.16.0
	github.com/fatih/structs v1.1.0
	github.com/felixge/fgprof v0.9.1
	github.com/felixge/httpsnoop v1.0.1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gchaincl/sqlhooks/v2 v2.0.1
	github.com/getsentry/raven-go v0.2.0
	github.com/ghodss/yaml v1.0.0
	github.com/gitchander/permutation v0.0.0-20181107151852-9e56b92e9909
	github.com/glycerine/go-unsnap-stream v0.0.0-20190901134440-81cf024a9e0a // indirect
	github.com/go-enry/go-enry/v2 v2.6.0
	github.com/go-git/go-git/v5 v5.13.0 // indirect
	github.com/go-openapi/runtime v0.19.21 // indirect
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/validate v0.19.11 // indirect
	github.com/go-redsync/redsync v1.4.2
	github.com/gobwas/glob v0.2.3
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/gddo v0.0.0-20200831202555-721e228c7686
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/gomodule/oauth1 v0.0.0-20181215000758-9a59ed3b0a84
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/go-cmp v0.6.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-github/v28 v28.1.1
	github.com/google/go-github/v31 v31.0.0
	github.com/google/go-querystring v1.0.0
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.2.0
	github.com/google/zoekt v0.0.0-20200720095054-b48e35d16e83
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/gorilla/context v1.1.1
	github.com/gorilla/csrf v1.7.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/schema v1.4.1
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/gosimple/slug v1.9.0 // indirect
	github.com/goware/urlx v0.3.1
	github.com/grafana-tools/sdk v0.0.0-20210709154219-f35c5af8140d
	github.com/graph-gophers/graphql-go v0.0.0-20201113091052-beb923fada29
	github.com/graphql-go/graphql v0.7.9
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/golang-lru v0.5.4
	github.com/hexops/autogold v1.3.0
	github.com/honeycombio/libhoney-go v1.14.0
	github.com/inconshreveable/log15 v0.0.0-20201112154412-8562bdadbbac
	github.com/jackc/pgconn v1.8.0
	github.com/jackc/pgx/v4 v4.10.0
	github.com/jmoiron/sqlx v1.2.1-0.20190826204134-d7d95172beb5
	github.com/joho/godotenv v1.3.0
	github.com/jordan-wright/email v4.0.1-0.20200824153738-3f5bafa1cd84+incompatible
	github.com/json-iterator/go v1.1.11
	github.com/karrick/godirwalk v1.16.1
	github.com/keegancsmith/rpc v1.3.0
	github.com/keegancsmith/sqlf v1.1.0
	github.com/keegancsmith/tmpfriend v0.0.0-20180423180255-86e88902a513
	github.com/kr/text v0.2.0
	github.com/kylelemons/godebug v1.1.0
	github.com/lib/pq v1.8.0
	github.com/machinebox/graphql v0.2.2
	github.com/matryer/is v1.4.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.8
	github.com/mcuadros/go-version v0.0.0-20190830083331-035f6764e8d2
	github.com/microcosm-cc/bluemonday v1.0.4
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f
	github.com/neelance/parallel v0.0.0-20160708114440-4de9ce63d14c
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/peterbourgon/ff v1.7.0
	github.com/peterhellberg/link v1.1.0
	github.com/pquerna/cachecontrol v0.0.0-20200819021114-67c6ae64274f // indirect
	github.com/prometheus/alertmanager v0.21.0
	github.com/prometheus/client_golang v1.9.0
	github.com/prometheus/common v0.15.0
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be
	github.com/russellhaering/gosaml2 v0.6.0
	github.com/russellhaering/goxmldsig v1.3.0
	github.com/russross/blackfriday v2.0.0+incompatible // indirect
	github.com/schollz/progressbar/v3 v3.5.0
	github.com/segmentio/fasthash v1.0.3
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3
	github.com/shurcooL/github_flavored_markdown v0.0.0-20181002035957-2122de532470
	github.com/shurcooL/go v0.0.0-20200502201357-93f07166e636 // indirect
	github.com/shurcooL/highlight_diff v0.0.0-20181222201841-111da2e7d480 // indirect
	github.com/shurcooL/highlight_go v0.0.0-20191220051317-782971ddf21b // indirect
	github.com/shurcooL/httpgzip v0.0.0-20190720172056-320755c1c1b0
	github.com/shurcooL/octicon v0.0.0-20191102190552-cbb32d6a785c // indirect
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/batch-change-utils v0.0.0-20210309183117-206c057cc03e
	github.com/sourcegraph/ctxvfs v0.0.0-20180418081416-2b65f1b1ea81
	github.com/sourcegraph/go-ctags v0.0.0-20210426132232-02b1941e7258
	github.com/sourcegraph/go-diff v0.6.1
	github.com/sourcegraph/go-jsonschema v0.0.0-20200907102109-d14e9f2f3a28
	github.com/sourcegraph/go-langserver v2.0.1-0.20181108233942-4a51fa2e1238+incompatible
	github.com/sourcegraph/go-lsp v0.0.0-20200429204803-219e11d77f5d
	github.com/sourcegraph/gosyntect v0.0.0-20210422223331-645353f16ddc
	github.com/sourcegraph/jsonx v0.0.0-20200629203448-1a936bd500cf
	github.com/sourcegraph/sourcegraph/enterprise/dev/ci/images v0.0.0-00010101000000-000000000000
	github.com/sourcegraph/sourcegraph/lib v0.0.0-00010101000000-000000000000
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/stretchr/testify v1.10.0
	github.com/stripe/stripe-go v70.15.0+incompatible
	github.com/temoto/robotstxt v1.1.1
	github.com/throttled/throttled/v2 v2.7.1
	github.com/tidwall/gjson v1.6.8
	github.com/tinylib/msgp v1.1.2 // indirect
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80
	github.com/uber/gonduit v0.11.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	github.com/willf/bitset v1.1.11 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonschema v1.2.0
	github.com/xeonx/timeago v1.0.0-rc4
	github.com/xhit/go-str2duration/v2 v2.0.0
	go.mongodb.org/mongo-driver v1.4.1 // indirect
	go.uber.org/atomic v1.7.0
	go.uber.org/automaxprocs v1.3.0
	go.uber.org/ratelimit v0.2.0
	golang.org/x/crypto v0.35.0
	golang.org/x/net v0.36.0
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
	golang.org/x/sync v0.11.0
	golang.org/x/sys v0.30.0
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	golang.org/x/tools v0.23.0
	google.golang.org/api v0.46.0
	google.golang.org/genproto v0.0.0-20210517163617-5e0236093d7a
	google.golang.org/protobuf v1.34.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/utils v0.0.0-20200729134348-d5654de09c73 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

// Permanent replace directives
// ============================
// These entries indicate permanent replace directives due to significant changes from upstream
// or intentional forks.
replace (
	// We maintain our own fork of Zoekt. Update with ./dev/zoekt/update
	github.com/google/zoekt => github.com/sourcegraph/zoekt v0.0.0-20210721121719-c54bd1fb8d2e
	// We use a fork of Alertmanager to allow prom-wrapper to better manipulate Alertmanager configuration.
	// See https://docs.sourcegraph.com/dev/background-information/observability/prometheus
	github.com/prometheus/alertmanager => github.com/sourcegraph/alertmanager v0.21.1-0.20200727091526-3e856a90b534
	// We publish 'enterprise/dev/ci/images' as a package for import in other tooling.
	// When developing Sourcegraph itself, this replace uses the local package instead of a pushed version.
	github.com/sourcegraph/sourcegraph/enterprise/dev/ci/images => ./enterprise/dev/ci/images
	// We publish 'lib' as a package for import in other tooling.
	// When developing Sourcegraph itself, this replace uses the local package instead of a pushed version.
	github.com/sourcegraph/sourcegraph/lib => ./lib
)

// Temporary replace directives
// ============================
// These entries indicate temporary replace directives due to a pending pull request upstream
// or issues with specific versions.
replace (
	// Pending: https://github.com/ghodss/yaml/pull/65
	github.com/ghodss/yaml => github.com/sourcegraph/yaml v1.0.1-0.20200714132230-56936252f152
	github.com/shurcooL/httpgzip => github.com/sourcegraph/httpgzip v0.0.0-20210213125624-48ebf036a6a1
)

// Status unclear replace directives
// =================================
// These entries indicate replace directives that are defined for unknown reasons.
replace (
	github.com/dghubble/gologin => github.com/sourcegraph/gologin v1.0.2-0.20181110030308-c6f1b62954d8
	github.com/golang/lint => golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f
	github.com/mattn/goreman => github.com/sourcegraph/goreman v0.1.2-0.20180928223752-6e9a2beb830d
	github.com/russellhaering/gosaml2 => github.com/sourcegraph/gosaml2 v0.6.1-0.20210128133756-84151d087b10
	github.com/russross/blackfriday => github.com/russross/blackfriday v1.5.2
	golang.org/x/oauth2 => github.com/sourcegraph/oauth2 v0.0.0-20201011192344-605770292164
)
