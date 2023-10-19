#!/bin/sh

mkdir -p /usr/share/nginx/html/app/
ENV_PATH="/usr/share/nginx/html/app/config.js"
echo "window.WG_API = '$WG_API';" > $ENV_PATH

npm ci
npm run dev
