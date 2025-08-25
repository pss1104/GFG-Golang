Standard Go Project Structure/
├── main.go
├── config/
│   └── db.go
├── models/
│   └── user.go
├── handlers/
│   └── user_handler.go
├── middleware/
│   ├── logging.go
│   └── security.go
└── utils/
    └── response.go


we will write all the db connection code in db.go
another option would be to store the config variables there 
and the db connection code is written someewhere else

db structure or model is stored in models folder
user.go -> car.go

user_handler.go -> car_handler.go
all the crud operations, hhtp method handling is done here

middlewares work on the request before going t the handnlers part or the actual operational logic

utils for sidecar helper code like json ecoding, decoding, random no. generation, hash generation etc
