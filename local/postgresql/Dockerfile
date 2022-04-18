FROM postgres:latest

ENV POSTGRES_USER HERDOGAN
ENV POSTGRES_PASSWORD ytjpzldm
ENV POSTGRES_DB basket-db

COPY migration.sql /docker-entrypoint-initdb.d/

EXPOSE 5432