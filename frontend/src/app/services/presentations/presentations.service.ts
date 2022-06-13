import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { TokenService } from '../token/token.service';

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
  presentationData: {answers:any,correctAnswer:any,questionNumber:any,question:any, id:any}[] = []
  constructor(private tokenService:TokenService,private http: HttpClient) { }
  user = this.tokenService.getUser()

  getAllPresentations(subjectID): Observable<any> {
    return this.http.get(API_URL +"getAll/" + subjectID, httpOptions)
  }

  getPresentationById(id): Observable<any> {
    return this.http.get(API_URL + id, httpOptionsPdf)
  }

  getPresentationInfo(id): Observable<any> {
    return this.http.get(API_URL + id + "/questions", httpOptionsPdf)
  }

  getClassificationByPresentation(username): Observable<any> {
    return this.http.get(API_URL + "getClassificationByPresentation/" + username , httpOptions)
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

  sendUserResponse(data, id): Observable<any> {
    return this.http.post(API_URL +"submitAnswer", {
      student_username: this.user.username,
      answer: data,
      question_id: id
    }, httpOptions);
  }
    
  deletePresentation(id): Observable<any> {
    return this.http.delete(API_URL + id, httpOptions);
  }

}
