docker run --name ipynb2py -v $PWD/.:/app --rm ipynb_convert:1.0 jupyter nbconvert --to script 6GDemo.ipynb
# docker run --name $1_train_job --rm -v $PWD/6GDemo.py:/app/6GDemo.py train_env:1.1 python 6GDemo.py $1
docker run --name $1_train_job --rm -v $PWD/6GDemo.py:/app/6GDemo.py 6g-demo:train python 6GDemo.py $1
rm -f $PWD/6GDemo.py
