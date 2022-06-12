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

  if(addUser(userName))
  {
    socketID[userName] = socket.id // ArrayList de userId com os respective socketsId's
    
    console.log("user connect " + userName + "socket: " + socketID[userName])

    socket.broadcast.emit('user-list', [...userList])
    socket.emit('user-list', [...userList])

    console.log(userList)

    io.emit("user connect",userList)
  }
  let indexDelete = index
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
    console.log("Inicio da apresentação")
    presentationData = data.data
  })
  
  socket.on("disconnect", () => {
     removeUser(userName,indexDelete)
     socket.broadcast.emit('user-list', [...userList])
     socket.emit('user-list', [...userList])
  });
});

function addUser(userName){

  let index = userList.findIndex(aux => aux.name == userName);
  if(index != -1){
    console.log("user já existe")
    return false
  }
  userList.push({
    name: userName,
  })
  index++
  return true
}

function removeUser(userName, indexDelete){ // elimina o user quando faz logout ou faz reload
  console.log("user disconnect " + userName)
  userList.splice(indexDelete,1)
  socketID.splice(userName,1)
  index--
}
