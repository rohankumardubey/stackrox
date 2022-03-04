import common.Constants
import groups.BAT
import groups.OpenShift
import objects.Deployment
import objects.Job
import org.junit.Assume
import org.junit.experimental.categories.Category
import spock.lang.Unroll

@OpenShift
class OpenShiftInternalImageScanningTest extends BaseSpecification {

    private static final String NAMESPACE = Constants.ORCHESTRATOR_NAMESPACE
    private static final String[] IMAGES = [
        "nginx:1.18.0@sha256:e90ac5331fe095cea01b121a3627174b2e33e06e83720e9a934c7b8ccc9c55a0",
        "quay.io/rhacs-eng/qa:sandbox-jenkins-agent-maven-35-rhel7",
        "quay.io/rhacs-eng/qa:sandbox-nodejs-10",
        "quay.io/rhacs-eng/qa:sandbox-log4j-2-12-2",
    ]
    private static final Deployment INIT_JOB = new Job()
            .setName("populate-internal-registry")
            .addLabel("app", "init-intenral-registry-job")
            .setImage("docker:20.10.12-alpine3.15")

    def setupSpec() {
        orchestrator.createDeployment(INIT_JOB)
        orchestrator.waitForAllPodsToBeRemoved(INIT_JOB.getNamespace(), ["app": "init-intenral-registry-job"], 30, 10)
    }

    def cleanupSpec() {

    }

    @Unroll
    @Category(BAT)
    def "Verify image scans are equal #imageName" {

    }
}
