FROM postgres:15-alpine

COPY db/migration /docker-entrypoint-initdb.d/migration

CMD ["postgres"]