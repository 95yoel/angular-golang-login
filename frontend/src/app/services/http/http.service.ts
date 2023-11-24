import { HttpHeaders,HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { UserLog } from '../../interfaces/user-log';
import { environment } from '../../../environments/environment.development';
import { UserReg } from '../../interfaces/user-reg';
import { Router } from '@angular/router';
import { AuthService } from '../auth/auth.service';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  private headers = new HttpHeaders({
    'Content-Type': 'application/json',
    'Authorization': 'Bearer tuToken',
  });

  constructor(private http:HttpClient,private auth:AuthService, private router:Router) { }

  /*
  POST THE FORM TO THE LOGIN ENDPOINT
  */
  login(formData:UserLog){

    return this.http.post(`${environment.URL_BACKEND}/${environment.LOGIN_ENDPOINT}`,formData,{
      headers:this.headers,
      responseType:'text'
    });

  }
  /*
  POST THE FORM TO THE REGISTER ENDPOINT
  */
  register(formData:UserReg){
    
    return this.http.post(`${environment.URL_BACKEND}/${environment.REGISTER_ENDPOINT}`,formData,{
      headers:this.headers,
      responseType:'text'
    })
  }

  /*
  HANDLE THE RESPONSE OF THE LOGIN
  */

  handleLogin(res:any):void{
    this.auth.setTokenJWT(res);
    this.auth.isValidJWT();
    this.router.navigate(['/private']);
  }

}
