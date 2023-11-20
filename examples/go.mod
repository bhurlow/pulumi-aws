module github.com/pulumi/pulumi-aws/examples/v6

go 1.21.0

require (
	github.com/aws/aws-sdk-go v1.47.7
	github.com/pulumi/pulumi-aws/provider/v6 v6.0.0-00010101000000-000000000000
	github.com/pulumi/pulumi-terraform-bridge/pf v0.19.1-0.20231116043842-89a4965045a1
	github.com/pulumi/pulumi-terraform-bridge/testing v0.0.2-0.20230927165309-e3fd9503f2d3
	github.com/pulumi/pulumi/pkg/v3 v3.93.0
	github.com/stretchr/testify v1.8.4
)

// Replace to allow for correctly linking the aws provider.
//
// We use this for gRPC based testing.
replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20230912190043-e6d96b3b8f7e

	github.com/hashicorp/terraform-provider-aws => ../upstream
	github.com/pulumi/pulumi-aws/provider/v6 => ../provider
)

// This replace is copied from upstream/go.mod, and should be maintained only as long as upstream maintains the same replace.
replace github.com/hashicorp/terraform-plugin-log => github.com/gdavison/terraform-plugin-log v0.0.0-20230928191232-6c653d8ef8fb

require (
	cloud.google.com/go v0.110.7 // indirect
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.1 // indirect
	cloud.google.com/go/kms v1.15.0 // indirect
	cloud.google.com/go/logging v1.7.0 // indirect
	cloud.google.com/go/longrunning v0.5.1 // indirect
	cloud.google.com/go/storage v1.30.1 // indirect
	dario.cat/mergo v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go v66.0.0+incompatible // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.28 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.22 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.11 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230923063757-afb1ddc0824c // indirect
	github.com/YakDriver/regexache v0.23.0 // indirect
	github.com/acomagu/bufpipe v1.0.4 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/armon/go-metrics v0.4.0 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aws/aws-sdk-go-v2 v1.22.2 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.15.2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.13.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/account v1.13.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/acm v1.21.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.34.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrock v1.3.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cleanrooms v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.14.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecatalyst v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.21.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdbelastic v1.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.25.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.130.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/eks v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrserverless v1.13.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/finspace v1.16.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/fis v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/glacier v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/iam v1.27.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internetmonitor v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivschat v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafka v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kendra v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kms v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lambda v1.44.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelsv2 v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/medialive v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/oam v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/pipes v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/pricing v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/qldb v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/rbin v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/rds v1.62.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/rolesanywhere v1.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.42.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3control v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/signer v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.28.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v1.42.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmincidents v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/swf v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/vpclattice v1.4.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.33.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/xray v1.22.1 // indirect
	github.com/aws/smithy-go v1.16.0 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/beevik/etree v1.2.0 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/charmbracelet/bubbles v0.16.1 // indirect
	github.com/charmbracelet/bubbletea v0.24.2 // indirect
	github.com/charmbracelet/lipgloss v0.7.1 // indirect
	github.com/cheggaaa/pb v1.0.29 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/containerd/console v1.0.4-0.20230313162750-1ae8d489ac81 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/djherbis/times v1.5.0 // indirect
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/gedex/inflector v0.0.0-20170307190818-16278e9db813 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/go-git/go-git/v5 v5.9.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/golang/glog v1.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/s2a-go v0.1.4 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.11.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645 // indirect
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.21.0 // indirect
	github.com/hashicorp/aws-sdk-go-base v1.1.0 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.39 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2 v2.0.0-beta.40 // indirect
	github.com/hashicorp/awspolicyequivalence v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200723130312-85980079f637 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/mlock v0.1.2 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.6 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hc-install v0.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.18.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.19.0 // indirect
	github.com/hashicorp/terraform-json v0.17.1 // indirect
	github.com/hashicorp/terraform-plugin-framework v1.4.2 // indirect
	github.com/hashicorp/terraform-plugin-framework-timeouts v0.4.1 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.12.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.19.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/hashicorp/terraform-plugin-mux v0.12.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.29.0 // indirect
	github.com/hashicorp/terraform-plugin-testing v1.5.1 // indirect
	github.com/hashicorp/terraform-provider-aws v1.60.1-0.20220923175450-ca71523cdc36 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.2 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/vault/api v1.8.2 // indirect
	github.com/hashicorp/vault/sdk v0.6.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20230413205102-771768614e91 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/muesli/ansi v0.0.0-20211018074035-2e021307bc4b // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.1 // indirect
	github.com/natefinch/atomic v1.0.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/opentracing/basictracer-go v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pgavlin/goldmark v1.1.33-0.20200616210433-b5eb04559386 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/term v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pulumi/esc v0.5.7-0.20231030195049-f71961c0d5fa // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.65.1-0.20231116043842-89a4965045a1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/x/muxer v0.0.7-0.20230801203955-5d215c892096 // indirect
	github.com/pulumi/pulumi/sdk/v3 v3.93.0 // indirect
	github.com/pulumi/terraform-diff-reader v0.0.2 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.0.0 // indirect
	github.com/segmentio/asm v1.1.3 // indirect
	github.com/segmentio/encoding v0.3.5 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/skeema/knownhosts v1.2.0 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/texttheater/golang-levenshtein v1.0.1 // indirect
	github.com/tweekmonster/luser v0.0.0-20161003172636-3fa38070dbd7 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/zclconf/go-cty v1.14.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.45.0 // indirect
	go.opentelemetry.io/otel v1.19.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/otel/trace v1.19.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	gocloud.dev v0.27.0 // indirect
	gocloud.dev/secrets/hashivault v0.27.0 // indirect
	golang.org/x/crypto v0.15.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/oauth2 v0.10.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/term v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.126.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20230913181813-007df8e322eb // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/grpc v1.58.2 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/frand v1.4.2 // indirect
	sourcegraph.com/sourcegraph/appdash v0.0.0-20211028080628-e2786a622600 // indirect
)
