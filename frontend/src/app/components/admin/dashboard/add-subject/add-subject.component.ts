import { Component, OnInit } from '@angular/core';
import { SubjectsService } from 'src/app/services/subjects/subjects.service';

@Component({
  selector: 'app-add-subject',
  templateUrl: './add-subject.component.html',
  styleUrls: ['./add-subject.component.scss']
})
export class AddSubjectComponent implements OnInit {
  subjectName: string
  constructor(private subjectsService : SubjectsService) { }

  ngOnInit(): void {
  }

  addSubject(){
    console.warn(this.subjectName)
    this.subjectsService.addSubject(this.subjectName).subscribe(
      data => {
       console.warn(data)
       this.getSubjectsData()
      },
      err => {
        alert(err)
      }
    );
  }

  getSubjectsData(){
    this.subjectsService.getAllSubjects().subscribe(data =>{
      this.subjectsService.subjects = data.data
      console.warn(data.data)
    })
  }

}
