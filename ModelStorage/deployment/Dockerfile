FROM 6g-demo:inference

RUN mkdir model
ADD models/model/. /app/model/.
ADD deployment/inference.py /app/main.py

CMD [ "python","main.py" ]