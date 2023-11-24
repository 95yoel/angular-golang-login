import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { UserLog } from '../../interfaces/user-log';
import { HttpService } from '../../services/http/http.service';
import { ValidationRegService } from '../../services/validation-reg/validation-reg.service';


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
  constructor(private http:HttpService,private validator:ValidationRegService){}
  
  onSubmit(){
    // Validate the form data using the validation service
    this.validator.validationLogin(this.formData).then(
      ()=>{
        // Call the HTTP service to log the user 
        this.http.login(this.formData).subscribe(
          (res)=>{
            console.log(res);
          }
        );
      }
    ).catch(
      //Log validation errors
      (error)=>{
        console.warn("Error :\n",error);
      }
    );
  }

}
