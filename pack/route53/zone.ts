// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Zone extends lumi.NamedResource implements ZoneArgs {
    public readonly comment?: string;
    public readonly delegationSetId?: string;
    public readonly forceDestroy?: boolean;
    public readonly zoneName: string;
    public /*out*/ readonly nameServers: string[];
    public readonly tags?: {[key: string]: any};
    public readonly vpcId?: string;
    public readonly vpcRegion: string;
    public /*out*/ readonly zoneId: string;

    public static get(id: lumi.ID): Zone {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Zone[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: ZoneArgs) {
        super(name);
        this.comment = <any>args.comment;
        this.delegationSetId = <any>args.delegationSetId;
        this.forceDestroy = <any>args.forceDestroy;
        this.zoneName = <any>args.zoneName;
        this.tags = <any>args.tags;
        this.vpcId = <any>args.vpcId;
        this.vpcRegion = <any>args.vpcRegion;
    }
}

export interface ZoneArgs {
    readonly comment?: string;
    readonly delegationSetId?: string;
    readonly forceDestroy?: boolean;
    readonly zoneName?: string;
    readonly tags?: {[key: string]: any};
    readonly vpcId?: string;
    readonly vpcRegion?: string;
}
