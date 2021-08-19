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

__all__ = ['PipelineArgs', 'Pipeline']

@pulumi.input_type
class PipelineArgs:
    def __init__(__self__, *,
                 artifact_store: pulumi.Input['PipelineArtifactStoreArgs'],
                 role_arn: pulumi.Input[str],
                 stages: pulumi.Input[Sequence[pulumi.Input['PipelineStageArgs']]],
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Pipeline resource.
        :param pulumi.Input['PipelineArtifactStoreArgs'] artifact_store: One or more artifact_store blocks. Artifact stores are documented below.
        :param pulumi.Input[str] role_arn: A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        :param pulumi.Input[Sequence[pulumi.Input['PipelineStageArgs']]] stages: A stage block. Stages are documented below.
        :param pulumi.Input[str] name: The name of the pipeline.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "artifact_store", artifact_store)
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "stages", stages)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="artifactStore")
    def artifact_store(self) -> pulumi.Input['PipelineArtifactStoreArgs']:
        """
        One or more artifact_store blocks. Artifact stores are documented below.
        """
        return pulumi.get(self, "artifact_store")

    @artifact_store.setter
    def artifact_store(self, value: pulumi.Input['PipelineArtifactStoreArgs']):
        pulumi.set(self, "artifact_store", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def stages(self) -> pulumi.Input[Sequence[pulumi.Input['PipelineStageArgs']]]:
        """
        A stage block. Stages are documented below.
        """
        return pulumi.get(self, "stages")

    @stages.setter
    def stages(self, value: pulumi.Input[Sequence[pulumi.Input['PipelineStageArgs']]]):
        pulumi.set(self, "stages", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the pipeline.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _PipelineState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 artifact_store: Optional[pulumi.Input['PipelineArtifactStoreArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stages: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineStageArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Pipeline resources.
        :param pulumi.Input[str] arn: The codepipeline ARN.
        :param pulumi.Input['PipelineArtifactStoreArgs'] artifact_store: One or more artifact_store blocks. Artifact stores are documented below.
        :param pulumi.Input[str] name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        :param pulumi.Input[Sequence[pulumi.Input['PipelineStageArgs']]] stages: A stage block. Stages are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if artifact_store is not None:
            pulumi.set(__self__, "artifact_store", artifact_store)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if stages is not None:
            pulumi.set(__self__, "stages", stages)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The codepipeline ARN.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="artifactStore")
    def artifact_store(self) -> Optional[pulumi.Input['PipelineArtifactStoreArgs']]:
        """
        One or more artifact_store blocks. Artifact stores are documented below.
        """
        return pulumi.get(self, "artifact_store")

    @artifact_store.setter
    def artifact_store(self, value: Optional[pulumi.Input['PipelineArtifactStoreArgs']]):
        pulumi.set(self, "artifact_store", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the pipeline.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def stages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineStageArgs']]]]:
        """
        A stage block. Stages are documented below.
        """
        return pulumi.get(self, "stages")

    @stages.setter
    def stages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineStageArgs']]]]):
        pulumi.set(self, "stages", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class Pipeline(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 artifact_store: Optional[pulumi.Input[pulumi.InputType['PipelineArtifactStoreArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineStageArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a CodePipeline.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.codestarconnections.Connection("example", provider_type="GitHub")
        codepipeline_bucket = aws.s3.Bucket("codepipelineBucket", acl="private")
        codepipeline_role = aws.iam.Role("codepipelineRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": "codepipeline.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        }
        \"\"\")
        s3kmskey = aws.kms.get_alias(name="alias/myKmsKey")
        codepipeline = aws.codepipeline.Pipeline("codepipeline",
            role_arn=codepipeline_role.arn,
            artifact_store=aws.codepipeline.PipelineArtifactStoreArgs(
                location=codepipeline_bucket.bucket,
                type="S3",
                encryption_key=aws.codepipeline.PipelineArtifactStoreEncryptionKeyArgs(
                    id=s3kmskey.arn,
                    type="KMS",
                ),
            ),
            stages=[
                aws.codepipeline.PipelineStageArgs(
                    name="Source",
                    actions=[aws.codepipeline.PipelineStageActionArgs(
                        name="Source",
                        category="Source",
                        owner="AWS",
                        provider="CodeStarSourceConnection",
                        version="1",
                        output_artifacts=["source_output"],
                        configuration={
                            "ConnectionArn": example.arn,
                            "FullRepositoryId": "my-organization/example",
                            "BranchName": "main",
                        },
                    )],
                ),
                aws.codepipeline.PipelineStageArgs(
                    name="Build",
                    actions=[aws.codepipeline.PipelineStageActionArgs(
                        name="Build",
                        category="Build",
                        owner="AWS",
                        provider="CodeBuild",
                        input_artifacts=["source_output"],
                        output_artifacts=["build_output"],
                        version="1",
                        configuration={
                            "ProjectName": "test",
                        },
                    )],
                ),
                aws.codepipeline.PipelineStageArgs(
                    name="Deploy",
                    actions=[aws.codepipeline.PipelineStageActionArgs(
                        name="Deploy",
                        category="Deploy",
                        owner="AWS",
                        provider="CloudFormation",
                        input_artifacts=["build_output"],
                        version="1",
                        configuration={
                            "ActionMode": "REPLACE_ON_FAILURE",
                            "Capabilities": "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM",
                            "OutputFileName": "CreateStackOutput.json",
                            "StackName": "MyStack",
                            "TemplatePath": "build_output::sam-templated.yaml",
                        },
                    )],
                ),
            ])
        codepipeline_policy = aws.iam.RolePolicy("codepipelinePolicy",
            role=codepipeline_role.id,
            policy=pulumi.Output.all(codepipeline_bucket.arn, codepipeline_bucket.arn, example.arn).apply(lambda codepipelineBucketArn, codepipelineBucketArn1, exampleArn: f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
              "Effect":"Allow",
              "Action": [
                "s3:GetObject",
                "s3:GetObjectVersion",
                "s3:GetBucketVersioning",
                "s3:PutObjectAcl",
                "s3:PutObject"
              ],
              "Resource": [
                "{codepipeline_bucket_arn}",
                "{codepipeline_bucket_arn1}/*"
              ]
            }},
            {{
              "Effect": "Allow",
              "Action": [
                "codestar-connections:UseConnection"
              ],
              "Resource": "{example_arn}"
            }},
            {{
              "Effect": "Allow",
              "Action": [
                "codebuild:BatchGetBuilds",
                "codebuild:StartBuild"
              ],
              "Resource": "*"
            }}
          ]
        }}
        \"\"\"))
        ```

        ## Import

        CodePipelines can be imported using the name, e.g.

        ```sh
         $ pulumi import aws:codepipeline/pipeline:Pipeline foo example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PipelineArtifactStoreArgs']] artifact_store: One or more artifact_store blocks. Artifact stores are documented below.
        :param pulumi.Input[str] name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineStageArgs']]]] stages: A stage block. Stages are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CodePipeline.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.codestarconnections.Connection("example", provider_type="GitHub")
        codepipeline_bucket = aws.s3.Bucket("codepipelineBucket", acl="private")
        codepipeline_role = aws.iam.Role("codepipelineRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": "codepipeline.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        }
        \"\"\")
        s3kmskey = aws.kms.get_alias(name="alias/myKmsKey")
        codepipeline = aws.codepipeline.Pipeline("codepipeline",
            role_arn=codepipeline_role.arn,
            artifact_store=aws.codepipeline.PipelineArtifactStoreArgs(
                location=codepipeline_bucket.bucket,
                type="S3",
                encryption_key=aws.codepipeline.PipelineArtifactStoreEncryptionKeyArgs(
                    id=s3kmskey.arn,
                    type="KMS",
                ),
            ),
            stages=[
                aws.codepipeline.PipelineStageArgs(
                    name="Source",
                    actions=[aws.codepipeline.PipelineStageActionArgs(
                        name="Source",
                        category="Source",
                        owner="AWS",
                        provider="CodeStarSourceConnection",
                        version="1",
                        output_artifacts=["source_output"],
                        configuration={
                            "ConnectionArn": example.arn,
                            "FullRepositoryId": "my-organization/example",
                            "BranchName": "main",
                        },
                    )],
                ),
                aws.codepipeline.PipelineStageArgs(
                    name="Build",
                    actions=[aws.codepipeline.PipelineStageActionArgs(
                        name="Build",
                        category="Build",
                        owner="AWS",
                        provider="CodeBuild",
                        input_artifacts=["source_output"],
                        output_artifacts=["build_output"],
                        version="1",
                        configuration={
                            "ProjectName": "test",
                        },
                    )],
                ),
                aws.codepipeline.PipelineStageArgs(
                    name="Deploy",
                    actions=[aws.codepipeline.PipelineStageActionArgs(
                        name="Deploy",
                        category="Deploy",
                        owner="AWS",
                        provider="CloudFormation",
                        input_artifacts=["build_output"],
                        version="1",
                        configuration={
                            "ActionMode": "REPLACE_ON_FAILURE",
                            "Capabilities": "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM",
                            "OutputFileName": "CreateStackOutput.json",
                            "StackName": "MyStack",
                            "TemplatePath": "build_output::sam-templated.yaml",
                        },
                    )],
                ),
            ])
        codepipeline_policy = aws.iam.RolePolicy("codepipelinePolicy",
            role=codepipeline_role.id,
            policy=pulumi.Output.all(codepipeline_bucket.arn, codepipeline_bucket.arn, example.arn).apply(lambda codepipelineBucketArn, codepipelineBucketArn1, exampleArn: f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
              "Effect":"Allow",
              "Action": [
                "s3:GetObject",
                "s3:GetObjectVersion",
                "s3:GetBucketVersioning",
                "s3:PutObjectAcl",
                "s3:PutObject"
              ],
              "Resource": [
                "{codepipeline_bucket_arn}",
                "{codepipeline_bucket_arn1}/*"
              ]
            }},
            {{
              "Effect": "Allow",
              "Action": [
                "codestar-connections:UseConnection"
              ],
              "Resource": "{example_arn}"
            }},
            {{
              "Effect": "Allow",
              "Action": [
                "codebuild:BatchGetBuilds",
                "codebuild:StartBuild"
              ],
              "Resource": "*"
            }}
          ]
        }}
        \"\"\"))
        ```

        ## Import

        CodePipelines can be imported using the name, e.g.

        ```sh
         $ pulumi import aws:codepipeline/pipeline:Pipeline foo example
        ```

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
                 artifact_store: Optional[pulumi.Input[pulumi.InputType['PipelineArtifactStoreArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineStageArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = PipelineArgs.__new__(PipelineArgs)

            if artifact_store is None and not opts.urn:
                raise TypeError("Missing required property 'artifact_store'")
            __props__.__dict__["artifact_store"] = artifact_store
            __props__.__dict__["name"] = name
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            if stages is None and not opts.urn:
                raise TypeError("Missing required property 'stages'")
            __props__.__dict__["stages"] = stages
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(Pipeline, __self__).__init__(
            'aws:codepipeline/pipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            artifact_store: Optional[pulumi.Input[pulumi.InputType['PipelineArtifactStoreArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            stages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineStageArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The codepipeline ARN.
        :param pulumi.Input[pulumi.InputType['PipelineArtifactStoreArgs']] artifact_store: One or more artifact_store blocks. Artifact stores are documented below.
        :param pulumi.Input[str] name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineStageArgs']]]] stages: A stage block. Stages are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PipelineState.__new__(_PipelineState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["artifact_store"] = artifact_store
        __props__.__dict__["name"] = name
        __props__.__dict__["role_arn"] = role_arn
        __props__.__dict__["stages"] = stages
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The codepipeline ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="artifactStore")
    def artifact_store(self) -> pulumi.Output['outputs.PipelineArtifactStore']:
        """
        One or more artifact_store blocks. Artifact stores are documented below.
        """
        return pulumi.get(self, "artifact_store")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the pipeline.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def stages(self) -> pulumi.Output[Sequence['outputs.PipelineStage']]:
        """
        A stage block. Stages are documented below.
        """
        return pulumi.get(self, "stages")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

