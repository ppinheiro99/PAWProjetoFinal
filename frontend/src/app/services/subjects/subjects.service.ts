import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

//const API_URL = 'http://18.130.231.194:8080/api/';
const API_URL = 'http://localhost:8080/api/v1/subject/';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})
export class SubjectsService {
  subjects : any[]
  constructor(private http: HttpClient) { }

  getAllSubjects(): Observable<any> {
    return this.http.get(API_URL, httpOptions)
  }

  addSubject(data): Observable<any> {
    return this.http.post(API_URL, {
      name: data,
    }, httpOptions);
  }
  
}
