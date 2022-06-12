import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Rx';
import { io, Socket } from 'socket.io-client';
import { TokenService } from "../token/token.service";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import {Observable} from 'rxjs';

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
  timer: number = 30

  constructor(private tokenService:TokenService, private http: HttpClient) {
   
   } 

  init(){
    const user = this.tokenService.getUser()
    this.userName = user.username
    this.socket = io(`http://localhost:3000?userName=${this.userName}`)
    
    this.socket.on('user connect', (id) =>{
    })

    this.socket.on('receive_data', (data, username) =>{
      if(this.userName != username)
        this.presentationData = data  
        setTimeout(() => {
          this.presentationData = undefined
        }, 30000) 
    })
  }

  sendPresentationData(data){
    console.warn("enviar dados da apresentação")
    this.socket.emit('send_presentation_data', { 
      data: data,
    })
    // this.myTimer()
  }

  sendPageNumber(pageNumber){
    this.socket.emit('page_number', { 
      data:pageNumber,
    })
  }

  myTimer() {
    console.log(this.timer)
    this.timer = this.timer -1;
    if(this.timer == 0){
      return;
    }
    setTimeout(() => {this.myTimer()}, 1000)
    // setInterval(this.myTimer, 1000);
  }


}

