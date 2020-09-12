package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	apixv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Example Controller:", func() {

	const (
		timeout  = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When creating a CRD", func() {

		ctx := context.Background()

		It("should get the CRD", func() {
			By("Creating CRD")
			testCRD := &apixv1beta1.CustomResourceDefinition{
				TypeMeta: metav1.TypeMeta{
					Kind:       "CustomResourceDefinition",
					APIVersion: "apiextensions.k8s.io/v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "tests.app.example.org",
				},
				Spec: apixv1beta1.CustomResourceDefinitionSpec{
					Group: "app.example.org",
					Versions: []apixv1beta1.CustomResourceDefinitionVersion{{
						Name:    "v1alpha1",
						Served:  true,
						Storage: true,
					}},
					Names: apixv1beta1.CustomResourceDefinitionNames{
						Plural: "tests",
						Kind:   "Test",
					},
					Scope: apixv1beta1.ClusterScoped,
				},
			}
			Expect(k8sClient.Create(ctx, testCRD)).Should(Succeed())

			testCRDLookupKey := client.ObjectKey{Name: "tests.app.example.org", Namespace: "default"}
			createdTestCRD := &apixv1beta1.CustomResourceDefinition{}

			By("Verifying Test CRD")
			// Retry getting newly created Test CRD
			Eventually(func() bool {
				// FIXME: `k8sClient` seems to be not working
				err := k8sClient.Get(ctx, testCRDLookupKey, createdTestCRD)
				//err := k8sClient2.Get(ctx, testCRDLookupKey, createdTestCRD)
				return err == nil
			}, timeout, interval).Should(BeTrue())

		})
	})
})
