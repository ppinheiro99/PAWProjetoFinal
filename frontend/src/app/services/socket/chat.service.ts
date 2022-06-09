import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Rx';
import { io, Socket } from 'socket.io-client';
import { TokenService } from "../token/token.service";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import {Observable} from 'rxjs';
import { UsersService } from '../user/users.service';

const API_URL = 'http://localhost:8080/api/';
//const API_URL = 'http://18.130.231.194:8080/api/';
const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})
export class ChatService {

  receiver_id : any
  socket: any
  user_id:any
  userName:any
  messageList: {message: string, userName: string, senderID: string, userId : string, mine: boolean}[] = []
  users_connected = []
  users: {userID:number, userName:string, online:boolean}[]= []
  userList = []
  messages : {message: string, userName: string, senderID: string, userId : string, mine: boolean}[] = []
  user_select: any
  page_number: any
  presentationData: any

  constructor(private tokenService:TokenService, private http: HttpClient, private userService: UsersService) {
   
   }

  init(){
    const user = this.tokenService.getUser()
    console.warn(user.username)
    // this.user_id = user.ID
    this.userName = user.username
    // //this.socket = io(`http://18.130.231.194:3000?userName=${this.userName}&id=${this.user_id}`)
    this.socket = io(`http://localhost:3000?userName=${this.userName}`)
    
    this.socket.on('user connect', (id) =>{
    })

    this.socket.on('receive_data', (data, username) =>{
      console.warn("ENTREI " + data.correctAnswer, username)
      if(this.userName != username)
        this.presentationData = data  
    })


    //  this.socket.on('user-list', (userList) =>{
    //     this.userList = userList
    //     this.userService.getData().subscribe(data =>{

    //     // quando recebo um user-list é pq alguem novo se conectou e como tal vou receber o array de novos users, tendo entao que limpar o antigo
    //     this.users = []  
    //     data.data.forEach( (currentValue, index) => { // para saber dos users todos quais estao online e quais estao offline
    //       if(userList.find(x => x.id == data.data[index].ID ) != null){ 
    //         this.users.push({
    //           userID: data.data[index].ID,
    //           userName: data.data[index].first_name,
    //           online: true
    //         })
    //       }else{
    //         this.users.push({
    //           userID: data.data[index].ID,
    //           userName: data.data[index].first_name,
    //           online: false
    //         })
    //       }
    //     })
    //   })
    // })

    // this.socket.on("message", (users) =>{
    //   this.messageList.push({
    //       message:users.message,
    //       userName:users.userName,
    //       senderID:users.sender,
    //       userId: users.received,
    //       mine:users.mine
    //   })
    //     let aux = 0
    //     this.messageList.forEach( (currentValue, index) => {
    //     if(this.messageList[index].senderID == this.user_select  || this.messageList[index].userId == this.user_select){
    //       this.messages[aux] = this.messageList[index]
    //       aux++
    //     }
    //   })
    // })
  }

  sendPresentationData(data){
    this.socket.emit('send_presentation_data', { 
      data: data,
    })
  }

  sendPageNumber(pageNumber){
    this.socket.emit('page_number', { 
      data:pageNumber,
    })
  }
}