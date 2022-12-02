def train_export_model():
    import inspect
    import requests
    import numpy as np
    from tensorflow.keras.models import Sequential, Model
    from tensorflow.keras.layers import Dense, LSTM
    import os
    import zipfile
    import paho.mqtt.client as mqtt
    import json
    TrainJobID = os.environ.get('JOB_ID')
    def ZipDir(path,trainJobId):
        zf = zipfile.ZipFile('{}.zip'.format(trainJobId), 'w', zipfile.ZIP_DEFLATED)
        for root, dirs, files in os.walk(path):
            for file_name in files:
                zf.write(os.path.join(root, file_name))
    class MQStatus:
        def __init__(self):
            self.topic = "/mqtt"
            self.client = mqtt.Client()
        def Publish(self,state):
            dic = {"uuid":TrainJobID , "state":state}

            message = json.dumps(dic)
            self.client.publish(self.topic , message)
        def Connect(self):
            self.client.connect("yen-test.com", 1883)
    class ModelStorageSDK:
        @staticmethod
        def UploadModel(path,trainJobId):
            ZipDir(path,trainJobId)
            requests.post("http://172.17.0.1:3501/model/upload/{0}".format(trainJobId),files={'file':open("{}.zip".format(trainJobId), 'rb')})
    class EasyLSTMModel:
        def __init__(self,X,Y,sw_width,features,epochs_num):
            self.X = X
            self.Y = Y
            self.sw_width = sw_width
            self.features = features
            self.epochs_num = epochs_num
            self.model = None
        def Training(self):
            self.model = Sequential()
            self.model.add(LSTM(50, activation='relu', input_shape=(self.sw_width, self.features)))
            self.model.add(Dense(1))
            self.model.compile(optimizer='adam', loss='mse', metrics=['accuracy'])
            self.model.summary()
            history = self.model.fit(self.X, self.Y, epochs=self.epochs_num, verbose=0)
            print('\ntrain_acc:%s'%np.mean(history.history['accuracy']), '\ntrain_loss:%s'%np.mean(history.history['loss']))

        def Predict(self,testSeq):
            print('yhat:%s'%(self.model.predict(testSeq)),'\n-----------------------------')

        def Save(self):
            self.model.save("./model")
    class DataExtractor:
        def __init__(self):
            pass
        def getDBData(self,n):
            rawDatas = []
            for i in range(1,n+1):
              rawDatas.append(i)
            return rawDatas
        def dataExtraction(self,rawDatas):
            extractedDatas = []
            cur = 0
            length = len(rawDatas)
            X,Y = [],[]
            while cur < (length-3):
              X.append(rawDatas[cur:cur+3])
              Y.append(rawDatas[cur+3:cur+4])
              cur = cur + 1
            X = np.array(X)
            Y = np.array(Y)
            # print("X Shape:{}",X.shape)
            # print(X)
            X = X.reshape((X.shape[0], X.shape[1], 1))
            return X , Y

        def Run(self):
            rawDatas = self.getDBData(10)
            return self.dataExtraction(rawDatas)
    mq = MQStatus()
    mq.Connect()
    mq.Publish(1)
    # 1. Data Extraction
    extractor = DataExtractor()
    # time.sleep(30)
    X , Y = extractor.Run()
    print("X Shape:{}",X.shape)
    print(X)
    # 2. Model Training
    # time.sleep(30)
    mq.Publish(2)
    model = EasyLSTMModel(X,Y,3,1,500)
    model.Training()
    # 3. Model Experiment
    testDatas = np.array([5,6,7])
    testDatas = testDatas.reshape((1, 3, 1))
    print("testDatas Shape:{}",testDatas.shape)
    print(testDatas)
    model.Predict(testDatas)
    model.Save()
    ModelStorageSDK.UploadModel("./model",TrainJobID)
    mq.Publish(3)
train_export_model()
