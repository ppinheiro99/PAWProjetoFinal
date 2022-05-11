import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Params, Router } from '@angular/router';
import { SubjectsService } from 'src/app/services/subjects/subjects.service';
import { TokenService } from 'src/app/services/token/token.service';
import { ActivatedRoute } from '@angular/router';
import { PresentationsService } from 'src/app/services/presentations/presentations.service';

@Component({
  selector: 'app-subject-dashboard',
  templateUrl: './subject-dashboard.component.html',
  styleUrls: ['./subject-dashboard.component.scss']
})

export class SubjectDashboardComponent implements OnInit {
  subjectId : any;
  subject: any
  presentation:any;

  constructor(private router: Router,public subjectsService : SubjectsService, private tokenStorage: TokenService, private dialog: MatDialog,private route: ActivatedRoute, public presentationService: PresentationsService) {
  }

  ngOnInit() {
    this.route.params.subscribe((params: Params) => this.subjectId = params['id']);
    this.getSubjectById(this.subjectId)
    console.warn(this.subjectId)
    this.getAllPresentationsData()
  }

  getSubjectById(id){
    this.subjectsService.getSubjectById(id).subscribe(data =>{
      console.warn(data.data)
      this.subject = data.data.name
      console.warn(data)
    })
  }

  getAllPresentationsData(){
    this.presentationService.getAllPresentations(this.subjectId).subscribe(data =>{
      this.presentationService.presentations = data.data
      console.warn(this.presentationService.presentations)
    })
  }

  presentationAvailable(){
    return this.presentationService.presentations!=undefined;
  }

  addPresentation(){
    // this.dialog.open(AddSubjectComponent)
  }

  backToDashboard(){
    this.router.navigate(['dashboard'])
  }

}
