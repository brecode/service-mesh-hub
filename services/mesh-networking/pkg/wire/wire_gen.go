// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"context"

	istio_networking "github.com/solo-io/mesh-projects/pkg/clients/istio/networking"
	"github.com/solo-io/mesh-projects/pkg/clients/istio/security"
	kubernetes_core "github.com/solo-io/mesh-projects/pkg/clients/kubernetes/core"
	zephyr_discovery "github.com/solo-io/mesh-projects/pkg/clients/zephyr/discovery"
	zephyr_networking "github.com/solo-io/mesh-projects/pkg/clients/zephyr/networking"
	zephyr_security "github.com/solo-io/mesh-projects/pkg/clients/zephyr/security"
	"github.com/solo-io/mesh-projects/pkg/security/certgen"
	mc_wire "github.com/solo-io/mesh-projects/services/common/multicluster/wire"
	csr_generator "github.com/solo-io/mesh-projects/services/csr-agent/pkg/csr-generator"
	"github.com/solo-io/mesh-projects/services/mesh-discovery/pkg/multicluster/controllers"
	access_control_poilcy "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/access/access-control-poilcy"
	networking_multicluster "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/multicluster"
	controller_factories "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/multicluster/controllers"
	traffic_policy_translator "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/routing/traffic-policy-translator"
	istio_translator "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/routing/traffic-policy-translator/istio-translator"
	"github.com/solo-io/mesh-projects/services/mesh-networking/pkg/routing/traffic-policy-translator/preprocess"
	cert_manager "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/security/cert-manager"
	cert_signer "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/security/cert-signer"
	group_validation "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/validation"
)

// Injectors from wire.go:

func InitializeMeshNetworking(ctx context.Context) (MeshNetworkingContext, error) {
	config, err := mc_wire.LocalKubeConfigProvider()
	if err != nil {
		return MeshNetworkingContext{}, err
	}
	asyncManager, err := mc_wire.LocalManagerProvider(ctx, config)
	if err != nil {
		return MeshNetworkingContext{}, err
	}
	asyncManagerController := mc_wire.AsyncManagerControllerProvider(ctx, asyncManager)
	asyncManagerStartOptionsFunc := mc_wire.LocalManagerStarterProvider(asyncManagerController)
	multiClusterDependencies := mc_wire.MulticlusterDependenciesProvider(ctx, asyncManager, asyncManagerController, asyncManagerStartOptionsFunc)
	meshGroupCSRControllerFactory := controller_factories.NewMeshGroupCSRControllerFactory()
	controllerFactories := NewControllerFactories(meshGroupCSRControllerFactory)
	meshGroupCSRClientFactory := zephyr_security.MeshGroupCSRClientFactoryProvider()
	clientFactories := NewClientFactories(meshGroupCSRClientFactory)
	client := mc_wire.DynamicClientProvider(asyncManager)
	secretsClient := kubernetes_core.NewSecretsClient(client)
	meshGroupClient := zephyr_networking.NewMeshGroupClient(client)
	meshGroupCertClient := cert_signer.NewMeshGroupCertClient(secretsClient, meshGroupClient)
	signer := certgen.NewSigner()
	meshGroupCSRDataSourceFactory := csr_generator.NewMeshGroupCSRDataSourceFactory()
	asyncManagerHandler, err := networking_multicluster.NewMeshNetworkingClusterHandler(asyncManager, controllerFactories, clientFactories, meshGroupCertClient, signer, meshGroupCSRDataSourceFactory)
	if err != nil {
		return MeshNetworkingContext{}, err
	}
	meshServiceClient := zephyr_discovery.NewMeshServiceClient(client)
	meshServiceSelector := preprocess.NewMeshServiceSelector(meshServiceClient)
	meshClient := zephyr_discovery.NewMeshClient(client)
	trafficPolicyClient := zephyr_networking.NewTrafficPolicyClient(client)
	trafficPolicyMerger := preprocess.NewTrafficPolicyMerger(meshServiceSelector, meshClient, trafficPolicyClient)
	trafficPolicyValidator := preprocess.NewTrafficPolicyValidator(meshServiceClient, meshServiceSelector)
	trafficPolicyPreprocessor := preprocess.NewTrafficPolicyPreprocessor(meshServiceSelector, trafficPolicyMerger, trafficPolicyValidator)
	dynamicClientGetter := mc_wire.DynamicClientGetterProvider(asyncManagerController)
	virtualServiceClientFactory := istio_networking.VirtualServiceClientFactoryProvider()
	istioTranslator := istio_translator.NewIstioTrafficPolicyTranslator(dynamicClientGetter, meshClient, meshServiceClient, meshServiceSelector, virtualServiceClientFactory)
	v := TrafficPolicyMeshTranslatorsProvider(istioTranslator)
	trafficPolicyController, err := LocalTrafficPolicyControllerProvider(asyncManager)
	if err != nil {
		return MeshNetworkingContext{}, err
	}
	meshServiceController, err := LocalMeshServiceControllerProvider(asyncManager)
	if err != nil {
		return MeshNetworkingContext{}, err
	}
	trafficPolicyTranslator := traffic_policy_translator.NewTrafficPolicyTranslator(ctx, trafficPolicyPreprocessor, v, meshClient, meshServiceClient, trafficPolicyClient, trafficPolicyController, meshServiceController)
	meshWorkloadControllerFactory := controllers.NewMeshWorkloadControllerFactory()
	meshServiceControllerFactory := controllers.NewMeshServiceControllerFactory()
	meshGroupControllerFactory := controller_factories.NewMeshGroupControllerFactory()
	groupMeshFinder := group_validation.NewGroupMeshFinder(meshClient)
	meshNetworkingSnapshotValidator := group_validation.NewMeshGroupValidator(groupMeshFinder, meshGroupClient)
	istioCertConfigProducer := cert_manager.NewIstioCertConfigProducer()
	meshGroupCertificateManager := cert_manager.NewMeshGroupCsrProcessor(dynamicClientGetter, meshClient, groupMeshFinder, meshGroupCSRClientFactory, istioCertConfigProducer)
	groupMgcsrSnapshotListener := cert_manager.NewGroupMgcsrSnapshotListener(meshGroupCertificateManager, meshGroupClient)
	meshNetworkingSnapshotContext := MeshNetworkingSnapshotContextProvider(meshWorkloadControllerFactory, meshServiceControllerFactory, meshGroupControllerFactory, meshNetworkingSnapshotValidator, groupMgcsrSnapshotListener)
	accessControlPolicyController, err := LocalAccessControlPolicyProvider(asyncManager)
	if err != nil {
		return MeshNetworkingContext{}, err
	}
	authorizationPolicyClient := security.NewAuthorizationPolicyClient(client)
	accessControlPolicyTranslator := access_control_poilcy.NewAccessControlPolicyTranslator(accessControlPolicyController, authorizationPolicyClient)
	meshNetworkingContext := MeshNetworkingContextProvider(multiClusterDependencies, asyncManagerHandler, trafficPolicyTranslator, meshNetworkingSnapshotContext, accessControlPolicyTranslator)
	return meshNetworkingContext, nil
}
