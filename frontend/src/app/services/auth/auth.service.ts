import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

//const API_URL = 'http://18.130.231.194:8080/api/';
const API_URL = 'http://localhost:8080/api/v1/';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(private http: HttpClient) { }

  login(credentials): Observable<any> {
    console.log(credentials.username, credentials.password)
    return this.http.post(API_URL + 'auth/login', {
      username: credentials.username,
      password: credentials.password
    }, httpOptions);
  }
  

  passRecoverVerififyToken(token): Observable<any> {
    return this.http.post(API_URL + 'auth/passRecover/checktoken', {
      token: token,
    }, httpOptions);
  }

}
