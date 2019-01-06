#! /bin/bash
sleep 40

/home/ubuntu/go/bin/horizon serve --ingest --db-url postgresql://horizonuser:password@localhost/horizondb --stellar-core-url http://127.0.0.1:8080 --stellar-core-db-url postgresql://postgres:password@localhost/stellar --friendbot-url http://localhost:8004/
