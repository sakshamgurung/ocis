<?php

# reference: https://owncloud.dev/clients/web/deployments/oc10-app/

function getWebConfigFromEnv()
{
    $config = [
        'web.baseUrl' => 'https://' . getenv('OC10_DOMAIN') . '/index.php/apps/web',
        'web.rewriteLinks' => getenv('OWNCLOUD_WEB_REWRITE_LINKS') == 'true',

    ];
    return $config;
}

$CONFIG = getWebConfigFromEnv();
