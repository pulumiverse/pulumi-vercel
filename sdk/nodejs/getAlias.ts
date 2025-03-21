// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAlias(args: GetAliasArgs, opts?: pulumi.InvokeOptions): Promise<GetAliasResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vercel:index/getAlias:getAlias", {
        "alias": args.alias,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAlias.
 */
export interface GetAliasArgs {
    alias: string;
    teamId?: string;
}

/**
 * A collection of values returned by getAlias.
 */
export interface GetAliasResult {
    readonly alias: string;
    readonly deploymentId: string;
    readonly id: string;
    readonly teamId: string;
}
export function getAliasOutput(args: GetAliasOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAliasResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vercel:index/getAlias:getAlias", {
        "alias": args.alias,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAlias.
 */
export interface GetAliasOutputArgs {
    alias: pulumi.Input<string>;
    teamId?: pulumi.Input<string>;
}
