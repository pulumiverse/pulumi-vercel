// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getEdgeConfigItem(args: GetEdgeConfigItemArgs, opts?: pulumi.InvokeOptions): Promise<GetEdgeConfigItemResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vercel:index/getEdgeConfigItem:getEdgeConfigItem", {
        "id": args.id,
        "key": args.key,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getEdgeConfigItem.
 */
export interface GetEdgeConfigItemArgs {
    id: string;
    key: string;
    teamId?: string;
}

/**
 * A collection of values returned by getEdgeConfigItem.
 */
export interface GetEdgeConfigItemResult {
    readonly id: string;
    readonly key: string;
    readonly teamId: string;
    readonly value: string;
}
export function getEdgeConfigItemOutput(args: GetEdgeConfigItemOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEdgeConfigItemResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vercel:index/getEdgeConfigItem:getEdgeConfigItem", {
        "id": args.id,
        "key": args.key,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getEdgeConfigItem.
 */
export interface GetEdgeConfigItemOutputArgs {
    id: pulumi.Input<string>;
    key: pulumi.Input<string>;
    teamId?: pulumi.Input<string>;
}
