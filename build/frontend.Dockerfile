
FROM nginx:stable

ARG app

COPY entrypoint.sh /entrypoint.sh
COPY .gitlab/docker/nginx.conf /etc/nginx/conf.d/default.conf

ADD dist/apps/$app /usr/share/nginx/html/app
RUN touch /var/run/nginx.pid && \
  chown -R nginx:nginx /usr/share/nginx/html/app /var/cache/nginx /var/log/nginx /etc/nginx/conf.d /var/run/nginx.pid && \
  chmod -R 755 /usr/share/nginx/html/app && \
  chmod 755 /entrypoint.sh

USER nginx

ENTRYPOINT [ "/entrypoint.sh" ]
