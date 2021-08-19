# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DefaultSecurityGroupArgs', 'DefaultSecurityGroup']

@pulumi.input_type
class DefaultSecurityGroupArgs:
    def __init__(__self__, *,
                 egress: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupEgressArgs']]]] = None,
                 ingress: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupIngressArgs']]]] = None,
                 revoke_rules_on_delete: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DefaultSecurityGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupEgressArgs']]] egress: Configuration block. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupIngressArgs']]] ingress: Configuration block. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] vpc_id: VPC ID. **Note that changing the `vpc_id` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
        """
        if egress is not None:
            pulumi.set(__self__, "egress", egress)
        if ingress is not None:
            pulumi.set(__self__, "ingress", ingress)
        if revoke_rules_on_delete is not None:
            pulumi.set(__self__, "revoke_rules_on_delete", revoke_rules_on_delete)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def egress(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupEgressArgs']]]]:
        """
        Configuration block. Detailed below.
        """
        return pulumi.get(self, "egress")

    @egress.setter
    def egress(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupEgressArgs']]]]):
        pulumi.set(self, "egress", value)

    @property
    @pulumi.getter
    def ingress(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupIngressArgs']]]]:
        """
        Configuration block. Detailed below.
        """
        return pulumi.get(self, "ingress")

    @ingress.setter
    def ingress(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupIngressArgs']]]]):
        pulumi.set(self, "ingress", value)

    @property
    @pulumi.getter(name="revokeRulesOnDelete")
    def revoke_rules_on_delete(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "revoke_rules_on_delete")

    @revoke_rules_on_delete.setter
    def revoke_rules_on_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "revoke_rules_on_delete", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID. **Note that changing the `vpc_id` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _DefaultSecurityGroupState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupEgressArgs']]]] = None,
                 ingress: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupIngressArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 revoke_rules_on_delete: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DefaultSecurityGroup resources.
        :param pulumi.Input[str] arn: ARN of the security group.
        :param pulumi.Input[str] description: Description of this rule.
        :param pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupEgressArgs']]] egress: Configuration block. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupIngressArgs']]] ingress: Configuration block. Detailed below.
        :param pulumi.Input[str] name: Name of the security group.
        :param pulumi.Input[str] owner_id: Owner ID.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] vpc_id: VPC ID. **Note that changing the `vpc_id` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if egress is not None:
            pulumi.set(__self__, "egress", egress)
        if ingress is not None:
            pulumi.set(__self__, "ingress", ingress)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if revoke_rules_on_delete is not None:
            pulumi.set(__self__, "revoke_rules_on_delete", revoke_rules_on_delete)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the security group.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of this rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def egress(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupEgressArgs']]]]:
        """
        Configuration block. Detailed below.
        """
        return pulumi.get(self, "egress")

    @egress.setter
    def egress(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupEgressArgs']]]]):
        pulumi.set(self, "egress", value)

    @property
    @pulumi.getter
    def ingress(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupIngressArgs']]]]:
        """
        Configuration block. Detailed below.
        """
        return pulumi.get(self, "ingress")

    @ingress.setter
    def ingress(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultSecurityGroupIngressArgs']]]]):
        pulumi.set(self, "ingress", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the security group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        Owner ID.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter(name="revokeRulesOnDelete")
    def revoke_rules_on_delete(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "revoke_rules_on_delete")

    @revoke_rules_on_delete.setter
    def revoke_rules_on_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "revoke_rules_on_delete", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID. **Note that changing the `vpc_id` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class DefaultSecurityGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupEgressArgs']]]]] = None,
                 ingress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupIngressArgs']]]]] = None,
                 revoke_rules_on_delete: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage a default security group. This resource can manage the default security group of the default or a non-default VPC.

        > **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ec2.DefaultSecurityGroup` resource behaves differently from normal resources. This provider does not _create_ this resource but instead attempts to "adopt" it into management.

        For EC2 Classic accounts, each region comes with a default security group. Additionally, each VPC created in AWS comes with a default security group that can be managed but not destroyed.

        When the provider first adopts the default security group, it **immediately removes all ingress and egress rules in the Security Group**. It then creates any rules specified in the configuration. This way only the rules specified in the configuration are created.

        This resource treats its inline rules as absolute; only the rules defined inline are created, and any additions/removals external to this resource will result in diff shown. For these reasons, this resource is incompatible with the `ec2.SecurityGroupRule` resource.

        For more information about default security groups, see the AWS documentation on [Default Security Groups][aws-default-security-groups]. To manage normal security groups, see the `ec2.SecurityGroup` resource.

        ## Example Usage

        The following config gives the default security group the same rules that AWS provides by default but under management by this provider. This means that any ingress or egress rules added or changed will be detected as drift.

        ```python
        import pulumi
        import pulumi_aws as aws

        mainvpc = aws.ec2.Vpc("mainvpc", cidr_block="10.1.0.0/16")
        default = aws.ec2.DefaultSecurityGroup("default",
            vpc_id=mainvpc.id,
            ingress=[aws.ec2.DefaultSecurityGroupIngressArgs(
                protocol="-1",
                self=True,
                from_port=0,
                to_port=0,
            )],
            egress=[aws.ec2.DefaultSecurityGroupEgressArgs(
                from_port=0,
                to_port=0,
                protocol="-1",
                cidr_blocks=["0.0.0.0/0"],
            )])
        ```
        ### Example Config To Deny All Egress Traffic, Allowing Ingress

        The following denies all Egress traffic by omitting any `egress` rules, while including the default `ingress` rule to allow all traffic.

        ```python
        import pulumi
        import pulumi_aws as aws

        mainvpc = aws.ec2.Vpc("mainvpc", cidr_block="10.1.0.0/16")
        default = aws.ec2.DefaultSecurityGroup("default",
            vpc_id=mainvpc.id,
            ingress=[aws.ec2.DefaultSecurityGroupIngressArgs(
                protocol="-1",
                self=True,
                from_port=0,
                to_port=0,
            )])
        ```
        ### Removing `ec2.DefaultSecurityGroup` From Your Configuration

        Removing this resource from your configuration will remove it from your statefile and management, but will not destroy the Security Group. All ingress or egress rules will be left as they are at the time of removal. You can resume managing them via the AWS Console.

        ## Import

        Security Groups can be imported using the `security group id`, e.g.

        ```sh
         $ pulumi import aws:ec2/defaultSecurityGroup:DefaultSecurityGroup default_sg sg-903004f8
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupEgressArgs']]]] egress: Configuration block. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupIngressArgs']]]] ingress: Configuration block. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] vpc_id: VPC ID. **Note that changing the `vpc_id` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DefaultSecurityGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage a default security group. This resource can manage the default security group of the default or a non-default VPC.

        > **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ec2.DefaultSecurityGroup` resource behaves differently from normal resources. This provider does not _create_ this resource but instead attempts to "adopt" it into management.

        For EC2 Classic accounts, each region comes with a default security group. Additionally, each VPC created in AWS comes with a default security group that can be managed but not destroyed.

        When the provider first adopts the default security group, it **immediately removes all ingress and egress rules in the Security Group**. It then creates any rules specified in the configuration. This way only the rules specified in the configuration are created.

        This resource treats its inline rules as absolute; only the rules defined inline are created, and any additions/removals external to this resource will result in diff shown. For these reasons, this resource is incompatible with the `ec2.SecurityGroupRule` resource.

        For more information about default security groups, see the AWS documentation on [Default Security Groups][aws-default-security-groups]. To manage normal security groups, see the `ec2.SecurityGroup` resource.

        ## Example Usage

        The following config gives the default security group the same rules that AWS provides by default but under management by this provider. This means that any ingress or egress rules added or changed will be detected as drift.

        ```python
        import pulumi
        import pulumi_aws as aws

        mainvpc = aws.ec2.Vpc("mainvpc", cidr_block="10.1.0.0/16")
        default = aws.ec2.DefaultSecurityGroup("default",
            vpc_id=mainvpc.id,
            ingress=[aws.ec2.DefaultSecurityGroupIngressArgs(
                protocol="-1",
                self=True,
                from_port=0,
                to_port=0,
            )],
            egress=[aws.ec2.DefaultSecurityGroupEgressArgs(
                from_port=0,
                to_port=0,
                protocol="-1",
                cidr_blocks=["0.0.0.0/0"],
            )])
        ```
        ### Example Config To Deny All Egress Traffic, Allowing Ingress

        The following denies all Egress traffic by omitting any `egress` rules, while including the default `ingress` rule to allow all traffic.

        ```python
        import pulumi
        import pulumi_aws as aws

        mainvpc = aws.ec2.Vpc("mainvpc", cidr_block="10.1.0.0/16")
        default = aws.ec2.DefaultSecurityGroup("default",
            vpc_id=mainvpc.id,
            ingress=[aws.ec2.DefaultSecurityGroupIngressArgs(
                protocol="-1",
                self=True,
                from_port=0,
                to_port=0,
            )])
        ```
        ### Removing `ec2.DefaultSecurityGroup` From Your Configuration

        Removing this resource from your configuration will remove it from your statefile and management, but will not destroy the Security Group. All ingress or egress rules will be left as they are at the time of removal. You can resume managing them via the AWS Console.

        ## Import

        Security Groups can be imported using the `security group id`, e.g.

        ```sh
         $ pulumi import aws:ec2/defaultSecurityGroup:DefaultSecurityGroup default_sg sg-903004f8
        ```

        :param str resource_name: The name of the resource.
        :param DefaultSecurityGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefaultSecurityGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupEgressArgs']]]]] = None,
                 ingress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupIngressArgs']]]]] = None,
                 revoke_rules_on_delete: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DefaultSecurityGroupArgs.__new__(DefaultSecurityGroupArgs)

            __props__.__dict__["egress"] = egress
            __props__.__dict__["ingress"] = ingress
            __props__.__dict__["revoke_rules_on_delete"] = revoke_rules_on_delete
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["description"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["owner_id"] = None
            __props__.__dict__["tags_all"] = None
        super(DefaultSecurityGroup, __self__).__init__(
            'aws:ec2/defaultSecurityGroup:DefaultSecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            egress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupEgressArgs']]]]] = None,
            ingress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupIngressArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            revoke_rules_on_delete: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'DefaultSecurityGroup':
        """
        Get an existing DefaultSecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the security group.
        :param pulumi.Input[str] description: Description of this rule.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupEgressArgs']]]] egress: Configuration block. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultSecurityGroupIngressArgs']]]] ingress: Configuration block. Detailed below.
        :param pulumi.Input[str] name: Name of the security group.
        :param pulumi.Input[str] owner_id: Owner ID.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] vpc_id: VPC ID. **Note that changing the `vpc_id` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DefaultSecurityGroupState.__new__(_DefaultSecurityGroupState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["egress"] = egress
        __props__.__dict__["ingress"] = ingress
        __props__.__dict__["name"] = name
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["revoke_rules_on_delete"] = revoke_rules_on_delete
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["vpc_id"] = vpc_id
        return DefaultSecurityGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the security group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of this rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def egress(self) -> pulumi.Output[Sequence['outputs.DefaultSecurityGroupEgress']]:
        """
        Configuration block. Detailed below.
        """
        return pulumi.get(self, "egress")

    @property
    @pulumi.getter
    def ingress(self) -> pulumi.Output[Sequence['outputs.DefaultSecurityGroupIngress']]:
        """
        Configuration block. Detailed below.
        """
        return pulumi.get(self, "ingress")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the security group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        Owner ID.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="revokeRulesOnDelete")
    def revoke_rules_on_delete(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "revoke_rules_on_delete")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        VPC ID. **Note that changing the `vpc_id` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
        """
        return pulumi.get(self, "vpc_id")

