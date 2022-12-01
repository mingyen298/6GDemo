# PA=$HOME/Desktop/SideProject/6GDemo/ModelStorage
unzip models/$1.zip -d models
docker build -t $1-model:1.0 -f deployment/Dockerfile .
rm -rf models/model
