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

__all__ = ['RegexMatchSetArgs', 'RegexMatchSet']

@pulumi.input_type
class RegexMatchSetArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 regex_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input['RegexMatchSetRegexMatchTupleArgs']]]] = None):
        """
        The set of arguments for constructing a RegexMatchSet resource.
        :param pulumi.Input[str] name: The name or description of the Regex Match Set.
        :param pulumi.Input[Sequence[pulumi.Input['RegexMatchSetRegexMatchTupleArgs']]] regex_match_tuples: The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regex_match_tuples is not None:
            pulumi.set(__self__, "regex_match_tuples", regex_match_tuples)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the Regex Match Set.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regexMatchTuples")
    def regex_match_tuples(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegexMatchSetRegexMatchTupleArgs']]]]:
        """
        The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
        """
        return pulumi.get(self, "regex_match_tuples")

    @regex_match_tuples.setter
    def regex_match_tuples(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegexMatchSetRegexMatchTupleArgs']]]]):
        pulumi.set(self, "regex_match_tuples", value)


@pulumi.input_type
class _RegexMatchSetState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regex_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input['RegexMatchSetRegexMatchTupleArgs']]]] = None):
        """
        Input properties used for looking up and filtering RegexMatchSet resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] name: The name or description of the Regex Match Set.
        :param pulumi.Input[Sequence[pulumi.Input['RegexMatchSetRegexMatchTupleArgs']]] regex_match_tuples: The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regex_match_tuples is not None:
            pulumi.set(__self__, "regex_match_tuples", regex_match_tuples)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the Regex Match Set.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regexMatchTuples")
    def regex_match_tuples(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegexMatchSetRegexMatchTupleArgs']]]]:
        """
        The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
        """
        return pulumi.get(self, "regex_match_tuples")

    @regex_match_tuples.setter
    def regex_match_tuples(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegexMatchSetRegexMatchTupleArgs']]]]):
        pulumi.set(self, "regex_match_tuples", value)


class RegexMatchSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regex_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexMatchSetRegexMatchTupleArgs']]]]] = None,
                 __props__=None):
        """
        Provides a WAF Regex Match Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_regex_pattern_set = aws.waf.RegexPatternSet("exampleRegexPatternSet", regex_pattern_strings=[
            "one",
            "two",
        ])
        example_regex_match_set = aws.waf.RegexMatchSet("exampleRegexMatchSet", regex_match_tuples=[aws.waf.RegexMatchSetRegexMatchTupleArgs(
            field_to_match=aws.waf.RegexMatchSetRegexMatchTupleFieldToMatchArgs(
                data="User-Agent",
                type="HEADER",
            ),
            regex_pattern_set_id=example_regex_pattern_set.id,
            text_transformation="NONE",
        )])
        ```

        ## Import

        terraform import {

         to = aws_waf_regex_match_set.example

         id = "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc" } Using `pulumi import`, import WAF Regex Match Set using their ID. For exampleconsole % pulumi import aws_waf_regex_match_set.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name or description of the Regex Match Set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexMatchSetRegexMatchTupleArgs']]]] regex_match_tuples: The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RegexMatchSetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF Regex Match Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_regex_pattern_set = aws.waf.RegexPatternSet("exampleRegexPatternSet", regex_pattern_strings=[
            "one",
            "two",
        ])
        example_regex_match_set = aws.waf.RegexMatchSet("exampleRegexMatchSet", regex_match_tuples=[aws.waf.RegexMatchSetRegexMatchTupleArgs(
            field_to_match=aws.waf.RegexMatchSetRegexMatchTupleFieldToMatchArgs(
                data="User-Agent",
                type="HEADER",
            ),
            regex_pattern_set_id=example_regex_pattern_set.id,
            text_transformation="NONE",
        )])
        ```

        ## Import

        terraform import {

         to = aws_waf_regex_match_set.example

         id = "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc" } Using `pulumi import`, import WAF Regex Match Set using their ID. For exampleconsole % pulumi import aws_waf_regex_match_set.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc

        :param str resource_name: The name of the resource.
        :param RegexMatchSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegexMatchSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regex_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexMatchSetRegexMatchTupleArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegexMatchSetArgs.__new__(RegexMatchSetArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["regex_match_tuples"] = regex_match_tuples
            __props__.__dict__["arn"] = None
        super(RegexMatchSet, __self__).__init__(
            'aws:waf/regexMatchSet:RegexMatchSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            regex_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexMatchSetRegexMatchTupleArgs']]]]] = None) -> 'RegexMatchSet':
        """
        Get an existing RegexMatchSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] name: The name or description of the Regex Match Set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexMatchSetRegexMatchTupleArgs']]]] regex_match_tuples: The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegexMatchSetState.__new__(_RegexMatchSetState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["name"] = name
        __props__.__dict__["regex_match_tuples"] = regex_match_tuples
        return RegexMatchSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN)
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name or description of the Regex Match Set.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="regexMatchTuples")
    def regex_match_tuples(self) -> pulumi.Output[Optional[Sequence['outputs.RegexMatchSetRegexMatchTuple']]]:
        """
        The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
        """
        return pulumi.get(self, "regex_match_tuples")

