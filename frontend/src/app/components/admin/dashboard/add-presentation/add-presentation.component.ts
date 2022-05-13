import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-add-presentation',
  templateUrl: './add-presentation.component.html',
  styleUrls: ['./add-presentation.component.scss']
})
export class AddPresentationComponent implements OnInit {
  presentationName: string;
  questionsNumber = '';
  responsesNumber = '';
  questions : any
  count  = 0;

  constructor() { }

  ngOnInit(): void {
  }

  addPresentation(){
  //  console.log(document.getElementById('questions1').getAttribute);
    console.log((<HTMLInputElement>document.getElementById("question1")).value);
  }

  counter(i: string) {
    return new Array(Number(i));
  }

}
