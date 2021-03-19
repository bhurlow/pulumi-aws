# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['MetricAlarm']


class MetricAlarm(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions_enabled: Optional[pulumi.Input[bool]] = None,
                 alarm_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 alarm_description: Optional[pulumi.Input[str]] = None,
                 comparison_operator: Optional[pulumi.Input[str]] = None,
                 datapoints_to_alarm: Optional[pulumi.Input[int]] = None,
                 dimensions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 evaluate_low_sample_count_percentiles: Optional[pulumi.Input[str]] = None,
                 evaluation_periods: Optional[pulumi.Input[int]] = None,
                 extended_statistic: Optional[pulumi.Input[str]] = None,
                 insufficient_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 metric_queries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricAlarmMetricQueryArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ok_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 statistic: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 threshold: Optional[pulumi.Input[float]] = None,
                 threshold_metric_id: Optional[pulumi.Input[str]] = None,
                 treat_missing_data: Optional[pulumi.Input[str]] = None,
                 unit: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudWatch Metric Alarm resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foobar = aws.cloudwatch.MetricAlarm("foobar",
            alarm_description="This metric monitors ec2 cpu utilization",
            comparison_operator="GreaterThanOrEqualToThreshold",
            evaluation_periods=2,
            insufficient_data_actions=[],
            metric_name="CPUUtilization",
            namespace="AWS/EC2",
            period=120,
            statistic="Average",
            threshold=80)
        ```
        ## Example in Conjunction with Scaling Policies

        ```python
        import pulumi
        import pulumi_aws as aws

        bat_policy = aws.autoscaling.Policy("batPolicy",
            scaling_adjustment=4,
            adjustment_type="ChangeInCapacity",
            cooldown=300,
            autoscaling_group_name=aws_autoscaling_group["bar"]["name"])
        bat_metric_alarm = aws.cloudwatch.MetricAlarm("batMetricAlarm",
            comparison_operator="GreaterThanOrEqualToThreshold",
            evaluation_periods=2,
            metric_name="CPUUtilization",
            namespace="AWS/EC2",
            period=120,
            statistic="Average",
            threshold=80,
            dimensions={
                "AutoScalingGroupName": aws_autoscaling_group["bar"]["name"],
            },
            alarm_description="This metric monitors ec2 cpu utilization",
            alarm_actions=[bat_policy.arn])
        ```

        ## Example with an Expression

        ```python
        import pulumi
        import pulumi_aws as aws

        foobar = aws.cloudwatch.MetricAlarm("foobar",
            alarm_description="Request error rate has exceeded 10%",
            comparison_operator="GreaterThanOrEqualToThreshold",
            evaluation_periods=2,
            insufficient_data_actions=[],
            metric_queries=[
                aws.cloudwatch.MetricAlarmMetricQueryArgs(
                    expression="m2/m1*100",
                    id="e1",
                    label="Error Rate",
                    return_data=True,
                ),
                aws.cloudwatch.MetricAlarmMetricQueryArgs(
                    id="m1",
                    metric=aws.cloudwatch.MetricAlarmMetricQueryMetricArgs(
                        dimensions={
                            "LoadBalancer": "app/web",
                        },
                        metric_name="RequestCount",
                        namespace="AWS/ApplicationELB",
                        period=120,
                        stat="Sum",
                        unit="Count",
                    ),
                ),
                aws.cloudwatch.MetricAlarmMetricQueryArgs(
                    id="m2",
                    metric=aws.cloudwatch.MetricAlarmMetricQueryMetricArgs(
                        dimensions={
                            "LoadBalancer": "app/web",
                        },
                        metric_name="HTTPCode_ELB_5XX_Count",
                        namespace="AWS/ApplicationELB",
                        period=120,
                        stat="Sum",
                        unit="Count",
                    ),
                ),
            ],
            threshold=10)
        ```

        ```python
        import pulumi
        import pulumi_aws as aws

        xx_anomaly_detection = aws.cloudwatch.MetricAlarm("xxAnomalyDetection",
            alarm_description="This metric monitors ec2 cpu utilization",
            comparison_operator="GreaterThanUpperThreshold",
            evaluation_periods=2,
            insufficient_data_actions=[],
            metric_queries=[
                aws.cloudwatch.MetricAlarmMetricQueryArgs(
                    expression="ANOMALY_DETECTION_BAND(m1)",
                    id="e1",
                    label="CPUUtilization (Expected)",
                    return_data=True,
                ),
                aws.cloudwatch.MetricAlarmMetricQueryArgs(
                    id="m1",
                    metric=aws.cloudwatch.MetricAlarmMetricQueryMetricArgs(
                        dimensions={
                            "InstanceId": "i-abc123",
                        },
                        metric_name="CPUUtilization",
                        namespace="AWS/EC2",
                        period=120,
                        stat="Average",
                        unit="Count",
                    ),
                    return_data=True,
                ),
            ],
            threshold_metric_id="e1")
        ```

        ## Example of monitoring Healthy Hosts on NLB using Target Group and NLB

        ```python
        import pulumi
        import pulumi_aws as aws

        nlb_healthyhosts = aws.cloudwatch.MetricAlarm("nlbHealthyhosts",
            comparison_operator="LessThanThreshold",
            evaluation_periods=1,
            metric_name="HealthyHostCount",
            namespace="AWS/NetworkELB",
            period=60,
            statistic="Average",
            threshold=var["logstash_servers_count"],
            alarm_description="Number of healthy nodes in Target Group",
            actions_enabled=True,
            alarm_actions=[aws_sns_topic["sns"]["arn"]],
            ok_actions=[aws_sns_topic["sns"]["arn"]],
            dimensions={
                "TargetGroup": aws_lb_target_group["lb-tg"]["arn_suffix"],
                "LoadBalancer": aws_lb["lb"]["arn_suffix"],
            })
        ```

        > **NOTE:**  You cannot create a metric alarm consisting of both `statistic` and `extended_statistic` parameters.
        You must choose one or the other

        ## Import

        CloudWatch Metric Alarm can be imported using the `alarm_name`, e.g.

        ```sh
         $ pulumi import aws:cloudwatch/metricAlarm:MetricAlarm test alarm-12345
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] actions_enabled: Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_actions: The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[str] alarm_description: The description for the alarm.
        :param pulumi.Input[str] comparison_operator: The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
        :param pulumi.Input[int] datapoints_to_alarm: The number of datapoints that must be breaching to trigger the alarm.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] dimensions: The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[str] evaluate_low_sample_count_percentiles: Used only for alarms
               based on percentiles. If you specify `ignore`, the alarm state will not
               change during periods with too few data points to be statistically significant.
               If you specify `evaluate` or omit this parameter, the alarm will always be
               evaluated and possibly change state no matter how many data points are available.
               The following values are supported: `ignore`, and `evaluate`.
        :param pulumi.Input[int] evaluation_periods: The number of periods over which data is compared to the specified threshold.
        :param pulumi.Input[str] extended_statistic: The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] insufficient_data_actions: The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[str] metric_name: The name for this metric.
               See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricAlarmMetricQueryArgs']]]] metric_queries: Enables you to create an alarm based on a metric math expression. You may specify at most 20.
        :param pulumi.Input[str] name: The descriptive name for the alarm. This name must be unique within the user's AWS account
        :param pulumi.Input[str] namespace: The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
               See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ok_actions: The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[int] period: The period in seconds over which the specified `stat` is applied.
        :param pulumi.Input[str] statistic: The statistic to apply to the alarm's associated metric.
               Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[float] threshold: The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
        :param pulumi.Input[str] threshold_metric_id: If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
        :param pulumi.Input[str] treat_missing_data: Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
        :param pulumi.Input[str] unit: The unit for this metric.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['actions_enabled'] = actions_enabled
            __props__['alarm_actions'] = alarm_actions
            __props__['alarm_description'] = alarm_description
            if comparison_operator is None and not opts.urn:
                raise TypeError("Missing required property 'comparison_operator'")
            __props__['comparison_operator'] = comparison_operator
            __props__['datapoints_to_alarm'] = datapoints_to_alarm
            __props__['dimensions'] = dimensions
            __props__['evaluate_low_sample_count_percentiles'] = evaluate_low_sample_count_percentiles
            if evaluation_periods is None and not opts.urn:
                raise TypeError("Missing required property 'evaluation_periods'")
            __props__['evaluation_periods'] = evaluation_periods
            __props__['extended_statistic'] = extended_statistic
            __props__['insufficient_data_actions'] = insufficient_data_actions
            __props__['metric_name'] = metric_name
            __props__['metric_queries'] = metric_queries
            __props__['name'] = name
            __props__['namespace'] = namespace
            __props__['ok_actions'] = ok_actions
            __props__['period'] = period
            __props__['statistic'] = statistic
            __props__['tags'] = tags
            __props__['threshold'] = threshold
            __props__['threshold_metric_id'] = threshold_metric_id
            __props__['treat_missing_data'] = treat_missing_data
            __props__['unit'] = unit
            __props__['arn'] = None
        super(MetricAlarm, __self__).__init__(
            'aws:cloudwatch/metricAlarm:MetricAlarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            actions_enabled: Optional[pulumi.Input[bool]] = None,
            alarm_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            alarm_description: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            comparison_operator: Optional[pulumi.Input[str]] = None,
            datapoints_to_alarm: Optional[pulumi.Input[int]] = None,
            dimensions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            evaluate_low_sample_count_percentiles: Optional[pulumi.Input[str]] = None,
            evaluation_periods: Optional[pulumi.Input[int]] = None,
            extended_statistic: Optional[pulumi.Input[str]] = None,
            insufficient_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            metric_name: Optional[pulumi.Input[str]] = None,
            metric_queries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricAlarmMetricQueryArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            ok_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            period: Optional[pulumi.Input[int]] = None,
            statistic: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            threshold: Optional[pulumi.Input[float]] = None,
            threshold_metric_id: Optional[pulumi.Input[str]] = None,
            treat_missing_data: Optional[pulumi.Input[str]] = None,
            unit: Optional[pulumi.Input[str]] = None) -> 'MetricAlarm':
        """
        Get an existing MetricAlarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] actions_enabled: Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_actions: The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[str] alarm_description: The description for the alarm.
        :param pulumi.Input[str] arn: The ARN of the CloudWatch Metric Alarm.
        :param pulumi.Input[str] comparison_operator: The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
        :param pulumi.Input[int] datapoints_to_alarm: The number of datapoints that must be breaching to trigger the alarm.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] dimensions: The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[str] evaluate_low_sample_count_percentiles: Used only for alarms
               based on percentiles. If you specify `ignore`, the alarm state will not
               change during periods with too few data points to be statistically significant.
               If you specify `evaluate` or omit this parameter, the alarm will always be
               evaluated and possibly change state no matter how many data points are available.
               The following values are supported: `ignore`, and `evaluate`.
        :param pulumi.Input[int] evaluation_periods: The number of periods over which data is compared to the specified threshold.
        :param pulumi.Input[str] extended_statistic: The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] insufficient_data_actions: The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[str] metric_name: The name for this metric.
               See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricAlarmMetricQueryArgs']]]] metric_queries: Enables you to create an alarm based on a metric math expression. You may specify at most 20.
        :param pulumi.Input[str] name: The descriptive name for the alarm. This name must be unique within the user's AWS account
        :param pulumi.Input[str] namespace: The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
               See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ok_actions: The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[int] period: The period in seconds over which the specified `stat` is applied.
        :param pulumi.Input[str] statistic: The statistic to apply to the alarm's associated metric.
               Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[float] threshold: The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
        :param pulumi.Input[str] threshold_metric_id: If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
        :param pulumi.Input[str] treat_missing_data: Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
        :param pulumi.Input[str] unit: The unit for this metric.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["actions_enabled"] = actions_enabled
        __props__["alarm_actions"] = alarm_actions
        __props__["alarm_description"] = alarm_description
        __props__["arn"] = arn
        __props__["comparison_operator"] = comparison_operator
        __props__["datapoints_to_alarm"] = datapoints_to_alarm
        __props__["dimensions"] = dimensions
        __props__["evaluate_low_sample_count_percentiles"] = evaluate_low_sample_count_percentiles
        __props__["evaluation_periods"] = evaluation_periods
        __props__["extended_statistic"] = extended_statistic
        __props__["insufficient_data_actions"] = insufficient_data_actions
        __props__["metric_name"] = metric_name
        __props__["metric_queries"] = metric_queries
        __props__["name"] = name
        __props__["namespace"] = namespace
        __props__["ok_actions"] = ok_actions
        __props__["period"] = period
        __props__["statistic"] = statistic
        __props__["tags"] = tags
        __props__["threshold"] = threshold
        __props__["threshold_metric_id"] = threshold_metric_id
        __props__["treat_missing_data"] = treat_missing_data
        __props__["unit"] = unit
        return MetricAlarm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actionsEnabled")
    def actions_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
        """
        return pulumi.get(self, "actions_enabled")

    @property
    @pulumi.getter(name="alarmActions")
    def alarm_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "alarm_actions")

    @property
    @pulumi.getter(name="alarmDescription")
    def alarm_description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the alarm.
        """
        return pulumi.get(self, "alarm_description")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the CloudWatch Metric Alarm.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="comparisonOperator")
    def comparison_operator(self) -> pulumi.Output[str]:
        """
        The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
        """
        return pulumi.get(self, "comparison_operator")

    @property
    @pulumi.getter(name="datapointsToAlarm")
    def datapoints_to_alarm(self) -> pulumi.Output[Optional[int]]:
        """
        The number of datapoints that must be breaching to trigger the alarm.
        """
        return pulumi.get(self, "datapoints_to_alarm")

    @property
    @pulumi.getter
    def dimensions(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        """
        return pulumi.get(self, "dimensions")

    @property
    @pulumi.getter(name="evaluateLowSampleCountPercentiles")
    def evaluate_low_sample_count_percentiles(self) -> pulumi.Output[str]:
        """
        Used only for alarms
        based on percentiles. If you specify `ignore`, the alarm state will not
        change during periods with too few data points to be statistically significant.
        If you specify `evaluate` or omit this parameter, the alarm will always be
        evaluated and possibly change state no matter how many data points are available.
        The following values are supported: `ignore`, and `evaluate`.
        """
        return pulumi.get(self, "evaluate_low_sample_count_percentiles")

    @property
    @pulumi.getter(name="evaluationPeriods")
    def evaluation_periods(self) -> pulumi.Output[int]:
        """
        The number of periods over which data is compared to the specified threshold.
        """
        return pulumi.get(self, "evaluation_periods")

    @property
    @pulumi.getter(name="extendedStatistic")
    def extended_statistic(self) -> pulumi.Output[Optional[str]]:
        """
        The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
        """
        return pulumi.get(self, "extended_statistic")

    @property
    @pulumi.getter(name="insufficientDataActions")
    def insufficient_data_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "insufficient_data_actions")

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name for this metric.
        See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        """
        return pulumi.get(self, "metric_name")

    @property
    @pulumi.getter(name="metricQueries")
    def metric_queries(self) -> pulumi.Output[Optional[Sequence['outputs.MetricAlarmMetricQuery']]]:
        """
        Enables you to create an alarm based on a metric math expression. You may specify at most 20.
        """
        return pulumi.get(self, "metric_queries")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The descriptive name for the alarm. This name must be unique within the user's AWS account
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
        See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="okActions")
    def ok_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "ok_actions")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        """
        The period in seconds over which the specified `stat` is applied.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter
    def statistic(self) -> pulumi.Output[Optional[str]]:
        """
        The statistic to apply to the alarm's associated metric.
        Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
        """
        return pulumi.get(self, "statistic")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def threshold(self) -> pulumi.Output[Optional[float]]:
        """
        The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
        """
        return pulumi.get(self, "threshold")

    @property
    @pulumi.getter(name="thresholdMetricId")
    def threshold_metric_id(self) -> pulumi.Output[Optional[str]]:
        """
        If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
        """
        return pulumi.get(self, "threshold_metric_id")

    @property
    @pulumi.getter(name="treatMissingData")
    def treat_missing_data(self) -> pulumi.Output[Optional[str]]:
        """
        Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
        """
        return pulumi.get(self, "treat_missing_data")

    @property
    @pulumi.getter
    def unit(self) -> pulumi.Output[Optional[str]]:
        """
        The unit for this metric.
        """
        return pulumi.get(self, "unit")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

