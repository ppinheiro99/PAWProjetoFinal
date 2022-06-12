import { Component, OnInit } from '@angular/core';
import { PresentationsService } from 'src/app/services/presentations/presentations.service';
import { SubjectsService } from 'src/app/services/subjects/subjects.service';
import * as $ from 'jquery' 

@Component({
  selector: 'app-grades',
  templateUrl: './grades.component.html',
  styleUrls: ['./grades.component.scss']
})
export class GradesComponent implements OnInit {

  constructor(public presentationService: PresentationsService, public subjectsService: SubjectsService ) { }

  ngOnInit(): void {
    $('#grades .wrap-subjects .subject').click(function(){
      console.warn("entrei")
        $(this).addClass('active')
    })
    
    this.presentationService.presentations = undefined
    this.subjectsService.getAllSubjects().subscribe(data =>{
      this.subjectsService.subjects = data.data
    })
  }

  getAllPresentations(id){
    console.warn(this)
    this.presentationService.getAllPresentations(id).subscribe(data =>{
      console.warn(data.data)
      console.warn(data)
      if(data.data != undefined)
        this.presentationService.presentations = data.data
     })
  }

  getGrades(id){
    this.presentationService.getClassificationByPresentation(id).subscribe(data =>{
      console.warn(data.data)
      console.warn(data)
    })
  }

  selectedSubject(event: any){
    console.warn(event.target.value)
    console.warn(event.target.classList)
    
  }

}
