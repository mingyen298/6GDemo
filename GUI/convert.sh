# docker rm -f ipynb2py
docker run --name ipynb2py -v $PWD/.:/app --rm ipynb_convert:1.0 jupyter nbconvert --to script 6GDemo.ipynb