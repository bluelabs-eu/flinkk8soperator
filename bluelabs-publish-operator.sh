current_version="1.0.8"
debug=false
while true; do
    read -p "Build as debug version? [no]: " yn
    case $yn in
        [Yy]* ) debug=true break;;
        [Nn]* ) break ;;
        * ) echo "Using default: Building as regular version" && break;;
    esac
done

read -p "Enter version [$current_version]: " version

build_version=${version:-$current_version}
if [ $debug == true ]; then
  build_version=${build_version}-debug
fi
echo "Building version $build_version, with_debug_flags=$debug"
docker buildx build --platform=linux/amd64 --build-arg with_debug_flags=$debug -t "eu.gcr.io/bluelabs-container-registry/data-lyft-flink-operator:v$build_version" .
docker push "eu.gcr.io/bluelabs-container-registry/data-lyft-flink-operator:v$build_version"