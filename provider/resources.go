// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vercel

import (
	_ "embed" // nolint: golint
	"fmt"
	"path/filepath"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumiverse/pulumi-vercel/provider/pkg/version"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/vercel/terraform-provider-vercel/vercel"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "vercel"
	// modules:
	mainMod   = "index" // the vercel module
	publisher = "pulumiverse"
)

//go:embed cmd/pulumi-resource-vercel/bridge-metadata.json
var bridgeMetadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pf.ShimProvider(vercel.New())

	caser := cases.Title(language.English)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              mainPkg,
		DisplayName:       caser.String(mainPkg),
		Publisher:         caser.String(publisher),
		Version:           version.Version,
		PluginDownloadURL: "github://api.github.com/pulumiverse",
		Keywords:          []string{"pulumi", "vercel", "category/cloud"},
		License:           "Apache-2.0",
		GitHubOrg:         "vercel",
		Repository:        "https://github.com/pulumiverse/pulumi-vercel",
		Resources: map[string]*tfbridge.ResourceInfo{
			"vercel_alias": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Alias"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"alias": {
						CSharpName: "DeploymentAlias",
					},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: fmt.Sprintf("@%s/%s", publisher, mainPkg),
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: fmt.Sprintf("%s_%s", publisher, mainPkg),
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/%s/pulumi-%s/sdk/", publisher, mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: caser.String(publisher),
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumiverse",
		},
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
	}

	// These are new API's that you may opt to use to automatically compute resource tokens,
	// and apply auto aliasing for full backwards compatibility.
	// For more information, please reference
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("vercel_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
