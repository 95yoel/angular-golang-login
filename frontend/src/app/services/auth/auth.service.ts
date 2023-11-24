import { Injectable } from '@angular/core';
import { jwtDecode } from 'jwt-decode';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private authTokenKey ="auth_token";
  private expTokenKey = "exp_token";
  private nameTokenKey = "user_name";
  private surnameTokenKey = "user_surname";
  private emailTokenKey = "user_email";

  constructor() { }

  setTokenJWT(token:string):void{
    localStorage.setItem(this.authTokenKey,token);
  }

  getTokenJWT():string|null{
    return localStorage.getItem(this.authTokenKey);
  }

  removeTokenJWT():void{
    localStorage.removeItem(this.authTokenKey);
  }

  decodeToken(): any | null {
    const token = this.getTokenJWT();
    return token ? jwtDecode(token) : null;
  }

  isValidJWT():boolean{
    const decodedToken = this.decodeToken();
    if (decodedToken) {
      const expirationTime = decodedToken.exp * 1000;
      localStorage.setItem(this.expTokenKey, String(new Date(expirationTime)));
      localStorage.setItem(this.nameTokenKey,decodedToken.name);
      localStorage.setItem(this.surnameTokenKey,decodedToken.surname);
      localStorage.setItem(this.emailTokenKey,decodedToken.email);
      return expirationTime > Date.now();
    }
    return false;
  }
  deauthorize(){
    localStorage.removeItem(this.authTokenKey);
    localStorage.removeItem(this.nameTokenKey);
    localStorage.removeItem(this.surnameTokenKey);
    localStorage.removeItem(this.emailTokenKey);
    localStorage.removeItem(this.expTokenKey);
  }
  
  getLocalStorage(item:string):string{
    return localStorage.getItem(item)||'';
  }

}
