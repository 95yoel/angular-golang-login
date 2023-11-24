import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { UserLog } from '../../interfaces/user-log';
import { HttpService } from '../../services/http/http.service';


@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule,FormsModule,RouterModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {

  formData:UserLog={
    email : '',
    password: ''
  }
  constructor(private http:HttpService){}
  
  onSubmit(){
    console.log(this.formData);
    this.http.login(this.formData).subscribe(
      (res)=>{
        console.log(res);
      }
    );

  }

}
