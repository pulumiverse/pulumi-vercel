package vercel

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumiverse/pulumi-vercel/provider/v3/pkg/version"
)

func TestDnsRecordSrvTransform_EmptyObjectBecomesNull(t *testing.T) {
	if version.Version == "" {
		version.Version = "3.0.0"
	}

	p := Provider()
	r := p.Resources["vercel_dns_record"]
	if r == nil {
		t.Fatalf("expected vercel_dns_record resource mapping")
	}
	f := r.Fields["srv"]
	if f == nil || f.Transform == nil {
		t.Fatalf("expected srv field transform")
	}

	out, err := f.Transform(resource.NewObjectProperty(resource.PropertyMap{}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.IsNull() {
		t.Fatalf("expected null property, got %#v", out)
	}
}
