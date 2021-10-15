import React from 'react';
import Labeled from 'Components/Labeled';
import { oidcCallbackValues } from 'constants/accessControl';

const CommonDetails = ({ name }) => (
    <>
        <Labeled label="Integration Name">{name}</Labeled>
    </>
);

const OidcDetails = ({ authProvider: { name, config } }) => {
    const callbackModeValue = oidcCallbackValues[config.mode];
    if (!callbackModeValue) {
        throw new Error(`Unknown callback mode "${config.mode}"`);
    }

    return (
        <>
            <CommonDetails name={name} />
            <Labeled label="Callback Mode">{callbackModeValue}</Labeled>
            <Labeled label="Issuer">{config.issuer}</Labeled>
            <Labeled label="Client ID">{config.client_id}</Labeled>
            <Labeled label="Client Secret">{config.client_secret ? '*****' : null}</Labeled>
        </>
    );
};

const Auth0Details = ({ authProvider: { name, config } }) => (
    <>
        <CommonDetails name={name} />
        <Labeled label="Auth0 Tenant">{config.issuer}</Labeled>
        <Labeled label="Client ID">{config.client_id}</Labeled>
    </>
);

const SamlDetails = ({ authProvider: { name, config } }) => {
    const idpDetails = config.idp_metadata_url ? (
        <Labeled label="Dynamically configured using IdP metadata URL">
            {config.idp_metadata_url}
        </Labeled>
    ) : (
        <>
            <Labeled label="IdP Issuer">{config.idp_issuer}</Labeled>
            <Labeled label="IdP SSO URL">{config.idp_sso_url}</Labeled>
            <Labeled label="Name/ID Format">{config.idp_nameid_format}</Labeled>
            <Labeled label="IdP Certificate(s) (PEM)">
                <pre className="font-500 whitespace-pre-line">{config.idp_cert_pem}</pre>
            </Labeled>
        </>
    );
    return (
        <>
            <CommonDetails name={name} />
            <Labeled label="ServiceProvider Issuer">{config.sp_issuer}</Labeled>
            {idpDetails}
        </>
    );
};

const UserPkiDetails = ({ authProvider: { name, config } }) => (
    <>
        <CommonDetails name={name} />
        <Labeled label="CA Certificates (PEM)">
            <pre className="font-500 whitespace-pre-line">{config.idp_cert_pem}</pre>
        </Labeled>
    </>
);

const IapDetails = ({ authProvider: { name, config } }) => (
    <>
        <CommonDetails name={name} />
        <Labeled label="Audience">
            <pre className="font-500 whitespace-pre-line">{config.audience}</pre>
        </Labeled>
    </>
);

const detailsComponents = {
    oidc: OidcDetails,
    auth0: Auth0Details,
    saml: SamlDetails,
    userpki: UserPkiDetails,
    iap: IapDetails,
};

const ConfigurationDetails = ({ authProvider }) => {
    const DetailsComponent = detailsComponents[authProvider.type];
    if (!DetailsComponent) {
        throw new Error(`Unknown auth provider type: ${JSON.stringify(authProvider)}`);
    }

    return <DetailsComponent authProvider={authProvider} />;
};

export default ConfigurationDetails;