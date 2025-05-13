---
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Vercel Provider
meta_desc: Provides an overview on how to configure the Pulumi Vercel provider.
layout: package
---
## Installation

The Vercel provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/vercel`](https://www.npmjs.com/package/@pulumi/vercel)
* Python: [`pulumi-vercel`](https://pypi.org/project/pulumi-vercel/)
* Go: [`github.com/pulumiverse/pulumi-vercel/sdk/v2/go/vercel`](https://github.com/pulumi/pulumi-vercel)
* .NET: [`Pulumi.Vercel`](https://www.nuget.org/packages/Pulumi.Vercel)
* Java: [`com.pulumi/vercel`](https://central.sonatype.com/artifact/com.pulumi/vercel)
## Overview

The Vercel provider is used to interact with resources supported by Vercel.
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

{{% /choosable %}}
{{% choosable language python %}}

{{% /choosable %}}
{{% choosable language csharp %}}

{{% /choosable %}}
{{% choosable language go %}}

{{% /choosable %}}
{{% choosable language yaml %}}

{{% /choosable %}}
{{% choosable language java %}}

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `apiToken` (String, Sensitive) The Vercel API Token to use. This can also be specified with the `VERCEL_API_TOKEN` shell environment variable. Tokens can be created from your [Vercel settings](https://vercel.com/account/tokens).
- `team` (String) The default Vercel Team to use when creating resources or reading functions. This can be provided as either a team slug, or team ID. The slug and ID are both available from the Team Settings page in the Vercel dashboard.