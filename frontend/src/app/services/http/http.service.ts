import { HttpHeaders,HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { UserLog } from '../../interfaces/user-log';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  private headers = new HttpHeaders({
    'Content-Type': 'application/json',
    'Authorization': 'Bearer tuToken',
  });

  constructor(private http:HttpClient) { }

  login(formData:UserLog){

    return this.http.post('http://localhost:8080/login',formData,{
      headers:this.headers,responseType:'text'
    });

  }

  register(){

  }


}
