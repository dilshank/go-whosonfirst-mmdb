#!/bin/sh

# This assumes 'ADD . /go-whosonfirst-mmdb' which is defined in the Dockerfile

MMDB_SERVER="/go-whosonfirst-mmdb/bin/wof-mmdb-server"
ARGS=""

CURL=`which curl`

if [ "${HOST}" != "" ]
then
    ARGS="${ARGS} -host ${HOST}"
fi

# FETCH MMDB database here...
# ${CURL} -s -o ${LOCAL}.bz2 ${REMOTE}.bz2

echo ${MMDB_SERVER} ${ARGS}
${MMDB_SERVER} ${ARGS}

if [ $? -ne 0 ]
then
   echo "command '${MMDB_SERVER} ${ARGS}' failed"
   exit 1
fi

exit 0
