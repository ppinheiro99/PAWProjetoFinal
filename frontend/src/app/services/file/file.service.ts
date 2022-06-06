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

@Injectable({
  providedIn: 'root'
})
export class FileService {
    // API url
    baseApiUrl = "https://file.io"
    
    constructor(private http:HttpClient) { }
    
    // Returns an observable
    upload(file):Observable<any> {
    
        // Create form data
        const formData = new FormData(); 
          
        // Store form name as "file" with file data
        formData.append("file", file, file.name);
          
        // Make http post request over api
        // with formData as req
        return this.http.post(this.baseApiUrl, formData)
    }
  
}
