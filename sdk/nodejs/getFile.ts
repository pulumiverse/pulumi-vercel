// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getFile(args: GetFileArgs, opts?: pulumi.InvokeOptions): Promise<GetFileResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vercel:index/getFile:getFile", {
        "path": args.path,
    }, opts);
}

/**
 * A collection of arguments for invoking getFile.
 */
export interface GetFileArgs {
    path: string;
}

/**
 * A collection of values returned by getFile.
 */
export interface GetFileResult {
    readonly file: {[key: string]: string};
    readonly id: string;
    readonly path: string;
}
export function getFileOutput(args: GetFileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFileResult> {
    return pulumi.output(args).apply((a: any) => getFile(a, opts))
}

/**
 * A collection of arguments for invoking getFile.
 */
export interface GetFileOutputArgs {
    path: pulumi.Input<string>;
}