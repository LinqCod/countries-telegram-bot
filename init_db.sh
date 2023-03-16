set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE countries_pg;
    CREATE USER linqcod WITH ENCRYPTED PASSWORD 'root';
    GRANT ALL PRIVILEGES ON DATABASE countries_pg TO linqcod;
EOSQL