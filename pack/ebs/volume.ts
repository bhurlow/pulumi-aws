// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Volume extends lumi.NamedResource implements VolumeArgs {
    public readonly availabilityZone: string;
    public readonly encrypted: boolean;
    public readonly iops: number;
    public readonly kmsKeyId: string;
    public readonly size: number;
    public readonly snapshotId: string;
    public readonly tags?: {[key: string]: any};
    public readonly type: string;

    public static get(id: lumi.ID): Volume {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Volume[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: VolumeArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.availabilityZone, "") === undefined) {
            throw new Error("Property argument 'availabilityZone' is required, but was missing");
        }
        this.availabilityZone = <any>args.availabilityZone;
        this.encrypted = <any>args.encrypted;
        this.iops = <any>args.iops;
        this.kmsKeyId = <any>args.kmsKeyId;
        this.size = <any>args.size;
        this.snapshotId = <any>args.snapshotId;
        this.tags = <any>args.tags;
        this.type = <any>args.type;
    }
}

export interface VolumeArgs {
    readonly availabilityZone: string;
    readonly encrypted?: boolean;
    readonly iops?: number;
    readonly kmsKeyId?: string;
    readonly size?: number;
    readonly snapshotId?: string;
    readonly tags?: {[key: string]: any};
    readonly type?: string;
}
