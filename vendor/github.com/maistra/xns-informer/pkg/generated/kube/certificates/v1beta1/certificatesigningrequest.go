// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/certificates/v1beta1"
	informers "k8s.io/client-go/informers/certificates/v1beta1"
	listers "k8s.io/client-go/listers/certificates/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type certificateSigningRequestInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.CertificateSigningRequestInformer = &certificateSigningRequestInformer{}

func NewCertificateSigningRequestInformer(f xnsinformers.SharedInformerFactory) informers.CertificateSigningRequestInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("certificatesigningrequests")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1beta1.CertificateSigningRequest{},
		&v1beta1.CertificateSigningRequestList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      true,
		ListWatchConverter: converter,
	})

	return &certificateSigningRequestInformer{informer: informer.Informer()}
}

func (i *certificateSigningRequestInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *certificateSigningRequestInformer) Lister() listers.CertificateSigningRequestLister {
	return listers.NewCertificateSigningRequestLister(i.informer.GetIndexer())
}
