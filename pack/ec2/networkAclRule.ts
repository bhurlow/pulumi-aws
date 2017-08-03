// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class NetworkAclRule extends lumi.NamedResource implements NetworkAclRuleArgs {
    public readonly cidrBlock?: string;
    public readonly egress?: boolean;
    public readonly fromPort?: number;
    public readonly icmpCode?: string;
    public readonly icmpType?: string;
    public readonly ipv6CidrBlock?: string;
    public readonly networkAclId: string;
    public readonly protocol: string;
    public readonly ruleAction: string;
    public readonly ruleNumber: number;
    public readonly toPort?: number;

    public static get(id: lumi.ID): NetworkAclRule {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): NetworkAclRule[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: NetworkAclRuleArgs) {
        super(name);
        this.cidrBlock = <any>args.cidrBlock;
        this.egress = <any>args.egress;
        this.fromPort = <any>args.fromPort;
        this.icmpCode = <any>args.icmpCode;
        this.icmpType = <any>args.icmpType;
        this.ipv6CidrBlock = <any>args.ipv6CidrBlock;
        if (lumirt.defaultIfComputed(args.networkAclId, "") === undefined) {
            throw new Error("Property argument 'networkAclId' is required, but was missing");
        }
        this.networkAclId = <any>args.networkAclId;
        if (lumirt.defaultIfComputed(args.protocol, "") === undefined) {
            throw new Error("Property argument 'protocol' is required, but was missing");
        }
        this.protocol = <any>args.protocol;
        if (lumirt.defaultIfComputed(args.ruleAction, "") === undefined) {
            throw new Error("Property argument 'ruleAction' is required, but was missing");
        }
        this.ruleAction = <any>args.ruleAction;
        if (lumirt.defaultIfComputed(args.ruleNumber, "") === undefined) {
            throw new Error("Property argument 'ruleNumber' is required, but was missing");
        }
        this.ruleNumber = <any>args.ruleNumber;
        this.toPort = <any>args.toPort;
    }
}

export interface NetworkAclRuleArgs {
    readonly cidrBlock?: string;
    readonly egress?: boolean;
    readonly fromPort?: number;
    readonly icmpCode?: string;
    readonly icmpType?: string;
    readonly ipv6CidrBlock?: string;
    readonly networkAclId: string;
    readonly protocol: string;
    readonly ruleAction: string;
    readonly ruleNumber: number;
    readonly toPort?: number;
}
