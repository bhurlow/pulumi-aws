// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amazon Lex Bot resource. For more information see
// [Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := lex.NewBot(ctx, "orderFlowersBot", &lex.BotArgs{
//				AbortStatement: &lex.BotAbortStatementArgs{
//					Messages: lex.BotAbortStatementMessageArray{
//						&lex.BotAbortStatementMessageArgs{
//							Content:     pulumi.String("Sorry, I am not able to assist at this time"),
//							ContentType: pulumi.String("PlainText"),
//						},
//					},
//				},
//				ChildDirected: pulumi.Bool(false),
//				ClarificationPrompt: &lex.BotClarificationPromptArgs{
//					MaxAttempts: pulumi.Int(2),
//					Messages: lex.BotClarificationPromptMessageArray{
//						&lex.BotClarificationPromptMessageArgs{
//							Content:     pulumi.String("I didn't understand you, what would you like to do?"),
//							ContentType: pulumi.String("PlainText"),
//						},
//					},
//				},
//				CreateVersion:           pulumi.Bool(false),
//				Description:             pulumi.String("Bot to order flowers on the behalf of a user"),
//				IdleSessionTtlInSeconds: pulumi.Int(600),
//				Intents: lex.BotIntentArray{
//					&lex.BotIntentArgs{
//						IntentName:    pulumi.String("OrderFlowers"),
//						IntentVersion: pulumi.String("1"),
//					},
//				},
//				Locale:          pulumi.String("en-US"),
//				Name:            pulumi.String("OrderFlowers"),
//				ProcessBehavior: pulumi.String("BUILD"),
//				VoiceId:         pulumi.String("Salli"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// terraform import {
//
//	to = aws_lex_bot.order_flowers_bot
//
//	id = "OrderFlowers" } Using `pulumi import`, import bots using their name. For exampleconsole % pulumi import aws_lex_bot.order_flowers_bot OrderFlowers
type Bot struct {
	pulumi.CustomResourceState

	// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
	AbortStatement BotAbortStatementOutput `pulumi:"abortStatement"`
	Arn            pulumi.StringOutput     `pulumi:"arn"`
	// Checksum identifying the version of the bot that was created. The checksum is not
	// included as an argument because the resource will add it automatically when updating the bot.
	Checksum pulumi.StringOutput `pulumi:"checksum"`
	// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
	ChildDirected pulumi.BoolOutput `pulumi:"childDirected"`
	// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
	ClarificationPrompt BotClarificationPromptPtrOutput `pulumi:"clarificationPrompt"`
	// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
	CreateVersion pulumi.BoolPtrOutput `pulumi:"createVersion"`
	// The date when the bot version was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
	DetectSentiment pulumi.BoolPtrOutput `pulumi:"detectSentiment"`
	// Set to `true` to enable access to natural language understanding improvements. When you set the `enableModelImprovements` parameter to true you can use the `nluIntentConfidenceThreshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enableModelImprovements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
	EnableModelImprovements pulumi.BoolPtrOutput `pulumi:"enableModelImprovements"`
	// If status is FAILED, Amazon Lex provides the reason that it failed to build the bot.
	FailureReason pulumi.StringOutput `pulumi:"failureReason"`
	// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
	IdleSessionTtlInSeconds pulumi.IntPtrOutput `pulumi:"idleSessionTtlInSeconds"`
	// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
	Intents BotIntentArrayOutput `pulumi:"intents"`
	// The date when the $LATEST version of this bot was updated.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
	Locale pulumi.StringPtrOutput `pulumi:"locale"`
	// The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enableModelImprovements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
	NluIntentConfidenceThreshold pulumi.Float64PtrOutput `pulumi:"nluIntentConfidenceThreshold"`
	// If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
	ProcessBehavior pulumi.StringPtrOutput `pulumi:"processBehavior"`
	// When you send a request to create or update a bot, Amazon Lex sets the status response
	// element to BUILDING. After Amazon Lex builds the bot, it sets status to READY. If Amazon Lex can't
	// build the bot, it sets status to FAILED. Amazon Lex returns the reason for the failure in the
	// failureReason response element.
	Status pulumi.StringOutput `pulumi:"status"`
	// The version of the bot.
	Version pulumi.StringOutput `pulumi:"version"`
	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
	VoiceId pulumi.StringOutput `pulumi:"voiceId"`
}

