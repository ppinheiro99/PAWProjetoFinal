import { Component, ViewChild, OnInit, Output, EventEmitter, ElementRef, AfterViewInit } from '@angular/core';
import { Subject } from 'rxjs';
import WebViewer, { WebViewerInstance } from '@pdftron/webviewer';
import { PresentationsService } from 'src/app/services/presentations/presentations.service';
import { Router } from '@angular/router';
import { ChatService } from 'src/app/services/socket/chat.service';

@Component({
  selector: 'app-presentation-pdf',
  templateUrl: './presentation-pdf.component.html',
  styleUrls: ['./presentation-pdf.component.scss']
})
export class PresentationPdfComponent implements OnInit, AfterViewInit {
  @ViewChild('viewer') viewer: ElementRef;
  wvInstance: WebViewerInstance;
  @Output() coreControlsEvent:EventEmitter<string> = new EventEmitter();

  private documentLoaded$: Subject<void>;

  constructor(private router: Router,public presentationService: PresentationsService, public chat: ChatService) { 
    this.documentLoaded$ = new Subject<void>();
  }

  ngAfterViewInit(): void {
    const base64 = this.presentationService.presentation.pdf_file;
    let base64String = this.presentationService.presentation.pdf_file;
    const source = `data:application/pdf;base64,${base64String}`;
    
    const link = document.createElement("a");
    link.href = source;
    console.log(link.href);
    console.log(source)
    // console.warn(this.base64ToBlob(base64String))
    
     WebViewer({
       path: "../../../../../assets/lib",
        // initialDoc: 'https://pdftron.s3.amazonaws.com/downloads/pl/webviewer-demo.pdf',
        initialDoc: 'http://localhost:8080/api/v1/presentations/'+this.presentationService.presentation.ID

     }, this.viewer.nativeElement)
     .then(instance => {
         instance.UI.setTheme('dark');
         const { documentViewer, annotationManager } = instance.Core;
         documentViewer.addEventListener('pageNumberUpdated', pageNumber => {
          console.log(pageNumber)
          this.chat.sendPageNumber(pageNumber)
        });
         
     });
  }

  ngOnInit(): void {
    if(this.presentationService.presentation == undefined)
      this.router.navigate([''])
    this.presentationService.getPresentationInfo(this.presentationService.presentation.ID).subscribe(data =>{
      console.log(data.data)
      for(var i = 0; i < data.data.length; i++){
        console.warn(data.data[i].answers)
        const split1 = data.data[i].answers.split(/[,_]+/);
        const result1 =  split1.filter(e =>  e);
        const split2 = data.data[i].question_number.split(/[e_]+/);
        const result2 =  split2.filter(e =>  e);
        console.log(result2[1])
        this.presentationService.presentationData.push({
          id: data.data[i].ID,
          answers: result1,
          correctAnswer: data.data[i].correct_answer,
          questionNumber: result2[1],
          question: data.data[i].question
        })
      }
      console.log(this.presentationService.presentationData)
      
      this.chat.sendPresentationData(this.presentationService.presentationData)
    })
  }

}
