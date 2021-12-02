wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

go build
./GoTraining