// NewBot registers a new resource with the given unique name, arguments, and options.
func NewBot(ctx *pulumi.Context,
	name string, args *BotArgs, opts ...pulumi.ResourceOption) (*Bot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AbortStatement == nil {
		return nil, errors.New("invalid value for required argument 'AbortStatement'")
	}
	if args.ChildDirected == nil {
		return nil, errors.New("invalid value for required argument 'ChildDirected'")
	}
	if args.Intents == nil {
		return nil, errors.New("invalid value for required argument 'Intents'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bot
	err := ctx.RegisterResource("aws:lex/bot:Bot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBot gets an existing Bot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotState, opts ...pulumi.ResourceOption) (*Bot, error) {
	var resource Bot
	err := ctx.ReadResource("aws:lex/bot:Bot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bot resources.
type botState struct {
	// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
	AbortStatement *BotAbortStatement `pulumi:"abortStatement"`
	Arn            *string            `pulumi:"arn"`
	// Checksum identifying the version of the bot that was created. The checksum is not
	// included as an argument because the resource will add it automatically when updating the bot.
	Checksum *string `pulumi:"checksum"`
	// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
	ChildDirected *bool `pulumi:"childDirected"`
	// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
	ClarificationPrompt *BotClarificationPrompt `pulumi:"clarificationPrompt"`
	// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
	CreateVersion *bool `pulumi:"createVersion"`
	// The date when the bot version was created.
	CreatedDate *string `pulumi:"createdDate"`
	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description *string `pulumi:"description"`
	// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
	DetectSentiment *bool `pulumi:"detectSentiment"`
	// Set to `true` to enable access to natural language understanding improvements. When you set the `enableModelImprovements` parameter to true you can use the `nluIntentConfidenceThreshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enableModelImprovements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
	EnableModelImprovements *bool `pulumi:"enableModelImprovements"`
	// If status is FAILED, Amazon Lex provides the reason that it failed to build the bot.
	FailureReason *string `pulumi:"failureReason"`
	// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
	IdleSessionTtlInSeconds *int `pulumi:"idleSessionTtlInSeconds"`
	// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
	Intents []BotIntent `pulumi:"intents"`
	// The date when the $LATEST version of this bot was updated.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
	Locale *string `pulumi:"locale"`
	// The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
	Name *string `pulumi:"name"`
	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enableModelImprovements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
	NluIntentConfidenceThreshold *float64 `pulumi:"nluIntentConfidenceThreshold"`
	// If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
	ProcessBehavior *string `pulumi:"processBehavior"`
	// When you send a request to create or update a bot, Amazon Lex sets the status response
	// element to BUILDING. After Amazon Lex builds the bot, it sets status to READY. If Amazon Lex can't
	// build the bot, it sets status to FAILED. Amazon Lex returns the reason for the failure in the
	// failureReason response element.
	Status *string `pulumi:"status"`
	// The version of the bot.
	Version *string `pulumi:"version"`
	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
	VoiceId *string `pulumi:"voiceId"`
}

type BotState struct {
	// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
	AbortStatement BotAbortStatementPtrInput
	Arn            pulumi.StringPtrInput
	// Checksum identifying the version of the bot that was created. The checksum is not
	// included as an argument because the resource will add it automatically when updating the bot.
	Checksum pulumi.StringPtrInput
	// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
	ChildDirected pulumi.BoolPtrInput
	// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
	ClarificationPrompt BotClarificationPromptPtrInput
	// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
	CreateVersion pulumi.BoolPtrInput
	// The date when the bot version was created.
	CreatedDate pulumi.StringPtrInput
	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description pulumi.StringPtrInput
	// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
	DetectSentiment pulumi.BoolPtrInput
	// Set to `true` to enable access to natural language understanding improvements. When you set the `enableModelImprovements` parameter to true you can use the `nluIntentConfidenceThreshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enableModelImprovements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
	EnableModelImprovements pulumi.BoolPtrInput
	// If status is FAILED, Amazon Lex provides the reason that it failed to build the bot.
	FailureReason pulumi.StringPtrInput
	// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
	IdleSessionTtlInSeconds pulumi.IntPtrInput
	// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
	Intents BotIntentArrayInput
	// The date when the $LATEST version of this bot was updated.
	LastUpdatedDate pulumi.StringPtrInput
	// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
	Locale pulumi.StringPtrInput
	// The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
	Name pulumi.StringPtrInput
	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enableModelImprovements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
	NluIntentConfidenceThreshold pulumi.Float64PtrInput
	// If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
	ProcessBehavior pulumi.StringPtrInput
	// When you send a request to create or update a bot, Amazon Lex sets the status response
	// element to BUILDING. After Amazon Lex builds the bot, it sets status to READY. If Amazon Lex can't
	// build the bot, it sets status to FAILED. Amazon Lex returns the reason for the failure in the
	// failureReason response element.
	Status pulumi.StringPtrInput
	// The version of the bot.
	Version pulumi.StringPtrInput
	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
	VoiceId pulumi.StringPtrInput
}

func (BotState) ElementType() reflect.Type {
	return reflect.TypeOf((*botState)(nil)).Elem()
}

type botArgs struct {
	// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
	AbortStatement BotAbortStatement `pulumi:"abortStatement"`
	// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
	ChildDirected bool `pulumi:"childDirected"`
	// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
	ClarificationPrompt *BotClarificationPrompt `pulumi:"clarificationPrompt"`
	// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
	CreateVersion *bool `pulumi:"createVersion"`
	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description *string `pulumi:"description"`
	// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
	DetectSentiment *bool `pulumi:"detectSentiment"`
	// Set to `true` to enable access to natural language understanding improvements. When you set the `enableModelImprovements` parameter to true you can use the `nluIntentConfidenceThreshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enableModelImprovements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
	EnableModelImprovements *bool `pulumi:"enableModelImprovements"`
	// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
	IdleSessionTtlInSeconds *int `pulumi:"idleSessionTtlInSeconds"`
	// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
	Intents []BotIntent `pulumi:"intents"`
	// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
	Locale *string `pulumi:"locale"`
	// The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
	Name *string `pulumi:"name"`
	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enableModelImprovements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
	NluIntentConfidenceThreshold *float64 `pulumi:"nluIntentConfidenceThreshold"`
	// If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
	ProcessBehavior *string `pulumi:"processBehavior"`
	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
	VoiceId *string `pulumi:"voiceId"`
}

// The set of arguments for constructing a Bot resource.
type BotArgs struct {
	// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
	AbortStatement BotAbortStatementInput
	// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
	ChildDirected pulumi.BoolInput
	// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
	ClarificationPrompt BotClarificationPromptPtrInput
	// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
	CreateVersion pulumi.BoolPtrInput
	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description pulumi.StringPtrInput
	// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
	DetectSentiment pulumi.BoolPtrInput
	// Set to `true` to enable access to natural language understanding improvements. When you set the `enableModelImprovements` parameter to true you can use the `nluIntentConfidenceThreshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enableModelImprovements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
	EnableModelImprovements pulumi.BoolPtrInput
	// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
	IdleSessionTtlInSeconds pulumi.IntPtrInput
	// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
	Intents BotIntentArrayInput
	// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
	Locale pulumi.StringPtrInput
	// The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
	Name pulumi.StringPtrInput
	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enableModelImprovements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
	NluIntentConfidenceThreshold pulumi.Float64PtrInput
	// If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
	ProcessBehavior pulumi.StringPtrInput
	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
	VoiceId pulumi.StringPtrInput
}

func (BotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*botArgs)(nil)).Elem()
}

type BotInput interface {
	pulumi.Input

	ToBotOutput() BotOutput
	ToBotOutputWithContext(ctx context.Context) BotOutput
}

func (*Bot) ElementType() reflect.Type {
	return reflect.TypeOf((**Bot)(nil)).Elem()
}

func (i *Bot) ToBotOutput() BotOutput {
	return i.ToBotOutputWithContext(context.Background())
}

func (i *Bot) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotOutput)
}

// BotArrayInput is an input type that accepts BotArray and BotArrayOutput values.
// You can construct a concrete instance of `BotArrayInput` via:
//
//	BotArray{ BotArgs{...} }
type BotArrayInput interface {
	pulumi.Input

	ToBotArrayOutput() BotArrayOutput
	ToBotArrayOutputWithContext(context.Context) BotArrayOutput
}

type BotArray []BotInput

func (BotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bot)(nil)).Elem()
}

