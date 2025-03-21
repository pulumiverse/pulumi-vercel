// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAccessGroupProject(args: GetAccessGroupProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessGroupProjectResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vercel:index/getAccessGroupProject:getAccessGroupProject", {
        "accessGroupId": args.accessGroupId,
        "projectId": args.projectId,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessGroupProject.
 */
export interface GetAccessGroupProjectArgs {
    accessGroupId: string;
    projectId: string;
    teamId?: string;
}

/**
 * A collection of values returned by getAccessGroupProject.
 */
export interface GetAccessGroupProjectResult {
    readonly accessGroupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId: string;
    readonly role: string;
    readonly teamId: string;
}
export function getAccessGroupProjectOutput(args: GetAccessGroupProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessGroupProjectResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vercel:index/getAccessGroupProject:getAccessGroupProject", {
        "accessGroupId": args.accessGroupId,
        "projectId": args.projectId,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessGroupProject.
 */
export interface GetAccessGroupProjectOutputArgs {
    accessGroupId: pulumi.Input<string>;
    projectId: pulumi.Input<string>;
    teamId?: pulumi.Input<string>;
}
