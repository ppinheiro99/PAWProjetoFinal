import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TokenService } from "../../../../services/token/token.service";
import { SubjectsService } from 'src/app/services/subjects/subjects.service';
import { MatDialog } from '@angular/material/dialog';
import { AddSubjectComponent } from '../add-subject/add-subject.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  isLoggedIn = false;
  constructor(private router: Router,public subjectsService : SubjectsService, private tokenStorage: TokenService, private dialog: MatDialog) {
  }

  ngOnInit() {
    if (this.tokenStorage.getToken()) {
      this.isLoggedIn = true;
      this.router.navigate(['/dashboard'])
    }
    this.getSubjectsData()
  }

  getSubjectsData(){
    this.subjectsService.getAllSubjects().subscribe(data =>{
      this.subjectsService.subjects = data.data
      console.warn(data.data)
    })
  }

  subjectsAvailable(){
    return this.subjectsService.subjects!=undefined;
  }

  addSubject(){
    console.warn("entrei")
    this.dialog.open(AddSubjectComponent)
  }

}


