#!/bin/bash
printenv
# replace oCIS domain in keycloak realm import
cp /opt/jboss/keycloak/ocis-realm.dist.json /opt/jboss/keycloak/ocis-realm.json
sed -i "s/cloud.owncloud.test/${CLOUD_DOMAIN}/g" /opt/jboss/keycloak/ocis-realm.json

# run original docker-entrypoint
/opt/jboss/tools/docker-entrypoint.sh
