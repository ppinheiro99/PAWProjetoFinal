var express = require("express");
var socket = require('socket.io');
var cors = require('cors');


var app = express();
app.use(cors());
var index= 0;

// Initialize our websocket server on port 5000
var server = app.listen(3000, () => {
  console.log("started on port 3000");
});


let userList = []
let socketID = []
var presentationData = []

var io = socket(server);
io.on("connection", (socket) => {

  let userName = socket.handshake.query.userName
  // let userID = socket.handshake.query.id

  socketID[userName] = socket.id // ArrayList de userId com os respectivos socketsId's
  let indexDelete = index
  console.log("user connect " + userName + "socket: " + socketID[userName])
  
  addUser(userName)

  socket.broadcast.emit('user-list', [...userList])
  socket.emit('user-list', [...userList])

  console.log(userList)

  io.emit("user connect",userList)

  socket.on("page_number", (pg) =>{ 
    for(let i = 0;  i < presentationData.length; i++){
      console.log(presentationData[i])
      if(Number(presentationData[i].questionNumber) == pg.data){
        console.log("enviar dados" + userName + "")
        io.emit("receive_data",presentationData[i], userName)
      }
    }

  })

  socket.on("send_presentation_data", (data) =>{ 
    presentationData = data.data
    console.log(presentationData.answer)
  })

  // socket.on("message", (msg) =>{ 
  //   console.log("Id sender: " + userID);
  //   console.log("Id received: " + msg.receivedID);
  //   console.log("Data Received: " + msg.data);
  //   // Mensagem enviada para o utilizador que pretendemos enviar msg
  //   io.to(socketID[msg.receivedID]).emit("message",{
  //     message: msg.data,
  //     userName:userName,
  //     sender: userID,
  //     received:msg.receivedID,
  //     mine: false
  //   })
  //   // Mensagem enviada para ele próprio
  //   io.to(socketID[userID]).emit("message",{
  //     message: msg.data,
  //     userName:userName,
  //     sender: userID,
  //     received:msg.receivedID,
  //     mine: true
  //   })

  // });

  socket.on("disconnect", () => {
     removeUser(userName,indexDelete)
     socket.broadcast.emit('user-list', [...userList])
     socket.emit('user-list', [...userList])
  });
});

function addUser(userName){
  userList.push({
    name: userName,
  })
  index++
}

function removeUser(userName, indexDelete){ // elimina o user quando faz logout ou faz reload
  console.log("user disconnect " + userName)
  userList.splice(indexDelete,1)
  socketID.splice(userName,1)
  index--
}
