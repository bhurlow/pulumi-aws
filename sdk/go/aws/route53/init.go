// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
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
	case "aws:route53/delegationSet:DelegationSet":
		r = &DelegationSet{}
	case "aws:route53/healthCheck:HealthCheck":
		r = &HealthCheck{}
	case "aws:route53/hostedZoneDnsSec:HostedZoneDnsSec":
		r = &HostedZoneDnsSec{}
	case "aws:route53/keySigningKey:KeySigningKey":
		r = &KeySigningKey{}
	case "aws:route53/queryLog:QueryLog":
		r = &QueryLog{}
	case "aws:route53/record:Record":
		r = &Record{}
	case "aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig":
		r = &ResolverDnsSecConfig{}
	case "aws:route53/resolverEndpoint:ResolverEndpoint":
		r = &ResolverEndpoint{}
	case "aws:route53/resolverFirewallDomainList:ResolverFirewallDomainList":
		r = &ResolverFirewallDomainList{}
	case "aws:route53/resolverFirewallRuleGroup:ResolverFirewallRuleGroup":
		r = &ResolverFirewallRuleGroup{}
	case "aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig":
		r = &ResolverQueryLogConfig{}
	case "aws:route53/resolverQueryLogConfigAssociation:ResolverQueryLogConfigAssociation":
		r = &ResolverQueryLogConfigAssociation{}
	case "aws:route53/resolverRule:ResolverRule":
		r = &ResolverRule{}
	case "aws:route53/resolverRuleAssociation:ResolverRuleAssociation":
		r = &ResolverRuleAssociation{}
	case "aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization":
		r = &VpcAssociationAuthorization{}
	case "aws:route53/zone:Zone":
		r = &Zone{}
	case "aws:route53/zoneAssociation:ZoneAssociation":
		r = &ZoneAssociation{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"route53/delegationSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/healthCheck",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/hostedZoneDnsSec",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/keySigningKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/queryLog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/record",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverDnsSecConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverFirewallDomainList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverFirewallRuleGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverQueryLogConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverQueryLogConfigAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverRuleAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/vpcAssociationAuthorization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/zone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/zoneAssociation",
		&module{version},
	)
}
