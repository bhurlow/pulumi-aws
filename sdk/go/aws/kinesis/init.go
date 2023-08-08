// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:kinesis/analyticsApplication:AnalyticsApplication":
		r = &AnalyticsApplication{}
	case "aws:kinesis/firehoseDeliveryStream:FirehoseDeliveryStream":
		r = &FirehoseDeliveryStream{}
	case "aws:kinesis/stream:Stream":
		r = &Stream{}
	case "aws:kinesis/streamConsumer:StreamConsumer":
		r = &StreamConsumer{}
	case "aws:kinesis/videoStream:VideoStream":
		r = &VideoStream{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws",
		"kinesis/analyticsApplication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"kinesis/firehoseDeliveryStream",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"kinesis/stream",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"kinesis/streamConsumer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"kinesis/videoStream",
		&module{version},
	)
}
