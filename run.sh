
if [ -f .env ]
then
   export $(cat .env | grep -v '#' | grep -E '.+=.+' | sed 's/\r$//')
fi

go build -o build/program

./build/program --nodesFilePath=${NODES_FILE_PATH} --edgesFilePath=${EDGES_FILE_PATH}
