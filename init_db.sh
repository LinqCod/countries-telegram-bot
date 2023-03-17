set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE countries_pg;
    CREATE USER linqcod WITH ENCRYPTED PASSWORD 'root';
    GRANT ALL PRIVILEGES ON DATABASE countries_pg TO linqcod;

    \c countries_pg

    CREATE TABLE country (
        country_id SERIAL PRIMARY KEY,
        information JSONB
    );
    GRANT ALL PRIVILEGES ON TABLE country TO linqcod;
    GRANT USAGE, SELECT ON SEQUENCE country_country_id_seq TO linqcod;
EOSQL