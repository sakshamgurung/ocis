<?php

# reference: https://owncloud.dev/clients/web/deployments/oc10-app/

function getWebConfigFromEnv()
{
    $config = [
        'web.baseUrl' => 'https://' . getenv('OC10_DOMAIN') . '/index.php/apps/web',
        'web.rewriteLinks' => true,
        'defaultapp' => 'web',

    ];
    return $config;
}

$CONFIG = getWebConfigFromEnv();