func (i BotArray) ToBotArrayOutput() BotArrayOutput {
	return i.ToBotArrayOutputWithContext(context.Background())
}

func (i BotArray) ToBotArrayOutputWithContext(ctx context.Context) BotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotArrayOutput)
}

// BotMapInput is an input type that accepts BotMap and BotMapOutput values.
// You can construct a concrete instance of `BotMapInput` via:
//
//	BotMap{ "key": BotArgs{...} }
type BotMapInput interface {
	pulumi.Input

	ToBotMapOutput() BotMapOutput
	ToBotMapOutputWithContext(context.Context) BotMapOutput
}

type BotMap map[string]BotInput

func (BotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bot)(nil)).Elem()
}

func (i BotMap) ToBotMapOutput() BotMapOutput {
	return i.ToBotMapOutputWithContext(context.Background())
}

func (i BotMap) ToBotMapOutputWithContext(ctx context.Context) BotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotMapOutput)
}

type BotOutput struct{ *pulumi.OutputState }

func (BotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bot)(nil)).Elem()
}

func (o BotOutput) ToBotOutput() BotOutput {
	return o
}

func (o BotOutput) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return o
}

// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
func (o BotOutput) AbortStatement() BotAbortStatementOutput {
	return o.ApplyT(func(v *Bot) BotAbortStatementOutput { return v.AbortStatement }).(BotAbortStatementOutput)
}

