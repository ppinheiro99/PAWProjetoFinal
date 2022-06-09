import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

//const API_URL = 'http://18.130.231.194:8080/api/';
const API_URL = 'http://localhost:8080/api/v1/presentations/';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*',
  })
};

const httpOptionsPdf = {
  headers: new HttpHeaders({
    'Content-Type': 'application/pdf',
    'Content-Transfer-Encoding': 'binary',
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})
export class PresentationsService {
  presentations : any[]
  presentation: any
  constructor(private http: HttpClient) { }

  getAllPresentations(subjectID): Observable<any> {
    return this.http.get(API_URL +"getAll/" + subjectID, httpOptions)
  }

  getPresentationById(id): Observable<any> {
    return this.http.get(API_URL + id, httpOptionsPdf)
  }

  addPresentation(name, pdf_file,questions,subjectId): Observable<any> {
    console.warn(name, pdf_file,questions,subjectId)
    
    let formDataBody = new FormData()
    formDataBody.append("name", name)
    formDataBody.append("pdf_file", pdf_file)
    formDataBody.append("questions", questions)
    formDataBody.append("subjectid", subjectId)
    
    console.warn(formDataBody)
    return this.http.post(API_URL, formDataBody);
  }
    
  deletePresentation(id): Observable<any> {
    return this.http.delete(API_URL + id, httpOptions);
  }

}
