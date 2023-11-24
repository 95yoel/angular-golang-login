import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UserStorage } from '../../interfaces/user-storage';
import { AuthService } from '../../services/auth/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-private',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './private.component.html',
  styleUrl: './private.component.css'
})
export class PrivateComponent {

  formData: UserStorage = {
    exp: '',
    name: '',
    surname: '',
    email: '',
  };

  constructor(private auth:AuthService,private router:Router){}

  ngOnInit(){
    this.jwtInfo();
    console.log(this.formData);
  }


  jwtInfo(){
    this.formData.exp = this.auth.getLocalStorage("exp_token");
    this.formData.name = this.auth.getLocalStorage("user_name");
    this.formData.surname = this.auth.getLocalStorage("user_surname");
    this.formData.email = this.auth.getLocalStorage("user_email"); 
  }

  exit(){
    this.auth.deauthorize();
    this.router.navigate(['/']);
  }

}
