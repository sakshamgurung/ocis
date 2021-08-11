#!/usr/bin/env bash

# ensure users exist, because first login in oCIS causes trouble...
curl http://localhost:8080/remote.php/webdav -u marie:radioactivity
curl http://localhost:8080/remote.php/webdav -u einstein:relativity
curl http://localhost:8080/remote.php/webdav -u richard:superfluidity

true
