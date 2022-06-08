import { Component, ViewChild, OnInit, Output, EventEmitter, ElementRef, AfterViewInit } from '@angular/core';
import { Subject } from 'rxjs';
import WebViewer, { WebViewerInstance } from '@pdftron/webviewer';
import { PresentationsService } from 'src/app/services/presentations/presentations.service';
import { Router } from '@angular/router';

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

  constructor(private router: Router,public presentationService: PresentationsService) { 
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
          
        });
         
     });
  }

  ngOnInit(): void {
    if(this.presentationService.presentation == undefined)
      this.router.navigate([''])
  }

}
