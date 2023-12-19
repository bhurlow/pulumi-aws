# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetLoadBalancerResult',
    'AwaitableGetLoadBalancerResult',
    'get_load_balancer',
    'get_load_balancer_output',
]

@pulumi.output_type
class GetLoadBalancerResult:
    """
    A collection of values returned by getLoadBalancer.
    """
    def __init__(__self__, access_logs=None, arn=None, arn_suffix=None, connection_logs=None, customer_owned_ipv4_pool=None, desync_mitigation_mode=None, dns_name=None, dns_record_client_routing_policy=None, drop_invalid_header_fields=None, enable_cross_zone_load_balancing=None, enable_deletion_protection=None, enable_http2=None, enable_tls_version_and_cipher_suite_headers=None, enable_waf_fail_open=None, enable_xff_client_port=None, enforce_security_group_inbound_rules_on_private_link_traffic=None, id=None, idle_timeout=None, internal=None, ip_address_type=None, load_balancer_type=None, name=None, preserve_host_header=None, security_groups=None, subnet_mappings=None, subnets=None, tags=None, vpc_id=None, xff_header_processing_mode=None, zone_id=None):
        if access_logs and not isinstance(access_logs, dict):
            raise TypeError("Expected argument 'access_logs' to be a dict")
        pulumi.set(__self__, "access_logs", access_logs)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if arn_suffix and not isinstance(arn_suffix, str):
            raise TypeError("Expected argument 'arn_suffix' to be a str")
        pulumi.set(__self__, "arn_suffix", arn_suffix)
        if connection_logs and not isinstance(connection_logs, list):
            raise TypeError("Expected argument 'connection_logs' to be a list")
        pulumi.set(__self__, "connection_logs", connection_logs)
        if customer_owned_ipv4_pool and not isinstance(customer_owned_ipv4_pool, str):
            raise TypeError("Expected argument 'customer_owned_ipv4_pool' to be a str")
        pulumi.set(__self__, "customer_owned_ipv4_pool", customer_owned_ipv4_pool)
        if desync_mitigation_mode and not isinstance(desync_mitigation_mode, str):
            raise TypeError("Expected argument 'desync_mitigation_mode' to be a str")
        pulumi.set(__self__, "desync_mitigation_mode", desync_mitigation_mode)
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if dns_record_client_routing_policy and not isinstance(dns_record_client_routing_policy, str):
            raise TypeError("Expected argument 'dns_record_client_routing_policy' to be a str")
        pulumi.set(__self__, "dns_record_client_routing_policy", dns_record_client_routing_policy)
        if drop_invalid_header_fields and not isinstance(drop_invalid_header_fields, bool):
            raise TypeError("Expected argument 'drop_invalid_header_fields' to be a bool")
        pulumi.set(__self__, "drop_invalid_header_fields", drop_invalid_header_fields)
        if enable_cross_zone_load_balancing and not isinstance(enable_cross_zone_load_balancing, bool):
            raise TypeError("Expected argument 'enable_cross_zone_load_balancing' to be a bool")
        pulumi.set(__self__, "enable_cross_zone_load_balancing", enable_cross_zone_load_balancing)
        if enable_deletion_protection and not isinstance(enable_deletion_protection, bool):
            raise TypeError("Expected argument 'enable_deletion_protection' to be a bool")
        pulumi.set(__self__, "enable_deletion_protection", enable_deletion_protection)
        if enable_http2 and not isinstance(enable_http2, bool):
            raise TypeError("Expected argument 'enable_http2' to be a bool")
        pulumi.set(__self__, "enable_http2", enable_http2)
        if enable_tls_version_and_cipher_suite_headers and not isinstance(enable_tls_version_and_cipher_suite_headers, bool):
            raise TypeError("Expected argument 'enable_tls_version_and_cipher_suite_headers' to be a bool")
        pulumi.set(__self__, "enable_tls_version_and_cipher_suite_headers", enable_tls_version_and_cipher_suite_headers)
        if enable_waf_fail_open and not isinstance(enable_waf_fail_open, bool):
            raise TypeError("Expected argument 'enable_waf_fail_open' to be a bool")
        pulumi.set(__self__, "enable_waf_fail_open", enable_waf_fail_open)
        if enable_xff_client_port and not isinstance(enable_xff_client_port, bool):
            raise TypeError("Expected argument 'enable_xff_client_port' to be a bool")
        pulumi.set(__self__, "enable_xff_client_port", enable_xff_client_port)
        if enforce_security_group_inbound_rules_on_private_link_traffic and not isinstance(enforce_security_group_inbound_rules_on_private_link_traffic, str):
            raise TypeError("Expected argument 'enforce_security_group_inbound_rules_on_private_link_traffic' to be a str")
        pulumi.set(__self__, "enforce_security_group_inbound_rules_on_private_link_traffic", enforce_security_group_inbound_rules_on_private_link_traffic)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idle_timeout and not isinstance(idle_timeout, int):
            raise TypeError("Expected argument 'idle_timeout' to be a int")
        pulumi.set(__self__, "idle_timeout", idle_timeout)
        if internal and not isinstance(internal, bool):
            raise TypeError("Expected argument 'internal' to be a bool")
        pulumi.set(__self__, "internal", internal)
        if ip_address_type and not isinstance(ip_address_type, str):
            raise TypeError("Expected argument 'ip_address_type' to be a str")
        pulumi.set(__self__, "ip_address_type", ip_address_type)
        if load_balancer_type and not isinstance(load_balancer_type, str):
            raise TypeError("Expected argument 'load_balancer_type' to be a str")
        pulumi.set(__self__, "load_balancer_type", load_balancer_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if preserve_host_header and not isinstance(preserve_host_header, bool):
            raise TypeError("Expected argument 'preserve_host_header' to be a bool")
        pulumi.set(__self__, "preserve_host_header", preserve_host_header)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if subnet_mappings and not isinstance(subnet_mappings, list):
            raise TypeError("Expected argument 'subnet_mappings' to be a list")
        pulumi.set(__self__, "subnet_mappings", subnet_mappings)
        if subnets and not isinstance(subnets, list):
            raise TypeError("Expected argument 'subnets' to be a list")
        pulumi.set(__self__, "subnets", subnets)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if xff_header_processing_mode and not isinstance(xff_header_processing_mode, str):
            raise TypeError("Expected argument 'xff_header_processing_mode' to be a str")
        pulumi.set(__self__, "xff_header_processing_mode", xff_header_processing_mode)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="accessLogs")
    def access_logs(self) -> 'outputs.GetLoadBalancerAccessLogsResult':
        return pulumi.get(self, "access_logs")

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="arnSuffix")
    def arn_suffix(self) -> str:
        return pulumi.get(self, "arn_suffix")

    @property
    @pulumi.getter(name="connectionLogs")
    def connection_logs(self) -> Sequence['outputs.GetLoadBalancerConnectionLogResult']:
        return pulumi.get(self, "connection_logs")

    @property
    @pulumi.getter(name="customerOwnedIpv4Pool")
    def customer_owned_ipv4_pool(self) -> str:
        return pulumi.get(self, "customer_owned_ipv4_pool")

    @property
    @pulumi.getter(name="desyncMitigationMode")
    def desync_mitigation_mode(self) -> str:
        return pulumi.get(self, "desync_mitigation_mode")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="dnsRecordClientRoutingPolicy")
    def dns_record_client_routing_policy(self) -> str:
        return pulumi.get(self, "dns_record_client_routing_policy")

    @property
    @pulumi.getter(name="dropInvalidHeaderFields")
    def drop_invalid_header_fields(self) -> bool:
        return pulumi.get(self, "drop_invalid_header_fields")

    @property
    @pulumi.getter(name="enableCrossZoneLoadBalancing")
    def enable_cross_zone_load_balancing(self) -> bool:
        return pulumi.get(self, "enable_cross_zone_load_balancing")

    @property
    @pulumi.getter(name="enableDeletionProtection")
    def enable_deletion_protection(self) -> bool:
        return pulumi.get(self, "enable_deletion_protection")

    @property
    @pulumi.getter(name="enableHttp2")
    def enable_http2(self) -> bool:
        return pulumi.get(self, "enable_http2")

    @property
    @pulumi.getter(name="enableTlsVersionAndCipherSuiteHeaders")
    def enable_tls_version_and_cipher_suite_headers(self) -> bool:
        return pulumi.get(self, "enable_tls_version_and_cipher_suite_headers")

    @property
    @pulumi.getter(name="enableWafFailOpen")
    def enable_waf_fail_open(self) -> bool:
        return pulumi.get(self, "enable_waf_fail_open")

    @property
    @pulumi.getter(name="enableXffClientPort")
    def enable_xff_client_port(self) -> bool:
        return pulumi.get(self, "enable_xff_client_port")

    @property
    @pulumi.getter(name="enforceSecurityGroupInboundRulesOnPrivateLinkTraffic")
    def enforce_security_group_inbound_rules_on_private_link_traffic(self) -> str:
        return pulumi.get(self, "enforce_security_group_inbound_rules_on_private_link_traffic")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> int:
        return pulumi.get(self, "idle_timeout")

    @property
    @pulumi.getter
    def internal(self) -> bool:
        return pulumi.get(self, "internal")

    @property
    @pulumi.getter(name="ipAddressType")
    def ip_address_type(self) -> str:
        return pulumi.get(self, "ip_address_type")

    @property
    @pulumi.getter(name="loadBalancerType")
    def load_balancer_type(self) -> str:
        return pulumi.get(self, "load_balancer_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="preserveHostHeader")
    def preserve_host_header(self) -> bool:
        return pulumi.get(self, "preserve_host_header")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence[str]:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="subnetMappings")
    def subnet_mappings(self) -> Sequence['outputs.GetLoadBalancerSubnetMappingResult']:
        return pulumi.get(self, "subnet_mappings")

    @property
    @pulumi.getter
    def subnets(self) -> Sequence[str]:
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="xffHeaderProcessingMode")
    def xff_header_processing_mode(self) -> str:
        return pulumi.get(self, "xff_header_processing_mode")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        return pulumi.get(self, "zone_id")


