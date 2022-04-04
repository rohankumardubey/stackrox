package services

import io.grpc.Status
import io.grpc.StatusRuntimeException
import io.stackrox.proto.api.v1.EmptyOuterClass
import io.stackrox.proto.api.v1.SignatureIntegrationServiceGrpc
import io.stackrox.proto.api.v1.SignatureIntegrationServiceOuterClass
import io.stackrox.proto.storage.ImageIntegrationOuterClass
import io.stackrox.proto.storage.SignatureIntegrationOuterClass
import util.Timer

class SignatureIntegrationService extends BaseService {
    static getSignatureIntegrationClient() {
        return SignatureIntegrationServiceGrpc.newBlockingStub(getChannel())
    }

    static createSignatureIntegration(SignatureIntegrationOuterClass.SignatureIntegration integration) {
        SignatureIntegrationOuterClass.SignatureIntegration createdIntegration
        Timer t = new Timer(15, 3)
        while (t.IsValid()) {
            try {
                createdIntegration = getSignatureIntegrationClient().postSignatureIntegration(integration)
                println "Signature integration created: ${createdIntegration.getName()}: ${createdIntegration.getId()}"
                break
            } catch (Exception e) {
                println "Unable to create signature integration ${integration.getName()}: ${e.getMessage()}"
            }
        }

        if (!createdIntegration || !createdIntegration.getId()) {
            println "Unable to create signature integration"
            return ""
        }

        SignatureIntegrationOuterClass.SignatureIntegration foundIntegration
        t = new Timer(15, 3)
        while (t.IsValid()) {
            try {
                foundIntegration =
                        getSignatureIntegrationClient().getSignatureIntegration(
                                getResourceByID(createdIntegration.getId()))
                if (foundIntegration) {
                    println "Integration found after creation: " +
                            "${foundIntegration.getName()}: ${foundIntegration.getId()}"
                    return foundIntegration.getId()
                }
            } catch (Exception e) {
                println "Unable to find the created signature integration ${integration.getName()}: ${e.message}"
            }
        }
        println "Unable to find the created signature integration: ${integration.getName()}"
        return ""
    }

    static deleteSignatureIntegration(String integrationId) {
        try {
            getSignatureIntegrationClient().deleteSignatureIntegration(getResourceByID(integrationId))
        } catch (Exception e) {
            println "Failed to delete signature integration with id ${integrationId}: ${e.message}"
        }

        Timer t = new Timer(15, 3)
        while(t.IsValid()) {
            try {
                ImageIntegrationOuterClass integration =
                        getSignatureIntegrationClient().getSignatureIntegration(getResourceByID(integrationId))
            } catch (StatusRuntimeException e) {
                if (e.status.code == Status.Code.NOT_FOUND) {
                    println "Signature integration deleted: ${integrationId}"
                    return true
                }
            }
        }
    }
}
