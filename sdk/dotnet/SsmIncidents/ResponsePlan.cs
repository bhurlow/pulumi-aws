// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SsmIncidents
{
    /// <summary>
    /// Provides a resource to manage response plans in AWS Systems Manager Incident Manager.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.SsmIncidents.ResponsePlan("example", new()
    ///     {
    ///         IncidentTemplate = new Aws.SsmIncidents.Inputs.ResponsePlanIncidentTemplateArgs
    ///         {
    ///             Title = "title",
    ///             Impact = 3,
    ///         },
    ///         Tags = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             aws_ssmincidents_replication_set.Example,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Usage With All Fields
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.SsmIncidents.ResponsePlan("example", new()
    ///     {
    ///         IncidentTemplate = new Aws.SsmIncidents.Inputs.ResponsePlanIncidentTemplateArgs
    ///         {
    ///             Title = "title",
    ///             Impact = 3,
    ///             DedupeString = "dedupe",
    ///             IncidentTags = 
    ///             {
    ///                 { "key", "value" },
    ///             },
    ///             NotificationTargets = new[]
    ///             {
    ///                 new Aws.SsmIncidents.Inputs.ResponsePlanIncidentTemplateNotificationTargetArgs
    ///                 {
    ///                     SnsTopicArn = aws_sns_topic.Example1.Arn,
    ///                 },
    ///                 new Aws.SsmIncidents.Inputs.ResponsePlanIncidentTemplateNotificationTargetArgs
    ///                 {
    ///                     SnsTopicArn = aws_sns_topic.Example2.Arn,
    ///                 },
    ///             },
    ///             Summary = "summary",
    ///         },
    ///         DisplayName = "display name",
    ///         ChatChannels = new[]
    ///         {
    ///             aws_sns_topic.Topic.Arn,
    ///         },
    ///         Engagements = new[]
    ///         {
    ///             "arn:aws:ssm-contacts:us-east-2:111122223333:contact/test1",
    ///         },
    ///         Action = new Aws.SsmIncidents.Inputs.ResponsePlanActionArgs
    ///         {
    ///             SsmAutomations = new[]
    ///             {
    ///                 new Aws.SsmIncidents.Inputs.ResponsePlanActionSsmAutomationArgs
    ///                 {
    ///                     DocumentName = aws_ssm_document.Document1.Name,
    ///                     RoleArn = aws_iam_role.Role1.Arn,
    ///                     DocumentVersion = "version1",
    ///                     TargetAccount = "RESPONSE_PLAN_OWNER_ACCOUNT",
    ///                     Parameters = new[]
    ///                     {
    ///                         new Aws.SsmIncidents.Inputs.ResponsePlanActionSsmAutomationParameterArgs
    ///                         {
    ///                             Name = "key",
    ///                             Values = new[]
    ///                             {
    ///                                 "value1",
    ///                                 "value2",
    ///                             },
    ///                         },
    ///                         new Aws.SsmIncidents.Inputs.ResponsePlanActionSsmAutomationParameterArgs
    ///                         {
    ///                             Name = "foo",
    ///                             Values = new[]
    ///                             {
    ///                                 "bar",
    ///                             },
    ///                         },
    ///                     },
    ///                     DynamicParameters = 
    ///                     {
    ///                         { "someKey", "INVOLVED_RESOURCES" },
    ///                         { "anotherKey", "INCIDENT_RECORD_ARN" },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Integration = new Aws.SsmIncidents.Inputs.ResponsePlanIntegrationArgs
    ///         {
    ///             Pagerduties = new[]
    ///             {
    ///                 new Aws.SsmIncidents.Inputs.ResponsePlanIntegrationPagerdutyArgs
    ///                 {
    ///                     Name = "pagerdutyIntergration",
    ///                     ServiceId = "example",
    ///                     SecretId = "example",
    ///                 },
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             aws_ssmincidents_replication_set.Example,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// To import an Incident Manager response plan, specify the response plan ARN. You can find the response plan ARN in the AWS Management Console. Use the following command to run the import operation
    /// 
    /// ```sh
    ///  $ pulumi import aws:ssmincidents/responsePlan:ResponsePlan responsePlanName ARNValue
    /// ```
    /// </summary>
    [AwsResourceType("aws:ssmincidents/responsePlan:ResponsePlan")]
    public partial class ResponsePlan : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The actions that the response plan starts at the beginning of an incident.
        /// </summary>
        [Output("action")]
        public Output<Outputs.ResponsePlanAction?> Action { get; private set; } = null!;

        /// <summary>
        /// The ARN of the response plan.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Chatbot chat channel used for collaboration during an incident.
        /// </summary>
        [Output("chatChannels")]
        public Output<ImmutableArray<string>> ChatChannels { get; private set; } = null!;

        /// <summary>
        /// The long format of the response plan name. This field can contain spaces.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        /// </summary>
        [Output("engagements")]
        public Output<ImmutableArray<string>> Engagements { get; private set; } = null!;

        [Output("incidentTemplate")]
        public Output<Outputs.ResponsePlanIncidentTemplate> IncidentTemplate { get; private set; } = null!;

        /// <summary>
        /// Information about third-party services integrated into the response plan. The following values are supported:
        /// </summary>
        [Output("integration")]
        public Output<Outputs.ResponsePlanIntegration?> Integration { get; private set; } = null!;

        /// <summary>
        /// The name of the response plan.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The tags applied to the response plan.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a ResponsePlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResponsePlan(string name, ResponsePlanArgs args, CustomResourceOptions? options = null)
            : base("aws:ssmincidents/responsePlan:ResponsePlan", name, args ?? new ResponsePlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResponsePlan(string name, Input<string> id, ResponsePlanState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssmincidents/responsePlan:ResponsePlan", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResponsePlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResponsePlan Get(string name, Input<string> id, ResponsePlanState? state = null, CustomResourceOptions? options = null)
        {
            return new ResponsePlan(name, id, state, options);
        }
    }

    public sealed class ResponsePlanArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The actions that the response plan starts at the beginning of an incident.
        /// </summary>
        [Input("action")]
        public Input<Inputs.ResponsePlanActionArgs>? Action { get; set; }

        [Input("chatChannels")]
        private InputList<string>? _chatChannels;

        /// <summary>
        /// The Chatbot chat channel used for collaboration during an incident.
        /// </summary>
        public InputList<string> ChatChannels
        {
            get => _chatChannels ?? (_chatChannels = new InputList<string>());
            set => _chatChannels = value;
        }

        /// <summary>
        /// The long format of the response plan name. This field can contain spaces.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("engagements")]
        private InputList<string>? _engagements;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        /// </summary>
        public InputList<string> Engagements
        {
            get => _engagements ?? (_engagements = new InputList<string>());
            set => _engagements = value;
        }

        [Input("incidentTemplate", required: true)]
        public Input<Inputs.ResponsePlanIncidentTemplateArgs> IncidentTemplate { get; set; } = null!;

        /// <summary>
        /// Information about third-party services integrated into the response plan. The following values are supported:
        /// </summary>
        [Input("integration")]
        public Input<Inputs.ResponsePlanIntegrationArgs>? Integration { get; set; }

        /// <summary>
        /// The name of the response plan.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags applied to the response plan.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ResponsePlanArgs()
        {
        }
        public static new ResponsePlanArgs Empty => new ResponsePlanArgs();
    }

    public sealed class ResponsePlanState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The actions that the response plan starts at the beginning of an incident.
        /// </summary>
        [Input("action")]
        public Input<Inputs.ResponsePlanActionGetArgs>? Action { get; set; }

        /// <summary>
        /// The ARN of the response plan.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("chatChannels")]
        private InputList<string>? _chatChannels;

        /// <summary>
        /// The Chatbot chat channel used for collaboration during an incident.
        /// </summary>
        public InputList<string> ChatChannels
        {
            get => _chatChannels ?? (_chatChannels = new InputList<string>());
            set => _chatChannels = value;
        }

        /// <summary>
        /// The long format of the response plan name. This field can contain spaces.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("engagements")]
        private InputList<string>? _engagements;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        /// </summary>
        public InputList<string> Engagements
        {
            get => _engagements ?? (_engagements = new InputList<string>());
            set => _engagements = value;
        }

        [Input("incidentTemplate")]
        public Input<Inputs.ResponsePlanIncidentTemplateGetArgs>? IncidentTemplate { get; set; }

        /// <summary>
        /// Information about third-party services integrated into the response plan. The following values are supported:
        /// </summary>
        [Input("integration")]
        public Input<Inputs.ResponsePlanIntegrationGetArgs>? Integration { get; set; }

        /// <summary>
        /// The name of the response plan.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags applied to the response plan.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public ResponsePlanState()
        {
        }
        public static new ResponsePlanState Empty => new ResponsePlanState();
    }
}
