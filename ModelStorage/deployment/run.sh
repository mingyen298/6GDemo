# PA=$HOME/Desktop/SideProject/6GDemo/ModelStorage
unzip models/$1.zip -d models
docker build -t $1-model:1.0 -f deployment/Dockerfile .
rm -rf models/model
# docker rm -f test_inference
# docker run -e MODEL_PORT=4000 -itd -p 4000:4000 --name test_inference $1-model:1.0