func (o BotOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Checksum identifying the version of the bot that was created. The checksum is not
// included as an argument because the resource will add it automatically when updating the bot.
func (o BotOutput) Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Checksum }).(pulumi.StringOutput)
}

// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
func (o BotOutput) ChildDirected() pulumi.BoolOutput {
	return o.ApplyT(func(v *Bot) pulumi.BoolOutput { return v.ChildDirected }).(pulumi.BoolOutput)
}

// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
func (o BotOutput) ClarificationPrompt() BotClarificationPromptPtrOutput {
	return o.ApplyT(func(v *Bot) BotClarificationPromptPtrOutput { return v.ClarificationPrompt }).(BotClarificationPromptPtrOutput)
}

// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
func (o BotOutput) CreateVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.BoolPtrOutput { return v.CreateVersion }).(pulumi.BoolPtrOutput)
}

// The date when the bot version was created.
func (o BotOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// A description of the bot. Must be less than or equal to 200 characters in length.
func (o BotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
func (o BotOutput) DetectSentiment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.BoolPtrOutput { return v.DetectSentiment }).(pulumi.BoolPtrOutput)
}

// Set to `true` to enable access to natural language understanding improvements. When you set the `enableModelImprovements` parameter to true you can use the `nluIntentConfidenceThreshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enableModelImprovements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
func (o BotOutput) EnableModelImprovements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.BoolPtrOutput { return v.EnableModelImprovements }).(pulumi.BoolPtrOutput)
}

// If status is FAILED, Amazon Lex provides the reason that it failed to build the bot.
func (o BotOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.FailureReason }).(pulumi.StringOutput)
}

// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
func (o BotOutput) IdleSessionTtlInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.IntPtrOutput { return v.IdleSessionTtlInSeconds }).(pulumi.IntPtrOutput)
}

// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
func (o BotOutput) Intents() BotIntentArrayOutput {
	return o.ApplyT(func(v *Bot) BotIntentArrayOutput { return v.Intents }).(BotIntentArrayOutput)
}

// The date when the $LATEST version of this bot was updated.
func (o BotOutput) LastUpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.LastUpdatedDate }).(pulumi.StringOutput)
}

// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
func (o BotOutput) Locale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringPtrOutput { return v.Locale }).(pulumi.StringPtrOutput)
}

// The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
func (o BotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enableModelImprovements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
func (o BotOutput) NluIntentConfidenceThreshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.Float64PtrOutput { return v.NluIntentConfidenceThreshold }).(pulumi.Float64PtrOutput)
}

// If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
func (o BotOutput) ProcessBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringPtrOutput { return v.ProcessBehavior }).(pulumi.StringPtrOutput)
}

// When you send a request to create or update a bot, Amazon Lex sets the status response
// element to BUILDING. After Amazon Lex builds the bot, it sets status to READY. If Amazon Lex can't
// build the bot, it sets status to FAILED. Amazon Lex returns the reason for the failure in the
// failureReason response element.
func (o BotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The version of the bot.
func (o BotOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
func (o BotOutput) VoiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.VoiceId }).(pulumi.StringOutput)
}

type BotArrayOutput struct{ *pulumi.OutputState }

func (BotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bot)(nil)).Elem()
}

func (o BotArrayOutput) ToBotArrayOutput() BotArrayOutput {
	return o
}

func (o BotArrayOutput) ToBotArrayOutputWithContext(ctx context.Context) BotArrayOutput {
	return o
}

func (o BotArrayOutput) Index(i pulumi.IntInput) BotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bot {
		return vs[0].([]*Bot)[vs[1].(int)]
	}).(BotOutput)
}

type BotMapOutput struct{ *pulumi.OutputState }

func (BotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bot)(nil)).Elem()
}

func (o BotMapOutput) ToBotMapOutput() BotMapOutput {
	return o
}

func (o BotMapOutput) ToBotMapOutputWithContext(ctx context.Context) BotMapOutput {
	return o
}

func (o BotMapOutput) MapIndex(k pulumi.StringInput) BotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bot {
		return vs[0].(map[string]*Bot)[vs[1].(string)]
	}).(BotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BotInput)(nil)).Elem(), &Bot{})
	pulumi.RegisterInputType(reflect.TypeOf((*BotArrayInput)(nil)).Elem(), BotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BotMapInput)(nil)).Elem(), BotMap{})
	pulumi.RegisterOutputType(BotOutput{})
	pulumi.RegisterOutputType(BotArrayOutput{})
	pulumi.RegisterOutputType(BotMapOutput{})
}
