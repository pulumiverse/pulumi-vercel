// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("vercel");

/**
 * The Vercel API Token to use. This can also be specified with the `VERCEL_API_TOKEN` shell environment variable. Tokens
 * can be created from your [Vercel settings](https://vercel.com/account/tokens).
 */
export declare const apiToken: string | undefined;
Object.defineProperty(exports, "apiToken", {
    get() {
        return __config.get("apiToken");
    },
    enumerable: true,
});

/**
 * The default Vercel Team to use when creating resources or reading data sources. This can be provided as either a team
 * slug, or team ID. The slug and ID are both available from the Team Settings page in the Vercel dashboard.
 */
export declare const team: string | undefined;
Object.defineProperty(exports, "team", {
    get() {
        return __config.get("team");
    },
    enumerable: true,
});

