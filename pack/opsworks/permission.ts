// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Permission extends lumi.NamedResource implements PermissionArgs {
    public readonly allowSsh: boolean;
    public readonly allowSudo: boolean;
    public /*out*/ readonly permissionId: string;
    public readonly level: string;
    public readonly stackId: string;
    public readonly userArn: string;

    public static get(id: lumi.ID): Permission {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Permission[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: PermissionArgs) {
        super(name);
        this.allowSsh = <any>args.allowSsh;
        this.allowSudo = <any>args.allowSudo;
        this.level = <any>args.level;
        this.stackId = <any>args.stackId;
        if (lumirt.defaultIfComputed(args.userArn, "") === undefined) {
            throw new Error("Property argument 'userArn' is required, but was missing");
        }
        this.userArn = <any>args.userArn;
    }
}

export interface PermissionArgs {
    readonly allowSsh?: boolean;
    readonly allowSudo?: boolean;
    readonly level?: string;
    readonly stackId?: string;
    readonly userArn: string;
}
