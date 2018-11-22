#! /bin/bash
sleep 40

~/go/bin/horizon db init --ingest --db-url postgresql://horizonuser:password@localhost/horizondb?sslmode=disable --stellar-core-url http://127.0.0.1:8080 --stellar-core-db-url postgresql://postgres:password@localhost/stellar?sslmode=disable