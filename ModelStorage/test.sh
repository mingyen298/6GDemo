
docker rm -f test_inference
# docker rmi fefad97d-44a5-40ea-b93c-22d59d293fd7-model:1.0
# docker 
docker run -e MODEL_PORT=4000 -itd -p 4000:4000 --name test_inference $1-model:1.0