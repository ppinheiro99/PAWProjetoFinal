import { Component, OnInit } from '@angular/core';
import { PresentationsService } from 'src/app/services/presentations/presentations.service';
import { SubjectsService } from 'src/app/services/subjects/subjects.service';
import * as $ from 'jquery' 
import { TokenService } from 'src/app/services/token/token.service';

@Component({
  selector: 'app-grades',
  templateUrl: './grades.component.html',
  styleUrls: ['./grades.component.scss']
})
export class GradesComponent implements OnInit {

  subjects: {studentUsername: string, presentationName: string, classification: number}[] = []
  students: {studentUsername: string, presentationName: string, classification: number, questions: number}[] = []
  presentationData: {studentUsername: string, presentationName: string, classification: number}
  questionsNumbers: number = 0
  role: any
  constructor(private tokenService:TokenService,public presentationService: PresentationsService, public subjectsService: SubjectsService ) { }

  ngOnInit(): void {
    // $('#grades .wrap-subjects .subject').click(function(){
    //   console.warn("entrei")
    //     $(this).addClass('active')
    // })
    
    // this.presentationService.presentations = undefined
    // this.subjectsService.getAllSubjects().subscribe(data =>{
    //   this.subjectsService.subjects = data.data
    // })
    let user = this.tokenService.getUser()
    this.role = user.role
    if(!this.role){ // aluno
      this.presentationService.getClassificationByPresentation(user.username).subscribe(data =>{
        console.warn(data.data)
        this.questionsNumbers = data.size
        for(var i = 0; i < data.data.length; i++){
          this.subjects.push({
            studentUsername: data.data[i].studentUsername,
            presentationName: data.data[i].presentationName,
            classification: data.data[i].classification
          })
          // this.classification = data.data.slice(-1).pop()        
        }
        console.warn(this.presentationData)
      })
    }
    if(this.role){ // professor
      this.presentationService.getClassificationByPresentation(user.username).subscribe(data =>{
        console.warn(data.data.SubjectGrades)
        for(var i = 0; i < data.data.SubjectGrades.length; i++){
          console.warn(data.data.SubjectGrades[i])
          for(var j = 0; j < data.data.SubjectGrades[i].Grades.length; j++){
            this.students.push({
              studentUsername: data.data.SubjectGrades[i].Grades[j].studentUsername,
              presentationName: data.data.SubjectGrades[i].Grades[j].presentationName,
              classification: data.data.SubjectGrades[i].Grades[j].classification,
              questions: data.data.SubjectGrades[i].Grades[j].questions
            })
          }
          // this.classification = data.data.slice(-1).pop()        
        }
        console.log(this.students)
        // this.questionsNumbers = data.size
        // for(var i = 0; i < data.data.length; i++){
        //   this.subjects.push({
        //     studentUsername: data.data[i].studentUsername,
        //     presentationName: data.data[i].presentationName,
        //     classification: data.data[i].classification
        //   })
        //   // this.classification = data.data.slice(-1).pop()        
        // }
        // console.warn(this.presentationData)
      })
    }
  }

  // getAllPresentations(id){
  //   console.warn(this)
  //   this.presentationService.getAllPresentations(id).subscribe(data =>{
  //     console.warn(data.data)
  //     console.warn(data)
  //     if(data.data != undefined)
  //       this.presentationService.presentations = data.data
  //    })
  // }

  // getGrades(){
  //   let user = this.tokenService.getUser()
  //   this.presentationService.getClassificationByPresentation(user.username).subscribe(data =>{
  //     console.warn(data.data)
  //     console.warn(data)
  //   })
  // }

  // selectedSubject(event: any){
  //   console.warn(event.target.value)
  //   console.warn(event.target.classList)
    
  // }

}
