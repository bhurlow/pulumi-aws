// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Rule extends lumi.NamedResource implements RuleArgs {
    public /*out*/ readonly arn: string;
    public readonly description?: string;
    public readonly inputParameters?: string;
    public readonly maximumExecutionFrequency?: string;
    public readonly ruleName: string;
    public /*out*/ readonly ruleId: string;
    public readonly scope?: { complianceResourceId?: string, complianceResourceTypes?: string[], tagKey?: string, tagValue?: string }[];
    public readonly source: { owner: string, sourceDetail?: { eventSource?: string, maximumExecutionFrequency?: string, messageType?: string }[], sourceIdentifier: string }[];

    public static get(id: lumi.ID): Rule {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Rule[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: RuleArgs) {
        super(name);
        this.description = <any>args.description;
        this.inputParameters = <any>args.inputParameters;
        this.maximumExecutionFrequency = <any>args.maximumExecutionFrequency;
        this.ruleName = <any>args.ruleName;
        this.scope = <any>args.scope;
        if (lumirt.defaultIfComputed(args.source, "") === undefined) {
            throw new Error("Property argument 'source' is required, but was missing");
        }
        this.source = <any>args.source;
    }
}

export interface RuleArgs {
    readonly description?: string;
    readonly inputParameters?: string;
    readonly maximumExecutionFrequency?: string;
    readonly ruleName?: string;
    readonly scope?: { complianceResourceId?: string, complianceResourceTypes?: string[], tagKey?: string, tagValue?: string }[];
    readonly source: { owner: string, sourceDetail?: { eventSource?: string, maximumExecutionFrequency?: string, messageType?: string }[], sourceIdentifier: string }[];
}
