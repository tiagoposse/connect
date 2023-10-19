#!/bin/sh

salt=$(openssl rand -base64 12)
password="portal"
hashed_password=$(echo -n "${password}${salt}" | sha256sum | awk '{print $1}')
user_id="b5785cb7-6c95-4817-abd1-24b96fbec905"


generate_user_cmd="INSERT INTO users (id, email, provider, password, salt, is_admin, photo_url, disabled, disabled_reason, firstname, lastname) VALUES ('$user_id', 'admin@local.io', 'local', '$hashed_password', '$salt', true, '', false, '', 'Admin', 'User')"

API_KEY="fa00afd0309ff0f52f760d60195ffbe5"

generate_api_key_cmd="INSERT INTO api_keys (user_id, key) VALUES ('$user_id', '$API_KEY')"

if [ "$DB_TYPE" == "MYSQL" ]; then
  mysql -h localhost -u wireguard -p -e "USE wireguard; $generate_user_cmd; $generate_api_key_cmd;"
else
  psql -h localhost -U wireguard -d wireguard -c "$generate_user_cmd; $generate_api_key_cmd;"
fi
