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

__all__ = ['PipelineArgs', 'Pipeline']

@pulumi.input_type
class PipelineArgs:
    def __init__(__self__, *,
                 pipeline_display_name: pulumi.Input[str],
                 pipeline_name: pulumi.Input[str],
                 parallelism_configuration: Optional[pulumi.Input['PipelineParallelismConfigurationArgs']] = None,
                 pipeline_definition: Optional[pulumi.Input[str]] = None,
                 pipeline_definition_s3_location: Optional[pulumi.Input['PipelinePipelineDefinitionS3LocationArgs']] = None,
                 pipeline_description: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Pipeline resource.
        :param pulumi.Input[str] pipeline_display_name: The display name of the pipeline.
        :param pulumi.Input[str] pipeline_name: The name of the pipeline.
        :param pulumi.Input['PipelineParallelismConfigurationArgs'] parallelism_configuration: This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
        :param pulumi.Input[str] pipeline_definition: The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
        :param pulumi.Input['PipelinePipelineDefinitionS3LocationArgs'] pipeline_definition_s3_location: The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
        :param pulumi.Input[str] pipeline_description: A description of the pipeline.
        :param pulumi.Input[str] role_arn: The name of the Pipeline (must be unique).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "pipeline_display_name", pipeline_display_name)
        pulumi.set(__self__, "pipeline_name", pipeline_name)
        if parallelism_configuration is not None:
            pulumi.set(__self__, "parallelism_configuration", parallelism_configuration)
        if pipeline_definition is not None:
            pulumi.set(__self__, "pipeline_definition", pipeline_definition)
        if pipeline_definition_s3_location is not None:
            pulumi.set(__self__, "pipeline_definition_s3_location", pipeline_definition_s3_location)
        if pipeline_description is not None:
            pulumi.set(__self__, "pipeline_description", pipeline_description)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="pipelineDisplayName")
    def pipeline_display_name(self) -> pulumi.Input[str]:
        """
        The display name of the pipeline.
        """
        return pulumi.get(self, "pipeline_display_name")

    @pipeline_display_name.setter
    def pipeline_display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "pipeline_display_name", value)

    @property
    @pulumi.getter(name="pipelineName")
    def pipeline_name(self) -> pulumi.Input[str]:
        """
        The name of the pipeline.
        """
        return pulumi.get(self, "pipeline_name")

    @pipeline_name.setter
    def pipeline_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "pipeline_name", value)

    @property
    @pulumi.getter(name="parallelismConfiguration")
    def parallelism_configuration(self) -> Optional[pulumi.Input['PipelineParallelismConfigurationArgs']]:
        """
        This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
        """
        return pulumi.get(self, "parallelism_configuration")

    @parallelism_configuration.setter
    def parallelism_configuration(self, value: Optional[pulumi.Input['PipelineParallelismConfigurationArgs']]):
        pulumi.set(self, "parallelism_configuration", value)

    @property
    @pulumi.getter(name="pipelineDefinition")
    def pipeline_definition(self) -> Optional[pulumi.Input[str]]:
        """
        The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
        """
        return pulumi.get(self, "pipeline_definition")

    @pipeline_definition.setter
    def pipeline_definition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_definition", value)

    @property
    @pulumi.getter(name="pipelineDefinitionS3Location")
    def pipeline_definition_s3_location(self) -> Optional[pulumi.Input['PipelinePipelineDefinitionS3LocationArgs']]:
        """
        The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
        """
        return pulumi.get(self, "pipeline_definition_s3_location")

    @pipeline_definition_s3_location.setter
    def pipeline_definition_s3_location(self, value: Optional[pulumi.Input['PipelinePipelineDefinitionS3LocationArgs']]):
        pulumi.set(self, "pipeline_definition_s3_location", value)

    @property
    @pulumi.getter(name="pipelineDescription")
    def pipeline_description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the pipeline.
        """
        return pulumi.get(self, "pipeline_description")

    @pipeline_description.setter
    def pipeline_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_description", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Pipeline (must be unique).
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _PipelineState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 parallelism_configuration: Optional[pulumi.Input['PipelineParallelismConfigurationArgs']] = None,
                 pipeline_definition: Optional[pulumi.Input[str]] = None,
                 pipeline_definition_s3_location: Optional[pulumi.Input['PipelinePipelineDefinitionS3LocationArgs']] = None,
                 pipeline_description: Optional[pulumi.Input[str]] = None,
                 pipeline_display_name: Optional[pulumi.Input[str]] = None,
                 pipeline_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Pipeline resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
        :param pulumi.Input['PipelineParallelismConfigurationArgs'] parallelism_configuration: This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
        :param pulumi.Input[str] pipeline_definition: The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
        :param pulumi.Input['PipelinePipelineDefinitionS3LocationArgs'] pipeline_definition_s3_location: The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
        :param pulumi.Input[str] pipeline_description: A description of the pipeline.
        :param pulumi.Input[str] pipeline_display_name: The display name of the pipeline.
        :param pulumi.Input[str] pipeline_name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: The name of the Pipeline (must be unique).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if parallelism_configuration is not None:
            pulumi.set(__self__, "parallelism_configuration", parallelism_configuration)
        if pipeline_definition is not None:
            pulumi.set(__self__, "pipeline_definition", pipeline_definition)
        if pipeline_definition_s3_location is not None:
            pulumi.set(__self__, "pipeline_definition_s3_location", pipeline_definition_s3_location)
        if pipeline_description is not None:
            pulumi.set(__self__, "pipeline_description", pipeline_description)
        if pipeline_display_name is not None:
            pulumi.set(__self__, "pipeline_display_name", pipeline_display_name)
        if pipeline_name is not None:
            pulumi.set(__self__, "pipeline_name", pipeline_name)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="parallelismConfiguration")
    def parallelism_configuration(self) -> Optional[pulumi.Input['PipelineParallelismConfigurationArgs']]:
        """
        This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
        """
        return pulumi.get(self, "parallelism_configuration")

    @parallelism_configuration.setter
    def parallelism_configuration(self, value: Optional[pulumi.Input['PipelineParallelismConfigurationArgs']]):
        pulumi.set(self, "parallelism_configuration", value)

    @property
    @pulumi.getter(name="pipelineDefinition")
    def pipeline_definition(self) -> Optional[pulumi.Input[str]]:
        """
        The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
        """
        return pulumi.get(self, "pipeline_definition")

    @pipeline_definition.setter
    def pipeline_definition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_definition", value)

    @property
    @pulumi.getter(name="pipelineDefinitionS3Location")
    def pipeline_definition_s3_location(self) -> Optional[pulumi.Input['PipelinePipelineDefinitionS3LocationArgs']]:
        """
        The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
        """
        return pulumi.get(self, "pipeline_definition_s3_location")

    @pipeline_definition_s3_location.setter
    def pipeline_definition_s3_location(self, value: Optional[pulumi.Input['PipelinePipelineDefinitionS3LocationArgs']]):
        pulumi.set(self, "pipeline_definition_s3_location", value)

    @property
    @pulumi.getter(name="pipelineDescription")
    def pipeline_description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the pipeline.
        """
        return pulumi.get(self, "pipeline_description")

    @pipeline_description.setter
    def pipeline_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_description", value)

    @property
    @pulumi.getter(name="pipelineDisplayName")
    def pipeline_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the pipeline.
        """
        return pulumi.get(self, "pipeline_display_name")

    @pipeline_display_name.setter
    def pipeline_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_display_name", value)

    @property
    @pulumi.getter(name="pipelineName")
    def pipeline_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the pipeline.
        """
        return pulumi.get(self, "pipeline_name")

    @pipeline_name.setter
    def pipeline_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Pipeline (must be unique).
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class Pipeline(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parallelism_configuration: Optional[pulumi.Input[pulumi.InputType['PipelineParallelismConfigurationArgs']]] = None,
                 pipeline_definition: Optional[pulumi.Input[str]] = None,
                 pipeline_definition_s3_location: Optional[pulumi.Input[pulumi.InputType['PipelinePipelineDefinitionS3LocationArgs']]] = None,
                 pipeline_description: Optional[pulumi.Input[str]] = None,
                 pipeline_display_name: Optional[pulumi.Input[str]] = None,
                 pipeline_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a SageMaker Pipeline resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        example = aws.sagemaker.Pipeline("example",
            pipeline_name="example",
            pipeline_display_name="example",
            role_arn=aws_iam_role["example"]["arn"],
            pipeline_definition=json.dumps({
                "Version": "2020-12-01",
                "Steps": [{
                    "Name": "Test",
                    "Type": "Fail",
                    "Arguments": {
                        "ErrorMessage": "test",
                    },
                }],
            }))
        ```

        ## Import

        You can use `pulumi import` to import pipelines using `pipeline_name`. For exampleterraform import {

         to = aws_sagemaker_pipeline.test_pipeline

         id = "pipeline" } Using `pulumi import`, import pipelines using the `pipeline_name`. For exampleconsole % pulumi import aws_sagemaker_pipeline.test_pipeline pipeline

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PipelineParallelismConfigurationArgs']] parallelism_configuration: This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
        :param pulumi.Input[str] pipeline_definition: The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
        :param pulumi.Input[pulumi.InputType['PipelinePipelineDefinitionS3LocationArgs']] pipeline_definition_s3_location: The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
        :param pulumi.Input[str] pipeline_description: A description of the pipeline.
        :param pulumi.Input[str] pipeline_display_name: The display name of the pipeline.
        :param pulumi.Input[str] pipeline_name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: The name of the Pipeline (must be unique).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a SageMaker Pipeline resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        example = aws.sagemaker.Pipeline("example",
            pipeline_name="example",
            pipeline_display_name="example",
            role_arn=aws_iam_role["example"]["arn"],
            pipeline_definition=json.dumps({
                "Version": "2020-12-01",
                "Steps": [{
                    "Name": "Test",
                    "Type": "Fail",
                    "Arguments": {
                        "ErrorMessage": "test",
                    },
                }],
            }))
        ```

        ## Import

        You can use `pulumi import` to import pipelines using `pipeline_name`. For exampleterraform import {

         to = aws_sagemaker_pipeline.test_pipeline

         id = "pipeline" } Using `pulumi import`, import pipelines using the `pipeline_name`. For exampleconsole % pulumi import aws_sagemaker_pipeline.test_pipeline pipeline

        :param str resource_name: The name of the resource.
        :param PipelineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parallelism_configuration: Optional[pulumi.Input[pulumi.InputType['PipelineParallelismConfigurationArgs']]] = None,
                 pipeline_definition: Optional[pulumi.Input[str]] = None,
                 pipeline_definition_s3_location: Optional[pulumi.Input[pulumi.InputType['PipelinePipelineDefinitionS3LocationArgs']]] = None,
                 pipeline_description: Optional[pulumi.Input[str]] = None,
                 pipeline_display_name: Optional[pulumi.Input[str]] = None,
                 pipeline_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipelineArgs.__new__(PipelineArgs)

            __props__.__dict__["parallelism_configuration"] = parallelism_configuration
            __props__.__dict__["pipeline_definition"] = pipeline_definition
            __props__.__dict__["pipeline_definition_s3_location"] = pipeline_definition_s3_location
            __props__.__dict__["pipeline_description"] = pipeline_description
            if pipeline_display_name is None and not opts.urn:
                raise TypeError("Missing required property 'pipeline_display_name'")
            __props__.__dict__["pipeline_display_name"] = pipeline_display_name
            if pipeline_name is None and not opts.urn:
                raise TypeError("Missing required property 'pipeline_name'")
            __props__.__dict__["pipeline_name"] = pipeline_name
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(Pipeline, __self__).__init__(
            'aws:sagemaker/pipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            parallelism_configuration: Optional[pulumi.Input[pulumi.InputType['PipelineParallelismConfigurationArgs']]] = None,
            pipeline_definition: Optional[pulumi.Input[str]] = None,
            pipeline_definition_s3_location: Optional[pulumi.Input[pulumi.InputType['PipelinePipelineDefinitionS3LocationArgs']]] = None,
            pipeline_description: Optional[pulumi.Input[str]] = None,
            pipeline_display_name: Optional[pulumi.Input[str]] = None,
            pipeline_name: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineParallelismConfigurationArgs']] parallelism_configuration: This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
        :param pulumi.Input[str] pipeline_definition: The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
        :param pulumi.Input[pulumi.InputType['PipelinePipelineDefinitionS3LocationArgs']] pipeline_definition_s3_location: The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
        :param pulumi.Input[str] pipeline_description: A description of the pipeline.
        :param pulumi.Input[str] pipeline_display_name: The display name of the pipeline.
        :param pulumi.Input[str] pipeline_name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: The name of the Pipeline (must be unique).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PipelineState.__new__(_PipelineState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["parallelism_configuration"] = parallelism_configuration
        __props__.__dict__["pipeline_definition"] = pipeline_definition
        __props__.__dict__["pipeline_definition_s3_location"] = pipeline_definition_s3_location
        __props__.__dict__["pipeline_description"] = pipeline_description
        __props__.__dict__["pipeline_display_name"] = pipeline_display_name
        __props__.__dict__["pipeline_name"] = pipeline_name
        __props__.__dict__["role_arn"] = role_arn
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="parallelismConfiguration")
    def parallelism_configuration(self) -> pulumi.Output[Optional['outputs.PipelineParallelismConfiguration']]:
        """
        This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
        """
        return pulumi.get(self, "parallelism_configuration")

    @property
    @pulumi.getter(name="pipelineDefinition")
    def pipeline_definition(self) -> pulumi.Output[Optional[str]]:
        """
        The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
        """
        return pulumi.get(self, "pipeline_definition")

    @property
    @pulumi.getter(name="pipelineDefinitionS3Location")
    def pipeline_definition_s3_location(self) -> pulumi.Output[Optional['outputs.PipelinePipelineDefinitionS3Location']]:
        """
        The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
        """
        return pulumi.get(self, "pipeline_definition_s3_location")

    @property
    @pulumi.getter(name="pipelineDescription")
    def pipeline_description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the pipeline.
        """
        return pulumi.get(self, "pipeline_description")

    @property
    @pulumi.getter(name="pipelineDisplayName")
    def pipeline_display_name(self) -> pulumi.Output[str]:
        """
        The display name of the pipeline.
        """
        return pulumi.get(self, "pipeline_display_name")

    @property
    @pulumi.getter(name="pipelineName")
    def pipeline_name(self) -> pulumi.Output[str]:
        """
        The name of the pipeline.
        """
        return pulumi.get(self, "pipeline_name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the Pipeline (must be unique).
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

