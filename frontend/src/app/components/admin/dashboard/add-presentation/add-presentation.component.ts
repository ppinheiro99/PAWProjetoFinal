import { Component, Inject, OnInit } from '@angular/core';
import { MatDialog, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { PresentationsService } from 'src/app/services/presentations/presentations.service';

@Component({
  selector: 'app-add-presentation',
  templateUrl: './add-presentation.component.html',
  styleUrls: ['./add-presentation.component.scss']
})
export class AddPresentationComponent implements OnInit {
  presentationName: string;
  questionsNumber = '';
  responsesNumber = '';
  questions : any;
  count  = 0;
  selectResponses = '';
  inputValue : any;
  inputValueId : any;
  varCorrectResponses = '';
  questionIdAndRespondes = [];
  correctResponseAndId = [];
  pdf: File
  subjectId : any;
  constructor(private route: ActivatedRoute, private presentationsService : PresentationsService, private dialog: MatDialog, @Inject(MAT_DIALOG_DATA) data) { 
    this.subjectId = data
  }
 

  ngOnInit(): void {}

  getPdf(event){
    console.warn(event[0])
    this.pdf = event[0];
    console.log(this.pdf)
  }

  responsesSelect(i){
    if(this.questionIdAndRespondes.length > 0){
      if(this.containsObject({id : i, n : this.selectResponses}, this.questionIdAndRespondes))
        this.questionIdAndRespondes.push({id : i, n : this.selectResponses});
    }else
      this.questionIdAndRespondes.push({id : i, n : this.selectResponses});
    this.selectResponses = '';
    console.log(this.questionIdAndRespondes)
  }

  correctResponse(i){
    console.warn(i)
    if(this.correctResponseAndId.length > 0){
      if(this.containsObject({id : i, n : this.varCorrectResponses},this.correctResponseAndId))
        this.correctResponseAndId.push({id : i, n : this.varCorrectResponses});
    }else
      this.correctResponseAndId.push({id : i, n : this.varCorrectResponses});
    this.varCorrectResponses = '';
    console.log(this.correctResponseAndId)
  }

  containsObject(obj, list) {
    for (var i = 0; i < list.length; i++) {
        if (list[i].id === obj.id) {
          list.n = obj.n;
          return false;
        }
    }
    return true;
  }
  
  counter(i: string) {
    return new Array(Number(i));
  }

  addPresentation(){
    var questions = ""
    for (var i = 0; i < this.questionIdAndRespondes.length; i++) {
      console.log((<HTMLInputElement>document.getElementById("page-number"+ i)).value);
      
      if(i == 0)
      questions += "questions.number=" +  ((<HTMLInputElement>document.getElementById("page-number"+ i)).value) + "e" +  (((<HTMLInputElement>document.getElementById("page-number"+ i)).value)+1);
      // questions += "questions.question=" + (<HTMLInputElement>document.getElementById("question"+ i)).value;
      else
      questions += "&questions.number=" +  ((<HTMLInputElement>document.getElementById("page-number"+ i)).value) + "e" +  (((<HTMLInputElement>document.getElementById("page-number"+ i)).value)+1);
      questions += "&"+"questions.question=" + (<HTMLInputElement>document.getElementById("question"+ i)).value;
      for (var j = 0; j < this.questionIdAndRespondes[i].n; j++) {
        // console.log((<HTMLInputElement>document.getElementById("response"+i + j)).value);
        questions += "&"+"questions.answer=" + (<HTMLInputElement>document.getElementById("response"+i + j)).value;
      }
      for (var j = 0; j < this.correctResponseAndId.length; j++) {
        if(this.correctResponseAndId[j].id == this.questionIdAndRespondes[i].id){
          // console.log("resposta:" , (<HTMLInputElement>document.getElementById("response"+i + this.correctResponseAndId[j].n)).value)
          questions += "&"+"questions.correct_answer=" + (<HTMLInputElement>document.getElementById("response"+i + this.correctResponseAndId[j].n)).value;
        } 
      }

    }

    console.log(this.subjectId)
    // console.warn(this.presentationName)
    // console.warn(this.pdf)
    // console.warn(questions)
    // console.warn(2)

    this.presentationsService.addPresentation(this.presentationName,this.pdf,questions,this.subjectId.id).subscribe(
      data => {
       console.log(data)
       this.getAllPresentationsData()
       this.dialog.closeAll()
      },
      err => {
        alert(err)
        this.dialog.closeAll()
      }
    );
  }

  getAllPresentationsData(){
    this.presentationsService.getAllPresentations(this.subjectId.id).subscribe(data =>{
      this.presentationsService.presentations = data.data
      console.warn(this.presentationsService.presentations)
    })
  }

}
