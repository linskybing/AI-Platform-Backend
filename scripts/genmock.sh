#! /bin/bash

sudo apt install mockgen

for file in ./src/repositories/*_repository.go; do
    base=$(basename $file .go)
    mockgen -source=$file -destination=./src/repositories/mock_repositories/${base}_mock.go -package=mock_repositories
done

for file in ./src/services/*_service.go; do
    base=$(basename $file .go)
    mockgen -source=$file -destination=./src/services/mock_services/${base}_mock.go -package=mock_services
done