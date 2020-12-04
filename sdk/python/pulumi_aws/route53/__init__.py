# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .delegation_set import *
from .get_delegation_set import *
from .get_resolver_endpoint import *
from .get_resolver_rule import *
from .get_resolver_rules import *
from .get_zone import *
from .health_check import *
from .query_log import *
from .record import *
from .resolver_endpoint import *
from .resolver_query_log_config import *
from .resolver_query_log_config_association import *
from .resolver_rule import *
from .resolver_rule_association import *
from .vpc_association_authorization import *
from .zone import *
from .zone_association import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:route53/delegationSet:DelegationSet":
                return DelegationSet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/healthCheck:HealthCheck":
                return HealthCheck(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/queryLog:QueryLog":
                return QueryLog(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/record:Record":
                return Record(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/resolverEndpoint:ResolverEndpoint":
                return ResolverEndpoint(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig":
                return ResolverQueryLogConfig(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/resolverQueryLogConfigAssociation:ResolverQueryLogConfigAssociation":
                return ResolverQueryLogConfigAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/resolverRule:ResolverRule":
                return ResolverRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/resolverRuleAssociation:ResolverRuleAssociation":
                return ResolverRuleAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization":
                return VpcAssociationAuthorization(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/zone:Zone":
                return Zone(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:route53/zoneAssociation:ZoneAssociation":
                return ZoneAssociation(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "route53/delegationSet", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/healthCheck", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/queryLog", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/record", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/resolverEndpoint", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/resolverQueryLogConfig", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/resolverQueryLogConfigAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/resolverRule", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/resolverRuleAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/vpcAssociationAuthorization", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/zone", _module_instance)
    pulumi.runtime.register_resource_module("aws", "route53/zoneAssociation", _module_instance)

_register_module()
