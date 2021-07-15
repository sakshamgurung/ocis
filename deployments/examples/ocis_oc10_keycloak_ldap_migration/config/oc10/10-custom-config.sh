#!/usr/bin/env bash
echo "Writing custom config files..."
gomplate \
  -f /etc/templates/oidc.config.php \
  -o ${OWNCLOUD_VOLUME_CONFIG}/oidc.config.php

gomplate \
  -f /etc/templates/web.config.php \
  -o ${OWNCLOUD_VOLUME_CONFIG}/web.config.php

gomplate \
  -f /etc/templates/web-config.tmpl.json \
  -o ${OWNCLOUD_VOLUME_CONFIG}/config.json


occ market:upgrade web

occ app:enable web

true


