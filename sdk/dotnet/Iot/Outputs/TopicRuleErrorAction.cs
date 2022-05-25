// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot.Outputs
{

    [OutputType]
    public sealed class TopicRuleErrorAction
    {
        public readonly Outputs.TopicRuleErrorActionCloudwatchAlarm? CloudwatchAlarm;
        public readonly Outputs.TopicRuleErrorActionCloudwatchLogs? CloudwatchLogs;
        public readonly Outputs.TopicRuleErrorActionCloudwatchMetric? CloudwatchMetric;
        public readonly Outputs.TopicRuleErrorActionDynamodb? Dynamodb;
        public readonly Outputs.TopicRuleErrorActionDynamodbv2? Dynamodbv2;
        public readonly Outputs.TopicRuleErrorActionElasticsearch? Elasticsearch;
        public readonly Outputs.TopicRuleErrorActionFirehose? Firehose;
        public readonly Outputs.TopicRuleErrorActionHttp? Http;
        public readonly Outputs.TopicRuleErrorActionIotAnalytics? IotAnalytics;
        public readonly Outputs.TopicRuleErrorActionIotEvents? IotEvents;
        public readonly Outputs.TopicRuleErrorActionKafka? Kafka;
        public readonly Outputs.TopicRuleErrorActionKinesis? Kinesis;
        public readonly Outputs.TopicRuleErrorActionLambda? Lambda;
        public readonly Outputs.TopicRuleErrorActionRepublish? Republish;
        public readonly Outputs.TopicRuleErrorActionS3? S3;
        public readonly Outputs.TopicRuleErrorActionSns? Sns;
        public readonly Outputs.TopicRuleErrorActionSqs? Sqs;
        public readonly Outputs.TopicRuleErrorActionStepFunctions? StepFunctions;
        public readonly Outputs.TopicRuleErrorActionTimestream? Timestream;

        [OutputConstructor]
        private TopicRuleErrorAction(
            Outputs.TopicRuleErrorActionCloudwatchAlarm? cloudwatchAlarm,

            Outputs.TopicRuleErrorActionCloudwatchLogs? cloudwatchLogs,

            Outputs.TopicRuleErrorActionCloudwatchMetric? cloudwatchMetric,

            Outputs.TopicRuleErrorActionDynamodb? dynamodb,

            Outputs.TopicRuleErrorActionDynamodbv2? dynamodbv2,

            Outputs.TopicRuleErrorActionElasticsearch? elasticsearch,

            Outputs.TopicRuleErrorActionFirehose? firehose,

            Outputs.TopicRuleErrorActionHttp? http,

            Outputs.TopicRuleErrorActionIotAnalytics? iotAnalytics,

            Outputs.TopicRuleErrorActionIotEvents? iotEvents,

            Outputs.TopicRuleErrorActionKafka? kafka,

            Outputs.TopicRuleErrorActionKinesis? kinesis,

            Outputs.TopicRuleErrorActionLambda? lambda,

            Outputs.TopicRuleErrorActionRepublish? republish,

            Outputs.TopicRuleErrorActionS3? s3,

            Outputs.TopicRuleErrorActionSns? sns,

            Outputs.TopicRuleErrorActionSqs? sqs,

            Outputs.TopicRuleErrorActionStepFunctions? stepFunctions,

            Outputs.TopicRuleErrorActionTimestream? timestream)
        {
            CloudwatchAlarm = cloudwatchAlarm;
            CloudwatchLogs = cloudwatchLogs;
            CloudwatchMetric = cloudwatchMetric;
            Dynamodb = dynamodb;
            Dynamodbv2 = dynamodbv2;
            Elasticsearch = elasticsearch;
            Firehose = firehose;
            Http = http;
            IotAnalytics = iotAnalytics;
            IotEvents = iotEvents;
            Kafka = kafka;
            Kinesis = kinesis;
            Lambda = lambda;
            Republish = republish;
            S3 = s3;
            Sns = sns;
            Sqs = sqs;
            StepFunctions = stepFunctions;
            Timestream = timestream;
        }
    }
}
