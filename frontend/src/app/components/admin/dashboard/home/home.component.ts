import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TokenService } from "../../../../services/token/token.service";
import { SubjectsService } from 'src/app/services/subjects/subjects.service';
import { MatDialog } from '@angular/material/dialog';
import { AddSubjectComponent } from '../add-subject/add-subject.component';
import { ChatService } from 'src/app/services/socket/chat.service';
import { PresentationsService } from 'src/app/services/presentations/presentations.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  isLoggedIn = false;
  userInfo: any
  role: any
  constructor(private tokenService:TokenService,private presentationsService : PresentationsService,public chat: ChatService,private router: Router,public subjectsService : SubjectsService, private tokenStorage: TokenService, private dialog: MatDialog) {
  }

  ngOnInit() {
    this.userInfo = this.tokenService.getUser()
    this.role = this.userInfo.role
    console.warn(this.role)
    if (this.tokenStorage.getToken()) {
      this.isLoggedIn = true;
      this.router.navigate(['/dashboard'])
    }
    this.getSubjectsData()
    this.chat.init()
  }

  getSubjectsData(){
    this.subjectsService.getAllSubjects().subscribe(data =>{
      this.subjectsService.subjects = data.data
    })
  }

  subjectsAvailable(){
    return this.subjectsService.subjects!=undefined;
  }

  addSubject(){
    this.dialog.open(AddSubjectComponent)
  }

  showSubject(subject){
    this.router.navigate(['subject', subject.ID])
  }

  deleteSubject(subject){
    event.preventDefault();
    console.warn(subject.ID)
    this.subjectsService.deleteSubject(subject.ID).subscribe(data =>{
      this.getSubjectsData()
    })
  }

  presentationAvailable(){
    if(this.chat.presentationData == undefined)
      return false
    else{
      this.chat.presentationData == undefined
      return true
    }
  }

  sendResponse(answer, id){
    console.warn(answer)
    this.presentationsService.sendUserResponse(answer, id).subscribe(data =>{
      console.warn("sucesso")
    })
  }

}


