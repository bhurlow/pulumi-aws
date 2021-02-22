// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sfn
{
    /// <summary>
    /// Provides a Step Function State Machine resource
    /// 
    /// ## Example Usage
    /// ### Basic (Standard Workflow)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // ...
    ///         var sfnStateMachine = new Aws.Sfn.StateMachine("sfnStateMachine", new Aws.Sfn.StateMachineArgs
    ///         {
    ///             RoleArn = aws_iam_role.Iam_for_sfn.Arn,
    ///             Definition = @$"{{
    ///   ""Comment"": ""A Hello World example of the Amazon States Language using an AWS Lambda Function"",
    ///   ""StartAt"": ""HelloWorld"",
    ///   ""States"": {{
    ///     ""HelloWorld"": {{
    ///       ""Type"": ""Task"",
    ///       ""Resource"": ""{aws_lambda_function.Lambda.Arn}"",
    ///       ""End"": true
    ///     }}
    ///   }}
    /// }}
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Basic (Express Workflow)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // ...
    ///         var sfnStateMachine = new Aws.Sfn.StateMachine("sfnStateMachine", new Aws.Sfn.StateMachineArgs
    ///         {
    ///             RoleArn = aws_iam_role.Iam_for_sfn.Arn,
    ///             Type = "EXPRESS",
    ///             Definition = @$"{{
    ///   ""Comment"": ""A Hello World example of the Amazon States Language using an AWS Lambda Function"",
    ///   ""StartAt"": ""HelloWorld"",
    ///   ""States"": {{
    ///     ""HelloWorld"": {{
    ///       ""Type"": ""Task"",
    ///       ""Resource"": ""{aws_lambda_function.Lambda.Arn}"",
    ///       ""End"": true
    ///     }}
    ///   }}
    /// }}
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Logging
    /// 
    /// &gt; *NOTE:* See the [AWS Step Functions Developer Guide](https://docs.aws.amazon.com/step-functions/latest/dg/welcome.html) for more information about enabling Step Function logging.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // ...
    ///         var sfnStateMachine = new Aws.Sfn.StateMachine("sfnStateMachine", new Aws.Sfn.StateMachineArgs
    ///         {
    ///             RoleArn = aws_iam_role.Iam_for_sfn.Arn,
    ///             Definition = @$"{{
    ///   ""Comment"": ""A Hello World example of the Amazon States Language using an AWS Lambda Function"",
    ///   ""StartAt"": ""HelloWorld"",
    ///   ""States"": {{
    ///     ""HelloWorld"": {{
    ///       ""Type"": ""Task"",
    ///       ""Resource"": ""{aws_lambda_function.Lambda.Arn}"",
    ///       ""End"": true
    ///     }}
    ///   }}
    /// }}
    /// ",
    ///             LoggingConfiguration = new Aws.Sfn.Inputs.StateMachineLoggingConfigurationArgs
    ///             {
    ///                 LogDestination = $"{aws_cloudwatch_log_group.Log_group_for_sfn.Arn}:*",
    ///                 IncludeExecutionData = true,
    ///                 Level = "ERROR",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// State Machines can be imported using the `arn`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:sfn/stateMachine:StateMachine foo arn:aws:states:eu-west-1:123456789098:stateMachine:bar
    /// ```
    /// </summary>
    [AwsResourceType("aws:sfn/stateMachine:StateMachine")]
    public partial class StateMachine : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the state machine.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The date the state machine was created.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The Amazon States Language definition of the state machine.
        /// </summary>
        [Output("definition")]
        public Output<string> Definition { get; private set; } = null!;

        /// <summary>
        /// Defines what execution history events are logged and where they are logged. The `logging_configuration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
        /// </summary>
        [Output("loggingConfiguration")]
        public Output<Outputs.StateMachineLoggingConfiguration> LoggingConfiguration { get; private set; } = null!;

        /// <summary>
        /// The name of the state machine.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The current status of the state machine. Either "ACTIVE" or "DELETING".
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StateMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StateMachine(string name, StateMachineArgs args, CustomResourceOptions? options = null)
            : base("aws:sfn/stateMachine:StateMachine", name, args ?? new StateMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StateMachine(string name, Input<string> id, StateMachineState? state = null, CustomResourceOptions? options = null)
            : base("aws:sfn/stateMachine:StateMachine", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StateMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StateMachine Get(string name, Input<string> id, StateMachineState? state = null, CustomResourceOptions? options = null)
        {
            return new StateMachine(name, id, state, options);
        }
    }

    public sealed class StateMachineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon States Language definition of the state machine.
        /// </summary>
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        /// <summary>
        /// Defines what execution history events are logged and where they are logged. The `logging_configuration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
        /// </summary>
        [Input("loggingConfiguration")]
        public Input<Inputs.StateMachineLoggingConfigurationArgs>? LoggingConfiguration { get; set; }

        /// <summary>
        /// The name of the state machine.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public StateMachineArgs()
        {
        }
    }

    public sealed class StateMachineState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the state machine.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The date the state machine was created.
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// The Amazon States Language definition of the state machine.
        /// </summary>
        [Input("definition")]
        public Input<string>? Definition { get; set; }

        /// <summary>
        /// Defines what execution history events are logged and where they are logged. The `logging_configuration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
        /// </summary>
        [Input("loggingConfiguration")]
        public Input<Inputs.StateMachineLoggingConfigurationGetArgs>? LoggingConfiguration { get; set; }

        /// <summary>
        /// The name of the state machine.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The current status of the state machine. Either "ACTIVE" or "DELETING".
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public StateMachineState()
        {
        }
    }
}
