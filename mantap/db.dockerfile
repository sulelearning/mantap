FROM postgres:15

COPY db/migration /docker-entrypoint-initdb.d/migration

CMD ["postgres"]