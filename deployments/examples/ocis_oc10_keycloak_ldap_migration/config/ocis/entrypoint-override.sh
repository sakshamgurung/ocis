#!/bin/sh

set -e

# start everything except glauth and idp https://github.com/owncloud/ocis/pull/2229
#ocis server --extensions="accounts, graph, graph-explorer, ocs, onlyoffice, proxy, settings, storage-authbasic, storage-authbearer, storage-frontend, storage-gateway, storage-groupsprovider, storage-home, storage-metadata, storage-public-link, storage-sharing, storage-users, storage-users-provider, store, thumbnails, web, webdav"

ocis server &
sleep 10

ocis kill idp
ocis kill glauth

ocis list

wait


