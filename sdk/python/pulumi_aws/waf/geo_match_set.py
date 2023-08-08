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
from ._inputs import *

__all__ = ['GeoMatchSetArgs', 'GeoMatchSet']

@pulumi.input_type
class GeoMatchSetArgs:
    def __init__(__self__, *,
                 geo_match_constraints: Optional[pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GeoMatchSet resource.
        :param pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]] geo_match_constraints: The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
        :param pulumi.Input[str] name: The name or description of the GeoMatchSet.
        """
        if geo_match_constraints is not None:
            pulumi.set(__self__, "geo_match_constraints", geo_match_constraints)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="geoMatchConstraints")
    def geo_match_constraints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]]]:
        """
        The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
        """
        return pulumi.get(self, "geo_match_constraints")

    @geo_match_constraints.setter
    def geo_match_constraints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]]]):
        pulumi.set(self, "geo_match_constraints", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the GeoMatchSet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GeoMatchSetState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 geo_match_constraints: Optional[pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GeoMatchSet resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]] geo_match_constraints: The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
        :param pulumi.Input[str] name: The name or description of the GeoMatchSet.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if geo_match_constraints is not None:
            pulumi.set(__self__, "geo_match_constraints", geo_match_constraints)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN)
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="geoMatchConstraints")
    def geo_match_constraints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]]]:
        """
        The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
        """
        return pulumi.get(self, "geo_match_constraints")

    @geo_match_constraints.setter
    def geo_match_constraints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]]]):
        pulumi.set(self, "geo_match_constraints", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the GeoMatchSet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class GeoMatchSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 geo_match_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoMatchSetGeoMatchConstraintArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a WAF Geo Match Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        geo_match_set = aws.waf.GeoMatchSet("geoMatchSet", geo_match_constraints=[
            aws.waf.GeoMatchSetGeoMatchConstraintArgs(
                type="Country",
                value="US",
            ),
            aws.waf.GeoMatchSetGeoMatchConstraintArgs(
                type="Country",
                value="CA",
            ),
        ])
        ```

        ## Import

        terraform import {

         to = aws_waf_geo_match_set.example

         id = "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc" } Using `pulumi import`, import WAF Geo Match Set using their ID. For exampleconsole % pulumi import aws_waf_geo_match_set.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoMatchSetGeoMatchConstraintArgs']]]] geo_match_constraints: The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
        :param pulumi.Input[str] name: The name or description of the GeoMatchSet.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GeoMatchSetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF Geo Match Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        geo_match_set = aws.waf.GeoMatchSet("geoMatchSet", geo_match_constraints=[
            aws.waf.GeoMatchSetGeoMatchConstraintArgs(
                type="Country",
                value="US",
            ),
            aws.waf.GeoMatchSetGeoMatchConstraintArgs(
                type="Country",
                value="CA",
            ),
        ])
        ```

        ## Import

        terraform import {

         to = aws_waf_geo_match_set.example

         id = "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc" } Using `pulumi import`, import WAF Geo Match Set using their ID. For exampleconsole % pulumi import aws_waf_geo_match_set.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc

        :param str resource_name: The name of the resource.
        :param GeoMatchSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GeoMatchSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 geo_match_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoMatchSetGeoMatchConstraintArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GeoMatchSetArgs.__new__(GeoMatchSetArgs)

            __props__.__dict__["geo_match_constraints"] = geo_match_constraints
            __props__.__dict__["name"] = name
            __props__.__dict__["arn"] = None
        super(GeoMatchSet, __self__).__init__(
            'aws:waf/geoMatchSet:GeoMatchSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            geo_match_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoMatchSetGeoMatchConstraintArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'GeoMatchSet':
        """
        Get an existing GeoMatchSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoMatchSetGeoMatchConstraintArgs']]]] geo_match_constraints: The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
        :param pulumi.Input[str] name: The name or description of the GeoMatchSet.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GeoMatchSetState.__new__(_GeoMatchSetState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["geo_match_constraints"] = geo_match_constraints
        __props__.__dict__["name"] = name
        return GeoMatchSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN)
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="geoMatchConstraints")
    def geo_match_constraints(self) -> pulumi.Output[Optional[Sequence['outputs.GeoMatchSetGeoMatchConstraint']]]:
        """
        The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
        """
        return pulumi.get(self, "geo_match_constraints")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name or description of the GeoMatchSet.
        """
        return pulumi.get(self, "name")

