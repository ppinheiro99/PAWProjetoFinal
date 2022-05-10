import { Component, OnInit } from '@angular/core';
import { SubjectsService } from 'src/app/services/subjects/subjects.service';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-add-subject',
  templateUrl: './add-subject.component.html',
  styleUrls: ['./add-subject.component.scss']
})
export class AddSubjectComponent implements OnInit {
  subjectName: string
  constructor(private subjectsService : SubjectsService, private dialog: MatDialog) { }

  ngOnInit(): void {
  }

  addSubject(){
    this.subjectsService.addSubject(this.subjectName).subscribe(
      data => {
       this.getSubjectsData()
       this.dialog.closeAll()
      },
      err => {
        alert(err)
        this.dialog.closeAll()
      }
    );
  }

  getSubjectsData(){
    this.subjectsService.getAllSubjects().subscribe(data =>{
      this.subjectsService.subjects = data.data
    })
  }

}
