docker build -t basketpostgresql ./local/postgresql/.
docker run --name basketpostgrescontainer -d -p 5434:5432 basketpostgresql

