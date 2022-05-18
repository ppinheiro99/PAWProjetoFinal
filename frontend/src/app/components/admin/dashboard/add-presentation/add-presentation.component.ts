import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
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
  constructor(private presentationsService : PresentationsService, private dialog: MatDialog) { }

  ngOnInit(): void {
  }

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
      //  console.log((<HTMLInputElement>document.getElementById("question"+ i)).value);
      if(i == 0)
      questions += "questions.question=" + (<HTMLInputElement>document.getElementById("question"+ i)).value;
      else
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
    console.warn(this.presentationName)
    console.warn(this.pdf)
    console.warn(questions)
    console.warn(2)

    this.presentationsService.addPresentation(this.presentationName,this.pdf,questions,2).subscribe(
      data => {
       console.log(data)
       this.dialog.closeAll()
      },
      err => {
        alert(err)
        this.dialog.closeAll()
      }
    );
  }

}
