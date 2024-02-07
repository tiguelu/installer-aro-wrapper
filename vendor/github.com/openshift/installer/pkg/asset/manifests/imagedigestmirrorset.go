package manifests

import (
	"path/filepath"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	apicfgv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
)

var imageDigestMirrorSetFilename = "image-digest-mirror-set.yaml"

// ImageDigestMirrorSet generates the image-digest-mirror-set.yaml file.
type ImageDigestMirrorSet struct {
	File *asset.File
}

var _ asset.WritableAsset = (*ImageDigestMirrorSet)(nil)

// Name returns a human-friendly name for the asset.
func (*ImageDigestMirrorSet) Name() string {
	return "Image Digest Mirror Set"
}

// Dependencies returns all of the dependencies directly needed to generate
// the asset.
func (*ImageDigestMirrorSet) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.InstallConfig{},
	}
}

// Generate generates the ImageDigestMirrorSet config and its CR.
func (p *ImageDigestMirrorSet) Generate(dependencies asset.Parents) error {
	installconfig := &installconfig.InstallConfig{}
	dependencies.Get(installconfig)

	if len(installconfig.Config.ImageDigestSources) > 0 {
		policy := apicfgv1.ImageDigestMirrorSet{
			TypeMeta: metav1.TypeMeta{
				APIVersion: apicfgv1.SchemeGroupVersion.String(),
				Kind:       "ImageDigestMirrorSet",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "image-digest-mirror",
				// not namespaced
			},
		}

		policy.Spec.ImageDigestMirrors = make([]apicfgv1.ImageDigestMirrors, len(installconfig.Config.ImageDigestSources))
		for gidx, group := range installconfig.Config.ImageDigestSources {
			mirrors := []apicfgv1.ImageMirror{}
			for _, m := range group.Mirrors {
				mirrors = append(mirrors, apicfgv1.ImageMirror(m))
			}
			policy.Spec.ImageDigestMirrors[gidx] = apicfgv1.ImageDigestMirrors{Source: group.Source, Mirrors: mirrors}
		}

		policyData, err := yaml.Marshal(policy)
		if err != nil {
			return errors.Wrapf(err, "failed to marshal ImageDigestMirrorSet")
		}
		p.File = &asset.File{
			Filename: filepath.Join(manifestDir, imageDigestMirrorSetFilename),
			Data:     policyData,
		}
	}
	return nil
}

// Files returns the files generated by the asset.
func (p *ImageDigestMirrorSet) Files() []*asset.File {
	if p.File == nil {
		return nil
	}
	return []*asset.File{p.File}
}

// Load loads the already-rendered files back from disk.
func (p *ImageDigestMirrorSet) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}
