echo 'build image ...'
docker build --tag github.com/wesleyburlani/go-routing-and-spectrum-allocation . &>/dev/null
echo 'starting container ...'
docker run github.com/wesleyburlani/go-routing-and-spectrum-allocation
