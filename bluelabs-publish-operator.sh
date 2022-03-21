version=1.0.6
docker buildx build --platform=linux/amd64 -t "eu.gcr.io/bluelabs-container-registry/data-lyft-flink-operator:v$version" .
docker push "eu.gcr.io/bluelabs-container-registry/data-lyft-flink-operator:v$version"