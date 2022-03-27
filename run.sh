
if [ -f .env ]
then
   export $(cat .env | grep -v '#' | grep -E '.+=.+' | sed 's/\r$//')
fi

go build -o build/program

./build/program \
  --nodesFilePath=${NODES_FILE_PATH} \
  --edgesFilePath=${EDGES_FILE_PATH} \
  --logType=${LOG_TYPE} \
  --logFilePath=${LOG_FILE_PATH} \
  --demandsSource=${DEMANDS_SOURCE} \
  --demandsFilePath=${DEMANDS_FILE_PATH} \
  --pathSearchAlgorithm=${PATH_SEARCH_ALGORITHM} \
  --disjointedPathPairSearchAlgorithm=${DISJOINTED_PATH_PAIR_SEARCH_ALGORITHM} \
  --tableFillAlgorithm=${TABLE_FILL_ALGORITHM} \
  --rsaType=${RSA_TYPE}
