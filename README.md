## 6GDemo環境設定
[install script](https://drive.google.com/file/d/1RdecE0FoUA4uepZnVQUk3kfHiDWW2vEN/view?usp=share_link)

### Install Go

```=
cd install/go/
bash install.sh
```

### Install Docker Images

```=
cd install/images/
bash install.sh
echo "PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "GOPATH=$HOME/go" >> ~/.bashrc
source ~/.bashrc
```

## 6GDemo 執行

進到FakeFlow、ModelStorage、GUI在裡面執行(3個都要執行)
```=
go run main.go
```


## 6GDemo 測試

進到FakeFlow、ModelStorage、GUI在裡面執行(3個都要執行)
```=
# 執行後會得到uuid
curl {IP}:3499/train_job/create
# 把ID帶到裡面 (記得等model上傳完)，建立inference model
curl {IP}:3501/model/deployment/{ID}
# 把ID帶到裡面，將image啟動
docker run -e MODEL_PORT=4000 -itd -p 4000:4000 --name test_inference {ID}-model:1.0
# predict
curl -X POST {IP}:4000/model/predict -H "Content-Type:application/json" -d '{"input":[4,5,6]}'

```