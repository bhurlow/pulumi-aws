// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class ReplicationTask extends lumi.NamedResource implements ReplicationTaskArgs {
    public readonly cdcStartTime?: string;
    public readonly migrationType: string;
    public readonly replicationInstanceArn: string;
    public /*out*/ readonly replicationTaskArn: string;
    public readonly replicationTaskId: string;
    public readonly replicationTaskSettings?: string;
    public readonly sourceEndpointArn: string;
    public readonly tableMappings: string;
    public readonly tags?: {[key: string]: any};
    public readonly targetEndpointArn: string;

    public static get(id: lumi.ID): ReplicationTask {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): ReplicationTask[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: ReplicationTaskArgs) {
        super(name);
        this.cdcStartTime = <any>args.cdcStartTime;
        if (lumirt.defaultIfComputed(args.migrationType, "") === undefined) {
            throw new Error("Property argument 'migrationType' is required, but was missing");
        }
        this.migrationType = <any>args.migrationType;
        if (lumirt.defaultIfComputed(args.replicationInstanceArn, "") === undefined) {
            throw new Error("Property argument 'replicationInstanceArn' is required, but was missing");
        }
        this.replicationInstanceArn = <any>args.replicationInstanceArn;
        if (lumirt.defaultIfComputed(args.replicationTaskId, "") === undefined) {
            throw new Error("Property argument 'replicationTaskId' is required, but was missing");
        }
        this.replicationTaskId = <any>args.replicationTaskId;
        this.replicationTaskSettings = <any>args.replicationTaskSettings;
        if (lumirt.defaultIfComputed(args.sourceEndpointArn, "") === undefined) {
            throw new Error("Property argument 'sourceEndpointArn' is required, but was missing");
        }
        this.sourceEndpointArn = <any>args.sourceEndpointArn;
        if (lumirt.defaultIfComputed(args.tableMappings, "") === undefined) {
            throw new Error("Property argument 'tableMappings' is required, but was missing");
        }
        this.tableMappings = <any>args.tableMappings;
        this.tags = <any>args.tags;
        if (lumirt.defaultIfComputed(args.targetEndpointArn, "") === undefined) {
            throw new Error("Property argument 'targetEndpointArn' is required, but was missing");
        }
        this.targetEndpointArn = <any>args.targetEndpointArn;
    }
}

export interface ReplicationTaskArgs {
    readonly cdcStartTime?: string;
    readonly migrationType: string;
    readonly replicationInstanceArn: string;
    readonly replicationTaskId: string;
    readonly replicationTaskSettings?: string;
    readonly sourceEndpointArn: string;
    readonly tableMappings: string;
    readonly tags?: {[key: string]: any};
    readonly targetEndpointArn: string;
}