class AwaitableGetLoadBalancerResult(GetLoadBalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancerResult(
            access_logs=self.access_logs,
            arn=self.arn,
            arn_suffix=self.arn_suffix,
            connection_logs=self.connection_logs,
            customer_owned_ipv4_pool=self.customer_owned_ipv4_pool,
            desync_mitigation_mode=self.desync_mitigation_mode,
            dns_name=self.dns_name,
            dns_record_client_routing_policy=self.dns_record_client_routing_policy,
            drop_invalid_header_fields=self.drop_invalid_header_fields,
            enable_cross_zone_load_balancing=self.enable_cross_zone_load_balancing,
            enable_deletion_protection=self.enable_deletion_protection,
            enable_http2=self.enable_http2,
            enable_tls_version_and_cipher_suite_headers=self.enable_tls_version_and_cipher_suite_headers,
            enable_waf_fail_open=self.enable_waf_fail_open,
            enable_xff_client_port=self.enable_xff_client_port,
            enforce_security_group_inbound_rules_on_private_link_traffic=self.enforce_security_group_inbound_rules_on_private_link_traffic,
            id=self.id,
            idle_timeout=self.idle_timeout,
            internal=self.internal,
            ip_address_type=self.ip_address_type,
            load_balancer_type=self.load_balancer_type,
            name=self.name,
            preserve_host_header=self.preserve_host_header,
            security_groups=self.security_groups,
            subnet_mappings=self.subnet_mappings,
            subnets=self.subnets,
            tags=self.tags,
            vpc_id=self.vpc_id,
            xff_header_processing_mode=self.xff_header_processing_mode,
            zone_id=self.zone_id)


def get_load_balancer(arn: Optional[str] = None,
                      name: Optional[str] = None,
                      tags: Optional[Mapping[str, str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancerResult:
    """
    > **Note:** `alb.LoadBalancer` is known as `lb.LoadBalancer`. The functionality is identical.

    Provides information about a Load Balancer.

    This data source can prove useful when a module accepts an LB as an input
    variable and needs to, for example, determine the security groups associated
    with it, etc.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    lb_arn = config.get("lbArn")
    if lb_arn is None:
        lb_arn = ""
    lb_name = config.get("lbName")
    if lb_name is None:
        lb_name = ""
    test = aws.lb.get_load_balancer(arn=lb_arn,
        name=lb_name)
    ```


    :param str arn: Full ARN of the load balancer.
    :param str name: Unique name of the load balancer.
    :param Mapping[str, str] tags: Mapping of tags, each pair of which must exactly match a pair on the desired load balancer.
           
           > **NOTE:** When both `arn` and `name` are specified, `arn` takes precedence. `tags` has lowest precedence.
    """
    __args__ = dict()
    __args__['arn'] = arn
    __args__['name'] = name
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:alb/getLoadBalancer:getLoadBalancer', __args__, opts=opts, typ=GetLoadBalancerResult).value

    return AwaitableGetLoadBalancerResult(
        access_logs=pulumi.get(__ret__, 'access_logs'),
        arn=pulumi.get(__ret__, 'arn'),
        arn_suffix=pulumi.get(__ret__, 'arn_suffix'),
        connection_logs=pulumi.get(__ret__, 'connection_logs'),
        customer_owned_ipv4_pool=pulumi.get(__ret__, 'customer_owned_ipv4_pool'),
        desync_mitigation_mode=pulumi.get(__ret__, 'desync_mitigation_mode'),
        dns_name=pulumi.get(__ret__, 'dns_name'),
        dns_record_client_routing_policy=pulumi.get(__ret__, 'dns_record_client_routing_policy'),
        drop_invalid_header_fields=pulumi.get(__ret__, 'drop_invalid_header_fields'),
        enable_cross_zone_load_balancing=pulumi.get(__ret__, 'enable_cross_zone_load_balancing'),
        enable_deletion_protection=pulumi.get(__ret__, 'enable_deletion_protection'),
        enable_http2=pulumi.get(__ret__, 'enable_http2'),
        enable_tls_version_and_cipher_suite_headers=pulumi.get(__ret__, 'enable_tls_version_and_cipher_suite_headers'),
        enable_waf_fail_open=pulumi.get(__ret__, 'enable_waf_fail_open'),
        enable_xff_client_port=pulumi.get(__ret__, 'enable_xff_client_port'),
        enforce_security_group_inbound_rules_on_private_link_traffic=pulumi.get(__ret__, 'enforce_security_group_inbound_rules_on_private_link_traffic'),
        id=pulumi.get(__ret__, 'id'),
        idle_timeout=pulumi.get(__ret__, 'idle_timeout'),
        internal=pulumi.get(__ret__, 'internal'),
        ip_address_type=pulumi.get(__ret__, 'ip_address_type'),
        load_balancer_type=pulumi.get(__ret__, 'load_balancer_type'),
        name=pulumi.get(__ret__, 'name'),
        preserve_host_header=pulumi.get(__ret__, 'preserve_host_header'),
        security_groups=pulumi.get(__ret__, 'security_groups'),
        subnet_mappings=pulumi.get(__ret__, 'subnet_mappings'),
        subnets=pulumi.get(__ret__, 'subnets'),
        tags=pulumi.get(__ret__, 'tags'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'),
        xff_header_processing_mode=pulumi.get(__ret__, 'xff_header_processing_mode'),
        zone_id=pulumi.get(__ret__, 'zone_id'))


@_utilities.lift_output_func(get_load_balancer)
def get_load_balancer_output(arn: Optional[pulumi.Input[Optional[str]]] = None,
                             name: Optional[pulumi.Input[Optional[str]]] = None,
                             tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLoadBalancerResult]:
    """
    > **Note:** `alb.LoadBalancer` is known as `lb.LoadBalancer`. The functionality is identical.

    Provides information about a Load Balancer.

    This data source can prove useful when a module accepts an LB as an input
    variable and needs to, for example, determine the security groups associated
    with it, etc.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    lb_arn = config.get("lbArn")
    if lb_arn is None:
        lb_arn = ""
    lb_name = config.get("lbName")
    if lb_name is None:
        lb_name = ""
    test = aws.lb.get_load_balancer(arn=lb_arn,
        name=lb_name)
    ```


    :param str arn: Full ARN of the load balancer.
    :param str name: Unique name of the load balancer.
    :param Mapping[str, str] tags: Mapping of tags, each pair of which must exactly match a pair on the desired load balancer.
           
           > **NOTE:** When both `arn` and `name` are specified, `arn` takes precedence. `tags` has lowest precedence.
    """
    ...
