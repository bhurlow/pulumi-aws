# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VocabularyFilterArgs', 'VocabularyFilter']

@pulumi.input_type
class VocabularyFilterArgs:
    def __init__(__self__, *,
                 language_code: pulumi.Input[str],
                 vocabulary_filter_name: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vocabulary_filter_file_uri: Optional[pulumi.Input[str]] = None,
                 words: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VocabularyFilter resource.
        :param pulumi.Input[str] language_code: The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
        :param pulumi.Input[str] vocabulary_filter_name: The name of the VocabularyFilter.
               
               The following arguments are optional:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] vocabulary_filter_file_uri: The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] words: A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
        """
        pulumi.set(__self__, "language_code", language_code)
        pulumi.set(__self__, "vocabulary_filter_name", vocabulary_filter_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vocabulary_filter_file_uri is not None:
            pulumi.set(__self__, "vocabulary_filter_file_uri", vocabulary_filter_file_uri)
        if words is not None:
            pulumi.set(__self__, "words", words)

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> pulumi.Input[str]:
        """
        The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
        """
        return pulumi.get(self, "language_code")

    @language_code.setter
    def language_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "language_code", value)

    @property
    @pulumi.getter(name="vocabularyFilterName")
    def vocabulary_filter_name(self) -> pulumi.Input[str]:
        """
        The name of the VocabularyFilter.

        The following arguments are optional:
        """
        return pulumi.get(self, "vocabulary_filter_name")

    @vocabulary_filter_name.setter
    def vocabulary_filter_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vocabulary_filter_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vocabularyFilterFileUri")
    def vocabulary_filter_file_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
        """
        return pulumi.get(self, "vocabulary_filter_file_uri")

    @vocabulary_filter_file_uri.setter
    def vocabulary_filter_file_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vocabulary_filter_file_uri", value)

    @property
    @pulumi.getter
    def words(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
        """
        return pulumi.get(self, "words")

    @words.setter
    def words(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "words", value)


@pulumi.input_type
class _VocabularyFilterState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 download_uri: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vocabulary_filter_file_uri: Optional[pulumi.Input[str]] = None,
                 vocabulary_filter_name: Optional[pulumi.Input[str]] = None,
                 words: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering VocabularyFilter resources.
        :param pulumi.Input[str] arn: ARN of the VocabularyFilter.
        :param pulumi.Input[str] download_uri: Generated download URI.
        :param pulumi.Input[str] language_code: The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] vocabulary_filter_file_uri: The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
        :param pulumi.Input[str] vocabulary_filter_name: The name of the VocabularyFilter.
               
               The following arguments are optional:
        :param pulumi.Input[Sequence[pulumi.Input[str]]] words: A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if download_uri is not None:
            pulumi.set(__self__, "download_uri", download_uri)
        if language_code is not None:
            pulumi.set(__self__, "language_code", language_code)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if vocabulary_filter_file_uri is not None:
            pulumi.set(__self__, "vocabulary_filter_file_uri", vocabulary_filter_file_uri)
        if vocabulary_filter_name is not None:
            pulumi.set(__self__, "vocabulary_filter_name", vocabulary_filter_name)
        if words is not None:
            pulumi.set(__self__, "words", words)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the VocabularyFilter.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="downloadUri")
    def download_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Generated download URI.
        """
        return pulumi.get(self, "download_uri")

    @download_uri.setter
    def download_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "download_uri", value)

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> Optional[pulumi.Input[str]]:
        """
        The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
        """
        return pulumi.get(self, "language_code")

    @language_code.setter
    def language_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language_code", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="vocabularyFilterFileUri")
    def vocabulary_filter_file_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
        """
        return pulumi.get(self, "vocabulary_filter_file_uri")

    @vocabulary_filter_file_uri.setter
    def vocabulary_filter_file_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vocabulary_filter_file_uri", value)

    @property
    @pulumi.getter(name="vocabularyFilterName")
    def vocabulary_filter_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VocabularyFilter.

        The following arguments are optional:
        """
        return pulumi.get(self, "vocabulary_filter_name")

    @vocabulary_filter_name.setter
    def vocabulary_filter_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vocabulary_filter_name", value)

    @property
    @pulumi.getter
    def words(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
        """
        return pulumi.get(self, "words")

    @words.setter
    def words(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "words", value)


class VocabularyFilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vocabulary_filter_file_uri: Optional[pulumi.Input[str]] = None,
                 vocabulary_filter_name: Optional[pulumi.Input[str]] = None,
                 words: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource for managing an AWS Transcribe VocabularyFilter.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.transcribe.VocabularyFilter("example",
            language_code="en-US",
            tags={
                "tag1": "value1",
                "tag2": "value3",
            },
            vocabulary_filter_name="example",
            words=[
                "cars",
                "bucket",
            ])
        ```

        ## Import

        Using `pulumi import`, import Transcribe VocabularyFilter using the `vocabulary_filter_name`. For example:

        ```sh
         $ pulumi import aws:transcribe/vocabularyFilter:VocabularyFilter example example-name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] language_code: The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] vocabulary_filter_file_uri: The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
        :param pulumi.Input[str] vocabulary_filter_name: The name of the VocabularyFilter.
               
               The following arguments are optional:
        :param pulumi.Input[Sequence[pulumi.Input[str]]] words: A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VocabularyFilterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing an AWS Transcribe VocabularyFilter.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.transcribe.VocabularyFilter("example",
            language_code="en-US",
            tags={
                "tag1": "value1",
                "tag2": "value3",
            },
            vocabulary_filter_name="example",
            words=[
                "cars",
                "bucket",
            ])
        ```

        ## Import

        Using `pulumi import`, import Transcribe VocabularyFilter using the `vocabulary_filter_name`. For example:

        ```sh
         $ pulumi import aws:transcribe/vocabularyFilter:VocabularyFilter example example-name
        ```

        :param str resource_name: The name of the resource.
        :param VocabularyFilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VocabularyFilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vocabulary_filter_file_uri: Optional[pulumi.Input[str]] = None,
                 vocabulary_filter_name: Optional[pulumi.Input[str]] = None,
                 words: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VocabularyFilterArgs.__new__(VocabularyFilterArgs)

            if language_code is None and not opts.urn:
                raise TypeError("Missing required property 'language_code'")
            __props__.__dict__["language_code"] = language_code
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vocabulary_filter_file_uri"] = vocabulary_filter_file_uri
            if vocabulary_filter_name is None and not opts.urn:
                raise TypeError("Missing required property 'vocabulary_filter_name'")
            __props__.__dict__["vocabulary_filter_name"] = vocabulary_filter_name
            __props__.__dict__["words"] = words
            __props__.__dict__["arn"] = None
            __props__.__dict__["download_uri"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(VocabularyFilter, __self__).__init__(
            'aws:transcribe/vocabularyFilter:VocabularyFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            download_uri: Optional[pulumi.Input[str]] = None,
            language_code: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vocabulary_filter_file_uri: Optional[pulumi.Input[str]] = None,
            vocabulary_filter_name: Optional[pulumi.Input[str]] = None,
            words: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'VocabularyFilter':
        """
        Get an existing VocabularyFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the VocabularyFilter.
        :param pulumi.Input[str] download_uri: Generated download URI.
        :param pulumi.Input[str] language_code: The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] vocabulary_filter_file_uri: The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
        :param pulumi.Input[str] vocabulary_filter_name: The name of the VocabularyFilter.
               
               The following arguments are optional:
        :param pulumi.Input[Sequence[pulumi.Input[str]]] words: A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VocabularyFilterState.__new__(_VocabularyFilterState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["download_uri"] = download_uri
        __props__.__dict__["language_code"] = language_code
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["vocabulary_filter_file_uri"] = vocabulary_filter_file_uri
        __props__.__dict__["vocabulary_filter_name"] = vocabulary_filter_name
        __props__.__dict__["words"] = words
        return VocabularyFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the VocabularyFilter.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="downloadUri")
    def download_uri(self) -> pulumi.Output[str]:
        """
        Generated download URI.
        """
        return pulumi.get(self, "download_uri")

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> pulumi.Output[str]:
        """
        The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
        """
        return pulumi.get(self, "language_code")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="vocabularyFilterFileUri")
    def vocabulary_filter_file_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
        """
        return pulumi.get(self, "vocabulary_filter_file_uri")

    @property
    @pulumi.getter(name="vocabularyFilterName")
    def vocabulary_filter_name(self) -> pulumi.Output[str]:
        """
        The name of the VocabularyFilter.

        The following arguments are optional:
        """
        return pulumi.get(self, "vocabulary_filter_name")

    @property
    @pulumi.getter
    def words(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
        """
        return pulumi.get(self, "words")

