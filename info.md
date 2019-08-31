mkdir -p ${HOME}/go/src/github.com/wallix/
C=${PWD}
cd ${HOME}/go/src/github.com/wallix/
ln -snf ${C} awless
cd -
go build
