// Package operator contains "operational logic" so Sensor is able to operate itself when it is
// not deployed by our operator
package operator

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/logging"
	"k8s.io/client-go/kubernetes"
)

var (
	log = logging.LoggerForModule()
)

// Operator performs some operator logic on deployments types that are not managed by our operator,
// like Helm or YAML bundle deployments
type Operator interface {
	Initialize(ctx context.Context) error
	Start(ctx context.Context) error
	Stopped() concurrency.ReadOnlyErrorSignal
	GetHelmReleaseRevision() uint64
}

type operatorImpl struct {
	initialized  bool
	initError    error
	monitor      sync.Mutex
	k8sClient    kubernetes.Interface
	appNamespace string
	// Zero value if not managed by Helm
	helmReleaseName string
	// Zero value if not managed by Helm
	helmReleaseRevision uint64
	stoppedC            concurrency.ErrorSignal
}

// New creates a new operator
func New(k8sClient kubernetes.Interface, appNamespace string) Operator {
	return &operatorImpl{
		k8sClient:    k8sClient,
		appNamespace: appNamespace,
		stoppedC:     concurrency.NewErrorSignal(),
	}
}

func (o *operatorImpl) Initialize(ctx context.Context) error {
	o.monitor.Lock()
	defer o.monitor.Unlock()
	if o.initialized {
		return o.initError
	}

	log.Infof("Initializing operator for namespace %s", o.appNamespace)
	if err := o.fetchHelmReleaseName(ctx); err != nil {
		return o.failInitialization(err)
	}

	if err := o.fetchCurrentSensorHelmReleaseRevision(ctx); err != nil {
		return o.failInitialization(err)
	}

	o.initialized = true
	o.initError = nil
	return nil
}

func (o *operatorImpl) failInitialization(err error) error {
	o.initError = err
	return errors.Wrap(err, "Operator initialization error")
}

func (o *operatorImpl) isInitialized() bool {
	o.monitor.Lock()
	defer o.monitor.Unlock()
	initialized := o.initialized

	return initialized
}

func (o *operatorImpl) GetHelmReleaseRevision() uint64 {
	return o.helmReleaseRevision
}

func (o *operatorImpl) Stopped() concurrency.ReadOnlyErrorSignal {
	return &o.stoppedC
}

func (o *operatorImpl) stop(err error) {
	o.stoppedC.SignalWithError(err)
}

func (o *operatorImpl) Start(ctx context.Context) error {
	log.Info("Starting embedded operator.")
	if !o.isInitialized() {
		if err := o.Initialize(ctx); err != nil {
			log.Error(err)
			return err
		}
	}

	if !o.isSensorHelmManaged() {
		log.Warn("Sensor is not managed by Helm, stopping the embedded operator as it only supports Helm.")
		return nil
	}

	go o.mainLoop()

	log.Info("Embedded operator started correctly.")
	return nil
}

func (o *operatorImpl) mainLoop() {
	// TODO: // Secret watch ... or informer ..
	//  // see https://pkg.go.dev/k8s.io/client-go/tools/watch#NewIndexerInformerWatcher
	// FIXME
	for i := 0; i < 1; i++ { // FIXME proper loop
		/*
			var secret *v1.Secret
			err := o.processSecret(secret)
			if err != nil {
				err := errors.Wrapf(err, "Error processing secret with name %s", secret.GetName())
				log.Error(err)
			} */
		log.Warn("Bye for now")
	}
}
