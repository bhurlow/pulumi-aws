# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReportDefinitionArgs', 'ReportDefinition']

@pulumi.input_type
class ReportDefinitionArgs:
    def __init__(__self__, *,
                 additional_schema_elements: pulumi.Input[Sequence[pulumi.Input[str]]],
                 compression: pulumi.Input[str],
                 format: pulumi.Input[str],
                 report_name: pulumi.Input[str],
                 s3_bucket: pulumi.Input[str],
                 s3_region: pulumi.Input[str],
                 time_unit: pulumi.Input[str],
                 additional_artifacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_closed_reports: Optional[pulumi.Input[bool]] = None,
                 report_versioning: Optional[pulumi.Input[str]] = None,
                 s3_prefix: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReportDefinition resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] additional_schema_elements: A list of schema elements. Valid values are: `RESOURCES`.
        :param pulumi.Input[str] compression: Compression format for report. Valid values are: `GZIP`, `ZIP`, `Parquet`. If `Parquet` is used, then format must also be `Parquet`.
        :param pulumi.Input[str] format: Format for report. Valid values are: `textORcsv`, `Parquet`. If `Parquet` is used, then Compression must also be `Parquet`.
        :param pulumi.Input[str] report_name: Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
        :param pulumi.Input[str] s3_bucket: Name of the existing S3 bucket to hold generated reports.
        :param pulumi.Input[str] s3_region: Region of the existing S3 bucket to hold generated reports.
        :param pulumi.Input[str] time_unit: The frequency on which report data are measured and displayed.  Valid values are: `HOURLY`, `DAILY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] additional_artifacts: A list of additional artifacts. Valid values are: `REDSHIFT`, `QUICKSIGHT`, `ATHENA`. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be `OVERWRITE_REPORT`.
        :param pulumi.Input[bool] refresh_closed_reports: Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
        :param pulumi.Input[str] report_versioning: Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: `CREATE_NEW_REPORT` and `OVERWRITE_REPORT`.
        :param pulumi.Input[str] s3_prefix: Report path prefix. Limited to 256 characters.
        """
        pulumi.set(__self__, "additional_schema_elements", additional_schema_elements)
        pulumi.set(__self__, "compression", compression)
        pulumi.set(__self__, "format", format)
        pulumi.set(__self__, "report_name", report_name)
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        pulumi.set(__self__, "s3_region", s3_region)
        pulumi.set(__self__, "time_unit", time_unit)
        if additional_artifacts is not None:
            pulumi.set(__self__, "additional_artifacts", additional_artifacts)
        if refresh_closed_reports is not None:
            pulumi.set(__self__, "refresh_closed_reports", refresh_closed_reports)
        if report_versioning is not None:
            pulumi.set(__self__, "report_versioning", report_versioning)
        if s3_prefix is not None:
            pulumi.set(__self__, "s3_prefix", s3_prefix)

    @property
    @pulumi.getter(name="additionalSchemaElements")
    def additional_schema_elements(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of schema elements. Valid values are: `RESOURCES`.
        """
        return pulumi.get(self, "additional_schema_elements")

    @additional_schema_elements.setter
    def additional_schema_elements(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "additional_schema_elements", value)

    @property
    @pulumi.getter
    def compression(self) -> pulumi.Input[str]:
        """
        Compression format for report. Valid values are: `GZIP`, `ZIP`, `Parquet`. If `Parquet` is used, then format must also be `Parquet`.
        """
        return pulumi.get(self, "compression")

    @compression.setter
    def compression(self, value: pulumi.Input[str]):
        pulumi.set(self, "compression", value)

    @property
    @pulumi.getter
    def format(self) -> pulumi.Input[str]:
        """
        Format for report. Valid values are: `textORcsv`, `Parquet`. If `Parquet` is used, then Compression must also be `Parquet`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: pulumi.Input[str]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="reportName")
    def report_name(self) -> pulumi.Input[str]:
        """
        Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
        """
        return pulumi.get(self, "report_name")

    @report_name.setter
    def report_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "report_name", value)

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> pulumi.Input[str]:
        """
        Name of the existing S3 bucket to hold generated reports.
        """
        return pulumi.get(self, "s3_bucket")

    @s3_bucket.setter
    def s3_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_bucket", value)

    @property
    @pulumi.getter(name="s3Region")
    def s3_region(self) -> pulumi.Input[str]:
        """
        Region of the existing S3 bucket to hold generated reports.
        """
        return pulumi.get(self, "s3_region")

    @s3_region.setter
    def s3_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_region", value)

    @property
    @pulumi.getter(name="timeUnit")
    def time_unit(self) -> pulumi.Input[str]:
        """
        The frequency on which report data are measured and displayed.  Valid values are: `HOURLY`, `DAILY`.
        """
        return pulumi.get(self, "time_unit")

    @time_unit.setter
    def time_unit(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_unit", value)

    @property
    @pulumi.getter(name="additionalArtifacts")
    def additional_artifacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of additional artifacts. Valid values are: `REDSHIFT`, `QUICKSIGHT`, `ATHENA`. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be `OVERWRITE_REPORT`.
        """
        return pulumi.get(self, "additional_artifacts")

    @additional_artifacts.setter
    def additional_artifacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "additional_artifacts", value)

    @property
    @pulumi.getter(name="refreshClosedReports")
    def refresh_closed_reports(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
        """
        return pulumi.get(self, "refresh_closed_reports")

    @refresh_closed_reports.setter
    def refresh_closed_reports(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "refresh_closed_reports", value)

    @property
    @pulumi.getter(name="reportVersioning")
    def report_versioning(self) -> Optional[pulumi.Input[str]]:
        """
        Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: `CREATE_NEW_REPORT` and `OVERWRITE_REPORT`.
        """
        return pulumi.get(self, "report_versioning")

    @report_versioning.setter
    def report_versioning(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "report_versioning", value)

    @property
    @pulumi.getter(name="s3Prefix")
    def s3_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Report path prefix. Limited to 256 characters.
        """
        return pulumi.get(self, "s3_prefix")

    @s3_prefix.setter
    def s3_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_prefix", value)


@pulumi.input_type
class _ReportDefinitionState:
    def __init__(__self__, *,
                 additional_artifacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 additional_schema_elements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 compression: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 refresh_closed_reports: Optional[pulumi.Input[bool]] = None,
                 report_name: Optional[pulumi.Input[str]] = None,
                 report_versioning: Optional[pulumi.Input[str]] = None,
                 s3_bucket: Optional[pulumi.Input[str]] = None,
                 s3_prefix: Optional[pulumi.Input[str]] = None,
                 s3_region: Optional[pulumi.Input[str]] = None,
                 time_unit: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReportDefinition resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] additional_artifacts: A list of additional artifacts. Valid values are: `REDSHIFT`, `QUICKSIGHT`, `ATHENA`. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be `OVERWRITE_REPORT`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] additional_schema_elements: A list of schema elements. Valid values are: `RESOURCES`.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the cur report.
        :param pulumi.Input[str] compression: Compression format for report. Valid values are: `GZIP`, `ZIP`, `Parquet`. If `Parquet` is used, then format must also be `Parquet`.
        :param pulumi.Input[str] format: Format for report. Valid values are: `textORcsv`, `Parquet`. If `Parquet` is used, then Compression must also be `Parquet`.
        :param pulumi.Input[bool] refresh_closed_reports: Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
        :param pulumi.Input[str] report_name: Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
        :param pulumi.Input[str] report_versioning: Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: `CREATE_NEW_REPORT` and `OVERWRITE_REPORT`.
        :param pulumi.Input[str] s3_bucket: Name of the existing S3 bucket to hold generated reports.
        :param pulumi.Input[str] s3_prefix: Report path prefix. Limited to 256 characters.
        :param pulumi.Input[str] s3_region: Region of the existing S3 bucket to hold generated reports.
        :param pulumi.Input[str] time_unit: The frequency on which report data are measured and displayed.  Valid values are: `HOURLY`, `DAILY`.
        """
        if additional_artifacts is not None:
            pulumi.set(__self__, "additional_artifacts", additional_artifacts)
        if additional_schema_elements is not None:
            pulumi.set(__self__, "additional_schema_elements", additional_schema_elements)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if compression is not None:
            pulumi.set(__self__, "compression", compression)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if refresh_closed_reports is not None:
            pulumi.set(__self__, "refresh_closed_reports", refresh_closed_reports)
        if report_name is not None:
            pulumi.set(__self__, "report_name", report_name)
        if report_versioning is not None:
            pulumi.set(__self__, "report_versioning", report_versioning)
        if s3_bucket is not None:
            pulumi.set(__self__, "s3_bucket", s3_bucket)
        if s3_prefix is not None:
            pulumi.set(__self__, "s3_prefix", s3_prefix)
        if s3_region is not None:
            pulumi.set(__self__, "s3_region", s3_region)
        if time_unit is not None:
            pulumi.set(__self__, "time_unit", time_unit)

    @property
    @pulumi.getter(name="additionalArtifacts")
    def additional_artifacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of additional artifacts. Valid values are: `REDSHIFT`, `QUICKSIGHT`, `ATHENA`. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be `OVERWRITE_REPORT`.
        """
        return pulumi.get(self, "additional_artifacts")

    @additional_artifacts.setter
    def additional_artifacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "additional_artifacts", value)

    @property
    @pulumi.getter(name="additionalSchemaElements")
    def additional_schema_elements(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of schema elements. Valid values are: `RESOURCES`.
        """
        return pulumi.get(self, "additional_schema_elements")

    @additional_schema_elements.setter
    def additional_schema_elements(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "additional_schema_elements", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) specifying the cur report.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def compression(self) -> Optional[pulumi.Input[str]]:
        """
        Compression format for report. Valid values are: `GZIP`, `ZIP`, `Parquet`. If `Parquet` is used, then format must also be `Parquet`.
        """
        return pulumi.get(self, "compression")

    @compression.setter
    def compression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compression", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Format for report. Valid values are: `textORcsv`, `Parquet`. If `Parquet` is used, then Compression must also be `Parquet`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="refreshClosedReports")
    def refresh_closed_reports(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
        """
        return pulumi.get(self, "refresh_closed_reports")

    @refresh_closed_reports.setter
    def refresh_closed_reports(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "refresh_closed_reports", value)

    @property
    @pulumi.getter(name="reportName")
    def report_name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
        """
        return pulumi.get(self, "report_name")

    @report_name.setter
    def report_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "report_name", value)

    @property
    @pulumi.getter(name="reportVersioning")
    def report_versioning(self) -> Optional[pulumi.Input[str]]:
        """
        Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: `CREATE_NEW_REPORT` and `OVERWRITE_REPORT`.
        """
        return pulumi.get(self, "report_versioning")

    @report_versioning.setter
    def report_versioning(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "report_versioning", value)

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the existing S3 bucket to hold generated reports.
        """
        return pulumi.get(self, "s3_bucket")

    @s3_bucket.setter
    def s3_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_bucket", value)

    @property
    @pulumi.getter(name="s3Prefix")
    def s3_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Report path prefix. Limited to 256 characters.
        """
        return pulumi.get(self, "s3_prefix")

    @s3_prefix.setter
    def s3_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_prefix", value)

    @property
    @pulumi.getter(name="s3Region")
    def s3_region(self) -> Optional[pulumi.Input[str]]:
        """
        Region of the existing S3 bucket to hold generated reports.
        """
        return pulumi.get(self, "s3_region")

    @s3_region.setter
    def s3_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_region", value)

    @property
    @pulumi.getter(name="timeUnit")
    def time_unit(self) -> Optional[pulumi.Input[str]]:
        """
        The frequency on which report data are measured and displayed.  Valid values are: `HOURLY`, `DAILY`.
        """
        return pulumi.get(self, "time_unit")

    @time_unit.setter
    def time_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_unit", value)


class ReportDefinition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_artifacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 additional_schema_elements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 compression: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 refresh_closed_reports: Optional[pulumi.Input[bool]] = None,
                 report_name: Optional[pulumi.Input[str]] = None,
                 report_versioning: Optional[pulumi.Input[str]] = None,
                 s3_bucket: Optional[pulumi.Input[str]] = None,
                 s3_prefix: Optional[pulumi.Input[str]] = None,
                 s3_region: Optional[pulumi.Input[str]] = None,
                 time_unit: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Cost and Usage Report Definitions.

        > *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.

        > *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_cur_report_definition = aws.cur.ReportDefinition("exampleCurReportDefinition",
            additional_artifacts=[
                "REDSHIFT",
                "QUICKSIGHT",
            ],
            additional_schema_elements=["RESOURCES"],
            compression="GZIP",
            format="textORcsv",
            report_name="example-cur-report-definition",
            s3_bucket="example-bucket-name",
            s3_region="us-east-1",
            time_unit="HOURLY")
        ```

        ## Import

        Report Definitions can be imported using the `report_name`, e.g.

        ```sh
         $ pulumi import aws:cur/reportDefinition:ReportDefinition example_cur_report_definition example-cur-report-definition
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] additional_artifacts: A list of additional artifacts. Valid values are: `REDSHIFT`, `QUICKSIGHT`, `ATHENA`. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be `OVERWRITE_REPORT`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] additional_schema_elements: A list of schema elements. Valid values are: `RESOURCES`.
        :param pulumi.Input[str] compression: Compression format for report. Valid values are: `GZIP`, `ZIP`, `Parquet`. If `Parquet` is used, then format must also be `Parquet`.
        :param pulumi.Input[str] format: Format for report. Valid values are: `textORcsv`, `Parquet`. If `Parquet` is used, then Compression must also be `Parquet`.
        :param pulumi.Input[bool] refresh_closed_reports: Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
        :param pulumi.Input[str] report_name: Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
        :param pulumi.Input[str] report_versioning: Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: `CREATE_NEW_REPORT` and `OVERWRITE_REPORT`.
        :param pulumi.Input[str] s3_bucket: Name of the existing S3 bucket to hold generated reports.
        :param pulumi.Input[str] s3_prefix: Report path prefix. Limited to 256 characters.
        :param pulumi.Input[str] s3_region: Region of the existing S3 bucket to hold generated reports.
        :param pulumi.Input[str] time_unit: The frequency on which report data are measured and displayed.  Valid values are: `HOURLY`, `DAILY`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReportDefinitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Cost and Usage Report Definitions.

        > *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.

        > *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_cur_report_definition = aws.cur.ReportDefinition("exampleCurReportDefinition",
            additional_artifacts=[
                "REDSHIFT",
                "QUICKSIGHT",
            ],
            additional_schema_elements=["RESOURCES"],
            compression="GZIP",
            format="textORcsv",
            report_name="example-cur-report-definition",
            s3_bucket="example-bucket-name",
            s3_region="us-east-1",
            time_unit="HOURLY")
        ```

        ## Import

        Report Definitions can be imported using the `report_name`, e.g.

        ```sh
         $ pulumi import aws:cur/reportDefinition:ReportDefinition example_cur_report_definition example-cur-report-definition
        ```

        :param str resource_name: The name of the resource.
        :param ReportDefinitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReportDefinitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_artifacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 additional_schema_elements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 compression: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 refresh_closed_reports: Optional[pulumi.Input[bool]] = None,
                 report_name: Optional[pulumi.Input[str]] = None,
                 report_versioning: Optional[pulumi.Input[str]] = None,
                 s3_bucket: Optional[pulumi.Input[str]] = None,
                 s3_prefix: Optional[pulumi.Input[str]] = None,
                 s3_region: Optional[pulumi.Input[str]] = None,
                 time_unit: Optional[pulumi.Input[str]] = None,
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
            __props__ = ReportDefinitionArgs.__new__(ReportDefinitionArgs)

            __props__.__dict__["additional_artifacts"] = additional_artifacts
            if additional_schema_elements is None and not opts.urn:
                raise TypeError("Missing required property 'additional_schema_elements'")
            __props__.__dict__["additional_schema_elements"] = additional_schema_elements
            if compression is None and not opts.urn:
                raise TypeError("Missing required property 'compression'")
            __props__.__dict__["compression"] = compression
            if format is None and not opts.urn:
                raise TypeError("Missing required property 'format'")
            __props__.__dict__["format"] = format
            __props__.__dict__["refresh_closed_reports"] = refresh_closed_reports
            if report_name is None and not opts.urn:
                raise TypeError("Missing required property 'report_name'")
            __props__.__dict__["report_name"] = report_name
            __props__.__dict__["report_versioning"] = report_versioning
            if s3_bucket is None and not opts.urn:
                raise TypeError("Missing required property 's3_bucket'")
            __props__.__dict__["s3_bucket"] = s3_bucket
            __props__.__dict__["s3_prefix"] = s3_prefix
            if s3_region is None and not opts.urn:
                raise TypeError("Missing required property 's3_region'")
            __props__.__dict__["s3_region"] = s3_region
            if time_unit is None and not opts.urn:
                raise TypeError("Missing required property 'time_unit'")
            __props__.__dict__["time_unit"] = time_unit
            __props__.__dict__["arn"] = None
        super(ReportDefinition, __self__).__init__(
            'aws:cur/reportDefinition:ReportDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_artifacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            additional_schema_elements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            compression: Optional[pulumi.Input[str]] = None,
            format: Optional[pulumi.Input[str]] = None,
            refresh_closed_reports: Optional[pulumi.Input[bool]] = None,
            report_name: Optional[pulumi.Input[str]] = None,
            report_versioning: Optional[pulumi.Input[str]] = None,
            s3_bucket: Optional[pulumi.Input[str]] = None,
            s3_prefix: Optional[pulumi.Input[str]] = None,
            s3_region: Optional[pulumi.Input[str]] = None,
            time_unit: Optional[pulumi.Input[str]] = None) -> 'ReportDefinition':
        """
        Get an existing ReportDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] additional_artifacts: A list of additional artifacts. Valid values are: `REDSHIFT`, `QUICKSIGHT`, `ATHENA`. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be `OVERWRITE_REPORT`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] additional_schema_elements: A list of schema elements. Valid values are: `RESOURCES`.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the cur report.
        :param pulumi.Input[str] compression: Compression format for report. Valid values are: `GZIP`, `ZIP`, `Parquet`. If `Parquet` is used, then format must also be `Parquet`.
        :param pulumi.Input[str] format: Format for report. Valid values are: `textORcsv`, `Parquet`. If `Parquet` is used, then Compression must also be `Parquet`.
        :param pulumi.Input[bool] refresh_closed_reports: Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
        :param pulumi.Input[str] report_name: Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
        :param pulumi.Input[str] report_versioning: Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: `CREATE_NEW_REPORT` and `OVERWRITE_REPORT`.
        :param pulumi.Input[str] s3_bucket: Name of the existing S3 bucket to hold generated reports.
        :param pulumi.Input[str] s3_prefix: Report path prefix. Limited to 256 characters.
        :param pulumi.Input[str] s3_region: Region of the existing S3 bucket to hold generated reports.
        :param pulumi.Input[str] time_unit: The frequency on which report data are measured and displayed.  Valid values are: `HOURLY`, `DAILY`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReportDefinitionState.__new__(_ReportDefinitionState)

        __props__.__dict__["additional_artifacts"] = additional_artifacts
        __props__.__dict__["additional_schema_elements"] = additional_schema_elements
        __props__.__dict__["arn"] = arn
        __props__.__dict__["compression"] = compression
        __props__.__dict__["format"] = format
        __props__.__dict__["refresh_closed_reports"] = refresh_closed_reports
        __props__.__dict__["report_name"] = report_name
        __props__.__dict__["report_versioning"] = report_versioning
        __props__.__dict__["s3_bucket"] = s3_bucket
        __props__.__dict__["s3_prefix"] = s3_prefix
        __props__.__dict__["s3_region"] = s3_region
        __props__.__dict__["time_unit"] = time_unit
        return ReportDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalArtifacts")
    def additional_artifacts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of additional artifacts. Valid values are: `REDSHIFT`, `QUICKSIGHT`, `ATHENA`. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be `OVERWRITE_REPORT`.
        """
        return pulumi.get(self, "additional_artifacts")

    @property
    @pulumi.getter(name="additionalSchemaElements")
    def additional_schema_elements(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of schema elements. Valid values are: `RESOURCES`.
        """
        return pulumi.get(self, "additional_schema_elements")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) specifying the cur report.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def compression(self) -> pulumi.Output[str]:
        """
        Compression format for report. Valid values are: `GZIP`, `ZIP`, `Parquet`. If `Parquet` is used, then format must also be `Parquet`.
        """
        return pulumi.get(self, "compression")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[str]:
        """
        Format for report. Valid values are: `textORcsv`, `Parquet`. If `Parquet` is used, then Compression must also be `Parquet`.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="refreshClosedReports")
    def refresh_closed_reports(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
        """
        return pulumi.get(self, "refresh_closed_reports")

    @property
    @pulumi.getter(name="reportName")
    def report_name(self) -> pulumi.Output[str]:
        """
        Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
        """
        return pulumi.get(self, "report_name")

    @property
    @pulumi.getter(name="reportVersioning")
    def report_versioning(self) -> pulumi.Output[Optional[str]]:
        """
        Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: `CREATE_NEW_REPORT` and `OVERWRITE_REPORT`.
        """
        return pulumi.get(self, "report_versioning")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> pulumi.Output[str]:
        """
        Name of the existing S3 bucket to hold generated reports.
        """
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="s3Prefix")
    def s3_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Report path prefix. Limited to 256 characters.
        """
        return pulumi.get(self, "s3_prefix")

    @property
    @pulumi.getter(name="s3Region")
    def s3_region(self) -> pulumi.Output[str]:
        """
        Region of the existing S3 bucket to hold generated reports.
        """
        return pulumi.get(self, "s3_region")

    @property
    @pulumi.getter(name="timeUnit")
    def time_unit(self) -> pulumi.Output[str]:
        """
        The frequency on which report data are measured and displayed.  Valid values are: `HOURLY`, `DAILY`.
        """
        return pulumi.get(self, "time_unit")

